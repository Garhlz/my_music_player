<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索专辑"
            class="search-input"
            :prefix-icon="Search"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
          />
        </div>

        <!-- 内容区域 -->
        <div class="content-area">
          <el-loading 
            v-loading="isLoading"
            element-loading-text="加载中..."
            element-loading-background="rgba(255, 255, 255, 0.8)"
          />

          <el-empty
            v-if="!isLoading && albums.length === 0"
            description="还没有收藏的专辑"
          />

          <div class="playlist-grid" v-if="albums.length > 0">
            <div 
              v-for="album in albums" 
              :key="album.id" 
              class="playlist-card"
            >
              <div class="playlist-cover">
                <el-image 
                  :src="album.cover || '/assets/default-cover.jpg'"
                  fit="cover"
                  @error="handleImageError(album.id)"
                />
                <div class="playlist-info">
                  <h3 class="playlist-name">{{ album.name }}</h3>
                  <p class="playlist-count">{{ album.artist }}</p>
                  <p class="playlist-creator">{{ album.songCount || 0 }}首歌曲</p>
                </div>
              </div>
              <div v-if="debugMode">
                <p>专辑 ID: {{ album.id }}</p>
                <p>封面: {{ album.cover }}</p>
                <p>艺术家: {{ album.artist }}</p>
                <p>歌曲数量: {{ album.songCount }}</p>
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
} from '@element-plus/icons-vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { usePlayerStore } from '@/stores/player'

// 响应式状态
const currentName = ref('收藏的专辑')
const albums = ref([])
const page = ref(1)
const pageSize = ref(12)
const totalAlbums = ref(0)
const isLoading = ref(false)
const searchQuery = ref('')
const sortBy = ref('latest')
const hoveredAlbum = ref(null)
const router = useRouter()
const debugMode = ref(true)
const playerStore = usePlayerStore()

// 加载数据
const loadData = async () => {
  isLoading.value = true
  try {
    // TODO: 实现获取收藏的专辑列表的接口
    isLoading.value = false
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败，请稍后重试')
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

// 播放专辑
const playAlbum = (album) => {
  if (album.songs && album.songs.length > 0) {
    playerStore.setPlaylist(album.songs)
    playerStore.play(0)
  }
}

// 取消收藏
const uncollectAlbum = (album) => {
  ElMessage.success(`已取消收藏专辑: ${album.name}`)
}

const handleImageError = (id) => {
  console.error(`专辑 ${id} 的封面加载失败`)
  // 可以在这里设置一个默认封面或其他处理
}

// 生命周期钩子
onMounted(() => {
  loadData()
})
</script>

<style scoped>
@import '../styles/common.css';

/* 主容器 */
.playlist-container {
  height: calc(100vh - 64px - 60px); /* 64px是顶部导航，60px是播放器高度 */
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
  gap: 16px;
  padding: 16px 24px;
  background: var(--el-bg-color);
  border-bottom: 1px solid var(--el-border-color-lighter);
}

.search-input {
  width: 300px;
}

/* 内容区域 */
.content-area {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
  /* 为分页器预留空间 */
  padding-bottom: calc(24px + 56px); /* 56px是分页器的高度 */
}

/* 专辑网格布局 */
.playlist-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 16px;
}

/* 专辑卡片 */
.playlist-card {
  background: #fff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

/* 封面区域 */
.playlist-cover {
  position: relative;
}

.playlist-info {
  padding: 12px;
}

.playlist-name {
  font-size: 14px;
  font-weight: bold;
}

.playlist-count,
.playlist-creator {
  font-size: 12px;
  color: #909399;
}

/* 分页器 */
.pagination {
  position: fixed;
  bottom: 60px; /* 播放器高度 */
  left: 200px; /* 侧边栏宽度 */
  right: 0;
  height: 56px;
  padding: 0 24px;
  background-color: var(--el-bg-color);
  display: flex;
  justify-content: center;
  align-items: center;
  border-top: 1px solid var(--el-border-color-lighter);
  z-index: 10;
}

/* 响应式设计 */
@media screen and (max-width: 1600px) {
  .playlist-grid {
    grid-template-columns: repeat(3, minmax(280px, 1fr));
  }
}

@media screen and (max-width: 1200px) {
  .playlist-grid {
    grid-template-columns: repeat(2, minmax(280px, 1fr));
  }
}

@media screen and (max-width: 768px) {
  .playlist-container {
    padding: 16px;
    height: calc(100vh - 160px);
  }

  .playlist-grid {
    grid-template-columns: repeat(1, 1fr);
    gap: 16px;
  }

  .playlist-card {
    height: 320px;
  }

  .playlist-cover {
    height: 240px;
  }

  .search-input {
    width: 100%;
  }

  .filter-section {
    flex-direction: column;
    gap: 12px;
  }

  .pagination {
    left: 0;
  }
}
</style>