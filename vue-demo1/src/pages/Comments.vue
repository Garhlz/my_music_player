<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索评论内容"
            class="search-input"
            :prefix-icon="Search"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
          />
          <el-select v-model="status" placeholder="评论状态" @change="handleSearch">
            <el-option label="全部" value="" />
            <el-option label="正常" value="normal" />
            <el-option label="已隐藏" value="hidden" />
          </el-select>
        </div>

        <!-- 加载状态 -->
        <el-loading 
          :visible="isLoading"
          element-loading-text="加载中..."
          element-loading-background="rgba(255, 255, 255, 0.8)"
        />

        <!-- 空状态 -->
        <el-empty
          v-if="!isLoading && comments.length === 0"
          description="暂无评论"
        />

        <!-- 评论列表 -->
        <div class="comment-list" v-if="comments.length > 0">
          <div 
            v-for="comment in comments" 
            :key="comment.id" 
            class="comment-item"
          >
            <div class="comment-header">
              <div class="user-info">
                <el-avatar :src="comment.user.avatar || '/assets/default-avatar.jpg'" :size="32" />
                <span class="username">{{ comment.user.name }}</span>
              </div>
              <div class="comment-time">{{ formatTime(comment.created_at) }}</div>
            </div>

            <div class="comment-content">
              <div class="song-info">
                <el-icon>
                  <Notebook />
                </el-icon>
                <span>评论歌曲：{{ comment.song.name }} - {{ comment.song.artist }}</span>
              </div>
              <div class="content-text">{{ comment.content }}</div>
            </div>

            <div class="comment-footer">
              <div class="comment-stats">
                <span class="likes">
                  <el-icon><Star /></el-icon>
                  {{ comment.likes }}
                </span>
                <span class="replies">
                  <el-icon><ChatDotRound /></el-icon>
                  {{ comment.replies }}
                </span>
              </div>
              <div class="comment-actions">
                <el-button-group>
                  <el-button 
                    :type="comment.status === 'hidden' ? 'success' : 'warning'"
                    size="small"
                    @click="toggleCommentStatus(comment)"
                  >
                    {{ comment.status === 'hidden' ? '显示' : '隐藏' }}
                  </el-button>
                  <el-button 
                    type="danger" 
                    size="small"
                    @click="deleteComment(comment)"
                  >
                    删除
                  </el-button>
                </el-button-group>
              </div>
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <el-pagination
          v-if="comments.length > 0"
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :total="totalComments"
          :page-sizes="[10, 20, 30, 50]"
          layout="total, sizes, prev, pager, next"
          class="pagination"
          @size-change="handlePageSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </template>
  </CommonLayout>
</template>

<script>
import { ref } from 'vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Search,
  Notebook,
  Star,
  ChatDotRound
} from '@element-plus/icons-vue'

export default {
  name: 'Comments',
  components: {
    CommonLayout,
    Search,
    Notebook,
    Star,
    ChatDotRound
  },
  data() {
    return {
      currentName: '评论管理',
      comments: [],
      page: 1,
      pageSize: 10,
      totalComments: 0,
      isLoading: false,
      searchQuery: '',
      status: ''
    }
  },
  mounted() {
    this.loadData()
  },
  methods: {
    formatTime(timestamp) {
      const date = new Date(timestamp)
      return date.toLocaleString()
    },

    async handleSearch() {
      this.page = 1
      await this.loadData()
    },

    async handlePageChange(newPage) {
      this.page = newPage
      await this.loadData()
    },

    async handlePageSizeChange(newSize) {
      this.pageSize = newSize
      this.page = 1
      await this.loadData()
    },

    async loadData() {
      this.isLoading = true
      try {
        // TODO: 实现获取评论列表的接口
        // const response = await getComments(this.page, this.pageSize, this.status)
        this.isLoading = false
      } catch (error) {
        console.error('加载数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
        this.isLoading = false
      }
    },

    async toggleCommentStatus(comment) {
      try {
        const action = comment.status === 'hidden' ? '显示' : '隐藏'
        await ElMessageBox.confirm(
          `确定要${action}这条评论吗？`,
          `${action}评论`,
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: comment.status === 'hidden' ? 'success' : 'warning'
          }
        )
        // TODO: 实现更新评论状态的接口
        ElMessage.success(`评论已${action}`)
        await this.loadData()
      } catch {
        // 用户取消操作
      }
    },

    async deleteComment(comment) {
      try {
        await ElMessageBox.confirm(
          '确定要删除这条评论吗？此操作不可恢复！',
          '删除评论',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        // TODO: 实现删除评论的接口
        ElMessage.success('评论已删除')
        await this.loadData()
      } catch {
        // 用户取消删除
      }
    }
  }
}
</script>

<style scoped>
@import '../styles/common.css';

.comment-list {
  margin: 20px 0;
}

.comment-item {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 16px;
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.comment-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.username {
  font-weight: 500;
}

.comment-time {
  color: #999;
  font-size: 12px;
}

.comment-content {
  margin-bottom: 12px;
}

.song-info {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #666;
  font-size: 13px;
  margin-bottom: 8px;
}

.content-text {
  font-size: 14px;
  line-height: 1.5;
}

.comment-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.comment-stats {
  display: flex;
  gap: 16px;
  color: #666;
  font-size: 13px;
}

.likes, .replies {
  display: flex;
  align-items: center;
  gap: 4px;
}

.comment-actions {
  display: flex;
  gap: 8px;
}
</style>