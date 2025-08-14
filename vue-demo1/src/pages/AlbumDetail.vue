<template>
  <CommonLayout :page-name="album?.name || '专辑详情'" :main-content-scrollable="false">
    <template #main>
      <DetailLayout :loading="isLoading">
        <template #info>
          <EntityInfo
            :image-url="album?.cover"
            :title="album?.name"
            :description="album?.description"
            :meta-items="albumMetaItems"
            :classification="classification"
          >
            <template #actions>
            </template>
          </EntityInfo>
        </template>

        <template #list>
          <SongList
            :fetch-function="fetchAlbumSongs"
            :sort-options="albumSortOptions"
            v-model:sortBy="sortBy"
            search-placeholder="在歌曲中搜索"
            empty-text="这张专辑还没有歌曲"
          />
        </template>
      </DetailLayout>
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { albumApi } from '@/api'; // 确保你的 API 实例是这样导入的
import type { ModelsAlbum, AlbumsIdSongsGetSortByEnum } from '@/api-client';

// 导入我们的新组件和已有的SongList
import CommonLayout from '@/layouts/CommonLayout.vue';
import DetailLayout from '@/layouts/DetailLayout.vue';
import EntityInfo from '@/components/EntityInfo.vue';
import SongList from '@/components/SongList.vue';

const route = useRoute();
const isLoading = ref(true);
const album = ref<ModelsAlbum>(null);
const albumId = Number(route.params.id);

const songNumber = ref<number>(0);
// --- 状态管理 ---
const sortBy = ref<AlbumsIdSongsGetSortByEnum>('latest');
const albumSortOptions = [
  { label: '最新发布', value: 'latest' },
  { label: '播放最多', value: 'play_count' },
  { label: '点赞最多', value: 'like_count' },
  { label: '最早发布', value: 'oldest' },
];
const classification = ref<string>('专辑');
// --- 计算属性，用于给 EntityInfo 传递数据 ---
const albumMetaItems = computed(() => {
  if (!album.value) return [];
  const items = [];
  if (songNumber.value != 0) {
    items.push(`${songNumber.value} 首歌曲`);
  }
  return items;
});

// --- 数据获取 ---

// 1. 获取专辑基本信息
const fetchAlbumInfo = async () => {
  if (!albumId) return;
  isLoading.value = true;
  try {
    const response = await albumApi.albumsIdGet(albumId);
    album.value = response.data;
  } catch (error) {
    console.error('获取专辑信息失败:', error);
    // 可以在这里处理错误，比如跳转到404页面
  } finally {
    isLoading.value = false;
  }
};

// 2. 准备一个函数，传递给 SongList 组件，让它自己去获取歌曲列表
const fetchAlbumSongs = async (page: number, pageSize: number, search: string, sortBy: string) => {
  if (!albumId) throw new Error('专辑ID无效');
  const response = await albumApi.albumsIdSongsGet(albumId, page, pageSize, search, sortBy as AlbumsIdSongsGetSortByEnum);
  // SongList 需要一个 { total: number, list: any[] } 格式的对象
  songNumber.value = response.data.total;
  return response;
};

// --- 生命周期钩子 ---
onMounted(() => {
  fetchAlbumInfo();
});
</script>

<style scoped>
</style>