<template>
  <div
    class="song-item"
    :class="{
      'song-item-playing': isPlaying,
      'song-item-hovered': isHovered
    }"
    @mouseenter="emit('update:hoveredId', song.id)"
    @mouseleave="emit('update:hoveredId', null)"
    @dblclick="handleDoubleClick"
    tabindex="0"
    @keydown="handleKeydown"
  >
    <!-- 序号/播放按钮 -->
    <div class="col-index">
      <span v-if="!isHovered && !isPlaying" class="index-number">{{ index }}</span>
      <el-icon v-else-if="isPlaying && !isHovered" class="playing-icon">
        <VideoPause />
      </el-icon>
      <el-icon v-else class="play-icon" @click.stop="handlePlayClick">
        <VideoPlay />
      </el-icon>
    </div>

    <!-- 歌曲信息（封面+标题+歌手） -->
    <div class="col-info">
      <div class="song-cover">
        <el-image
          :src="song.cover || '/assets/default-cover.jpg'"
          fit="cover"
          lazy
          @error="handleImageError"
        />
        <!-- hover时的播放遮罩 -->
        <div v-if="isHovered" class="cover-overlay" @click.stop="handlePlayClick">
          <el-icon>
            <VideoPlay />
          </el-icon>
        </div>
      </div>
      <div class="song-details">
        <div class="song-name" @click.stop="handlePlayClick">{{ song.name }}</div>
        <div class="artist-name" @click.stop="emit('goToArtist', song.artist_id)">
          {{ song.artist_name || '未知歌手' }}
        </div>
      </div>
    </div>

    <!-- 专辑名 -->
    <div class="col-album">
      <span class="album-name" @click.stop="emit('goToAlbum', song.album_id)">{{ song.album_name || '未知专辑' }}</span>
    </div>

    <!-- 操作按钮（hover时显示） -->
    <div class="col-actions">
      <transition name="fade">
        <div v-if="isHovered" class="action-buttons">
          <!-- 喜欢按钮 -->
          <el-tooltip content="喜欢" placement="top">
            <el-icon
              class="action-icon"
              :class="{ 'liked': isLiked }"
              @click.stop="emit('like', song)"
            >
              <StarFilled v-if="isLiked" />
              <Star v-else />
            </el-icon>
          </el-tooltip>

          <!-- 添加到歌单 -->
          <el-tooltip content="添加到歌单" placement="top">
            <el-icon class="action-icon" @click.stop="emit('addToPlaylist', song)">
              <Plus />
            </el-icon>
          </el-tooltip>

          <!-- 更多选项 -->
          <el-tooltip content="更多选项" placement="top">
            <el-icon class="action-icon" @click.stop="handleMoreOptions">
              <MoreFilled />
            </el-icon>
          </el-tooltip>
        </div>
      </transition>
    </div>

    <!-- 时长 -->
    <div class="col-duration">
      <span>{{ formatDuration(song.duration) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ModelsSongDetailDTO } from '@/api-client';
import {
  VideoPlay, VideoPause, Star, StarFilled, Plus, MoreFilled,
} from '@element-plus/icons-vue';
import { formatDuration } from '@/utils/format';

// Props定义
interface Props {
  song: ModelsSongDetailDTO;
  index: number;
  isLiked: boolean;
  isHovered: boolean;
  isPlaying?: boolean;
}

const props = defineProps<Props>();

// 事件定义
const emit = defineEmits<{
  (e: 'update:hoveredId', id: number | null): void;
  (e: 'play', song: ModelsSongDetailDTO): void;
  (e: 'like', song: ModelsSongDetailDTO): void;
  (e: 'addToPlaylist', song: ModelsSongDetailDTO): void;
  (e: 'showMoreOptions', song: ModelsSongDetailDTO): void;
  (e: 'goToArtist', artistId?: number | null): void;
  (e: 'goToAlbum', albumId?: number | null): void;
}>();

// 处理播放点击
const handlePlayClick = () => {
  emit('play', props.song);
};

// 处理双击播放
const handleDoubleClick = () => {
  emit('play', props.song);
};

// 处理更多选项
const handleMoreOptions = () => {
  emit('showMoreOptions', props.song);
};

// 处理图片加载错误
const handleImageError = () => {
  console.warn(`Failed to load cover for song: ${props.song.name}`);
};

// 键盘事件处理
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    emit('play', props.song);
  }
};
</script>

<style scoped>
/* Spotify风格的歌曲项 */
.song-item {
  display: grid;
  grid-template-columns: 48px 1fr minmax(120px, 1fr) 180px 60px;
  gap: 16px;
  padding: 8px 16px;
  align-items: center;
  border-radius: 4px;
  transition: background-color 0.2s ease;
  cursor: default;
  user-select: none;
}

/* hover效果 */
.song-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

/* 正在播放的歌曲 */
.song-item-playing {
  background-color: rgba(255, 255, 255, 0.05);
}

/* 序号列 */
.col-index {
  display: flex;
  justify-content: center;
  align-items: center;
  color: #b3b3b3;
  font-size: 14px;
  font-variant-numeric: tabular-nums;
}

.index-number {
  font-weight: 400;
}

.play-icon,
.playing-icon {
  font-size: 16px;
  cursor: pointer;
  transition: transform 0.1s ease;
}

.play-icon:hover {
  transform: scale(1.1);
}

.playing-icon {
  color: #1db954;
}

/* 歌曲信息列 */
.col-info {
  display: flex;
  align-items: center;
  gap: 12px;
  min-width: 0;
}

.song-cover {
  position: relative;
  width: 40px;
  height: 40px;
  flex-shrink: 0;
  border-radius: 2px;
  overflow: hidden;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.2);
}

.song-cover .el-image {
  width: 100%;
  height: 100%;
}

/* 封面hover播放遮罩 */
.cover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: opacity 0.2s ease;
}

.cover-overlay .el-icon {
  color: white;
  font-size: 18px;
}

/* 歌曲详情 */
.song-details {
  display: flex;
  flex-direction: column;
  min-width: 0;
  gap: 2px;
}

.song-name {
  font-size: 14px;
  font-weight: 400;
  color: #fff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  cursor: pointer;
  line-height: 1.4;
}

.song-name:hover {
  text-decoration: underline;
}

.artist-name {
  font-size: 12px;
  color: #b3b3b3;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  cursor: pointer;
  line-height: 1.4;
}

.artist-name:hover {
  color: #fff;
  text-decoration: underline;
}

/* 专辑列 */
.col-album {
  color: #b3b3b3;
  font-size: 13px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 新增或修改 .album-name 的样式 */
.album-name {
  /* display: inline-block; */ /* 可选，如果需要更精确的盒模型控制 */
  cursor: pointer; /* 将手型光标应用到 span 上 */
  transition: color 0.2s ease; /* 平滑颜色过渡 */
}

/* 将 :hover 伪类选择器直接作用于 .album-name */
.album-name:hover {
  color: #fff;
  text-decoration: underline;
}

/* 操作按钮列 */
.col-actions {
  display: flex;
  justify-content: center;
  align-items: center;
  min-width: 120px;
}

.action-buttons {
  display: flex;
  gap: 12px;
  align-items: center;
}

.action-icon {
  font-size: 24px;
  color: #b3b3b3;
  cursor: pointer;
  padding: 2px;
  transition: all 0.2s ease;
}

.action-icon:hover {
  color: #fff;
  transform: scale(1.1);
}

.action-icon.liked {
  color: #1db954;
}

/* 时长列 */
.col-duration {
  color: #b3b3b3;
  font-size: 13px;
  text-align: right;
  font-variant-numeric: tabular-nums;
}

/* 过渡动画 */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .song-item {
    grid-template-columns: 40px 1fr 50px;
    padding: 8px 12px;
  }

  .col-album,
  .col-actions {
    display: none;
  }

  .song-cover {
    width: 36px;
    height: 36px;
  }

  .song-name {
    font-size: 13px;
  }

  .artist-name {
    font-size: 11px;
  }
}
</style>