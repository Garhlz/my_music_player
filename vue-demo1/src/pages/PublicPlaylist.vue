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
import { getLikedSongs} from '@/api/axiosFile'
// import { getLikedSongs, addLikedSong, removeLikedSong } from '@/api/axiosFile'
//需要添加
import CommonLayout from '@/layouts/CommonLayout.vue'
import {
  Search,
  VideoPlay,
  Star,
  Plus,
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
    const likedResponse = await getLikedSongs()
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
@import '../styles/common.css';
</style>
