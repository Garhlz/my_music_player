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
  <el-empty v-else description="还没有创建歌单" />
</template>

<script setup lang="ts">
import type { PropType } from 'vue';
import type { ModelsPlaylistInfoDTO } from '@/api-client';
import PlaylistCard from './PlaylistCard.vue';

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
// 将从子组件 PlaylistCard 收到的事件，原封不动地再向上抛出
const emit = defineEmits<{
  (e: 'view', id: number): void;
  (e: 'play', id: number): void;
  (e: 'edit', playlist: ModelsPlaylistInfoDTO): void;
  (e: 'delete', playlist: ModelsPlaylistInfoDTO): void;
}>();
</script>

<style scoped>
.playlist-grid {
  display: grid;
  /* 响应式网格布局，每列最小200px，最大1fr，自动填充 */
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 24px;
}
</style>