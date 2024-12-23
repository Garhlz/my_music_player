<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索喜欢的音乐"
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

        <!-- 加载状态 -->
        <el-loading 
          :visible="isLoading"
          element-loading-text="加载中..."
          element-loading-background="rgba(255, 255, 255, 0.8)"
        />

        <!-- 空状态 -->
        <el-empty
          v-if="!isLoading && songs.length === 0"
          description="还没有喜欢的音乐"
        />

        <!-- 歌曲列表 -->
        <div class="song-list" v-if="songs.length > 0">
          <div class="song-header">
            <div class="col-index">#</div>
            <div class="col-title">标题</div>
            <div class="col-duration">时长</div>
            <div class="col-artist">歌手</div>
            <div class="col-album">专辑</div>
          </div>

          <div 
            v-for="(song, index) in songs" 
            :key="song.id" 
            class="song-item"
            @mouseenter="hoveredSong = song.id"
            @mouseleave="hoveredSong = null"
          >
            <div class="col-index">
              <span v-if="hoveredSong !== song.id">{{ (page - 1) * pageSize + index + 1 }}</span>
              <el-icon v-else class="play-icon" @click="handlePlaySong(song)">
                <VideoPlay />
              </el-icon>
            </div>

            <div class="col-title">
              <div class="song-cover" @click="goToPlayer(song)">
                <el-image 
                  :src="song.cover || '/assets/default-cover.jpg'"
                  fit="cover"
                />
              </div>
              <span class="song-name">{{ song.name }}</span>
            </div>

            <div class="col-duration">
              <template v-if="hoveredSong !== song.id">
                {{ formatDuration(song.duration) }}
              </template>
              <div v-else class="action-buttons">
                <el-tooltip content="取消喜欢" placement="top">
                  <el-icon 
                    @click="unlikeSong(song)"
                    :style="{ color: '#ffcc00' }"
                  >
                    <Star />
                  </el-icon>
                </el-tooltip>
                <el-tooltip content="添加到播放列表" placement="top">
                  <el-icon @click="addToPlaylist(song)"><Plus /></el-icon>
                </el-tooltip>
                <el-tooltip content="收藏专辑" placement="top">
                  <el-icon @click="goToAlbum(song.album_id)"><FolderAdd /></el-icon>
                </el-tooltip>
                <el-tooltip content="评论" placement="top">
                  <el-icon @click="showComments(song)"><ChatDotRound /></el-icon>
                </el-tooltip>
              </div>
            </div>

            <div class="col-artist" @click="goToArtist(song.author_id)">
              {{ song.artist }}
            </div>
            <div class="col-album" @click="goToAlbum(song.album_id)">
              {{ song.album }}
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <el-pagination
          v-if="songs.length > 0"
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :total="totalSongs"
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

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage } from 'element-plus'
import {
  Search,
  VideoPlay,
  Star,
  Plus,
  FolderAdd,
  ChatDotRound,
} from '@element-plus/icons-vue'
import { getLikedSongs, removeLikedSong } from '@/api/axiosFile'

// 响应式状态
const currentName = ref('我喜欢的音乐')
const songs = ref([])
const page = ref(1)
const pageSize = ref(10)
const totalSongs = ref(0)
const isLoading = ref(false)
const searchQuery = ref('')
const sortBy = ref('latest')
const hoveredSong = ref(null)

const router = useRouter()
const store = useStore()

// 加载数据
const loadData = async () => {
  isLoading.value = true
  try {
    const response = await getLikedSongs({
      page: page.value,
      pageSize: pageSize.value,
      search: searchQuery.value,
      sortBy: sortBy.value
    })
    console.log(response);
    if (response.data.message) {
      songs.value = response.data.data.list
      totalSongs.value = response.data.data.total
    } else {
      throw new Error(response.data.error || '获取数据失败')
    }
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败，请稍后重试')
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

// 播放歌曲
const handlePlaySong = (song) => {
  store.dispatch('player/playSong', song)
}

// 取消喜欢
const unlikeSong = async (song) => {
  try {
    await removeLikedSong(song.id)
    ElMessage.success(`已取消喜欢: ${song.name}`)
    await loadData() // 重新加载数据
  } catch (error) {
    console.error('取消喜欢失败:', error)
    ElMessage.error('操作失败，请稍后重试')
  }
}

// 添加到播放列表
const addToPlaylist = (song) => {
  store.dispatch('addToPlaylist', song)
  ElMessage.success(`已添加到播放列表: ${song.name}`)
}

const addAlbum = (song) => {
  console.log(song);
  store.dispatch('addToPlaylist', song)
  ElMessage.success(`已收藏专辑: ${song.album.name}`)
}

// 显示评论
const showComments = (song) => {
  router.push(`/song/${song.id}/comments`)
}

// 格式化时长
const formatDuration = (seconds) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

// 添加路由跳转方法
const goToArtist = (artistId) => {
  if (artistId) {
    router.push(`/artist/${artistId}`)
  }
}

const goToAlbum = (albumId) => {
  if (albumId) {
    router.push(`/album/${albumId}`)
  }
}

const goToPlayer = (song) => {
  handlePlaySong(song) // 点击封面时同时开始播放
  router.push('/player')
}

// 组件挂载时加载数据
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.playlist-container {
  padding: 24px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
  height: 100%;
  overflow-y: auto;
  padding-bottom: 100px;
}

.filter-section {
  display: flex;
  gap: 20px;
  margin-bottom: 24px;
  align-items: center;
}

.search-input {
  width: 300px;
}

.song-list {
  margin-top: 20px;
  border: 1px solid #ebeef5;
  border-radius: 4px;
  overflow: hidden;
}

.song-header {
  position: sticky;
  top: 0;
  display: grid;
  grid-template-columns: 60px 3fr 100px 1.5fr 1.5fr 120px;
  padding: 12px 20px;
  background-color: #f5f7fa;
  border-bottom: 1px solid #ebeef5;
  font-weight: 600;
  color: #606266;
  align-items: center;
  z-index: 1;
}

.song-item {
  display: grid;
  grid-template-columns: 60px 3fr 100px 1.5fr 1.5fr 120px;
  padding: 12px 20px;
  border-bottom: 1px solid #ebeef5;
  align-items: center;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.song-item:hover {
  background-color: #f5f7fa;
  transform: translateX(4px);
}

.col-index {
  text-align: center;
  color: #909399;
  font-size: 14px;
}

.col-title {
  display: flex;
  align-items: center;
  gap: 16px;
  min-width: 0;
}

.song-cover {
  width: 48px;
  height: 48px;
  border-radius: 6px;
  overflow: hidden;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

.song-cover:hover {
  transform: scale(1.05) rotate(2deg);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.song-name {
  font-weight: 500;
  color: #303133;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.col-duration {
  text-align: center;
  color: #606266;
  font-size: 14px;
}

.col-artist, .col-album {
  color: #606266;
  cursor: pointer;
  transition: color 0.3s ease;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  padding: 0 10px;
}

.col-artist:hover, .col-album:hover {
  color: var(--el-color-primary);
  text-decoration: underline;
}

.action-buttons {
  display: flex;
  gap: 16px;
  justify-content: center;
  align-items: center;
}

.el-icon {
  font-size: 18px;
  color: #909399;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  padding: 6px;
  border-radius: 50%;
}

.el-icon:hover {
  transform: scale(1.2);
  color: var(--el-color-primary);
  background-color: rgba(64, 158, 255, 0.1);
}

.liked {
  color: #ffcc00 !important;
  transform: scale(1.1);
  animation: heartBeat 0.3s ease-in-out;
}

.play-icon {
  color: var(--el-color-primary);
  font-size: 20px;
}

.play-icon:hover {
  animation: pulse 1s infinite;
}

.pagination {
  position: sticky;
  bottom: 0;
  margin-top: 24px;
  padding: 16px 0;
  background-color: #fff;
  display: flex;
  justify-content: center;
  box-shadow: 0 -2px 12px 0 rgba(0, 0, 0, 0.05);
}

/* 滚动条样式 */
.playlist-container::-webkit-scrollbar {
  width: 6px;
}

.playlist-container::-webkit-scrollbar-thumb {
  background-color: #dcdfe6;
  border-radius: 3px;
}

.playlist-container::-webkit-scrollbar-track {
  background-color: #f5f7fa;
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
@media screen and (max-width: 1200px) {
  .song-header, .song-item {
    grid-template-columns: 60px 2.5fr 100px 1.2fr 1.2fr 120px;
  }
}

@media screen and (max-width: 768px) {
  .song-header, .song-item {
    grid-template-columns: 50px 3fr 1.5fr 100px;
  }
  
  .col-album, .col-duration {
    display: none;
  }
  
  .filter-section {
    flex-direction: column;
    align-items: stretch;
  }
  
  .search-input {
    width: 100%;
  }
}
</style>