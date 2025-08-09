<template>
  <div
    class="song-item"
    :class="{ 'song-item-selected': isPlaying }"
    @mouseenter="emit('update:hoveredId', song.id)"
    @mouseleave="emit('update:hoveredId', null)"
    @click="emit('select', song)"
    @dblclick="emit('play', song)"
    @contextmenu.prevent="openContextMenu($event, song)"
    tabindex="0"
    @keydown="handleKeydown"
  >
    <div class="col-index">
      <span v-if="!isHovered">{{ index }}</span>
      <el-icon v-else class="play-icon" @click.stop="emit('play', song)">
        <VideoPlay />
      </el-icon>
    </div>

    <div class="col-title" @click.stop="emit('play', song)">
      <div class="song-cover">
        <el-image
          :src="song.cover || '/assets/default-cover.jpg'"
          fit="cover"
          lazy
          @error="handleImageError"
        />
      </div>
      <span class="song-name">{{ song.name }}</span>
    </div>

    <div class="col-duration">
      <template v-if="!isHovered">
        {{ formatDuration(song.duration) }}
      </template>
      <div v-else class="action-buttons">
        <el-tooltip
          v-for="action in actions"
          :key="action.event"
          :content="action.content"
          placement="top"
        >
          <el-icon
            @click.stop="emit(action.event, song)"
            :class="{ 'liked': action.event === 'like' && isLiked }"
          >
            <component :is="action.icon" />
          </el-icon>
        </el-tooltip>
      </div>
    </div>

    <div class="col-artist" @click.stop="emit('goToArtist', song.artist_id)">
      {{ song.artist_name || '未知歌手' }}
    </div>

    <div class="col-album" @click.stop="emit('goToAlbum', song.album_id)">
      {{ song.album_name || '未知专辑' }}
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ModelsSongDetailDTO } from '@/api-client';
import { VideoPlay, Star, Plus, ChatDotRound, Download } from '@element-plus/icons-vue';
import { formatDuration } from '@/utils/format';

// 定义 Props
interface Props {
  song: ModelsSongDetailDTO;
  index: number;
  isLiked: boolean;
  isHovered: boolean;
  isPlaying?: boolean;
  columnLayout?: string;
}

const props = defineProps<Props>();

// 定义 Emits
const emit = defineEmits<{
  (e: 'update:hoveredId', id: number | null): void;
  (e: 'play', song: ModelsSongDetailDTO): void;
  (e: 'select', song: ModelsSongDetailDTO): void;
  (e: 'like', song: ModelsSongDetailDTO): void;
  (e: 'addToPlaylist', song: ModelsSongDetailDTO): void;
  (e: 'comment', song: ModelsSongDetailDTO): void;
  (e: 'download', song: ModelsSongDetailDTO): void;
  (e: 'goToArtist', artistId?: number | null): void;
  (e: 'goToAlbum', albumId?: number | null): void;
  (e: 'contextmenu', payload: { song: ModelsSongDetailDTO; x: number; y: number }): void;
}>();

// 操作按钮配置
const actions = [
  { event: 'like', content: '喜欢', icon: Star },
  { event: 'addToPlaylist', content: '添加到歌单', icon: Plus },
  { event: 'comment', content: '评论', icon: ChatDotRound },
  { event: 'download', content: '下载', icon: Download },
];

// 方法
const handleImageError = () => {
  console.warn(`Failed to load cover for song: ${props.song.name}`);
};

const openContextMenu = (event: MouseEvent, song: ModelsSongDetailDTO) => {
  emit('contextmenu', { song, x: event.clientX, y: event.clientY });
};

const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    emit('play', props.song);
  }
};

</script>

<style scoped>
.song-item {
  display: grid;
  grid-template-columns: v-bind(columnLayout || '60px minmax(300px, 2.5fr) 180px minmax(160px, 1fr) minmax(160px, 1fr)');
  padding: 12px 24px;
  border-bottom: 1px solid #ebeef5;
  align-items: center;
  transition: all 0.2s ease-out;
  animation: fadeIn 0.5s ease-in;
  height: 72px; /* 固定高度，配合虚拟列表 */
}

.song-item:hover {
  background-color: #f5f7fa;
  transform: translateX(4px);
}

.song-item-selected {
  border-left: 4px solid var(--el-color-primary);
  background-color: var(--el-color-primary-light-9);
}

.col-index {
  text-align: center;
  color: #909399;
  font-size: 14px;
}

.col-title {
  display: flex;
  align-items: center;
  gap: 16px;
  min-width: 0;
  cursor: pointer;
}

.song-cover {
  width: 48px;
  height: 48px;
  border-radius: 6px;
  overflow: hidden;
  flex-shrink: 0;
  transition: all 0.3s;
}

.song-cover:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.song-name {
  font-weight: 500;
  color: #303133;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.col-duration {
  min-width: 180px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding-right: 24px;
}

.col-artist,
.col-album {
  padding: 0 12px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  cursor: pointer;
  color: #606266;
}

.col-artist:hover,
.col-album:hover {
  color: #409EFF;
}

.action-buttons {
  display: flex;
  gap: 12px;
  justify-content: flex-start;
  align-items: center;
}

.action-buttons .el-icon {
  font-size: 16px;
  cursor: pointer;
  color: #606266;
  padding: 4px;
  border-radius: 50%;
  transition: all 0.2s;
}

.action-buttons .el-icon:hover {
  color: #409EFF;
  background-color: var(--el-fill-color-light);
  transform: scale(1.2);
}

.liked {
  color: #ffcc00 !important;
  animation: heartBeat 0.3s ease-in-out;
}

.play-icon {
  color: var(--el-color-primary);
  font-size: 20px;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes heartBeat {
  0% { transform: scale(1); }
  50% { transform: scale(1.3); }
  100% { transform: scale(1); }
}

@media screen and (max-width: 768px) {
  .song-item {
    grid-template-columns: 50px 3fr 100px;
    padding: 8px 16px;
  }
  .col-album,
  .col-duration {
    display: none;
  }
  .song-cover {
    width: 40px;
    height: 40px;
  }
  .action-buttons {
    gap: 12px;
  }
}
</style>
