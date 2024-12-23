<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <div class="left-section">
            <el-input
              v-model="searchQuery"
              placeholder="搜索歌单"
              class="search-input"
              :prefix-icon="Search"
              clearable
              @clear="handleSearch"
              @keyup.enter="handleSearch"
            />
            <el-select v-model="sortBy" placeholder="排序方式" @change="handleSearch">
              <el-option label="最新" value="latest" />
              <el-option label="最早" value="oldest" />
            </el-select>
          </div>
          <el-button type="primary" @click="createDialogVisible = true">
            新建歌单
          </el-button>
        </div>

        <!-- 内容区域 -->
        <div class="content-area">
          <!-- 加载状态 -->
          <el-loading 
            v-loading="isLoading"
            element-loading-text="加载中..."
            element-loading-background="rgba(255, 255, 255, 0.8)"
          />

          <!-- 空状态 -->
          <el-empty
            v-if="!isLoading && playlists?.length === 0"
            description="还没有创建歌单"
          />

          <!-- 歌单网格 -->
          <div v-if="playlists?.length > 0" class="playlist-grid">
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
                  class="cover-image"
                />
                <div class="playlist-hover-overlay" v-show="hoveredPlaylist === playlist.id">
                  <div class="hover-buttons">
                    <el-button 
                      circle 
                      :icon="View"
                      @click.stop="navigateToDetail(playlist.id)"
                      title="查看详情"
                    />
                    <el-button 
                      circle 
                      :icon="Edit"
                      @click.stop="handleEdit(playlist)"
                      title="编辑歌单"
                    />
                    <el-button 
                      circle 
                      :icon="Delete"
                      @click.stop="handleDelete(playlist)"
                      title="删除歌单"
                    />
                  </div>
                </div>
              </div>
              <div class="playlist-info">
                <h3 class="playlist-name" :title="playlist.name">{{ playlist.name }}</h3>
                <p class="playlist-count">{{ playlist.song_count}}首歌曲</p>
                <p class="playlist-creator">创建者：{{ playlist.creator_name }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <el-pagination
          v-if="playlists?.length > 0"
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
            <el-form-item label="私密">
              <el-switch v-model="newPlaylist.isPublic" :active-value="false" :inactive-value="true" />
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="createDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitNewPlaylist">创建</el-button>
          </template>
        </el-dialog>

        <!-- 编辑歌单对话框 -->
        <el-dialog
          v-model="editDialogVisible"
          title="编辑歌单"
          width="400px"
          @close="resetEditData"
        >
          <el-form :model="editPlaylistData" label-width="80px">
            <el-form-item label="歌单名称">
              <el-input v-model="editPlaylistData.name" placeholder="请输入歌单名称" />
            </el-form-item>
            <el-form-item label="描述">
              <el-input
                v-model="editPlaylistData.description"
                type="textarea"
                placeholder="请输入歌单描述"
              />
            </el-form-item>
            <el-form-item label="公开">
              <el-switch v-model="editPlaylistData.isPublic" />
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="editDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitEditPlaylist">保存</el-button>
          </template>
        </el-dialog>
      </div>
    </template>
  </CommonLayout>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Search,
  VideoPlay,
  Plus,
  Edit,
  Delete,
  View
} from '@element-plus/icons-vue'
import { getMyPlaylists, createPlaylist, deletePlaylist, updatePlaylist } from '@/api/axiosFile'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { usePlayerStore } from '@/stores/player'

// 响应式状态
const currentName = ref('我的歌单')
const playlists = ref([])
const page = ref(1)
const pageSize = ref(12)
const totalPlaylists = ref(0)
const isLoading = ref(false)
const searchQuery = ref('')
const hoveredPlaylist = ref(null)
const createDialogVisible = ref(false)
const editDialogVisible = ref(false)
const editPlaylistData = ref({
  name: '',
  description: '',
  isPublic: false
})
const newPlaylist = ref({
  name: '',
  description: '',
  isPublic: false
})
const sortBy = ref('latest')

const router = useRouter()
const playerStore = usePlayerStore()

// 加载数据
const loadData = async () => {
  isLoading.value = true
  try {
    const response = await getMyPlaylists({
      page: page.value,
      pageSize: pageSize.value,
      search: searchQuery.value
    })
    if (response.data.message) {
      playlists.value = response.data.data.playlists.map(playlist => ({
        ...playlist,
        songCount: playlist.songCount || 0,
        creator: {
          name: playlist.creatorName || '未知'
        }
      }))
      totalPlaylists.value = response.data.data.total || 0;
    } else {
      throw new Error(response.data.error || '获取数据失败')
    }
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败，请稍后重试')
    playlists.value = []
    totalPlaylists.value = 0
  } finally {
    isLoading.value = false
  }
}

// 搜索处理
const handleSearch = async () => {
  page.value = 1
  await loadData()
}

// 分页处理
const handlePageChange = async (newPage) => {
  page.value = newPage
  await loadData()
}

const handlePageSizeChange = async (newSize) => {
  pageSize.value = newSize
  page.value = 1
  await loadData()
}

// 创建歌单
const createNewPlaylist = () => {
  createDialogVisible.value = true
}

const submitNewPlaylist = async () => {
  if (!newPlaylist.value.name.trim()) {
    ElMessage.warning('请输入歌单名称')
    return
  }
  try {
    await createPlaylist(newPlaylist.value)
    ElMessage.success('歌单创建成功')
    createDialogVisible.value = false
    newPlaylist.value = { name: '', description: '', isPublic: false }
    await loadData()
  } catch (error) {
    ElMessage.error('创建歌单失败')
  }
}

// 删除歌单
const handleDelete = (playlist) => {
  ElMessageBox.confirm(
    `确定要删除歌单"${playlist.name}"吗？`,
    '删除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async () => {
    try {
      const response = await deletePlaylist(playlist.id)
      if (response.data.message) {
        ElMessage.success('删除成功')
        loadData() // 重新加载数据
      } else {
        throw new Error(response.data.error || '删除失败')
      }
    } catch (error) {
      console.error('删除失败:', error)
      ElMessage.error(error.message || '删除失败，请稍后重试')
    }
  }).catch(() => {
    // 用户取消删除
  })
}

// 编辑歌单
const handleEdit = (playlist) => {
  editPlaylistData.value = { ...playlist }
  editDialogVisible.value = true
}

// 提交编辑的歌单
const submitEditPlaylist = async () => {
  try {
    const playlistId = playlists.value.find(p => p.name === editPlaylistData.value.name).id;
    await updatePlaylist(playlistId, editPlaylistData.value)
    ElMessage.success('歌单已更新')
    editDialogVisible.value = false
    await loadData()
  } catch (error) {
    ElMessage.error('更新歌单失败')
  }
}

// 重置编辑数据
const resetEditData = () => {
  editPlaylistData.value = { name: '', description: '', isPublic: false };
}

// 播放歌单
const playPlaylist = (playlist) => {
  if (playlist.songs && playlist.songs.length > 0) {
    playerStore.setPlaylist(playlist.songs)
    playerStore.play(0)
  } else {
    router.push(`/playlist/${playlist.id}`)
  }
}

// 跳转到详情页
const navigateToDetail = (id) => {
  router.push(`/playlist/${id}`)
}

// 生命周期钩子
onMounted(() => {
  loadData()
})

onBeforeUnmount(() => {
  // 这里可以添加任何需要在组件卸载时清理的逻辑
})
</script>

<style scoped>
@import '../styles/common.css';

/* 主容器 */
.playlist-container {
  height: calc(100vh - 64px - 72px); /* 减去顶部导航和底部播放器的高度 */
  display: flex;
  flex-direction: column;
  position: relative;
  background-color: var(--el-bg-color);
}

/* 搜索和筛选区域 */
.filter-section {
  position: sticky;
  top: 0;
  z-index: 10;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
  padding: 16px 24px;
  background: var(--el-bg-color);
  border-bottom: 1px solid var(--el-border-color-lighter);
}

.left-section {
  display: flex;
  gap: 16px;
  align-items: center;
}

.search-input {
  width: 300px;
}

/* 内容区域 */
.content-area {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
}

/* 歌单网格布局 */
.playlist-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 24px;
  padding: 24px;
}

/* 歌单卡片 */
.playlist-card {
  background: var(--el-bg-color);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  transition: transform 0.3s;
}

.playlist-card:hover {
  transform: translateY(-5px);
}

/* 封面区域 */
.playlist-cover {
  position: relative;
  width: 100%;
  padding-top: 100%; /* 1:1 aspect ratio */
}

.cover-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.playlist-hover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  opacity: 0;
  transition: opacity 0.3s;
}

.playlist-card:hover .playlist-hover-overlay {
  opacity: 1;
}

.hover-buttons {
  display: flex;
  gap: 12px;
  justify-content: center;
  align-items: center;
}

.hover-buttons .el-button {
  --el-button-bg-color: rgba(255, 255, 255, 0.9);
  --el-button-hover-bg-color: var(--el-color-primary);
  --el-button-hover-text-color: #fff;
  transform: scale(0.9);
  transition: all 0.3s;
}

.hover-buttons .el-button:hover {
  transform: scale(1);
}

/* 删除按钮特殊样式 */
.hover-buttons .el-button:last-child:hover {
  --el-button-hover-bg-color: var(--el-color-danger);
}

/* 歌单信息 */
.playlist-info {
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.playlist-name {
  font-size: 14px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.playlist-count,
.playlist-creator {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  line-height: 1.5;
}

/* 分页器 */
.pagination {
  position: sticky;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 16px 24px;
  background-color: var(--el-bg-color);
  display: flex;
  justify-content: center;
  border-top: 1px solid var(--el-border-color-lighter);
  z-index: 10;
}

/* 动画效果 */
@keyframes heartBeat {
  0% { transform: scale(1); }
  50% { transform: scale(1.3); }
  100% { transform: scale(1.1); }
}

@keyframes pulse {
  0% { transform: scale(1); }
  50% { transform: scale(1.1); }
  100% { transform: scale(1); }
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .filter-section {
    flex-direction: column;
    align-items: stretch;
  }

  .left-section {
    flex-direction: column;
  }

  .search-input {
    width: 100%;
  }

  .playlist-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 16px;
    padding: 16px;
  }
}
</style>