<template>
  <div v-if="playlists && playlists.length > 0" class="playlist-grid">
    <PlaylistCard
      v-for="playlist in playlists"
      :key="playlist.id"
      :playlist="playlist"
      :is-owner="isOwner"
      @view="emit('view', $event)"
      @play="emit('play', $event)"
      @edit="emit('edit', $event)"
      @delete="emit('delete', $event)"
    />
  </div>

  <!-- 空状态 - 使用深色主题样式 -->
  <!-- <div v-else class="empty-grid">
    <div class="empty-content">
      <el-icon class="empty-icon">
        <FolderOpened />
      </el-icon>
      <h3 class="empty-title">还没有创建歌单</h3>
      <p class="empty-description">开始创建你的第一个歌单，收集喜欢的音乐</p>
    </div>
  </div> -->
</template>

<script setup lang="ts">
import type { PropType } from 'vue';
import type { ModelsPlaylistInfoDTO } from '@/api-client';
import PlaylistCard from './PlaylistCard.vue';
import { FolderOpened } from '@element-plus/icons-vue';

// --- Props ---
defineProps({
  playlists: {
    type: Array as PropType<ModelsPlaylistInfoDTO[]>,
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
/* 歌单网格布局 */
.playlist-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 24px;
  padding: 0;
  animation: fadeInUp 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

/* 空状态样式 */
.empty-grid {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  width: 100%;
}

.empty-content {
  text-align: center;
  max-width: 400px;
  padding: 48px 32px;
  background: rgba(255, 255, 255, 0.05);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease;
}

.empty-content:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(255, 255, 255, 0.15);
  transform: translateY(-2px);
}

.empty-icon {
  font-size: 72px;
  color: #5e5e5e;
  margin-bottom: 24px;
  display: block;
  transition: color 0.3s ease;
}

.empty-content:hover .empty-icon {
  color: #1db954;
}

.empty-title {
  font-size: 24px;
  color: #fff;
  margin: 0 0 12px 0;
  font-weight: 700;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
}

.empty-description {
  font-size: 16px;
  color: #b3b3b3;
  margin: 0;
  line-height: 1.6;
}

/* 网格动画效果 */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 响应式网格布局 */
@media screen and (max-width: 1400px) {
  .playlist-grid {
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 20px;
  }
}

@media screen and (max-width: 1200px) {
  .playlist-grid {
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 18px;
  }
}

@media screen and (max-width: 992px) {
  .playlist-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
    gap: 16px;
  }

  .empty-content {
    padding: 40px 24px;
  }

  .empty-icon {
    font-size: 64px;
  }

  .empty-title {
    font-size: 20px;
  }

  .empty-description {
    font-size: 14px;
  }
}

@media screen and (max-width: 768px) {
  .playlist-grid {
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
    gap: 12px;
  }

  .empty-grid {
    min-height: 300px;
  }

  .empty-content {
    padding: 32px 20px;
    border-radius: 12px;
  }

  .empty-icon {
    font-size: 56px;
    margin-bottom: 20px;
  }

  .empty-title {
    font-size: 18px;
    margin-bottom: 8px;
  }

  .empty-description {
    font-size: 13px;
  }
}

@media screen and (max-width: 480px) {
  .playlist-grid {
    grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
    gap: 8px;
  }

  .empty-content {
    padding: 24px 16px;
  }

  .empty-icon {
    font-size: 48px;
    margin-bottom: 16px;
  }

  .empty-title {
    font-size: 16px;
  }

  .empty-description {
    font-size: 12px;
  }
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .empty-content {
    border: 2px solid rgba(255, 255, 255, 0.3);
  }

  .empty-title {
    text-shadow: 0 0 1px rgba(0, 0, 0, 0.8);
  }
}

/* 减少动画的用户偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .playlist-grid {
    animation: none;
  }

  .empty-content {
    transition: none !important;
  }

  .empty-content:hover {
    transform: none !important;
  }

  .empty-icon {
    transition: none !important;
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