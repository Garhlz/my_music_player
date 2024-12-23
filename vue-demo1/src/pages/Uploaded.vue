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

<script setup>
import { ref, onMounted } from 'vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import { usePlayerStore } from '@/stores/player'

const playerStore = usePlayerStore()

const currentName = ref('我的上传')
const songs = ref([])
const isLoading = ref(false)
const searchQuery = ref('')

const handlePlaySong = (song) => {
  playerStore.setPlaylist([song])
  playerStore.play(0)
}

const loadData = async () => {
  isLoading.value = true
  try {
    // TODO: 实现获取上传音乐列表的接口
    isLoading.value = false
  } catch (error) {
    console.error('加载数据失败:', error)
    ElMessage.error('加载数据失败，请稍后重试')
    isLoading.value = false
  }
}

const handleSearch = () => {
  loadData()
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
@import '../styles/common.css';

.upload-demo {
  width: 100%;
}
</style>