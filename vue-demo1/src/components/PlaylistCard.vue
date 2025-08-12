<template>
  <div class="playlist-card" @click="emit('view', playlist.id!)">
    <div class="playlist-cover">
      <el-image :src="playlist.cover || '/assets/default-cover.jpg'" fit="cover" class="cover-image" lazy />
      <div class="playlist-hover-overlay">
        <div class="hover-buttons">
          <el-tooltip content="播放全部" placement="top">
            <el-button circle :icon="VideoPlay" @click.stop="emit('play', playlist.id!)" />
          </el-tooltip>
          <template v-if="isOwner">
            <el-tooltip content="编辑歌单" placement="top">
              <el-button circle :icon="Edit" @click.stop="emit('edit', playlist)" />
            </el-tooltip>
            <el-tooltip content="删除歌单" placement="top">
              <el-button circle :icon="Delete" type="danger" plain @click.stop="emit('delete', playlist)" />
            </el-tooltip>
          </template>
        </div>
      </div>
    </div>
    <div class="playlist-info">
      <h3 class="playlist-name" :title="playlist.name">{{ playlist.name }}</h3>

      <div class="playlist-info-row">
        <p class="playlist-count">{{ playlist.song_count || 0 }} 首歌曲</p>
        <p class="is-public">{{ playlist.is_public ? '公开' : '私密' }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { PropType } from 'vue';
import type { ModelsPlaylistInfoDTO } from '@/api-client';
import { VideoPlay, Edit, Delete } from '@element-plus/icons-vue';

// --- Props ---
// 定义组件接收的属性
const props = defineProps({
  // 歌单数据对象，使用 PropType 来提供更精确的类型定义
  playlist: {
    type: Object as PropType<ModelsPlaylistInfoDTO>,
    required: true,
  },
  // 是否为当前登录用户的歌单，用于控制编辑/删除按钮的显示
  isOwner: {
    type: Boolean,
    default: false,
  },
});

// --- Emits ---
// 定义组件可以发出的事件
const emit = defineEmits<{
  (e: 'view', id: number): void; // 查看详情
  (e: 'play', id: number): void; // 播放歌单
  (e: 'edit', playlist: ModelsPlaylistInfoDTO): void; // 编辑歌单
  (e: 'delete', playlist: ModelsPlaylistInfoDTO): void; // 删除歌单
}>();
</script>

<style scoped>
.playlist-card {
  background: var(--el-bg-color);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease-in-out;
  cursor: pointer;
  position: relative;
}

.playlist-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.playlist-cover {
  position: relative;
  width: 100%;
  padding-top: 100%; /* 保证1:1的宽高比 */
}

.cover-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.playlist-hover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  opacity: 0;
  transition: opacity 0.3s;
}

.playlist-card:hover .playlist-hover-overlay {
  opacity: 1;
}

.hover-buttons {
  display: flex;
  gap: 12px;
}

.hover-buttons .el-button {
  transform: scale(0.9);
  transition: all 0.2s ease-out;
}

.playlist-card:hover .hover-buttons .el-button {
  transform: scale(1);
}

.hover-buttons .el-button:nth-child(2) {
  transition-delay: 0.05s;
}

.hover-buttons .el-button:nth-child(3) {
  transition-delay: 0.1s;
}


.playlist-info {
  padding: 12px 16px;
}

.playlist-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.playlist-info-row {
  display: flex; /* 启用 Flexbox 布局 */
  align-items: center; /* 垂直居中对齐，如果高度不一致 */
  gap: 20px; /* 设置子元素之间的间隔，例如 20px */

}

.playlist-count {
  font-size: 14px;
  color: var(--el-text-color-secondary);
  margin-top: 4px;
}

.is-public {
  font-size: 14px;
  color: var(--el-text-color-secondary);
  margin-top: 4px;
}
</style>