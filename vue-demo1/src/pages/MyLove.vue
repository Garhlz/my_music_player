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
                <!-- <el-tooltip content="收藏专辑" placement="top">
                  <el-icon @click="goToAlbum(song.album_id)"><FolderAdd /></el-icon>
                </el-tooltip> -->
                <!-- 暂时废弃这个功能不显示 -->
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

        <div class="player-controls" v-if="showPlayerControls">
          <el-button circle @click="previousSong">
            <el-icon><ArrowLeft /></el-icon>
          </el-button>
          <el-button circle @click="togglePlay">
            <el-icon>
              <VideoPlay v-if="!isPlaying" />
              <VideoPause v-else />
            </el-icon>
          </el-button>
          <el-button circle @click="nextSong">
            <el-icon><ArrowRight /></el-icon>
          </el-button>
        </div>

          <el-dialog
          v-model="playlistDialogVisible"
          title="添加到歌单"
          width="600px"
          @close="selectedPlaylist = null"
        >
          <div class="playlist-dialog-content">
            <!-- 现有歌单列表，设置为可见之后弹出 -->
            <div class="existing-playlists">
              <el-scrollbar height="400px">
                <div class="playlist-list">
                  <div 
                    v-for="playlist in userPlaylists" 
                    :key="playlist.id" 
                    class="playlist-item"
                    :class="{ 'playlist-item-selected': selectedPlaylist === playlist.id }"
                    @click="selectedPlaylist = playlist.id"
                  >
                    <div class="playlist-info-row">
                      <el-image 
                        :src="playlist.cover || '/assets/default-cover.jpg'" 
                        class="playlist-cover-img"
                      />
                      <div class="playlist-details">
                        <div class="playlist-name">{{ playlist.name }}</div>
                        <div class="playlist-description">{{ playlist.description || '暂无描述' }}</div>
                        <div class="playlist-count">{{ playlist.song_count || 0 }}首歌曲</div>
                      </div>
                    </div>
                  </div>
                </div>
              </el-scrollbar>
            </div>

            <!-- 创建新歌单 -->
            <div class="create-playlist">
              <el-divider>创建新歌单</el-divider>
              <div class="create-playlist-form">
                <el-input
                  v-model="newPlaylistName"
                  placeholder="请输入歌单名称"
                  clearable
                />
                <el-input
                  v-model="newPlaylistDescription"
                  type="textarea"
                  :rows="2"
                  placeholder="添加歌单描述（选填）"
                  class="mt-3"
                />
                <el-button type="primary" @click="createNewPlaylist" class="mt-3">
                  创建
                </el-button>
              </div>
            </div>
          </div>
          <template #footer>
            <span class="dialog-footer">
              <el-button @click="playlistDialogVisible = false">取消</el-button>
              <el-button type="primary" @click="confirmAddToPlaylist">
                确定
              </el-button>
            </span>
          </template>
        </el-dialog>
      </div>
    </template>
  </CommonLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  Search, 
  VideoPlay, 
  VideoPause, 
  Star,
  Plus,
  FolderAdd,
  ChatDotRound,
  ArrowLeft,
  ArrowRight
} from '@element-plus/icons-vue'
import { 
  getLikedSongs, 
  removeLikedSong,
  getMyPlaylists, // 直接都采用相同的接口函数即可
  createPlaylist,
  addSongToPlaylist 
} from '@/api/axiosFile'
import { usePlayerStore } from '@/stores/player'
import CommonLayout from '@/layouts/CommonLayout.vue'

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
const playerStore = usePlayerStore()

// 歌单相关的响应式变量
const playlistDialogVisible = ref(false)
const showCreateForm = ref(false)
const selectedPlaylist = ref(null)
const newPlaylistName = ref('')
const newPlaylistDescription = ref('')
const userPlaylists = ref([])
const selectedSong = ref(null)

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
  playerStore.setPlaylist([song])
  playerStore.play(0)
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
const addToPlaylist = async (song) => {
  selectedSong.value = song
  playlistDialogVisible.value = true
  //在这里修改显示状态，然后弹出一个弹窗
  try {
    const response = await getMyPlaylists()
    if (response.data.message) {
      // console.log(response.data.data);
      userPlaylists.value = response.data.data.playlists
    } else {
      throw new Error(response.data.error || '获取歌单失败')
    }
  } catch (error) {
    console.error('获取歌单失败:', error)
    ElMessage.error(error.message || '获取歌单失败，请稍后重试')
  }
}

const createNewPlaylist = async () => {
  if (!newPlaylistName.value.trim()) {
    ElMessage.warning('请输入歌单名称')
    return
  }

  try {
    const response = await createPlaylist({
      name: newPlaylistName.value,
      description: newPlaylistDescription.value,
      isPublic: false
    })

    if (response.data.message) {
      userPlaylists.value.unshift({
        ...response.data.data,
        songCount: 0
      })
      selectedPlaylist.value = response.data.data.id
      newPlaylistName.value = ''
      newPlaylistDescription.value = ''
      ElMessage.success('歌单创建成功')
    } else {
      throw new Error(response.data.error || '创建歌单失败')
    }
  } catch (error) {
    console.error('创建歌单失败:', error)
    ElMessage.error(error.message || '创建歌单失败，请稍后重试')
  }
}

const confirmAddToPlaylist = async () => {
  if (!selectedPlaylist.value) {
    ElMessage.warning('请选择歌单')
    return
  }

  try {
    const response = await addSongToPlaylist(selectedPlaylist.value, selectedSong.value.id)
    if (response.data.message) {
      ElMessage.success('添加成功')
      playlistDialogVisible.value = false
      selectedPlaylist.value = null
    } else {
      throw new Error(response.data.error || '添加失败')
    }
  } catch (error) {
    console.error('添加失败:', error)
    ElMessage.error(error.message || '添加失败，请稍后重试')
  }
}

const addAlbum = (song) => {
  console.log(song);
  // playerStore.setPlaylist([song])
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
  else{
    ElMessage.error("歌手不存在!")
  }
}

const goToAlbum = (albumId) => {
  if (albumId) {
    router.push(`/album/${albumId}`)
  }
  else{
    ElMessage.error("专辑不存在!")
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
  padding: 24px 32px;
  margin-right: 16px;
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
  border: 1px solid #EBEEF5;
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: 24px;
}

.song-header {
  position: sticky;
  top: 0;
  display: grid;
  grid-template-columns: 60px minmax(300px, 2.5fr) 180px minmax(160px, 1fr) minmax(160px, 1fr);
  padding: 12px 24px;
  background-color: #f5f7fa;
  border-bottom: 1px solid #ebeef5;
  font-weight: 600;
  color: #606266;
  align-items: center;
  z-index: 1;
  gap: 12px;
}

.song-item {
  display: grid;
  grid-template-columns: 60px minmax(300px, 2.5fr) 180px minmax(160px, 1fr) minmax(160px, 1fr);
  padding: 12px 24px;
  border-bottom: 1px solid #ebeef5;
  align-items: center;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.song-item:hover {
  background-color: #F5F7FA;
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
  min-width: 180px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding-right: 24px;
  color: #606266;
  font-size: 14px;
}

.col-artist, .col-album {
  padding: 0 12px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  cursor: pointer;
  color: #606266;
}

.col-artist:hover, .col-album:hover {
  color: #409EFF;
}

.action-buttons {
  display: flex;
  gap: 20px;
  justify-content: flex-start;
  align-items: center;
}

.action-buttons .el-icon {
  font-size: 16px;
  cursor: pointer;
  color: #606266;
  transition: color 0.3s;
  padding: 4px;
}

.action-buttons .el-icon:hover {
  color: #409EFF;
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

.playlist-dialog-content {
  padding: 20px 0;
}

.existing-playlists {
  max-height: 400px;
  overflow: hidden;
  margin: 0 -20px; /* 扩展到对话框边缘 */
}

.playlist-list {
  display: flex;
  flex-direction: column;
}

.playlist-item {
  padding: 16px 20px;
  border-bottom: 1px solid var(--el-border-color-lighter);
  cursor: pointer;
  transition: all 0.3s;
  position: relative;
}

.playlist-item:hover {
  background-color: var(--el-fill-color-light);
}

.playlist-item-selected {
  background-color: var(--el-color-primary-light-9);
}

.playlist-item-selected::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 4px;
  background-color: var(--el-color-primary);
}

.playlist-info-row {
  display: flex;
  gap: 16px;
  width: 100%;
}

.playlist-cover-img {
  width: 64px;
  height: 64px;
  border-radius: 6px;
  object-fit: cover;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s;
}

.playlist-item:hover .playlist-cover-img {
  transform: scale(1.05);
}

/* 自定义滚动条样式 */
.existing-playlists :deep(.el-scrollbar__wrap) {
  overflow-x: hidden;
}

.existing-playlists :deep(.el-scrollbar__bar.is-vertical) {
  width: 6px;
}

.existing-playlists :deep(.el-scrollbar__thumb) {
  background-color: var(--el-border-color);
  border-radius: 3px;
}

.existing-playlists :deep(.el-scrollbar__thumb:hover) {
  background-color: var(--el-border-color-darker);
}

/* 选中状态的额外样式 */
.playlist-item-selected .playlist-name {
  color: var(--el-color-primary);
  font-weight: 600;
}

.playlist-item-selected .playlist-description,
.playlist-item-selected .playlist-count {
  color: var(--el-color-primary-light-3);
}

/* 添加动画效果 */
.playlist-item {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.playlist-item:active {
  transform: scale(0.98);
}

.create-playlist {
  margin-top: 20px;
  padding: 0 20px;
}

.create-playlist-form {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 16px;
}

.mt-3 {
  margin-top: 12px;
}

/* 自定义滚动条样式 */
:deep(.el-scrollbar__wrap) {
  overflow-x: hidden !important;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .playlist-item {
    padding: 12px 16px;
  }

  .playlist-cover-img {
    width: 48px;
    height: 48px;
  }
}
</style>