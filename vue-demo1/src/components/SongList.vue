<template>
  <div class="song-list-container" v-loading="isLoading">
    <div class="filter-section">
      <el-input
        v-model="searchQuery"
        placeholder="搜索歌曲、歌手、专辑"
        class="search-input"
        :prefix-icon="Search"
        clearable
        @clear="handleSearch"
        @keyup.enter="handleSearch"
      />
      <el-select
        :model-value="props.sortBy"
        @update:modelValue="emit('update:sortBy', $event)"
        placeholder="排序方式"
        class="sort-select"
        @change="handleSearch"
        popper-class="highest-z-index-popper"
      >
        <el-option
          v-for="option in props.sortOptions"
          :key="option.value"
          :label="option.label"
          :value="option.value"
        />
      </el-select>
    </div>

    <el-empty v-if="!isLoading && songs.length === 0" description="暂无歌曲" />

    <div v-else class="song-list">
      <div class="song-header">
        <div class="col-index">#</div>
        <div class="col-title">标题</div>
        <div class="col-duration">时长</div>
        <div class="col-artist">歌手</div>
        <div class="col-album">专辑</div>
      </div>

      <SongItem
        v-for="(song, index) in songs"
        :key="song.id"
        :song="song"
        :index="index"
        :is-liked="likedSongIds.has(song.id!)"
        :is-hovered="hoveredSong === song.id"
        :is-playing="playerStore.currentSong?.id === song.id"
        @update:hoveredId="hoveredSong = $event"
        @play="handlePlaySong"
        @select="handleSelectSong"
        @like="likeSong"
        @addToPlaylist="openAddToPlaylistDialog"
        @comment="goToComment"
        @download="downloadSong"
        @goToArtist="goToArtist"
        @goToAlbum="goToAlbum"
      />
    </div>

    <el-pagination
      v-if="!isLoading && songs.length > 0"
      v-model:current-page="page"
      v-model:page-size="pageSize"
      :total="totalSongs"
      :page-sizes="[10, 20, 30, 50]"
      layout="total, sizes, prev, pager, next, jumper"
      background
      class="pagination"
      @size-change="handlePageSizeChange"
      @current-change="handlePageChange"
    />

    <AddToPlaylistDialog
      v-model="playlistDialogVisible"
      :song-to-add="selectedSong"
      @confirm="handlePlaylistAddConfirm"
    />
  </div>
</template>
<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { Search, VideoPlay, Star, Plus, ChatDotRound, Download } from '@element-plus/icons-vue';
import { likeApi } from '@/api';
import type { ModelsSongDetailDTO, SongsGetSortByEnum, ModelsPaginatedResponseDTO } from '@/api-client';
import { usePlayerStore } from '@/stores/player';
import { useUserStore } from '@/stores/user';
import SongItem from '@/components/SongItem.vue';
import AddToPlaylistDialog from '@/components/AddToPlaylistDialog.vue';

// --- 1. 定义 Props (组件的对外契约) ---
interface SortOption {
  label: string;
  value: string;
}

const props = defineProps<{
  fetchFunction: (page: number, pageSize: number, search: string, sortBy: string) => Promise<{
    data: ModelsPaginatedResponseDTO
  }>;
  sortOptions: SortOption[];
  sortBy: string; // <--- 新增：接收父组件的 v-model 值
  searchPlaceholder?: string;
  emptyText?: string;
}>();


// 定义 Emits：明确告诉 Vue 这个组件会发出 update:sortBy 事件
const emit = defineEmits<{
  (e: 'update:sortBy', value: string): void; // <--- 新增：为 v-model 配套的事件
}>();

// 初始化
const router = useRouter();
const playerStore = usePlayerStore();
const userStore = useUserStore();

import { storeToRefs } from 'pinia'; // 引入一个重要工具 storeToRefs
// const contextMenu = ref(null);
const { userId, isLoggedIn } = storeToRefs(userStore);

const songs = ref<ModelsSongDetailDTO[]>([]);
const page = ref(1);
const pageSize = ref(10);
const totalSongs = ref(0);
const isLoading = ref(false);
const searchQuery = ref('');

const hoveredSong = ref<number | null>(null);
const likedSongIds = ref<Set<number>>(new Set());
// 似乎是展示专辑封面的窗口
// const coverDialogVisible = ref(false);
const playlistDialogVisible = ref(false);
const selectedSong = ref<ModelsSongDetailDTO | null>(null);
// const contextSong = ref<ModelsSongDetailDTO | null>(null);
// const contextMenuPos = reactive<{ x: number; y: number }>({ x: 0, y: 0 });

// 动态对话框宽度
// const dialogWidth = computed(() => {
//   return window.innerWidth <= 768 ? '90%' : '400px';
// });

// 动态网格布局
const columnLayout = computed(() => {
  return window.innerWidth <= 768
    ? '50px 3fr 100px'
    : window.innerWidth <= 1200
      ? '60px 2.5fr 100px 1.2fr 1.2fr'
      : '60px minmax(300px, 2.5fr) 180px minmax(160px, 1fr) minmax(160px, 1fr)';
});


// --- 3. 核心数据加载与交互逻辑 ---
const loadData = async (refreshLikes: boolean = true) => {
  isLoading.value = true;
  try {
    const fetchTasks = [
      props.fetchFunction(page.value, pageSize.value, searchQuery.value, props.sortBy),
    ];

    // 仅在需要时刷新喜欢列表（例如，初次加载或页面大小改变时）
    if (refreshLikes) {
      // todo 可能需要修改这里的逻辑, 以后微调
      if (isLoggedIn) {
        fetchTasks.push(likeApi.meLikedSongsGet(1, 1000));
      }
    }

    const [songsResponse, likedSongsResponse] = await Promise.all(fetchTasks);

    songs.value = (songsResponse.data.list as ModelsSongDetailDTO[]) || [];
    totalSongs.value = songsResponse.data.total || 0;

    if (likedSongsResponse) {
      const likedData = (likedSongsResponse.data.list as ModelsSongDetailDTO[]) || [];
      likedSongIds.value = new Set(likedData.map(song => song.id!));
    }
  } catch (error) {
    console.error('加载数据失败:', error);
    ElMessage.error('加载数据失败，请稍后重试');
  } finally {
    isLoading.value = false;
  }
};
// 事件处理
const handlePlaySong = (song: ModelsSongDetailDTO) => {
  playerStore.setPlaylist([song]);
  playerStore.play(0);
};

const handleSelectSong = (song: ModelsSongDetailDTO) => {
  selectedSong.value = song;
  coverDialogVisible.value = true;
};

const likeSong = async (song: ModelsSongDetailDTO) => {
  if (!isLoggedIn.value) { // 使用从 store 来的 getter
    ElMessage.warning('请先登录');
    return;
  }
  try {
    const songId = song.id!;
    if (likedSongIds.value.has(songId)) {
      await likeApi.meLikedSongsSongIdDelete(songId);
      likedSongIds.value.delete(songId);
      ElMessage.success(`已取消喜欢: ${song.name}`);
    } else {
      await likeApi.meLikedSongsSongIdPost(songId);
      likedSongIds.value.add(songId);
      ElMessage.success(`已添加到我喜欢: ${song.name}`);
    }
  } catch (error) {
    console.error('操作失败:', error);
    ElMessage.error('操作失败，请稍后重试');
  }
};

const openAddToPlaylistDialog = (song: ModelsSongDetailDTO) => {
  if (!isLoggedIn.value) {
    ElMessage.warning('请先登录');
    return;
  }
  selectedSong.value = song;
  playlistDialogVisible.value = true;
};

const handlePlaylistAddConfirm = (data: { playlistId: number; songName: string }) => {
  ElMessage.success(`已将歌曲《${data.songName}》成功添加到歌单`);
};

const handleSearch = () => {
  page.value = 1;
  loadData();
};

const handlePageChange = (newPage: number) => {
  page.value = newPage;
  loadData();
};

const handlePageSizeChange = (newSize: number) => {
  pageSize.value = newSize;
  page.value = 1;
  loadData();
};

const goToArtist = (artistId?: number | null) => {
  if (artistId) {
    router.push(`/artist/${artistId}`);
  } else {
    ElMessage.error('歌手不存在!');
  }
};

const goToAlbum = (albumId?: number | null) => {
  if (albumId) {
    router.push(`/album/${albumId}`);
  } else {
    ElMessage.error('专辑不存在!');
  }
};

const goToComment = (song: ModelsSongDetailDTO) => {
  router.push(`/comment/${song.id}`);
};

// todo 下载功能
const downloadSong = (song: ModelsSongDetailDTO) => {
  ElMessage.info(`开始下载: ${song.name}`);
};

// const handleContextMenu = ({ song, x, y }: { song: ModelsSongDetailDTO; x: number; y: number }) => {
//   contextSong.value = song;
//   contextMenuPos.x = x;
//   contextMenuPos.y = y;
//   nextTick(() => {
//     (contextMenu.value as any)?.show();
//   });
// };

// const handleContextMenuVisible = (visible: boolean) => {
//   if (!visible) {
//     contextSong.value = null;
//   }
// };

// 生命周期
onMounted(() => {
  loadData();
});

// defineExpose({
//   loadData,
// });
</script>

<style scoped>
/* 容器样式 */
.song-list-container {
  margin-right: 16px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  padding: 24px;
  //height: 100%;
  //overflow-y: auto;
  animation: fadeIn 0.5s ease-in;
}

/* 过滤区域 */
.filter-section {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
  align-items: center;
  position: relative; /* 必须设置 position 才能让 z-index 生效 */
  z-index: 2; /* 设置一个比 song-header 的 z-index: 1 更高的值 */
}

.search-input {
  width: 300px;
  transition: all 0.3s;
}

:deep(.search-input .el-input__wrapper) {
  border-radius: 20px;
}

:deep(.search-input .el-input__wrapper:hover) {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.sort-select {
  width: 160px;
}

:deep(.sort-select .el-select__wrapper) {
  border-radius: 20px;
}

:deep(.sort-select .el-select__wrapper:hover) {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 歌曲列表 */
.song-list {
  border: 1px solid #ebeef5;
  border-radius: 8px;
  overflow: hidden;
}

/* 表头样式 */
.song-header {
  position: sticky;
  top: 0;
  display: grid;
  grid-template-columns: v-bind(columnLayout);
  padding: 12px 24px;
  background: linear-gradient(180deg, #f5f7fa, #ffffff);
  border-bottom: 1px solid #ebeef5;
  font-weight: 600;
  color: #303133;
  align-items: center;
  z-index: 1;
}

.col-index {
  text-align: center;
  font-size: 14px;
}

.col-title {
  padding-left: 16px;
}

.col-duration,
.col-artist,
.col-album {
  padding: 0 12px;
}

/* 分页器 */
.pagination {
  margin-top: 24px;
  padding: 16px 0;
  display: flex;
  justify-content: center;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 -2px 12px rgba(0, 0, 0, 0.05);
}

:deep(.el-pagination.is-background .el-pager li.is-active) {
  background-color: var(--el-color-primary);
  color: #fff;
}

:deep(.el-pagination.is-background .btn-prev, .el-pagination.is-background .btn-next) {
  border-radius: 8px;
}

/* 封面预览 */
.cover-preview {
  display: flex;
  justify-content: center;
}

.cover-preview .el-image {
  width: 100%;
  max-width: 300px;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

/* 滚动条样式 */
.song-list-container::-webkit-scrollbar {
  width: 6px;
}

.song-list-container::-webkit-scrollbar-thumb {
  background-color: #dcdfe6;
  border-radius: 3px;
}

.song-list-container::-webkit-scrollbar-track {
  background-color: #f5f7fa;
}

/* 动画 */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .filter-section {
    gap: 12px;
  }

  .search-input {
    width: 250px;
  }

  .sort-select {
    width: 140px;
  }
}

@media screen and (max-width: 768px) {
  .song-list-container {
    margin-right: 0;
    padding: 16px;
  }

  .filter-section {
    flex-direction: column;
    align-items: stretch;
    gap: 8px;
  }

  .search-input,
  .sort-select {
    width: 100%;
  }

  .song-header {
    display: none; /* 移动端隐藏表头 */
  }
}

:deep(.highest-z-index-popper) {
  z-index: 9999 !important;
}
</style>