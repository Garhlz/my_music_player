<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索歌曲、歌手、专辑"
            class="search-input"
            :prefix-icon="Search"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
          />
          <el-select v-model="sortBy" placeholder="排序方式" @change="handleSearch">
            <el-option label="最新" value="latest" />
            <el-option label="最热" value="popular" />
          </el-select>
        </div>

        <!-- 添加加载状态 -->
        <el-loading 
          :visible="isLoading"
          element-loading-text="加载中..."
          element-loading-background="rgba(255, 255, 255, 0.8)"
        />

        <!-- 当没有数据时显示空状态 -->
        <el-empty
          v-if="!isLoading && songs.length === 0"
          description="暂无歌曲"
        />

        <!-- 歌曲列表 -->
        <div class="song-list">
          <!-- 表头 -->
          <div class="song-header">
            <div class="col-index">#</div>
            <div class="col-title">标题</div>
            <div class="col-duration">时长</div>
            <div class="col-artist">歌手</div>
            <div class="col-album">专辑</div>
          </div>

          <!-- 歌曲列表项 -->
          <div 
            v-for="(song, index) in songs" 
            :key="song.id" 
            class="song-item"
            @mouseenter="hoveredSong = song.id"
            @mouseleave="hoveredSong = null"
            @click="handlePlaySong(song)"
          >
            <!-- 序号/播放按钮 -->
            <div class="col-index">
              <span v-if="hoveredSong !== song.id">{{ (page - 1) * pageSize + index + 1 }}</span>
              <el-icon v-else class="play-icon" @click="handlePlaySong(song)"><VideoPlay /></el-icon>
            </div>

            <!-- 歌曲信息 -->
            <div class="col-title">
              <div class="song-cover" @click="showCover(song)">
                <el-image 
                  :src="song.cover || '/assets/default-cover.jpg'"
                  fit="cover"
                />
              </div>
              <span class="song-name">{{ song.name }}</span>
            </div>

            <!-- 时长/操作按钮 -->
            <div class="col-duration">
              <template v-if="hoveredSong !== song.id">
                {{ formatDuration(song.duration) }}
              </template>
              <div v-else class="action-buttons">
                <el-tooltip content="喜欢" placement="top">
                  <el-icon 
                    @click="likeSong(song)" 
                    :class="{ 'liked': likedSongs.includes(song.id) }" 
                  >
                    <Star />
                  </el-icon>
                </el-tooltip>
                <el-tooltip content="添加到播放列表" placement="top">
                  <el-icon @click="addToPlaylist(song)"><Plus /></el-icon>
                </el-tooltip>
                <!-- <el-tooltip content="收藏专辑" placement="top">
                  <el-icon @click="addAlbum(song)"><FolderAdd /></el-icon>
                </el-tooltip> -->
                <el-tooltip content="评论" placement="top">
                  <el-icon @click="showComments(song)"><ChatDotRound /></el-icon>
                </el-tooltip>
                <el-tooltip content="下载" placement="top">
                  <el-icon @click="downloadSong(song)"><Download /></el-icon>
                </el-tooltip>
              </div>
            </div>

            <!-- 歌手 -->
            <div class="col-artist" @click="goToArtist(song.author_id)">
              {{ song.artist }}
            </div>

            <!-- 专辑 -->
            <div class="col-album" @click="goToAlbum(song.album_id)">
              {{ song.album }}
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <el-pagination
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

      <!-- 封面预览弹窗 -->
      <el-dialog
        v-model="coverDialogVisible"
        :title="selectedSong?.name"
        width="400px"
        align-center
      >
        <div class="cover-preview">
          <el-image
            :src="selectedSong?.cover || '/assets/default-cover.jpg'"
            fit="cover"
          />
        </div>
      </el-dialog>

      <!-- 添加到播放列表弹窗 -->
      <el-dialog
        v-model="playlistDialogVisible"
        title="添加到播放列表"
        width="400px"
      >
        <div class="playlist-dialog-content">
          <!-- 播放列表选择 -->
          <el-radio-group v-model="selectedPlaylist">
            <el-radio 
              v-for="list in userPlaylists" 
              :key="list.id" 
              :label="list.id"
            >
              {{ list.name }}
            </el-radio>
          </el-radio-group>

          <!-- 创建新播放列表 -->
          <div class="create-playlist">
            <el-input
              v-model="newPlaylistName"
              placeholder="新建播放列表"
            />
            <el-button type="primary" @click="createPlaylist">
              创建
            </el-button>
          </div>
        </div>
      </el-dialog>
    </template>
  </CommonLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getSongs } from '@/api/axiosFile'
import { getAlbums } from '@/api/axiosFile'
import { getArtists } from '@/api/axiosFile'
import { getLikedSongsById, addLikedSong, removeLikedSong } from '@/api/axiosFile'
//需要添加
import CommonLayout from '@/layouts/CommonLayout.vue'
import {
  Search,
  VideoPlay,
  Star,
  Plus,
  FolderAdd,
  ChatDotRound,
  Download
} from '@element-plus/icons-vue'
import { useStore } from 'vuex'

// 状态定义
const store = useStore()
const songs = ref([])
const page = ref(1)
const pageSize = ref(10)
const totalSongs = ref(0)
const isLoading = ref(false)
const searchQuery = ref('')
const sortBy = ref('latest')
const hoveredSong = ref(null)
const likedSongs = ref([])
const coverDialogVisible = ref(false)
const playlistDialogVisible = ref(false)
const selectedSong = ref(null)
const selectedPlaylist = ref(null)
const newPlaylistName = ref('')
const userPlaylists = ref([])

// 方法定义
const loadData = async () => {
  isLoading.value = true
  try {
    // 1. 获取歌曲列表
    const songsResponse = await getSongs({
      page: page.value,
      pageSize: pageSize.value,
      search: searchQuery.value,
      sortBy: sortBy.value
    })

    if (!songsResponse.data.message) {
      throw new Error(songsResponse.data.error || '获取歌曲列表失败')
    }

    const songsData = songsResponse.data.data
    totalSongs.value = songsData.total
    const songsList = songsData.list || []

    // 2. 获取喜欢列表
    const userId = localStorage.getItem('userId');
    console.log(userId);
    const likedResponse = await getLikedSongsById(userId);
    if (likedResponse.data.message) {
      likedSongs.value = likedResponse.data.data || []
    }

    // 3. 获取歌手和专辑信息
    const artistIds = [...new Set(songsList.map(song => song.author_id).filter(Boolean))]
    const albumIds = [...new Set(songsList.map(song => song.album_id).filter(Boolean))]

    const [artistsResponse, albumsResponse] = await Promise.all([
      artistIds.length > 0 ? getArtists(artistIds) : Promise.resolve({ data: { message: true, data: [] } }),
      albumIds.length > 0 ? getAlbums(albumIds) : Promise.resolve({ data: { message: true, data: [] } })
    ])

    const artists = artistsResponse.data.message ? artistsResponse.data.data : []
    const albums = albumsResponse.data.message ? albumsResponse.data.data : []
    
    // 4. 组装完整的歌曲信息
    songs.value = songsList.map(song => {
      const artist = artists.find(a => a.id === song.author_id)
      const album = albums.find(a => a.id === song.album_id)
      
      return {
        ...song,
        artist: artist ? artist.name : '未知歌手',
        album: album ? album.name : '未知专辑',
        cover: album ? album.cover : '/assets/default-cover.jpg',
        duration: song.duration || 0
      }
    })
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error(error.message || '加载数据失败，请稍后重试')
  } finally {
    isLoading.value = false
  }
}



const handlePlaySong = (song) => {
  store.dispatch('player/playSong', song)
}

const likeSong = async (song) => {
  try {
    if (likedSongs.value.includes(song.id)) {
      await removeLikedSong(song.id)
      likedSongs.value = likedSongs.value.filter(id => id !== song.id)
      ElMessage.success(`已取消喜欢: ${song.name}`)
    } else {
      await addLikedSong(song.id)
      likedSongs.value.push(song.id)
      ElMessage.success(`已添加到我喜欢: ${song.name}`)
    }
  } catch (error) {
    console.error('操作失败:', error)
    ElMessage.error('操作失败，请稍后重试')
  }
}

// const addAlbum = (song) => {
//   console.log(song);
//   store.dispatch('addToPlaylist', song)
//   ElMessage.success(`已收藏专辑: ${song.album.name}`)
// }

const formatDuration = (seconds) => {
  if (!seconds) return '00:00'
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = Math.floor(seconds % 60)
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

const handlePageChange = (newPage) => {
  page.value = newPage
  loadData()
}

const handlePageSizeChange = (newSize) => {
  pageSize.value = newSize
  page.value = 1
  loadData()
}

const handleSearch = () => {
  page.value = 1
  loadData()
}

const handleSortChange = () => {
  page.value = 1
  loadData()
}

// 生命周期钩子
onMounted(() => {
  loadData()
})

// 导出组件配置
defineExpose({
  loadData
})
</script>

<style scoped>
.playlist-container {
  padding: 24px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
  /* 添加滚动容器 */
  height: 100%;
  overflow-y: auto;
  /* 确保内容不会被播放器遮挡 */
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
  overflow: hidden; /* 确保圆角生效 */
}

/* 表头样式 */
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

/* 歌曲项样式 */
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

/* 序号列 */
.col-index {
  text-align: center;
  color: #909399;
  font-size: 14px;
}

/* 标题列 */
.col-title {
  display: flex;
  align-items: center;
  gap: 16px;
  min-width: 0; /* 防止内容溢出 */
}

.song-cover {
  width: 48px;
  height: 48px;
  border-radius: 6px;
  overflow: hidden;
  flex-shrink: 0; /* 防止图片被压缩 */
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
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

/* 时长列 */
.col-duration {
  text-align: center;
  color: #606266;
  font-size: 14px;
}

/* 歌手和专辑列 */
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

/* 操作按钮 */
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

/* 喜欢图标特殊样式 */
.liked {
  color: #ffcc00 !important;
  transform: scale(1.1);
  animation: heartBeat 0.3s ease-in-out;
}

/* 播放图标 */
.play-icon {
  color: var(--el-color-primary);
  font-size: 20px;
}

.play-icon:hover {
  animation: pulse 1s infinite;
}

/* 分页器 */
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

/* 添加滚动条样式 */
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
</style>
