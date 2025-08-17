<template>
  <div class="comment-section" v-loading="isLoading">
    <div v-if="isLoggedIn && currentUserProfile" class="comment-input-area">
      <el-avatar :size="40" :src="currentUserProfile.avatar" />
      <div class="input-wrapper">
        <div v-if="replyTarget" class="reply-to-bar">
          <span>正在回复 @{{ replyTarget.name || replyTarget.username }}</span>
          <el-button text :icon="Close" circle @click="cancelReply" class="cancel-reply-btn" />
        </div>
        <el-input
          v-model="newCommentText"
          type="textarea"
          :autosize="{ minRows: 2, maxRows: 5 }"
          :placeholder="replyTarget ? `回复TA...` : '发表一条友善的评论...'"
          class="comment-textarea"
          ref="commentInputRef"
        />
        <div class="actions">
          <el-button type="primary" @click="submitComment" :disabled="!newCommentText.trim()" :loading="isSubmitting"
                     round>
            {{ replyTarget ? '回复' : '评论' }}
          </el-button>
        </div>
      </div>
    </div>
    <div v-else class="login-prompt">
      <el-button text @click="goToLogin">请登录后发表评论</el-button>
    </div>

    <div v-if="comments.length > 0" class="comment-list">
      <div v-for="comment in comments" :key="comment.id" class="comment-item">
        <el-avatar :size="40" :src="comment.avatar" class="comment-avatar clickable"
                   @click="goToUserProfile(comment.user_id)" />
        <div class="comment-body">
          <div class="comment-header">
            <span class="username clickable"
                  @click="goToUserProfile(comment.user_id)">{{ comment.name || comment.username
              }}</span>
            <span class="timestamp">{{ formatTime(comment.created_at) }}</span>
          </div>
          <p class="comment-text">{{ comment.text }}</p>
          <div class="comment-footer">
            <el-button text :icon="ArrowUpBold" @click="toggleLikeComment(comment)">
              {{ comment.like_count || 0 }}
            </el-button>
            <el-button text @click="setReplyTarget(comment)">回复</el-button>
            <el-button v-if="comment.user_id === currentUserId" text @click="deleteComment(comment.id)"
                       class="delete-btn">
              删除
            </el-button>
          </div>

          <div v-if="comment.replies && comment.replies.length > 0" class="reply-list">
            <div v-for="reply in comment.replies" :key="reply.id" class="comment-item reply-item">
              <el-avatar :size="32" :src="reply.avatar" class="comment-avatar clickable"
                         @click="goToUserProfile(reply.user_id)" />
              <div class="comment-body">
                <div class="comment-header">
                  <span class="username clickable"
                        @click="goToUserProfile(reply.user_id)">{{ reply.name || reply.username
                    }}</span>
                  <el-icon class="reply-arrow">
                    <CaretRight />
                  </el-icon>
                  <span class="username reply-to-user clickable"
                        @click="goToUserProfile(reply.reply_to_user_id)">{{ reply.reply_to_name || reply.reply_to_username
                    }}</span>
                  <span class="timestamp">{{ formatTime(reply.created_at) }}</span>
                </div>
                <p class="comment-text">{{ reply.text }}</p>
                <div class="comment-footer">
                  <el-button text :icon="ArrowUpBold" @click="toggleLikeComment(reply)">{{ reply.like_count || 0 }}
                  </el-button>
                  <el-button text @click="setReplyTarget(comment, reply)">回复</el-button>
                  <el-button v-if="reply.user_id === currentUserId" text @click="deleteComment(reply.id)"
                             class="delete-btn">删除
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <el-pagination
      v-if="comments.length > 0"
      v-model:current-page="page"
      v-model:page-size="pageSize"
      :total="total"
      :page-sizes="[10,20,30,50]"
      layout="total, sizes, prev, pager, next, jumper"
      background
      class="pagination"
      @size-change="handlePageSizeChange"
      @current-change="handlePageChange"
    />

    <el-empty v-if="!isLoading && comments.length === 0" description="还没有评论，快来抢沙发吧！" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { commentApi, userApi } from '@/api';
import type { ModelsCommentDTO, ModelsUserProfile } from '@/api-client';
import { ArrowUpBold, Close, CaretRight } from '@element-plus/icons-vue';
// 假设你有一个时间格式化工具
// import { formatDistanceToNow } from 'date-fns';
import { zhCN } from 'date-fns/locale';
import { formatRelativeTime } from '@/utils/format';


const formatTime = formatRelativeTime;
const props = defineProps<{ songId: number }>();
const router = useRouter();
const userStore = useUserStore();
const { userId: currentUserId, isLoggedIn } = storeToRefs(userStore); // 从 store 获取 isLoggedIn


const isLoading = ref(true);
const isSubmitting = ref(false);
const comments = ref<ModelsCommentDTO[]>([]);
const page = ref(1);
const pageSize = ref(20);
const total = ref(0);

const newCommentText = ref('');
const commentInputRef = ref();
const replyTarget = ref<{ parentId: number; replyToUserId: number; name: string; username: string; } | null>(null);

const currentUserProfile = ref<ModelsUserProfile | null>(null);
// const formatTime = (timeStr: string) => formatDistanceToNow(new Date(timeStr), { addSuffix: true, locale: zhCN });

const fetchCurrentUserProfile = async () => {
  const userResponse = await userApi.usersIdGet(currentUserId.value);
  currentUserProfile.value = userResponse.data;
};

const loadData = async () => {
  if (!props.songId) return;
  isLoading.value = true;
  try {
    const res = await commentApi.songsIdCommentsGet(props.songId, page.value, pageSize.value);
    const rootComments: ModelsCommentDTO[] = res.data.list || [];
    total.value = res.data.total || 0;

    // 获取每个根评论的回复
    for (const comment of rootComments) {
      const repliesRes = await commentApi.commentsCommentIdRepliesGet(comment.id!);
      comment.replies = repliesRes.data || [];
    }
    comments.value = rootComments;
  } catch (error) {
    console.error('获取评论失败:', error);
    ElMessage.error('获取评论失败');
  } finally {
    isLoading.value = false;
  }
};

const submitComment = async () => {
  isSubmitting.value = true;
  try {
    await commentApi.songsIdCommentsPost(props.songId, {
      text: newCommentText.value,
      parent_id: replyTarget.value?.parentId,
      reply_to_user_id: replyTarget.value?.replyToUserId,
    });
    ElMessage.success(replyTarget.value ? '回复成功' : '评论成功');
    newCommentText.value = '';
    replyTarget.value = null;
    await loadData(); // 重新加载第一页
  } catch (error) {
    ElMessage.error('操作失败');
  } finally {
    isSubmitting.value = false;
  }
};

const toggleLikeComment = async (comment: ModelsCommentDTO) => {
  try {
    const res = await commentApi.commentsCommentIdLikePost(comment.id!);
    // 其实应该设置不同的状态的图标
    comment.like_count = res.data.like_count;
  } catch (error) {
    ElMessage.error('点赞失败');
  }
};

const deleteComment = async (commentId: number) => {
  try {
    await commentApi.commentsCommentIdDelete(commentId);
    ElMessage.success('删除成功');
    await loadData();
  } catch (error) {
    ElMessage.error('删除失败');
  }
};

const setReplyTarget = (parentComment: ModelsCommentDTO, replyToComment: ModelsCommentDTO | null = null) => {
  const target = replyToComment || parentComment;
  replyTarget.value = {
    parentId: parentComment.id!,
    replyToUserId: target.user_id!,
    name: target.name!,
    username: target.username!,
  };
  nextTick(() => commentInputRef.value?.focus());
};

const cancelReply = () => {
  replyTarget.value = null;
};

const handlePageChange = (newPage: number) => {
  page.value = newPage;
  loadData(false);
};

const handlePageSizeChange = (newSize: number) => {
  pageSize.value = newSize;
  page.value = 1;
  loadData();
};

const goToUserProfile = (userId?: number) => {
  if (userId) {
    router.push(`/profile/${userId}`);
  }
};

const goToLogin = () => {
  router.push('/auth');
}

onMounted(async () => {
  // 使用 Promise.all 并行执行，提升速度
  await Promise.all([
    fetchCurrentUserProfile(),
    loadData(),
  ]);
  console.log(total.value);
  console.log(pageSize.value);
});
</script>

<style scoped>
/* Spotify风格评论区 */
.comment-section {
  padding: 16px;
}

/* 评论输入 */
.comment-input-area {
  display: flex;
  gap: 16px;
  margin-bottom: 32px;
}

.input-wrapper {
  flex-grow: 1;
}

/* --- [THE FIX #1] 优化文本编辑区样式 --- */
.comment-textarea :deep(.el-textarea__inner) {
  background-color: #282828; /* 使用组件背景色 */
  box-shadow: none;
  border: 1px solid #535353; /* 增加边框 */
  border-radius: 6px; /* 增加圆角 */
  padding: 10px 12px; /* 增加内边距 */
  color: #fff;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}

.comment-textarea :deep(.el-textarea__inner:focus) {
  border-color: #1db954; /* 聚焦时使用 Spotify 绿 */
  box-shadow: 0 0 0 1px rgba(29, 185, 84, 0.3); /* 增加外发光效果 */
}

.reply-to-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  color: #b3b3b3;
  margin-bottom: 8px;
}

.cancel-reply-btn {
  color: #b3b3b3;
}

.actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 12px;
}

.actions .el-button {
  border-radius: 500px;
  font-weight: 700;
}

.actions .el-button--primary {
  background-color: #1db954;
  border-color: #1db954;
  color: #000;
}

/* 评论列表 */
.comment-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.comment-item {
  display: flex;
  gap: 16px;
}

.comment-avatar {
  flex-shrink: 0;
}

.comment-body {
  flex-grow: 1;
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.username {
  font-size: 14px;
  font-weight: 500;
  color: #fff;
}

.reply-to-user {
  color: #b3b3b3;
}

.reply-arrow {
  color: #b3b3b3;
}

.timestamp {
  font-size: 12px;
  color: #b3b3b3;
}

.comment-text {
  font-size: 14px;
  color: #b3b3b3;
  line-height: 1.6;
  white-space: pre-wrap;
  margin: 0;
}

.comment-footer {
  margin-top: 8px;
}

.comment-footer .el-button {
  color: #b3b3b3;
  font-size: 12px;
}

.comment-footer .el-button:hover {
  color: #fff;
}

.delete-btn {
  color: #f15e6c;
}

.delete-btn:hover {
  color: #f15e6c;
  background-color: rgba(241, 94, 108, 0.1);
}

/* 回复列表 */
.reply-list {
  margin-top: 16px;
  border-left: 2px solid #282828;
  padding-left: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 分页器 */
.pagination {
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  justify-content: center;
  --el-pagination-bg-color: transparent;
  --el-pagination-text-color: #b3b3b3;
  --el-pagination-hover-color: #fff;
  --el-pagination-button-disabled-bg-color: transparent;
}

:deep(.el-pagination.is-background .btn-prev),
:deep(.el-pagination.is-background .btn-next),
:deep(.el-pagination.is-background .el-pager li) {
  background-color: transparent;
  color: #b3b3b3;
}

:deep(.el-pagination.is-background .el-pager li:hover) {
  color: #fff;
}

:deep(.el-pagination.is-background .el-pager li.is-active) {
  background-color: #1db954;
  color: #000;
  font-weight: 600;
}

/* [THE FIX] Add styles for clickable elements */
.clickable {
  cursor: pointer;
  transition: opacity 0.2s ease;
}

.clickable:hover {
  opacity: 0.8;
}

.username.clickable:hover {
  text-decoration: underline;
}
</style>