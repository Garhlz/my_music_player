<template>
  <el-dialog
    v-model="visible"
    :title="`${song?.name || '歌曲'} - 更多选项`"
    width="400px"
    :before-close="handleClose"
    class="more-options-dialog"
    :append-to-body="true"
    :close-on-click-modal="true"
  >
    <div class="song-info" v-if="song">
      <!-- 歌曲基本信息 -->
      <div class="info-header">
        <el-image
          :src="song.cover || '/assets/default-cover.jpg'"
          fit="cover"
          class="song-cover"
        />
        <div class="song-details">
          <h3>{{ song.name }}</h3>
          <p class="artist">{{ song.artist_name || '未知歌手' }}</p>
          <p class="album">{{ song.album_name || '未知专辑' }}</p>
        </div>
      </div>

      <!-- 操作选项列表 -->
      <div class="options-list">
        <div class="option-item" @click="handleAddToQueue">
          <el-icon>
            <Plus />
          </el-icon>
          <span>添加到播放队列</span>
        </div>

        <div class="option-item" @click="handleAddToPlaylist">
          <el-icon>
            <FolderAdd />
          </el-icon>
          <span>添加到歌单</span>
        </div>

        <div class="option-item" @click="handleGoToArtist">
          <el-icon>
            <User />
          </el-icon>
          <span>查看歌手</span>
        </div>

        <div class="option-item" @click="handleGoToAlbum">
          <el-icon>
            <Collection />
          </el-icon>
          <span>查看专辑</span>
        </div>

        <div class="option-item" @click="handleViewComments">
          <el-icon>
            <ChatDotRound />
          </el-icon>
          <span>查看评论</span>
        </div>

        <div class="option-item" @click="handleShare">
          <el-icon>
            <Share />
          </el-icon>
          <span>分享歌曲</span>
        </div>

        <div class="option-item" @click="handleDownload">
          <el-icon>
            <Download />
          </el-icon>
          <span>下载歌曲</span>
        </div>

        <div class="option-item" @click="handleViewLyrics">
          <el-icon>
            <Document />
          </el-icon>
          <span>查看歌词</span>
        </div>
      </div>

      <!-- 歌曲统计信息 -->
      <div class="song-stats">
        <div class="stat-item">
          <el-icon>
            <VideoPlay />
          </el-icon>
          <span>{{ song.play_count || 0 }} 次播放</span>
        </div>
        <div class="stat-item">
          <el-icon>
            <Star />
          </el-icon>
          <span>{{ song.like_count || 0 }} 人喜欢</span>
        </div>
        <div class="stat-item">
          <el-icon>
            <Clock />
          </el-icon>
          <span>{{ formatDuration(song.duration) }}</span>
        </div>
      </div>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import { ElMessage } from 'element-plus';
import type { ModelsSongDetailDTO } from '@/api-client';
import {
  Plus, FolderAdd, User, Collection, ChatDotRound,
  Share, Download, Document, VideoPlay, Star, Clock,
} from '@element-plus/icons-vue';
import { formatDuration } from '@/utils/format';
import { usePlayerStore } from '@/stores/player';

// Props
const props = defineProps<{
  modelValue: boolean;
  song: ModelsSongDetailDTO | null;
}>();

// Emits
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void;
  (e: 'addToPlaylist', song: ModelsSongDetailDTO): void;
  (e: 'comment', song: ModelsSongDetailDTO): void;
  (e: 'download', song: ModelsSongDetailDTO): void;
  (e: 'goToArtist', artistId?: number | null): void;
  (e: 'goToAlbum', albumId?: number | null): void;
}>();

const playerStore = usePlayerStore();

// 控制对话框显示
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value),
});

// 处理关闭
const handleClose = () => {
  visible.value = false;
};

// 操作处理函数
const handleAddToQueue = () => {
  if (props.song) {
    // TODO: 实现添加到队列功能
    ElMessage.success('已添加到播放队列');
    handleClose();
  }
};

const handleAddToPlaylist = () => {
  if (props.song) {
    emit('addToPlaylist', props.song);
    handleClose();
  }
};

const handleGoToArtist = () => {
  if (props.song?.artist_id) {
    emit('goToArtist', props.song.artist_id);
    handleClose();
  }
};

const handleGoToAlbum = () => {
  if (props.song?.album_id) {
    emit('goToAlbum', props.song.album_id);
    handleClose();
  }
};

const handleViewComments = () => {
  if (props.song) {
    emit('comment', props.song);
    handleClose();
  }
};

const handleShare = () => {
  if (props.song) {
    // 复制分享链接到剪贴板
    const shareUrl = `${window.location.origin}/song/${props.song.id}`;
    navigator.clipboard.writeText(shareUrl).then(() => {
      ElMessage.success('分享链接已复制到剪贴板');
    }).catch(() => {
      ElMessage.error('复制失败，请手动复制');
    });
    handleClose();
  }
};

const handleDownload = () => {
  if (props.song) {
    emit('download', props.song);
    handleClose();
  }
};

const handleViewLyrics = () => {
  if (props.song) {
    // TODO: 实现查看歌词功能
    ElMessage.info('歌词功能开发中...');
    handleClose();
  }
};
</script>

<style scoped>
/* 对话框深色主题 */
:deep(.more-options-dialog) {
  --el-dialog-bg-color: #282828;
  --el-dialog-title-text-color: #fff;
}

:deep(.el-dialog__header) {
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

/* 歌曲信息区域 */
.song-info {
  color: #fff;
}

.info-header {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.song-cover {
  width: 80px;
  height: 80px;
  border-radius: 4px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.song-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 4px;
}

.song-details h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #fff;
}

.song-details .artist {
  font-size: 14px;
  color: #b3b3b3;
  margin: 0;
}

.song-details .album {
  font-size: 13px;
  color: #999;
  margin: 0;
}

/* 选项列表 */
.options-list {
  display: flex;
  flex-direction: column;
  gap: 2px;
  margin-bottom: 20px;
}

.option-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s ease;
  color: #b3b3b3;
  font-size: 14px;
}

.option-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
  color: #fff;
}

.option-item .el-icon {
  font-size: 20px;
}

/* 统计信息 */
.song-stats {
  display: flex;
  justify-content: space-around;
  padding-top: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #b3b3b3;
  font-size: 13px;
}

.stat-item .el-icon {
  font-size: 16px;
}

/* 响应式 */
@media screen and (max-width: 480px) {
  :deep(.more-options-dialog) {
    width: 90% !important;
  }
}
</style>