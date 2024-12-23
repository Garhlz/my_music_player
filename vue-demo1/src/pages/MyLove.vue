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
                  <el-icon @click="unlikeSong(song)"><Star /></el-icon>
                </el-tooltip>
                <el-tooltip content="添加到播放列表" placement="top">
                  <el-icon @click="addToPlaylist(song)"><Plus /></el-icon>
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

<script>
import { ref } from 'vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage } from 'element-plus'
import {
  Search,
  VideoPlay,
  Star,
  Plus,
  ChatDotRound,
} from '@element-plus/icons-vue'
import { mapActions } from 'vuex'

export default {
  name: 'MyLove',
  components: {
    CommonLayout,
    Search,
    VideoPlay,
    Star,
    Plus,
    ChatDotRound,
  },
  data() {
    return {
      currentName: '我喜欢的音乐',
      songs: [],
      page: 1,
      pageSize: 10,
      totalSongs: 0,
      isLoading: false,
      searchQuery: '',
      sortBy: 'latest',
      hoveredSong: null,
    }
  },
  mounted() {
    this.loadData()
  },
  methods: {
    ...mapActions(['playSong']),

    handlePlaySong(song) {
      this.playSong(song)
    },

    formatDuration(seconds) {
      const minutes = Math.floor(seconds / 60)
      const remainingSeconds = seconds % 60
      return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`
    },

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
        // TODO: 实现获取喜欢的音乐列表的接口
        // const response = await getLikedSongs(this.page, this.pageSize)
        this.isLoading = false
      } catch (error) {
        console.error('加载数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
        this.isLoading = false
      }
    },

    unlikeSong(song) {
      // TODO: 实现取消喜欢音乐的接口
      ElMessage.success(`已取消喜欢: ${song.name}`)
    },

    addToPlaylist(song) {
      // TODO: 实现添加到播放列表的功能
      ElMessage.success(`已添加到播放列表: ${song.name}`)
    },

    showComments(song) {
      this.$router.push(`/comments/${song.id}`)
    },
  }
}
</script>

<style scoped>
@import '../styles/common.css';
</style>