<template>
  <CommonLayout :page-name="artist?.name || '歌手详情'" :main-content-scrollable="false">
    <template #main>
      <DetailLayout :loading="isLoading">
        <template #info>
          <EntityInfo
            :image-url="artist?.avatar"
            :title="artist?.name"
            :description="artist?.description"
            :meta-items="artistMetaItems"
            :classification="classification"
          >
            <template #actions>
            </template>
          </EntityInfo>
        </template>

        <template #list>
          <SongList
            :fetch-function="fetchArtistSongs"
            :sort-options="artistSortOptions"
            v-model:sortBy="sortBy"
            search-placeholder="在歌曲中搜索"
            empty-text="这位艺术家还没有歌曲"
          />
        </template>
      </DetailLayout>
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { artistApi } from '@/api'; // 确保你的 API 实例是这样导入的
import type { ModelsArtist, ArtistsIdSongsGetSortByEnum } from '@/api-client';

// 导入我们的新组件和已有的SongList
import CommonLayout from '@/layouts/CommonLayout.vue';
import DetailLayout from '@/layouts/DetailLayout.vue';
import EntityInfo from '@/components/EntityInfo.vue';
import SongList from '@/components/SongList.vue';

const route = useRoute();
const isLoading = ref(true);
const artist = ref<ModelsArtist>(null);
const artistId = Number(route.params.id);

const songNumber = ref<number>(0);
// --- 状态管理 ---
const sortBy = ref<ArtistsIdSongsGetSortByEnum>('latest');
const artistSortOptions = [
  { label: '最新发布', value: 'latest' },
  { label: '播放最多', value: 'play_count' },
  { label: '点赞最多', value: 'like_count' },
  { label: '最早发布', value: 'oldest' },
];

const classification = ref<string>('艺术家');

// --- 计算属性，用于给 EntityInfo 传递数据 ---
const artistMetaItems = computed(() => {
  if (!artist.value) return [];
  const items = [];
  if (songNumber.value != 0) {
    items.push(`${songNumber.value} 首歌曲`);
  }
  if (artist.value.album_count) {
    // todo 需要创建专辑的视图
    items.push(`${artist.value.album_count} 张专辑`);
  }
  return items;
});

// --- 数据获取 ---

// 1. 获取歌手基本信息
const fetchArtistInfo = async () => {
  if (!artistId) return;
  isLoading.value = true;
  try {
    const response = await artistApi.artistsIdGet(artistId);
    artist.value = response.data;
  } catch (error) {
    console.error('获取歌手信息失败:', error);
    // 可以在这里处理错误，比如跳转到404页面
  } finally {
    isLoading.value = false;
  }
};

// 2. 准备一个函数，传递给 SongList 组件，让它自己去获取歌曲列表
const fetchArtistSongs = async (page: number, pageSize: number, search: string, sortBy: string) => {
  if (!artistId) throw new Error('艺术家ID无效');
  const response = await artistApi.artistsIdSongsGet(artistId, page, pageSize, search, sortBy as ArtistsIdSongsGetSortByEnum);
  // SongList 需要一个 { total: number, list: any[] } 格式的对象
  songNumber.value = response.data.total;
  return response;
};

// --- 生命周期钩子 ---
onMounted(() => {
  fetchArtistInfo();
});
</script>

<style scoped>
/* 页面级的样式可以写在这里，但现在大部分样式都封装在子组件里了，这里会很干净 */
</style>