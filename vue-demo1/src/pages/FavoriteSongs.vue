<template>
  <CommonLayout page-name="我喜欢的音乐">
    <template #main>
      <SongList
        :fetch-function="fetchMyLikedSongs"
        :sort-options="likedSortOptions"
        v-model:sortBy="sortBy"
        search-placeholder="在喜欢的音乐中搜索"
        empty-text="你还没有喜欢的音乐，快去发现吧！"
        :user-id="currentUserId"
      />
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { computed } from 'vue';
import { likeApi } from '@/api';
import type { MeLikedSongsGetSortByEnum } from '@/api-client';
import CommonLayout from '@/layouts/CommonLayout.vue';
import SongList from '@/components/SongList.vue';
import { useUserStore } from '@/stores/user';

const sortBy = ref('latest');

const userStore = useUserStore();
const currentUserId = userStore.userId;

// 定义此页面专属的排序选项
const likedSortOptions = [
  { label: '最近喜欢', value: 'latest' },
  { label: '最早喜欢', value: 'oldest' },
  { label: '歌曲名', value: 'name' },
  { label: '时长', value: 'duration' },
];

// 定义获取数据的函数
const fetchMyLikedSongs = (page: number, pageSize: number, search: string, sortBy: string) => {
  return likeApi.meLikedSongsGet(page, pageSize, search, sortBy as MeLikedSongsGetSortByEnum);
};
</script>

<style scoped>
/* 页面级样式可以保留，组件级样式已移除 */
</style>