<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <!-- 搜索和筛选区域 -->
        <div class="filter-section">
          <el-input
            v-model="searchQuery"
            placeholder="搜索上传的音乐"
            class="search-input"
            :prefix-icon="Search"
            clearable
            @clear="handleSearch"
            @keyup.enter="handleSearch"
          />
          <el-button type="primary" @click="uploadMusic">
            <el-icon><Upload /></el-icon>上传音乐
          </el-button>
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
          description="还没有上传音乐"
        >
          <el-button type="primary" @click="uploadMusic">上传音乐</el-button>
        </el-empty>

        <!-- 歌曲列表 -->
        <div class="song-list" v-if="songs.length > 0">
          <div class="song-header">
            <div class="col-index">#</div>
            <div class="col-title">标题</div>
            <div class="col-duration">时长</div>
            <div class="col-artist">歌手</div>
            <div class="col-album">专辑</div>
            <div class="col-status">状态</div>
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
                <el-tooltip content="编辑" placement="top">
                  <el-icon @click="editSong(song)"><Edit /></el-icon>
                </el-tooltip>
                <el-tooltip content="删除" placement="top">
                  <el-icon @click="deleteSong(song)"><Delete /></el-icon>
                </el-tooltip>
              </div>
            </div>

            <div class="col-artist">{{ song.artist }}</div>
            <div class="col-album">{{ song.album }}</div>
            <div class="col-status">
              <el-tag :type="song.status === 'approved' ? 'success' : 'warning'">
                {{ song.status === 'approved' ? '已通过' : '审核中' }}
              </el-tag>
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

        <!-- 上传音乐对话框 -->
        <el-dialog
          v-model="uploadDialogVisible"
          title="上传音乐"
          width="500px"
        >
          <el-form :model="uploadForm" label-width="80px">
            <el-form-item label="音乐文件">
              <el-upload
                class="upload-demo"
                drag
                action="/upload"
                :auto-upload="false"
                :on-change="handleFileChange"
              >
                <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                <div class="el-upload__text">
                  将文件拖到此处，或<em>点击上传</em>
                </div>
                <template #tip>
                  <div class="el-upload__tip">
                    支持 mp3 格式，单个文件不超过 20MB
                  </div>
                </template>
              </el-upload>
            </el-form-item>
            <el-form-item label="歌曲名称">
              <el-input v-model="uploadForm.name" />
            </el-form-item>
            <el-form-item label="歌手">
              <el-input v-model="uploadForm.artist" />
            </el-form-item>
            <el-form-item label="专辑">
              <el-input v-model="uploadForm.album" />
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="uploadDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitUpload">上传</el-button>
          </template>
        </el-dialog>
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
  Upload,
  Edit,
  Delete,
  UploadFilled
} from '@element-plus/icons-vue'
import { mapActions } from 'vuex'

export default {
  name: 'Uploaded',
  components: {
    CommonLayout,
    Search,
    VideoPlay,
    Upload,
    Edit,
    Delete,
    UploadFilled
  },
  data() {
    return {
      currentName: '上传的音乐',
      songs: [],
      page: 1,
      pageSize: 10,
      totalSongs: 0,
      isLoading: false,
      searchQuery: '',
      hoveredSong: null,
      uploadDialogVisible: false,
      uploadForm: {
        name: '',
        artist: '',
        album: '',
        file: null
      }
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
        // TODO: 实现获取上传音乐列表的接口
        // const response = await getUploadedSongs(this.page, this.pageSize)
        this.isLoading = false
      } catch (error) {
        console.error('加载数据失败:', error)
        ElMessage.error('加载数据失败，请稍后重试')
        this.isLoading = false
      }
    },

    uploadMusic() {
      this.uploadDialogVisible = true
    },

    handleFileChange(file) {
      this.uploadForm.file = file.raw
    },

    async submitUpload() {
      if (!this.uploadForm.file) {
        ElMessage.warning('请选择要上传的音乐文件')
        return
      }
      if (!this.uploadForm.name) {
        ElMessage.warning('请输入歌曲名称')
        return
      }
      // TODO: 实现上传音乐的接口
      ElMessage.success('音乐上传成功，等待审核')
      this.uploadDialogVisible = false
      this.uploadForm = {
        name: '',
        artist: '',
        album: '',
        file: null
      }
      await this.loadData()
    },

    async editSong(song) {
      // TODO: 实现编辑音乐的功能
      ElMessage.success(`编辑音乐: ${song.name}`)
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

.upload-demo {
  width: 100%;
}
</style>