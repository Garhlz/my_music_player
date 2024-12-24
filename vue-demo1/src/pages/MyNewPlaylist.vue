<template>
  <CommonLayout>
    <template #main>
      <div class="main-container">
        <!-- 左侧歌单信息 -->
        <div class="playlist-aside">
          <div class="playlist-info-card">
            <div class="playlist-cover">
              <el-image 
                :src="currentPlaylist?.cover || '/assets/default-cover.jpg'"
                fit="cover"
              />
            </div>
            <h2>{{ currentPlaylist?.name || '我的歌单' }}</h2>
            <div class="playlist-info">
              <span>创建者：{{ currentPlaylist?.creator_name || '未知' }}</span>
              <span>|</span>
              <span>{{ currentPlaylist?.is_public ? '公开' : '私密' }}歌单</span>
              <span>|</span>
              <span>{{ songs.length }}首歌曲</span>
            </div>
            <p class="playlist-description">{{ currentPlaylist?.description || '暂无描述' }}</p>
          </div>
        </div>

        <!-- 右侧歌曲列表 -->
        <div class="playlist-container" v-loading="isLoading">
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

          <!-- 歌曲列表 -->
          <div class="song-list">
            <div class="song-header">
              <div class="col-index">#</div>
              <div class="col-title">标题</div>
              <div class="col-duration">时长</div>
              <div class="col-artist">歌手</div>
              <div class="col-album">专辑</div>
              <div class="col-actions">操作</div>
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
                <el-icon v-else class="play-icon" @click="handlePlaySong(song)"><VideoPlay /></el-icon>
              </div>
              <div class="col-title">
                <div class="song-cover">
                  <el-image 
                    :src="song.cover || '/assets/default-cover.jpg'"
                    fit="cover"
                  />
                </div>
                <span class="song-name">{{ song.name }}</span>
              </div>
              <div class="col-duration">{{ formatDuration(song.duration) }}</div>
              <div class="col-artist">{{ song.artist }}</div>
              <div class="col-album">{{ song.album }}</div>
              <div class="col-actions">
                <el-button 
                  type="text" 
                  size="small" 
                  @click="removeSong(song)"
                >
                  删除
                </el-button>
              </div>
            </div>
          </div>
        </div>
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
  Star, 
  Plus, 
  FolderAdd, 
  ChatDotRound, 
  Download 
} from '@element-plus/icons-vue'
import { 
  getSongFromPlaylistById, // 其实这个页面直接从pub修改即可
  removeSongFromPlaylist,
  getLikedSongsById, 
  addLikedSong, 
  removeLikedSong,
  getMyPlaylists, 
  createPlaylist, 
  addSongToPlaylist 
} from '@/api/axiosFile'
import { usePlayerStore } from '@/stores/player'
import CommonLayout from '@/layouts/CommonLayout.vue'

const router = useRouter()
const playerStore = usePlayerStore()

// 基础数据
const currentName = ref('公共歌单')
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
const newPlaylistDescription = ref('')
const userPlaylists = ref([])
const currentPlaylist = ref(null)
// 方法定义
const loadData = async () => {
  isLoading.value = true
  try {
    // 1. 获取歌曲列表
    //注意这里的route文件获取当前路径id的方式
    const songsResponse = await getSongFromPlaylistById(router.currentRoute.value.params.id,{
      page: page.value,
      pageSize: pageSize.value,
      search: searchQuery.value,
      sortBy: sortBy.value
    });
    //此处已经在后端处理了多表查询了，所以这里不需要再处理
    if (!songsResponse.data.message) {
      throw new Error(songsResponse.data.error || '获取歌曲列表失败')
    }
    console.log(songsResponse);
    const songsData = songsResponse.data.data;
    totalSongs.value = songsData.total;
    const songsList = songsData.songs || []
    currentPlaylist.value = songsData.playlist;
    // 2. 获取喜欢列表
    const userId = localStorage.getItem('userId')
    const likedResponse = await getLikedSongsById(userId)
    if (likedResponse.data.message) {
      likedSongs.value = likedResponse.data.data || []
    }
    // 3. 组装完整的歌曲信息,只是重新处理罢了
    songs.value = songsList.map(song => {
      return {
        ...song,
        artist: song.artist_name ? song.artist_name : '未知歌手',
        album: song.album_name ? song.album_name : '未知专辑',
        cover: song.album_cover ? song.album_cover : '/assets/default-cover.jpg',
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
  playerStore.setPlaylist([song])
  playerStore.play(0)
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

const addToPlaylist = async (song) => {
  selectedSong.value = song
  playlistDialogVisible.value = true
  
  try {
    const response = await getMyPlaylists()
    if (response.data.message) {
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

const formatDuration = (seconds) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

const handleSearch = () => {
  page.value = 1
  loadData()
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

const showCover = (song) => {
  selectedSong.value = song
  coverDialogVisible.value = true
}

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

const removeSong = async (song) => {
  try {
    const response = await removeSongFromPlaylist(currentPlaylist.value.id, song.id)
    if (response.data.message) {
      ElMessage.success('删除成功')
      loadData()
    } else {
      throw new Error(response.data.error || '删除失败')
    }
  } catch (error) {
    console.error('删除失败:', error)
    ElMessage.error(error.message || '删除失败，请稍后重试')
  }
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
.main-container {
  display: flex;
  gap: 24px;
  height: 100%;
  padding: 20px;
}

.playlist-aside {
  flex: 0 0 300px;
}

.playlist-info-card {
  padding: 24px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
}

.playlist-cover {
  width: 200px;
  height: 200px;
  margin: 0 auto 20px;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.playlist-cover .el-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.playlist-info-card h2 {
  margin: 0;
  font-size: 24px;
  color: #303133;
  text-align: center;
}

.playlist-info {
  margin: 16px 0;
  color: #606266;
  font-size: 14px;
  text-align: center;
}

.playlist-info span {
  margin: 0 8px;
}

.playlist-description {
  color: #909399;
  font-size: 14px;
  line-height: 1.6;
  margin: 16px 0 0;
  padding-top: 16px;
  border-top: 1px solid #EBEEF5;
}

.playlist-container {
  flex: 1;
  background-color: #fff;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
}

.filter-section {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

.search-input {
  width: 300px;
}

.song-list {
  border-radius: 8px;
  overflow: hidden;
}

.song-header,
.song-item {
  display: grid;
  grid-template-columns: 60px 3fr 120px 2fr 2fr 100px;
  align-items: center;
  padding: 12px 16px;
  gap: 16px;
}

.song-header {
  background-color: #F5F7FA;
  color: #909399;
  font-size: 14px;
  font-weight: 500;
}

.song-item {
  border-bottom: 1px solid #EBEEF5;
  transition: background-color 0.3s;
}

.song-item:hover {
  background-color: #F5F7FA;
}

.song-cover {
  width: 40px;
  height: 40px;
  border-radius: 4px;
  overflow: hidden;
  margin-right: 12px;
}

.col-title {
  display: flex;
  align-items: center;
}

.song-name {
  font-weight: 500;
  color: #303133;
}

.play-icon {
  cursor: pointer;
  color: #409EFF;
  font-size: 16px;
}

.col-actions .el-button {
  padding: 4px 0;
}

.col-actions .el-button:hover {
  color: #F56C6C;
}
</style>
  