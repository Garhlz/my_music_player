<template>
    <CommonLayout :page-name="currentName">
      <template #main>
        <div class="main-container">
          <!-- 左侧歌单信息 -->
          <div class="playlist-aside">
            <div class="playlist-info-card">
              <div class="playlist-cover">
                <el-image 
                  :src="currnetAlbum?.cover || '/assets/default-cover.jpg'"
                  fit="cover"
                />
              </div>
              <h2>{{ currnetAlbum?.name || '我的歌单' }}</h2>
              <div class="playlist-info">
                <span>{{ songs.length }}首歌曲</span>
              </div>
              <p class="playlist-description">{{ currnetAlbum?.description || '暂无描述' }}</p>
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
                    :src="currnetAlbum?.cover || '/assets/default-cover.jpg'"
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
                    <el-tooltip content="评论" placement="top">
                      <el-icon @click="goToComment(song)"><ChatDotRound /></el-icon>
                    </el-tooltip>
                    <el-tooltip content="下载" placement="top">
                      <el-icon @click="downloadSong(song)"><Download /></el-icon>
                    </el-tooltip>
                  </div>
                </div>
                <div class="col-artist" @click="goToArtist(song.author_id)">{{ song.artist }}</div>
                <div class="col-album" @click="goToAlbum(song.album_id)">{{ song.album }}</div>
              </div>
            </div>
          </div>
        </div>
  
         <!-- 添加到播放列表弹窗 -->
        <el-dialog
          v-model="playlistDialogVisible"
          title="添加到歌单"
          width="400px"
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
    getAlbums,
    getAlbumById,
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
  const currentName = ref('专辑详情')
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
  const currnetAlbum = ref(null)
  // 方法定义
  const loadData = async () => {
    isLoading.value = true
    try {
      // 1. 获取歌曲列表
      //注意这里的route文件获取当前路径id的方式
      const songsResponse = await getAlbumById(router.currentRoute.value.params.id,{
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
      currnetAlbum.value = songsData.album;
      // 2. 获取喜欢列表
      const userId = localStorage.getItem('userId')
      const likedResponse = await getLikedSongsById(userId)
      if (likedResponse.data.message) {
        likedSongs.value = likedResponse.data.data || []
      }
      // 3. 组装完整的歌曲信息,只是重新处理罢了
      console.log(songsList);
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
      const response = await getMyPlaylists({
      id: parseInt(localStorage.getItem('userId')),
    })
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
  const goToComment = (song) => {
    router.push(`/comment/${song.id}`)
  }
  //TODO 管理员可以对专辑页面进行增删改
//   const removeSong = async (song) => {
//     try {
//       const response = await removeSongFromPlaylist(currnetAlbum.value.id, song.id)
//       if (response.data.message) {
//         ElMessage.success('删除成功')
//         loadData()
//       } else {
//         throw new Error(response.data.error || '删除失败')
//       }
//     } catch (error) {
//       console.error('删除失败:', error)
//       ElMessage.error(error.message || '删除失败，请稍后重试')
//     }
//   }
  
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
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 12px;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
    transition: all 0.3s ease;
  }
  
  .playlist-info-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.12);
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
    transition: transform 0.6s ease;
  }
  
  .playlist-cover:hover .el-image {
    transform: scale(1.1);
  }
  
  .playlist-info-card h2 {
    margin: 0;
    font-size: 24px;
    color: #303133;
    text-align: center;
    transition: color 0.3s ease;
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
    transition: all 0.3s ease;
  }
  
  .search-input:focus-within {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
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
    color: var(--el-color-primary);
    transition: transform 0.3s ease;
  }
  
  .play-icon:hover {
    transform: scale(1.2);
  }
  
  .col-actions .el-button {
    padding: 4px 0;
  }
  
  .col-actions .el-button:hover {
    color: #F56C6C;
  }
  
  .action-buttons {
    display: flex;
    gap: 16px;
    justify-content: flex-start;
    align-items: center;
  }
  
  .action-buttons .el-icon {
    font-size: 16px;
    cursor: pointer;
    color: var(--el-text-color-secondary);
    transition: all 0.3s ease;
  }
  
  .action-buttons .el-icon:hover {
    color: var(--el-color-primary);
    transform: scale(1.2);
  }
  
  .action-buttons .liked {
    color: #ff4757;
    animation: heartBeat 0.3s ease-in-out;
  }
  
  @keyframes heartBeat {
    0% { transform: scale(1); }
    50% { transform: scale(1.3); }
    100% { transform: scale(1.1); }
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
  
  /* 添加跳转提示动画 */
  .col-artist,
  .col-album {
    cursor: pointer;
    position: relative;
    transition: color 0.3s ease;
  }
  
  .col-artist:hover,
  .col-album:hover {
    color: var(--el-color-primary);
  }
  
  .col-artist::after,
  .col-album::after {
    content: '→';
    position: absolute;
    right: -20px;
    opacity: 0;
    transform: translateX(-10px);
    transition: all 0.3s ease;
  }
  
  .col-artist:hover::after,
  .col-album:hover::after {
    opacity: 1;
    transform: translateX(0);
  }
  
  /* 添加点击反馈动画 */
  .col-artist:active,
  .col-album:active {
    transform: scale(0.98);
  }
  </style>
    