<template>
  <div class="song-list-container" v-loading="isLoading">
    <!-- 搜索和筛选栏 -->
    <div class="filter-section">
      <el-input
        v-model="searchQuery"
        placeholder="搜索歌曲、歌手或专辑"
        class="search-input"
        :prefix-icon="Search"
        clearable
        @clear="handleSearch"
        @keyup.enter="handleSearch"
      />
      <el-select
        :model-value="props.sortBy"
        @update:modelValue="emit('update:sortBy', $event)"
        placeholder="排序"
        class="sort-select"
        @change="handleSearch"
      >
        <el-option
          v-for="option in props.sortOptions"
          :key="option.value"
          :label="option.label"
          :value="option.value"
        />
      </el-select>
    </div>

    <!-- 空状态 -->
    <el-empty v-if="!isLoading && songs.length === 0" description="暂无歌曲" />

    <!-- 歌曲列表 -->
    <div v-else class="song-list">
      <!-- Spotify风格的简洁表头 -->
      <div class="song-header">
        <div class="col-index">#</div>
        <div class="col-title">标题</div>
        <div class="col-album">专辑</div>
        <div class="col-actions"></div>
        <div class="col-duration">
          <el-icon>
            <Clock />
          </el-icon>
        </div>
      </div>

      <!-- 歌曲项 -->
      <SongItem
        v-for="(song, index) in songs"
        :key="song.id"
        :song="song"
        :index="(page - 1) * pageSize + index + 1"
        :is-liked="likeStore.isLiked(song.id!)"
        :is-hovered="hoveredSong === song.id"
        :is-playing="playerStore.currentSong?.id === song.id"
        @update:hoveredId="hoveredSong = $event"
        @play="handlePlaySong"
        @like="handleLikeToggle"
        @addToPlaylist="openAddToPlaylistDialog"
        @showMoreOptions="openMoreOptions"
        @goToArtist="goToArtist"
        @goToAlbum="goToAlbum"
      />
    </div>

    <!-- 分页器 -->
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

    <!-- 添加到歌单对话框 -->
    <AddToPlaylistDialog
      v-model="playlistDialogVisible"
      :song-to-add="selectedSong"
      @confirm="handlePlaylistAddConfirm"
    />

    <!-- 更多选项对话框 -->
    <MoreOptionsDialog
      v-model="moreOptionsVisible"
      :song="selectedSong"
      @comment="goToComment"
      @download="downloadSong"
      @goToArtist="goToArtist"
      @goToAlbum="goToAlbum"
      @addToPlaylist="openAddToPlaylistDialog"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { Search, Clock } from '@element-plus/icons-vue';
import { likeApi } from '@/api';
import type { ModelsSongDetailDTO, ModelsPaginatedResponseDTO } from '@/api-client';
import { usePlayerStore } from '@/stores/player';
import { useUserStore } from '@/stores/user';
import { useLikeStore } from '@/stores/like';
import { storeToRefs } from 'pinia';
import SongItem from '@/components/SongItem.vue';
import AddToPlaylistDialog from '@/components/AddToPlaylistDialog.vue';
import MoreOptionsDialog from '@/components/MoreOptionsDialog.vue';

// Props定义
interface SortOption {
  label: string;
  value: string;
}

const props = defineProps<{
  fetchFunction: (page: number, pageSize: number, search: string, sortBy: string) => Promise<{
    data: ModelsPaginatedResponseDTO
  }>;
  sortOptions: SortOption[];
  sortBy: string;
  searchPlaceholder?: string;
  emptyText?: string;
  maxSongsForPlaylist?: number;
}>();

// Emits定义
const emit = defineEmits<{
  (e: 'update:sortBy', value: string): void;
}>();

// 初始化
const router = useRouter();
const playerStore = usePlayerStore();
const userStore = useUserStore();
const likeStore = useLikeStore();
const { userId, isLoggedIn } = storeToRefs(userStore);

// 响应式数据
const songs = ref<ModelsSongDetailDTO[]>([]);
const page = ref(1);
const pageSize = ref(20); // 默认改为20，更符合Spotify风格
const totalSongs = ref(0);
const isLoading = ref(false);
const searchQuery = ref('');
const hoveredSong = ref<number | null>(null);
const playlistDialogVisible = ref(false);
const moreOptionsVisible = ref(false);
const selectedSong = ref<ModelsSongDetailDTO | null>(null);

// 缓存相关
const allSongsCache = ref<ModelsSongDetailDTO[]>([]);
const isCacheValid = ref(false);

// 加载所有歌曲（用于播放列表）
const loadAllSongs = async (): Promise<ModelsSongDetailDTO[]> => {
  if (isCacheValid.value && allSongsCache.value.length > 0) {
    return allSongsCache.value;
  }

  const maxSongs = props.maxSongsForPlaylist || 500;
  const allSongs: ModelsSongDetailDTO[] = [];
  let currentPage = 1;
  const batchSize = 100;

  try {
    const loadingMessage = ElMessage({
      message: '正在准备播放列表...',
      type: 'info',
      duration: 0,
      showClose: false,
    });

    while (allSongs.length < maxSongs) {
      const response = await props.fetchFunction(
        currentPage,
        batchSize,
        searchQuery.value,
        props.sortBy,
      );

      const songs = (response.data.list as ModelsSongDetailDTO[]) || [];
      allSongs.push(...songs);

      if (songs.length < batchSize || allSongs.length >= response.data.total!) {
        break;
      }
      currentPage++;
    }

    loadingMessage.close();
    allSongsCache.value = allSongs;
    isCacheValid.value = true;
    return allSongs;
  } catch (error) {
    console.error('加载播放列表失败:', error);
    ElMessage.error('加载播放列表失败');
    return [];
  }
};

// 加载数据
const loadData = async () => {
  isLoading.value = true;
  try {
    const songsResponse = await props.fetchFunction(page.value, pageSize.value, searchQuery.value, props.sortBy);
    songs.value = (songsResponse.data.list as ModelsSongDetailDTO[]) || [];
    totalSongs.value = songsResponse.data.total || 0;
  } catch (error) {
    console.error('加载数据失败:', error);
    ElMessage.error('加载数据失败，请稍后重试');
  } finally {
    isLoading.value = false;
  }
};

// 处理播放歌曲
const handlePlaySong = async (song: ModelsSongDetailDTO) => {
  const needLoadAll = totalSongs.value > pageSize.value;

  if (needLoadAll) {
    const allSongs = await loadAllSongs();
    if (allSongs.length > 0) {
      const songIndex = allSongs.findIndex(s => s.id === song.id);
      if (songIndex !== -1) {
        playerStore.setPlaylist(allSongs, songIndex);
        ElMessage.success(`已准备 ${allSongs.length} 首歌曲`);
      }
    }
  } else {
    const songIndex = songs.value.findIndex(s => s.id === song.id);
    if (songIndex !== -1) {
      playerStore.setPlaylist(songs.value, songIndex);
    }
  }
};

// 喜欢/取消喜欢歌曲
const handleLikeToggle = (song: ModelsSongDetailDTO) => {
  if (!isLoggedIn.value) {
    ElMessage.warning('请先登录');
    return;
  }
  const songId = song.id!;
  if (likeStore.isLiked(songId)) {
    likeStore.unlikeSong(songId);
    ElMessage.success('已取消喜欢');
  } else {
    likeStore.likeSong(songId);
    ElMessage.success('已添加到喜欢的音乐');
  }
};

// 打开添加到歌单对话框
const openAddToPlaylistDialog = (song: ModelsSongDetailDTO) => {
  if (!isLoggedIn.value) {
    ElMessage.warning('请先登录');
    return;
  }
  selectedSong.value = song;
  playlistDialogVisible.value = true;
};

// 打开更多选项对话框
const openMoreOptions = (song: ModelsSongDetailDTO) => {
  selectedSong.value = song;
  moreOptionsVisible.value = true;
};

// 其他事件处理
const handlePlaylistAddConfirm = (data: { playlistId: number; songName: string }) => {
  ElMessage.success(`已添加到歌单`);
  playlistDialogVisible.value = false;
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
  }
};

const goToAlbum = (albumId?: number | null) => {
  if (albumId) {
    router.push(`/album/${albumId}`);
  }
};

const goToComment = (song: ModelsSongDetailDTO) => {
  if (song?.id) {
    router.push(`/comment/${song.id}`);
  }
};

const downloadSong = (song: ModelsSongDetailDTO) => {
  ElMessage.info(`开始下载: ${song.name}`);
  // TODO: 实现下载功能
};

// 监听搜索和排序变化
watch([searchQuery, () => props.sortBy], () => {
  isCacheValid.value = false;
  allSongsCache.value = [];
});

// 生命周期
onMounted(() => {
  loadData();
});
</script>

<style scoped>
/* Spotify风格的容器 */
.song-list-container {
  background-color: #121212;
  color: #fff;
  border-radius: 8px;
  padding: 24px;
  min-height: 100%;
}

/* 搜索和筛选栏 */
.filter-section {
  display: flex;
  gap: 16px;
  margin-bottom: 32px;
  align-items: center;
}

.search-input {
  width: 350px;
  --el-input-bg-color: #242424;
  --el-input-border-color: transparent;
  --el-input-text-color: #fff;
  --el-input-placeholder-color: #b3b3b3;
}

:deep(.search-input .el-input__wrapper) {
  background-color: #242424;
  border-radius: 500px;
  padding: 6px 14px;
  box-shadow: none;
}

:deep(.search-input .el-input__wrapper:hover) {
  background-color: #2a2a2a;
  border-color: #535353;
}

:deep(.search-input .el-input__wrapper.is-focus) {
  border-color: #fff;
  background-color: #2a2a2a;
}

.sort-select {
  width: 160px;
  --el-select-input-color: #fff;
  --el-select-border-color-hover: #535353;
}

:deep(.sort-select .el-select__wrapper) {
  background-color: #242424;
  border-radius: 4px;
  border-color: transparent;
}

:deep(.sort-select .el-select__wrapper:hover) {
  background-color: #2a2a2a;
  border-color: #535353;
}

/* 歌曲列表 */
.song-list {
  margin-top: 16px;
}

/* Spotify风格的表头 */
.song-header {
  display: grid;
  grid-template-columns: 48px 1fr minmax(120px, 1fr) 180px 60px;
  gap: 16px;
  padding: 8px 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  margin-bottom: 16px;
  color: #b3b3b3;
  font-size: 12px;
  font-weight: 400;
  letter-spacing: 0.1px;
  text-transform: uppercase;
  user-select: none;
}

.col-index {
  text-align: center;
}

.col-title {
  padding-left: 52px;
}

.col-album {
  padding-left: 4px;
}

.col-duration {
  text-align: right;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.col-duration .el-icon {
  font-size: 16px;
}

/* 分页器 */
.pagination {
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  justify-content: center;
  --el-pagination-bg-color: transparent;
  --el-pagination-text-color: #b3b3b3;
  --el-pagination-hover-color: #fff;
  --el-pagination-button-disabled-bg-color: transparent;
}

:deep(.el-pagination.is-background .btn-prev),
:deep(.el-pagination.is-background .btn-next),
:deep(.el-pagination.is-background .el-pager li) {
  background-color: transparent;
  color: #b3b3b3;
}

:deep(.el-pagination.is-background .el-pager li:hover) {
  color: #fff;
}

:deep(.el-pagination.is-background .el-pager li.is-active) {
  background-color: #1db954;
  color: #000;
  font-weight: 600;
}

/* 滚动条样式 */
.song-list-container::-webkit-scrollbar {
  width: 12px;
}

.song-list-container::-webkit-scrollbar-thumb {
  background-color: #535353;
  border-radius: 6px;
}

.song-list-container::-webkit-scrollbar-thumb:hover {
  background-color: #b3b3b3;
}

.song-list-container::-webkit-scrollbar-track {
  background-color: transparent;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .song-list-container {
    padding: 16px;
  }

  .filter-section {
    flex-direction: column;
    align-items: stretch;
  }

  .search-input,
  .sort-select {
    width: 100%;
  }

  .song-header {
    grid-template-columns: 40px 1fr 50px;
  }

  .col-album {
    display: none;
  }
}
</style>