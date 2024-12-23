<template>
  <CommonLayout :page-name="currentPlaylist.name">
    <template #main>
      <div class="playlist-container">
        <!-- 歌单头部信息 -->
        <div class="playlist-header">
          <div class="playlist-info">
            <el-image 
              :src="currentPlaylist.cover || '/assets/default-cover.jpg'"
              class="playlist-cover"
              fit="cover"
            />
            <div class="playlist-details">
              <h2 class="playlist-name">{{ currentPlaylist.name }}</h2>
              <p class="playlist-description">{{ currentPlaylist.description || '暂无描述' }}</p>
              <p class="playlist-meta">
                创建者：{{ currentPlaylist.creatorName }} | 
                {{ songs.length }}首歌曲
              </p>
            </div>
          </div>
        </div>

        <!-- 操作栏 -->
        <div class="operation-bar">
          <el-button type="primary" @click="playAll">
            <el-icon><CaretRight /></el-icon>播放全部
          </el-button>
        </div>

        <!-- 歌曲列表 -->
        <div class="song-list">
          <el-table
            v-loading="isLoading"
            :data="songs"
            style="width: 100%"
            @row-dblclick="playSong"
          >
            <el-table-column type="index" width="50" />
            <el-table-column label="歌曲" min-width="200">
              <template #default="{ row }">
                <div class="song-info">
                  <span class="song-name">{{ row.name }}</span>
                  <span v-if="row.mark" class="song-mark">{{ row.mark }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="artist" label="歌手" min-width="120" />
            <el-table-column prop="album" label="专辑" min-width="120" />
            <el-table-column label="操作" width="150" fixed="right">
              <template #default="{ row }">
                <div class="song-actions">
                  <el-button 
                    circle 
                    :icon="CaretRight"
                    @click="playSong(row)"
                  />
                  <el-button 
                    circle 
                    :icon="Delete"
                    @click="removeSong(row)"
                  />
                </div>
              </template>
            </el-table-column>
          </el-table>

          <el-empty
            v-if="!isLoading && songs.length === 0"
            description="该歌单没有歌曲"
          />
        </div>
      </div>
    </template>
  </CommonLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { CaretRight, Delete, ArrowLeft, ArrowRight, VideoPlay, VideoPause } from '@element-plus/icons-vue'
import { getPlaylistSongs, removeFromPlaylist } from '@/api/axiosFile'
import { usePlayerStore } from '@/stores/player'

const route = useRoute()
const router = useRouter()
const playerStore = usePlayerStore()

const currentPlaylist = ref({})
const songs = ref([])
const isLoading = ref(false)

// 加载歌单数据
const loadData = async () => {
  isLoading.value = true
  try {
    const response = await getPlaylistSongs(route.params.id)
    if (response.data.message) {
      currentPlaylist.value = response.data.data.playlist
      songs.value = response.data.data.songs || []
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

// 播放全部
const playAll = () => {
  if (songs.value.length === 0) {
    ElMessage.warning('歌单中没有歌曲')
    return
  }
  playerStore.setPlaylist(songs.value)
  playerStore.play(0)
}

// 播放单首歌曲
const playSong = (song) => {
  playerStore.setPlaylist(songs.value)
  const index = songs.value.findIndex(s => s.id === song.id)
  playerStore.play(index)
}

// 从歌单中移除歌曲
const removeSong = (song) => {
  ElMessageBox.confirm(
    `确定要从歌单中移除 "${song.name}" 吗？`,
    '移除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async () => {
    try {
      const response = await removeFromPlaylist(currentPlaylist.value.id, song.id)
      if (response.data.message) {
        ElMessage.success('移除成功')
        loadData() // 重新加载数据
      } else {
        throw new Error(response.data.error || '移除失败')
      }
    } catch (error) {
      console.error('移除失败:', error)
      ElMessage.error(error.message || '移除失败，请稍后重试')
    }
  }).catch(() => {})
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.playlist-container {
  padding: 20px;
}

.playlist-header {
  margin-bottom: 24px;
}

.playlist-info {
  display: flex;
  gap: 24px;
}

.playlist-cover {
  width: 200px;
  height: 200px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.playlist-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.playlist-name {
  font-size: 24px;
  font-weight: 600;
  margin: 0 0 12px;
}

.playlist-description {
  font-size: 14px;
  color: var(--el-text-color-secondary);
  margin: 0 0 16px;
}

.playlist-meta {
  font-size: 13px;
  color: var(--el-text-color-secondary);
  margin: 0;
}

.operation-bar {
  margin-bottom: 24px;
}

.song-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.song-name {
  font-weight: 500;
}

.song-mark {
  font-size: 12px;
  color: var(--el-color-primary);
  border: 1px solid var(--el-color-primary);
  padding: 0 4px;
  border-radius: 4px;
}

.song-actions {
  display: flex;
  gap: 8px;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .playlist-info {
    flex-direction: column;
  }

  .playlist-cover {
    width: 150px;
    height: 150px;
    margin: 0 auto;
  }

  .playlist-details {
    text-align: center;
  }
}
</style> 