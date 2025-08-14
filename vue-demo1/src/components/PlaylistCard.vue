<template>
  <div class="playlist-card" @click="emit('view', playlist.id!)">
    <!-- 封面容器 -->
    <div class="playlist-cover">
      <el-image
        :src="playlist.cover || '/assets/default-cover.jpg'"
        fit="cover"
        class="cover-image"
        lazy
      />

      <!-- 悬浮遮罩层 -->
      <div class="playlist-hover-overlay">
        <!-- 播放按钮（主要） -->
        <div class="main-play-button" @click.stop="emit('play', playlist.id!)">
          <el-icon>
            <VideoPlay />
          </el-icon>
        </div>

        <!-- 其他操作按钮 -->
        <div class="action-buttons" v-if="isOwner">
          <el-tooltip content="编辑歌单" placement="top" :hide-after="0">
            <el-button
              circle
              size="small"
              class="action-btn edit-btn"
              @click.stop="emit('edit', playlist)"
            >
              <el-icon>
                <Edit />
              </el-icon>
            </el-button>
          </el-tooltip>

          <el-tooltip content="删除歌单" placement="top" :hide-after="0">
            <el-button
              circle
              size="small"
              class="action-btn delete-btn"
              @click.stop="emit('delete', playlist)"
            >
              <el-icon>
                <Delete />
              </el-icon>
            </el-button>
          </el-tooltip>
        </div>
      </div>

      <!-- 渐变遮罩 -->
      <div class="cover-gradient"></div>
    </div>

    <!-- 歌单信息 -->
    <div class="playlist-info">
      <h3 class="playlist-name" :title="playlist.name">
        {{ playlist.name }}
      </h3>

      <div class="playlist-meta">
        <span class="playlist-count">
          {{ playlist.song_count || 0 }} 首歌曲
        </span>
        <span class="playlist-visibility" :class="{ 'public': playlist.is_public }">
          {{ playlist.is_public ? '公开' : '私密' }}
        </span>
      </div>

      <!-- 描述（如果有） -->
      <p v-if="playlist.description" class="playlist-description" :title="playlist.description">
        {{ playlist.description }}
      </p>
    </div>

    <!-- 装饰性边框 -->
    <div class="card-border"></div>
  </div>
</template>

<script setup lang="ts">
import type { PropType } from 'vue';
import type { ModelsPlaylistInfoDTO } from '@/api-client';
import { VideoPlay, Edit, Delete } from '@element-plus/icons-vue';

// --- Props ---
const props = defineProps({
  playlist: {
    type: Object as PropType<ModelsPlaylistInfoDTO>,
    required: true,
  },
  isOwner: {
    type: Boolean,
    default: false,
  },
});

// --- Emits ---
const emit = defineEmits<{
  (e: 'view', id: number): void;
  (e: 'play', id: number): void;
  (e: 'edit', playlist: ModelsPlaylistInfoDTO): void;
  (e: 'delete', playlist: ModelsPlaylistInfoDTO): void;
}>();
</script>

<style scoped>
.playlist-card {
  background: rgba(18, 18, 18, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  position: relative;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.05);
  animation: fadeInScale 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.playlist-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.6);
  border-color: rgba(255, 255, 255, 0.1);
  background: rgba(18, 18, 18, 0.98);
}

/* 装饰性边框 */
.card-border {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(29, 185, 84, 0.5),
    transparent
  );
  opacity: 0;
  transition: opacity 0.3s ease;
}

.playlist-card:hover .card-border {
  opacity: 1;
}

/* 封面区域 */
.playlist-cover {
  position: relative;
  width: 100%;
  padding-top: 100%; /* 1:1 宽高比 */
  overflow: hidden;
  background: #000;
}

.cover-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.playlist-card:hover .cover-image {
  transform: scale(1.05);
}

/* 渐变遮罩 */
.cover-gradient {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 50%;
  background: linear-gradient(
    to top,
    rgba(0, 0, 0, 0.8),
    transparent
  );
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}

.playlist-card:hover .cover-gradient {
  opacity: 1;
}

/* 悬浮遮罩层 */
.playlist-hover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  opacity: 0;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.playlist-card:hover .playlist-hover-overlay {
  opacity: 1;
}

/* 主播放按钮 */
.main-play-button {
  width: 56px;
  height: 56px;
  background: #1db954;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #000;
  font-size: 24px;
  cursor: pointer;
  transition: all 0.2s ease;
  transform: scale(0.8) translateY(-10px);
  opacity: 0;
  box-shadow: 0 8px 24px rgba(29, 185, 84, 0.4);
  margin-top: auto;
  margin-bottom: auto;
}

.playlist-card:hover .main-play-button {
  transform: scale(1) translateY(0);
  opacity: 1;
}

.main-play-button:hover {
  transform: scale(1.1);
  background: #1ed760;
  box-shadow: 0 12px 32px rgba(29, 185, 84, 0.6);
}

.main-play-button:active {
  transform: scale(0.95);
}

/* 操作按钮组 */
.action-buttons {
  display: flex;
  gap: 8px;
  align-self: flex-end;
  transform: translateY(10px);
  opacity: 0;
  transition: all 0.3s ease;
  transition-delay: 0.1s;
}

.playlist-card:hover .action-buttons {
  transform: translateY(0);
  opacity: 1;
}

.action-btn {
  width: 32px;
  height: 32px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
  transform: scale(1.1);
}

.edit-btn:hover {
  background: rgba(29, 185, 84, 0.8);
  border-color: #1db954;
  color: #000;
}

.delete-btn:hover {
  background: rgba(242, 60, 60, 0.8);
  border-color: #f23c3c;
  color: #fff;
}

/* 歌单信息区域 */
.playlist-info {
  padding: 16px;
  color: #fff;
  position: relative;
}

.playlist-name {
  font-size: 16px;
  font-weight: 700;
  color: #fff;
  margin: 0 0 8px 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  transition: color 0.2s ease;
  line-height: 1.3;
}

.playlist-card:hover .playlist-name {
  color: #1db954;
}

.playlist-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
  flex-wrap: wrap;
}

.playlist-count {
  font-size: 13px;
  color: #b3b3b3;
  font-weight: 400;
}

.playlist-visibility {
  font-size: 12px;
  color: #999;
  padding: 2px 8px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  font-weight: 500;
}

.playlist-visibility.public {
  color: #1db954;
  background: rgba(29, 185, 84, 0.2);
}

.playlist-description {
  font-size: 12px;
  color: #b3b3b3;
  line-height: 1.4;
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  opacity: 0.8;
}

/* 卡片动画 */
@keyframes fadeInScale {
  from {
    opacity: 0;
    transform: scale(0.9) translateY(20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .playlist-info {
    padding: 14px;
  }

  .playlist-name {
    font-size: 15px;
  }

  .main-play-button {
    width: 48px;
    height: 48px;
    font-size: 20px;
  }
}

@media screen and (max-width: 992px) {
  .playlist-info {
    padding: 12px;
  }

  .playlist-name {
    font-size: 14px;
  }

  .playlist-meta {
    gap: 8px;
  }

  .main-play-button {
    width: 44px;
    height: 44px;
    font-size: 18px;
  }

  .action-btn {
    width: 28px;
    height: 28px;
  }
}

@media screen and (max-width: 768px) {
  .playlist-card {
    border-radius: 8px;
  }

  .playlist-info {
    padding: 10px;
  }

  .playlist-name {
    font-size: 13px;
    margin-bottom: 6px;
  }

  .playlist-count,
  .playlist-description {
    font-size: 11px;
  }

  .playlist-visibility {
    font-size: 10px;
    padding: 1px 6px;
  }

  .playlist-hover-overlay {
    padding: 12px;
  }

  .main-play-button {
    width: 40px;
    height: 40px;
    font-size: 16px;
  }

  .action-buttons {
    gap: 6px;
  }

  .action-btn {
    width: 24px;
    height: 24px;
  }
}

@media screen and (max-width: 480px) {
  .playlist-info {
    padding: 8px;
  }

  .playlist-name {
    font-size: 12px;
  }

  .playlist-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }

  .main-play-button {
    width: 36px;
    height: 36px;
    font-size: 14px;
  }
}

/* 工具提示样式覆盖 */
:deep(.el-tooltip__popper) {
  background: #282828 !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  color: #fff !important;
  font-size: 12px !important;
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .playlist-card {
    border: 2px solid rgba(255, 255, 255, 0.2);
  }

  .playlist-name {
    text-shadow: 0 0 1px rgba(0, 0, 0, 0.8);
  }
}

/* 减少动画的用户偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .playlist-card {
    animation: none;
    transition: none !important;
  }

  .playlist-card:hover {
    transform: none !important;
  }

  .cover-image,
  .main-play-button,
  .action-buttons,
  .action-btn {
    transition: none !important;
  }

  .playlist-card:hover .cover-image {
    transform: none;
  }

  .main-play-button:hover,
  .action-btn:hover {
    transform: none !important;
  }
}

/* 深色主题下的文本选择效果 */
::selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}

::-moz-selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}
</style>