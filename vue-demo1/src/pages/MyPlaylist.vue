<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索歌单"
            class="search-input"
            :prefix-icon="Search"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
          />
          <el-button type="primary" @click="createNewPlaylist">
            <el-icon><Plus /></el-icon>创建歌单
          </el-button>
        </div>

        <!-- 加载状态 -->
        <el-loading 
          :visible="isLoading"
          element-loading-text="加载中..."
          element-loading-background="rgba(255, 255, 255, 0.8)"
        />

        <!-- 空状态 -->
        <el-empty
          v-if="!isLoading && playlists.length === 0"
          description="还没有创建歌单"
        >
          <el-button type="primary" @click="createNewPlaylist">创建歌单</el-button>
        </el-empty>

        <!-- 歌单网格 -->
        <div class="playlist-grid" v-if="playlists.length > 0">
          <div 
            v-for="playlist in playlists" 
            :key="playlist.id" 
            class="playlist-card"
            @mouseenter="hoveredPlaylist = playlist.id"
            @mouseleave="hoveredPlaylist = null"
          >
            <div class="playlist-cover">
              <el-image 
                :src="playlist.cover || '/assets/default-cover.jpg'"
                fit="cover"
              />
              <div class="playlist-actions" v-show="hoveredPlaylist === playlist.id">
                <el-icon @click="playPlaylist(playlist)"><VideoPlay /></el-icon>
                <el-icon @click="editPlaylist(playlist)"><Edit /></el-icon>
                <el-icon @click="deletePlaylist(playlist)"><Delete /></el-icon>
              </div>
            </div>
            <div class="playlist-info">
              <h3 class="playlist-name">{{ playlist.name }}</h3>
              <p class="playlist-count">{{ playlist.songCount }}首歌曲</p>
              <p class="playlist-creator">创建者：{{ playlist.creator }}</p>
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <el-pagination
          v-if="playlists.length > 0"
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :total="totalPlaylists"
          :page-sizes="[12, 24, 36, 48]"
          layout="total, sizes, prev, pager, next"
          class="pagination"
          @size-change="handlePageSizeChange"
          @current-change="handlePageChange"
        />

        <!-- 创建歌单对话框 -->
        <el-dialog
          v-model="createDialogVisible"
          title="创建歌单"
          width="400px"
        >
          <el-form :model="newPlaylist" label-width="80px">
            <el-form-item label="歌单名称">
              <el-input v-model="newPlaylist.name" placeholder="请输入歌单名称" />
            </el-form-item>
            <el-form-item label="描述">
              <el-input
                v-model="newPlaylist.description"
                type="textarea"
                placeholder="请输入歌单描述"
              />
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="createDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitNewPlaylist">创建</el-button>
          </template>
        </el-dialog>
      </div>
    </template>
  </CommonLayout>
</template>

<script>
import { ref } from 'vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage } from 'element-plus'
import {
  Search,
  VideoPlay,
  Plus,
  Edit,
  Delete
} from '@element-plus/icons-vue'

export default {
  name: 'MyPlaylist',
  components: {
    CommonLayout,
    Search,
    VideoPlay,
    Plus,
    Edit,
    Delete
  },
  data() {
    return {
      currentName: '我的歌单',
      playlists: [],
      page: 1,
      pageSize: 12,
      totalPlaylists: 0,
      isLoading: false,
      searchQuery: '',
      hoveredPlaylist: null,
      createDialogVisible: false,
      newPlaylist: {
        name: '',
        description: ''
      }
    }
  },
  mounted() {
    this.loadData()
  },
  methods: {
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
        // TODO: 实现获取歌单列表的接口
        // const response = await getMyPlaylists(this.page, this.pageSize)
        this.isLoading = false
      } catch (error) {
        console.error('加载数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
        this.isLoading = false
      }
    },

    createNewPlaylist() {
      this.createDialogVisible = true
    },

    async submitNewPlaylist() {
      if (!this.newPlaylist.name.trim()) {
        ElMessage.warning('请输入歌单名称')
        return
      }
      // TODO: 实现创建歌单的接口
      ElMessage.success('歌单创建成功')
      this.createDialogVisible = false
      this.newPlaylist = { name: '', description: '' }
      await this.loadData()
    },

    playPlaylist(playlist) {
      // TODO: 实现播放歌单的功能
      ElMessage.success(`开始播放歌单: ${playlist.name}`)
    },

    editPlaylist(playlist) {
      // TODO: 实现编辑歌单的功能
      ElMessage.success(`编辑歌单: ${playlist.name}`)
    },

    async deletePlaylist(playlist) {
      try {
        await ElMessageBox.confirm(
          '确定要删除这个歌单吗？',
          '删除歌单',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        // TODO: 实现删除歌单的接口
        ElMessage.success(`歌单已删除: ${playlist.name}`)
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

.playlist-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
  padding: 20px;
}

.playlist-card {
  background: #fff;
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.3s ease;
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.playlist-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0,0,0,0.2);
}

.playlist-cover {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
}

.playlist-cover .el-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.playlist-actions {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
}

.playlist-actions .el-icon {
  font-size: 24px;
  color: white;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.playlist-actions .el-icon:hover {
  transform: scale(1.2);
}

.playlist-info {
  padding: 12px;
}

.playlist-name {
  margin: 0;
  font-size: 16px;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.playlist-count {
  margin: 4px 0;
  font-size: 14px;
  color: #666;
}

.playlist-creator {
  margin: 4px 0 0;
  font-size: 12px;
  color: #999;
}
</style>