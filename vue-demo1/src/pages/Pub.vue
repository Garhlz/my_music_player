<template>
  <CommonLayout page-name="公共曲库">
    <template #main>
      <SongList
        :fetch-function="fetchAllSongs"
        :sort-options="publicSortOptions"
        v-model:sortBy="sortBy"
        search-placeholder="搜索歌曲、歌手、专辑"
        empty-text="曲库中暂无歌曲"
      />
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { songApi } from '@/api';
import type { SongsGetSortByEnum } from '@/api-client';
import CommonLayout from '@/layouts/CommonLayout.vue';
import SongList from '@/components/SongList.vue';

const sortBy = ref('latest');
// 定义此页面专属的排序选项
const publicSortOptions = [
  { label: '最新发布', value: 'latest' },
  { label: '播放最多', value: 'play_count' },
  { label: '点赞最多', value: 'like_count' },
];

// 定义获取数据的函数，其签名必须与 SongList 组件 prop 的要求一致
const fetchAllSongs = (page: number, pageSize: number, search: string, sortBy: string) => {
  // 调用 songApi 的方法，并将 sortBy 转换为正确的枚举类型
  return songApi.songsGet(page, pageSize, search, sortBy as SongsGetSortByEnum);
};
</script>

<style scoped>
</style>