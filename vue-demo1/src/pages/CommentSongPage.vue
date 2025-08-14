<template>
  <CommonLayout :page-name="song?.name || '评论详情'" :main-content-scrollable="false">
    <template #main>
      <DetailLayout :loading="isLoading">
        <template #info>
          <EntityInfo
            :image-url="song?.cover"
            :title="song?.name"
            :meta-items="songMetaItems"
            :classification="classification"
          />
        </template>
        <template #list>
          <CommentSection v-if="songId" :song-id="songId" />
        </template>
      </DetailLayout>
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { songApi } from '@/api';
import type { ModelsSongDetailDTO } from '@/api-client';

// 导入布局和新创建的组件
import CommonLayout from '@/layouts/CommonLayout.vue';
import DetailLayout from '@/layouts/DetailLayout.vue';
import EntityInfo from '@/components/EntityInfo.vue';
import CommentSection from '@/components/CommentSection.vue';

const route = useRoute();
const isLoading = ref(true);
const song = ref<ModelsSongDetailDTO | null>(null);
const songId = Number(route.params.id);
const classification = ref<string>('歌曲');
// 计算属性，用于给 EntityInfo 传递元数据
const songMetaItems = computed(() => {
  if (!song.value) return [];
  const items = [];
  if (song.value.artist_name) {
    items.push(`艺人: ${song.value.artist_name}`);
  }
  if (song.value.album_name) {
    items.push(`专辑: ${song.value.album_name}`);
  }
  return items;
});

// 获取歌曲基本信息
const fetchSongInfo = async () => {
  if (!songId) return;
  isLoading.value = true;
  try {
    const response = await songApi.songsIdGet(songId);
    song.value = response.data;
  } catch (error) {
    console.error('获取歌曲信息失败:', error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchSongInfo();
});
</script>

<style scoped>
/* 这个页面组件本身不需要特定的样式 */
</style>