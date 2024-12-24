<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="main-container">
        <!-- 左侧歌曲信息 -->
        <div class="song-aside">
          <div class="song-info-card">
            <div class="song-cover">
              <el-image 
                :src="songInfo?.cover || '/assets/default-cover.jpg'"
                fit="cover"
              />
            </div>
            <h2>{{ songInfo?.name }}</h2>
            <div class="song-info">
              <span>歌手：{{ songInfo?.artist_name }}</span>
              <span>专辑：{{ songInfo?.album_name }}</span>
            </div>
          </div>
        </div>

        <!-- 右侧评论区 -->
        <div class="comment-container">
          <!-- 统一的评论输入框 -->
          <el-card class="comment-input">
            <div class="input-area">
              <!-- 添加回复对象提示 -->
              <div v-if="activeReplyTarget" class="reply-target">
                <el-tag size="small" class="reply-tag" type="info" closable @close="cancelReply">
                  <el-icon><ChatDotRound /></el-icon>
                  回复 @{{ activeReplyTarget.username }}
                </el-tag>
              </div>
              <el-input
                v-model="newComment"
                type="textarea"
                :rows="3"
                :placeholder="activeReplyTarget ? `回复 @${activeReplyTarget.username}...` : '发表评论...'"
              />
              <div class="comment-actions">
                <el-button v-if="activeReplyTarget" @click="cancelReply">取消回复</el-button>
                <el-button 
                  type="primary" 
                  @click="activeReplyTarget ? submitReply() : submitComment()" 
                  :disabled="!newComment.trim()"
                >
                  {{ activeReplyTarget ? '回复' : '发表评论' }}
                </el-button>
              </div>
            </div>
          </el-card>

          <!-- 评论列表 -->
          <el-card class="comment-list" v-loading="isLoading">
            <div v-for="comment in comments" :key="comment.id" class="comment-item">
              <!-- 主评论 -->
              <div class="comment-main">
                <el-avatar :src="comment.avatar" />
                <div class="comment-content">
                  <div class="comment-header">
                    <span class="username">{{ comment.username }}</span>
                    <span class="time">{{ formatTime(comment.created_at) }}</span>
                  </div>
                  <p class="text">{{ comment.text }}</p>
                  <div class="comment-actions">
                    <el-button type="text" @click="handleLikeComment(comment)">
                      <el-icon><ArrowUpBold /></el-icon>
                      {{ comment.like_count || 0 }}
                    </el-button>
                    <el-button type="text" @click="showReplyInput(comment)">
                      <el-icon><ChatLineRound /></el-icon>
                      回复
                    </el-button>
                    <el-button 
                      v-if="comment.user_id === currentUser.id" 
                      type="text" 
                      @click="deleteComment(comment.id)"
                    >
                      删除
                    </el-button>
                  </div>

                  <!-- 嵌套的回复列表 -->
                  <div v-if="comment.replies && comment.replies.length > 0" class="replies">
                    <template v-for="reply in comment.replies" :key="reply.id">
                      <div class="reply-item">
                        <el-avatar :size="30" :src="reply.avatar" />
                        <div class="reply-content">
                          <div class="reply-header">
                            <span class="username">{{ reply.username }}</span>
                            <span class="time">{{ formatTime(reply.created_at) }}</span>
                          </div>
                          <p class="text">
                            <el-tag size="small" class="reply-tag" type="info">
                              <el-icon><ChatDotRound /></el-icon>
                              @{{ reply.reply_to_username || comment.username }}
                            </el-tag>
                            {{ reply.text }}
                          </p>
                          <div class="reply-actions">
                            <el-button type="text" @click="handleLikeComment(reply)">
                              <el-icon><ArrowUpBold /></el-icon>
                              {{ reply.like_count || 0 }}
                            </el-button>
                            <el-button type="text" @click="showReplyInput(comment, reply)">
                              <el-icon><ChatLineRound /></el-icon>
                              回复
                            </el-button>
                            <el-button 
                              v-if="reply.user_id === currentUser.id" 
                              type="text" 
                              @click="handleDeleteComment(reply.id)"
                            >
                              删除
                            </el-button>
                          </div>
                        </div>
                      </div>
                    </template>
                  </div>
                </div>
              </div>
            </div>

            <!-- 分页 -->
            <div class="pagination">
              <el-pagination
                v-model:current-page="currentPage"
                v-model:page-size="pageSize"
                :total="total"
                :page-sizes="[5, 10, 20, 50]"
                layout="total, sizes, prev, pager, next"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
              />
            </div>
          </el-card>
        </div>
      </div>
    </template>
  </CommonLayout>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowUpBold, ChatLineRound, ChatDotRound } from '@element-plus/icons-vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { 
  getComments, 
  getCommentReplies, 
  createComment, 
  likeComment, 
  deleteComment,
  getSongById  // 添加获取歌曲信息的接口
} from '@/api/axiosFile'
import { useUserStore } from '@/stores/user'

const currentName = ref('评论详情')
const route = useRoute()
const userStore = useUserStore()
const currentUser = userStore.userInfo
const isLoading = ref(false)

// 状态变量
const songInfo = ref(null)
const comments = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const newComment = ref('')
const activeReplyTarget = ref(null)

// 获取歌曲信息
const fetchSongInfo = async () => {
  try {
    const resp = await getSongById(route.params.id)
    if (resp.data.message) {
      songInfo.value = resp.data.data
    }
  } catch (error) {
    console.error('获取歌曲信息失败:', error)
    ElMessage.error('获取歌曲信息失败')
  }
}

// 获取评论列表
const fetchComments = async () => {
  try {
    isLoading.value = true
    const params = {
      page: currentPage.value,
      pageSize: pageSize.value
    }
    const resp = await getComments(route.params.id, params)
    if (resp.data.message) {
      comments.value = resp.data.data.comments
      total.value = resp.data.data.total
      
      // 获取每个评论的回复
      for (const comment of comments.value) {
        const repliesResp = await getCommentReplies(comment.id)
        if (repliesResp.data.message) {
          comment.replies = repliesResp.data.data.replies || []
          console.log('评论回复:', comment.replies) // 添加调试日志
        }
      }
    }
    console.log('所有评论:', comments.value) // 添加调试日志
  } catch (error) {
    console.error('获取评论失败:', error)
    ElMessage.error('获取评论失败')
  } finally {
    isLoading.value = false
  }
}

// 提交评论
const submitComment = async () => {
  try {
    const resp = await createComment({
      songId: route.params.id,
      text: newComment.value
    })
    if (resp.data.message) {
      ElMessage.success('评论成功')
      newComment.value = ''
      await fetchComments()
    }
  } catch (error) {
    console.error('发表评论失败:', error)
    ElMessage.error('发表评论失败')
  }
}

// 提交回复
const submitReply = async () => {
  if (!activeReplyTarget.value) return
  
  try {
    const data = {
      songId: route.params.id,
      text: newComment.value,
      parentId: activeReplyTarget.value.commentId,
      replyToUserId: activeReplyTarget.value.id
    }
    
    const resp = await createComment(data)
    if (resp.data.message) {
      ElMessage.success('回复成功')
      newComment.value = ''
      activeReplyTarget.value = null
      await fetchComments()
    }
  } catch (error) {
    console.error('回复评论失败:', error)
    ElMessage.error('回复评论失败')
  }
}

// 点赞评论
const handleLikeComment = async (comment) => {
  try {
    // console.log(comment);
    await likeComment(comment.id)
    await fetchComments() // 重新获取评论列表以更新点赞状态
  } catch (error) {
    console.error('点赞失败:', error)
    ElMessage.error('点赞失败')
  }
}

// 删除评论
const handleDeleteComment = async (commentId) => {
  try {
    const resp = await deleteComment(commentId)
    if (resp.message) {
      ElMessage.success('删除成功')
      fetchComments()
    }
  } catch (error) {
    console.error('删除评论失败:', error)
    ElMessage.error('删除评论失败')
  }
}

// 修改显示回复输入框的逻辑
const showReplyInput = (comment, replyTo = null) => {
  // 如果点击的是同一个回复目标，则取消回复
  if (activeReplyTarget.value?.id === (replyTo?.user_id || comment.user_id)) {
    cancelReply()
    return
  }
  
  if (replyTo) {
    // 回复某条回复
    activeReplyTarget.value = {
      id: replyTo.user_id,
      username: replyTo.username,
      commentId: comment.id // 保存父评论ID
    }
  } else {
    // 回复主评论
    activeReplyTarget.value = {
      id: comment.user_id,
      username: comment.username,
      commentId: comment.id
    }
  }
  
  // 滚动到输入框
  document.querySelector('.comment-input')?.scrollIntoView({ behavior: 'smooth' })
}

// 修改取消回复的逻辑
const cancelReply = () => {
  activeReplyTarget.value = null
  newComment.value = ''
}

// 格式化时间
const formatTime = (time) => {
  return new Date(time).toLocaleString()
}

// 分页处理
const handleSizeChange = (val) => {
  pageSize.value = val
  fetchComments()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  fetchComments()
}

onMounted(() => {
  fetchSongInfo()
  fetchComments()
})
</script>

<style scoped>
.main-container {
  display: flex;
  gap: 24px;
  height: 100%;
  padding: 20px;
}

.song-aside {
  flex: 0 0 300px;
}

.song-info-card {
  padding: 24px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
}

.song-cover {
  width: 200px;
  height: 200px;
  margin: 0 auto 20px;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.song-cover .el-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.song-info-card h2 {
  margin: 0;
  font-size: 24px;
  color: #303133;
  text-align: center;
}

.song-info {
  margin: 16px 0;
  color: #606266;
  font-size: 14px;
  text-align: center;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.comment-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 20px;
  overflow: hidden;
}

.comment-input {
  position: sticky;
  top: 20px;
  z-index: 10;
  margin-bottom: 20px;
  background: var(--el-bg-color);
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.input-area {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.reply-target {
  display: flex;
  align-items: center;
  gap: 8px;
}

.reply-tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background-color: var(--el-color-info-light-9);
  border-color: var(--el-color-info-light-8);
  color: var(--el-color-info-dark-2);
}

.comment-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
}

.comment-list {
  flex: 1;
  overflow-y: auto;
}

.comment-item {
  padding: 20px 0;
  border-bottom: 1px solid #EBEEF5;
}

.comment-main {
  display: flex;
  gap: 15px;
}

.comment-content {
  flex: 1;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.username {
  font-weight: bold;
  color: #303133;
}

.time {
  color: #909399;
  font-size: 0.9em;
}

.comment-actions {
  margin-top: 8px;
  display: flex;
  gap: 16px;
}

.reply-input.nested {
  margin: 8px 0 8px 30px;
  padding: 12px;
  background: var(--el-color-primary-light-9);
  border-radius: 4px;
}

.replies {
  margin: 10px 0 0 40px;
  padding: 10px;
  background: var(--el-fill-color-light);
  border-radius: 4px;
}

.reply-item {
  display: flex;
  gap: 10px;
  padding: 12px;
  margin-bottom: 8px;
  background: var(--el-bg-color);
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.reply-item:last-child {
  margin-bottom: 0;
}

.reply-content {
  flex: 1;
}

.reply-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.username {
  font-weight: 500;
  color: var(--el-text-color-primary);
}

.time {
  font-size: 12px;
  color: var(--el-text-color-secondary);
}

.reply-tag {
  margin-right: 8px;
  font-size: 12px;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  background-color: var(--el-color-info-light-9);
  border-color: var(--el-color-info-light-8);
  color: var(--el-color-info);
}

.reply-tag .el-icon {
  font-size: 14px;
}

.text {
  margin: 8px 0;
  line-height: 1.6;
  display: flex;
  align-items: flex-start;
  gap: 8px;
}

.reply-input.nested {
  margin: 8px 0;
  padding: 8px;
  background: var(--el-color-primary-light-9);
  border-radius: 4px;
}

.replies {
  margin: 10px 0 0 40px;
  padding: 10px;
  background: var(--el-fill-color-light);
  border-radius: 4px;
}

.reply-item {
  padding: 12px;
  margin-bottom: 8px;
  background: var(--el-bg-color);
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.reply-item:last-child {
  margin-bottom: 0;
}

.reply-actions {
  margin-top: 8px;
  display: flex;
  gap: 16px;
  opacity: 0.8;
}

.reply-actions:hover {
  opacity: 1;
}

.reply-content {
  flex: 1;
}

.reply-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.username {
  font-weight: 500;
  color: var(--el-text-color-primary);
}

.time {
  font-size: 12px;
  color: var(--el-text-color-secondary);
}
</style> 