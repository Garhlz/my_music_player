<template>
  <CommonLayout :page-name="playlist?.name || '歌单详情'" :main-content-scrollable="false">
    <template #main>
      <DetailLayout :loading="isLoading">
        <template #info>
          <EntityInfo
            :image-url="playlist?.cover"
            :title="playlist?.name"
            :description="playlist?.description"
            :meta-items="playlistMetaItems"
          >
            <template #actions>
              <el-button v-if="isOwner" type="primary" @click="openEditDialog" round>
                编辑歌单信息
              </el-button>
            </template>
          </EntityInfo>
        </template>

        <template #list>
          <SongList
            :fetch-function="fetchPlaylistSongs"
            :sort-options="playlistSortOptions"
            v-model:sortBy="sortBy"
            search-placeholder="在歌曲中搜索"
            empty-text="当前歌单还没有歌曲"
          />
        </template>
      </DetailLayout>
    </template>
  </CommonLayout>
  <PlaylistFormDialog
    v-if="isOwner"
    v-model:visible="dialogVisible"
    mode="edit"
    :initial-data="playlist"
    @success="handleFormSuccess"
  />
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { playlistApi } from '@/api'; // 确保你的 API 实例是这样导入的
import type { ModelsPlaylistInfoDTO, PlaylistsPlaylistIdSongsGetSortByEnum } from '@/api-client';

// 导入我们的新组件和已有的SongList
import CommonLayout from '@/layouts/CommonLayout.vue';
import DetailLayout from '@/layouts/DetailLayout.vue';
import EntityInfo from '@/components/EntityInfo.vue';
import SongList from '@/components/SongList.vue';
import PlaylistFormDialog from '@/components/PlaylistFormDialog.vue';
import { storeToRefs } from 'pinia';

const route = useRoute();
const userStore = useUserStore();
const { userId: currentUserId } = storeToRefs(userStore);


const targetUserId = ref<Number | null>(0);

const isLoading = ref(true);
const playlist = ref<ModelsPlaylistInfoDTO>(null);
const playlistId = Number(route.params.id);

const dialogVisible = ref(false); // 创建一个 ref 来控制弹窗的显示/隐藏

const songNumber = ref<number>(0);
// --- 状态管理 ---
const sortBy = ref<PlaylistsPlaylistIdSongsGetSortByEnum>('latest');
const playlistSortOptions = [
  { label: '最新发布', value: 'latest' },
  { label: '播放最多', value: 'play_count' },
  { label: '点赞最多', value: 'like_count' },
  { label: '最早发布', value: 'oldest' },
];

const isOwner = computed(() => {
  // 确保两个值都不是 null 或 undefined
  if (targetUserId.value == null || currentUserId.value == null) {
    return false;
  }
  return Number(targetUserId.value) === Number(currentUserId.value);
});

// 定义打开编辑弹窗的函数
const openEditDialog = () => {
  dialogVisible.value = true;
};

// 定义表单成功提交后的回调函数
const handleFormSuccess = () => {
  // 非常简单：只需要重新调用获取歌单信息的方法，页面就会自动更新
  fetchPlaylistInfo();
};

// --- 计算属性，用于给 EntityInfo 传递数据 ---
const playlistMetaItems = computed(() => {
  if (!playlist.value) return [];
  const items = [];
  if (playlist.value.song_count != 0) {
    items.push(`${playlist.value.song_count} 首歌曲`);
  }
  items.push(`${playlist.value.is_public ? '公开' : '私密'}`);
  if (playlist.value.creator_name != '') {
    items.push(`歌单创建者: ${playlist.value.creator_name}`);
  }
  return items;
});

// --- 数据获取 ---

// 1. 获取歌单基本信息
const fetchPlaylistInfo = async () => {
  if (!playlistId) return;
  isLoading.value = true;


  try {
    const response = await playlistApi.playlistsPlaylistIdGet(playlistId);
    playlist.value = response.data;
    targetUserId.value = response.data.user_id;
  } catch (error) {
    console.error('获取歌单信息失败:', error);
    // 可以在这里处理错误，比如跳转到404页面
  } finally {
    isLoading.value = false;
  }
};

// 2. 准备一个函数，传递给 SongList 组件，让它自己去获取歌曲列表
const fetchPlaylistSongs = async (page: number, pageSize: number, search: string, sortBy: string) => {
  if (!playlistId) throw new Error('歌单ID无效');
  const response = await playlistApi.playlistsPlaylistIdSongsGet(playlistId, page, pageSize, search, sortBy as PlaylistsIdSongsGetSortByEnum);
  // SongList 需要一个 { total: number, list: any[] } 格式的对象
  songNumber.value = response.data.total;
  return response;
};

// --- 生命周期钩子 ---
onMounted(() => {
  fetchPlaylistInfo();
});
</script>

<style scoped>
</style>