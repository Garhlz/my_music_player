<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索收藏的专辑"
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
          v-if="!isLoading && albums.length === 0"
          description="还没有收藏的专辑"
        />

        <!-- 专辑网格 -->
        <div class="album-grid" v-if="albums.length > 0">
          <div 
            v-for="album in albums" 
            :key="album.id" 
            class="album-card"
            @mouseenter="hoveredAlbum = album.id"
            @mouseleave="hoveredAlbum = null"
          >
            <div class="album-cover">
              <el-image 
                :src="album.cover || '/assets/default-cover.jpg'"
                fit="cover"
              />
              <div class="album-actions" v-show="hoveredAlbum === album.id">
                <el-icon @click="playAlbum(album)"><VideoPlay /></el-icon>
                <el-icon @click="uncollectAlbum(album)"><Star /></el-icon>
              </div>
            </div>
            <div class="album-info">
              <h3 class="album-name">{{ album.name }}</h3>
              <p class="album-artist">{{ album.artist }}</p>
              <p class="album-count">{{ album.songCount }}首歌曲</p>
            </div>
          </div>
        </div>

        <!-- 分页器 -->
        <el-pagination
          v-if="albums.length > 0"
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :total="totalAlbums"
          :page-sizes="[12, 24, 36, 48]"
          layout="total, sizes, prev, pager, next"
          class="pagination"
          @size-change="handlePageSizeChange"
          @current-change="handlePageChange"
        />
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
  Star,
} from '@element-plus/icons-vue'

export default {
  name: 'MyAlbum',
  components: {
    CommonLayout,
    Search,
    VideoPlay,
    Star,
  },
  data() {
    return {
      currentName: '收藏的专辑',
      albums: [],
      page: 1,
      pageSize: 12,
      totalAlbums: 0,
      isLoading: false,
      searchQuery: '',
      sortBy: 'latest',
      hoveredAlbum: null,
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
        // TODO: 实现获取收藏的专辑列表的接口
        // const response = await getCollectedAlbums(this.page, this.pageSize)
        this.isLoading = false
      } catch (error) {
        console.error('加载数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
        this.isLoading = false
      }
    },

    playAlbum(album) {
      // TODO: 实现播放专辑的功能
      ElMessage.success(`开始播放专辑: ${album.name}`)
    },

    uncollectAlbum(album) {
      // TODO: 实现取消收藏专辑的接口
      ElMessage.success(`已取消收藏专辑: ${album.name}`)
    },
  }
}
</script>

<style scoped>
@import '../styles/common.css';

.album-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
  padding: 20px;
}

.album-card {
  background: #fff;
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.3s ease;
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.album-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0,0,0,0.2);
}

.album-cover {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
}

.album-cover .el-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.album-actions {
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

.album-actions .el-icon {
  font-size: 24px;
  color: white;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.album-actions .el-icon:hover {
  transform: scale(1.2);
}

.album-info {
  padding: 12px;
}

.album-name {
  margin: 0;
  font-size: 16px;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.album-artist {
  margin: 4px 0;
  font-size: 14px;
  color: #666;
}

.album-count {
  margin: 4px 0 0;
  font-size: 12px;
  color: #999;
}
</style>