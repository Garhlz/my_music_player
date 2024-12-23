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
                  <el-icon @click="addAlbum(song)"><FolderAdd /></el-icon>
                </el-tooltip>
                <el-tooltip content="评论" placement="top">
                  <el-icon @click="showComments(song)"><ChatDotRound /></el-icon>
                </el-tooltip>
              </div>
            </div>

            <div class="col-artist">{{ song.artist }}</div>
            <div class="col-album">{{ song.album }}</div>
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
  store.dispatch('playSong', song)
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
  router.push(`/comments/${song.id}`)
}

// 格式化时长
const formatDuration = (seconds) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
}

// 组件挂载时加载数据
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.playlist-container {
  padding: 20px;
}

.filter-section {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
}

.search-input {
  width: 300px;
}

.song-list {
  margin-top: 20px;
}

.song-header {
  display: flex;
  padding: 10px 0;
  border-bottom: 1px solid #eee;
  font-weight: bold;
}

.song-item {
  display: flex;
  padding: 10px 0;
  border-bottom: 1px solid #eee;
  align-items: center;
}

.col-index {
  width: 50px;
  text-align: center;
}

.col-title {
  flex: 2;
  display: flex;
  align-items: center;
  gap: 10px;
}

.col-duration {
  width: 120px;
  text-align: center;
}

.col-artist {
  flex: 1;
}

.col-album {
  flex: 1;
}

.song-cover {
  width: 40px;
  height: 40px;
  overflow: hidden;
  border-radius: 4px;
}

.action-buttons {
  display: flex;
  gap: 15px;
  justify-content: center;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}

.play-icon {
  cursor: pointer;
  color: var(--el-color-primary);
}

.el-icon {
  cursor: pointer;
}

.liked {
  color: #ffcc00;
}
</style>