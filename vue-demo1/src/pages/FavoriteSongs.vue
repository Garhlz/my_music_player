<template>
  <CommonLayout :page-name="pageTitle">
    <template #main>
      <SongList
        v-if="targetUserId"
        :key="targetUserId"
        :fetch-function="fetchLikedSongs"
        :sort-options="likedSortOptions"
        v-model:sortBy="sortBy"
        search-placeholder="在喜欢的音乐中搜索"
        :empty-text="emptyText"
      />
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'; // 引入 watch
import { useRoute } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { likeApi, userApi } from '@/api';
import type { MeLikedSongsGetSortByEnum } from '@/api-client';

// 组件导入
import CommonLayout from '@/layouts/CommonLayout.vue';
import SongList from '@/components/SongList.vue';

const route = useRoute();
const userStore = useUserStore();
const { userId: currentUserId } = storeToRefs(userStore);

const sortBy = ref<MeLikedSongsGetSortByEnum>('latest');
const targetUserId = ref<number | null>(null);
const targetUserName = ref<string>(''); // 用于存储目标用户的名字

const pageTitle = computed(() => {
  if (!targetUserId.value) return '喜欢的音乐';
  return isOwner.value ? '我喜欢的音乐' : `${targetUserName.value} 喜欢的音乐`;
});

const isOwner = computed(() => {
  if (!targetUserId.value || !currentUserId.value) return false;
  return targetUserId.value === currentUserId.value;
});

const emptyText = computed(() => {
  return isOwner.value ? '你还没有喜欢的音乐，快去发现吧！' : 'TA 还没有公开喜欢的音乐';
});

const likedSortOptions = [
  { label: '最近喜欢', value: 'latest' },
  { label: '最早喜欢', value: 'oldest' },
  { label: '歌曲名', value: 'name' },
  { label: '时长', value: 'duration' },
];

const fetchLikedSongs = (page: number, pageSize: number, search: string, sortBy: string) => {
  if (isOwner.value) {
    return likeApi.meLikedSongsGet(page, pageSize, search, sortBy as MeLikedSongsGetSortByEnum);
  } else if (targetUserId.value) {
    // [THE FIX] 确保 targetUserId.value 存在
    return likeApi.usersIdLikedSongsGet(targetUserId.value, page, pageSize, search, sortBy as MeLikedSongsGetSortByEnum);
  }
  // 如果没有有效的 targetUserId，返回一个空的 Promise 来避免错误
  return Promise.resolve({ data: { list: [], total: 0 } });
};

// [THE FIX] 用一个函数来处理所有与 ID 相关的逻辑
const handleIdChange = async (id: number) => {
  targetUserId.value = id;
  // 如果不是查看自己的页面，则获取目标用户的名字来更新标题
  if (!isOwner.value) {
    try {
      const res = await userApi.usersIdNameGet(id);
      targetUserName.value = res.data.name || res.data.username || '用户';
    } catch {
      targetUserName.value = '某用户';
    }
  }
};

// [THE FIX] 使用 watch 监听路由参数的变化
watch(
  () => route.params.id,
  (newId) => {
    if (newId) {
      handleIdChange(Number(newId));
    }
  },
  { immediate: true }, // immediate: true 确保组件首次加载时也会执行
);

</script>
<style scoped>
/* 页面级样式可以保留，组件级样式已移除 */
</style>