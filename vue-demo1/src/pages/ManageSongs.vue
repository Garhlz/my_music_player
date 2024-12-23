<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索音乐"
            class="search-input"
            :prefix-icon="Search"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
          />
          <el-select v-model="status" placeholder="审核状态" @change="handleSearch">
            <el-option label="全部" value="" />
            <el-option label="待审核" value="pending" />
            <el-option label="已通过" value="approved" />
            <el-option label="已拒绝" value="rejected" />
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
          description="暂无音乐"
        />

        <!-- 歌曲列表 -->
        <div class="song-list" v-if="songs.length > 0">
          <div class="song-header">
            <div class="col-index">#</div>
            <div class="col-title">标题</div>
            <div class="col-duration">时长</div>
            <div class="col-artist">歌手</div>
            <div class="col-album">专辑</div>
            <div class="col-uploader">上传者</div>
            <div class="col-status">状态</div>
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
            <div class="col-uploader">{{ song.uploader }}</div>
            
            <div class="col-status">
              <el-tag :type="getStatusType(song.status)">
                {{ getStatusText(song.status) }}
              </el-tag>
            </div>

            <div class="col-actions">
              <el-button-group>
                <el-button 
                  v-if="song.status === 'pending'"
                  type="success" 
                  size="small"
                  @click="approveSong(song)"
                >
                  通过
                </el-button>
                <el-button 
                  v-if="song.status === 'pending'"
                  type="danger" 
                  size="small"
                  @click="rejectSong(song)"
                >
                  拒绝
                </el-button>
                <el-button 
                  type="danger" 
                  size="small"
                  @click="deleteSong(song)"
                >
                  删除
                </el-button>
              </el-button-group>
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

<script>
import { ref } from 'vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Search,
  VideoPlay,
} from '@element-plus/icons-vue'
import { mapActions } from 'vuex'

export default {
  name: 'ManageSongs',
  components: {
    CommonLayout,
    Search,
    VideoPlay,
  },
  data() {
    return {
      currentName: '音乐管理',
      songs: [],
      page: 1,
      pageSize: 10,
      totalSongs: 0,
      isLoading: false,
      searchQuery: '',
      status: '',
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

    getStatusType(status) {
      const types = {
        pending: 'warning',
        approved: 'success',
        rejected: 'danger'
      }
      return types[status] || 'info'
    },

    getStatusText(status) {
      const texts = {
        pending: '待审核',
        approved: '已通过',
        rejected: '已拒绝'
      }
      return texts[status] || '未知'
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
        // TODO: 实现获取音乐列表的接口
        // const response = await getSongs(this.page, this.pageSize, this.status)
        this.isLoading = false
      } catch (error) {
        console.error('加载数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
        this.isLoading = false
      }
    },

    async approveSong(song) {
      try {
        await ElMessageBox.confirm(
          '确定通过这首音乐的审核吗？',
          '审核音乐',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }
        )
        // TODO: 实现通过审核的接口
        ElMessage.success(`已通过音乐: ${song.name}`)
        await this.loadData()
      } catch {
        // 用户取消操作
      }
    },

    async rejectSong(song) {
      try {
        await ElMessageBox.confirm(
          '确定拒绝这首音乐的审核吗？',
          '审核音乐',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        // TODO: 实现拒绝审核的接口
        ElMessage.success(`已拒绝音乐: ${song.name}`)
        await this.loadData()
      } catch {
        // 用户取消操作
      }
    },

    async deleteSong(song) {
      try {
        await ElMessageBox.confirm(
          '确定要删除这首音乐吗？',
          '删除音乐',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        // TODO: 实现删除音乐的接口
        ElMessage.success(`音乐已删除: ${song.name}`)
        await this.loadData()
      } catch {
        // 用户取消删除
      }
    }
  }
}
</script>

<style scoped>
@import '../styles/common.css';

.col-uploader {
  flex: 1;
  padding: 0 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.col-status {
  width: 80px;
  text-align: center;
}

.col-actions {
  width: 200px;
  text-align: center;
}
</style>