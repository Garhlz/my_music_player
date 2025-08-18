<template>
  <CommonLayout :page-name="pageTitle" main-content-scrollable>
    <template #main>
      <div class="playlists-page-container">
        <!-- 页面头部 -->
        <div class="page-header">
          <div class="header-content">
            <h1 class="page-title">{{ pageTitle }}</h1>
            <p class="page-subtitle" v-if="!isLoading">
              {{ total }} 个歌单
            </p>
          </div>

          <!-- 操作栏 -->
          <div class="filter-section">
            <div class="left-section">
              <el-input
                v-model="searchQuery"
                placeholder="搜索歌单"
                class="search-input"
                :prefix-icon="Search"
                clearable
                @clear="handleSearch"
                @keyup.enter="handleSearch"
              />
              <el-button
                @click="handleSearch"
                class="search-btn"
                :icon="Search"
              >
                搜索
              </el-button>
            </div>

            <el-button
              v-if="isOwner"
              type="primary"
              @click="handleOpenCreateDialog"
              :icon="Plus"
              class="create-btn"
            >
              新建歌单
            </el-button>
          </div>
        </div>

        <!-- 内容区域 -->
        <div class="content-area" v-loading="isLoading">
          <PlaylistGrid
            :playlists="playlists"
            :is-owner="isOwner"
            @view="navigateToDetail"
            @edit="handleOpenEditDialog"
            @delete="handleDelete"
            @play="handleSongPlay"
          />

          <!-- 空状态 -->
          <div v-if="!isLoading && playlists.length === 0" class="empty-state">
            <div class="empty-content">
              <el-icon>
                <FolderOpened />
              </el-icon>
              <h3>{{ searchQuery ? '没有找到匹配的歌单' : (isOwner ? '创建你的第一个歌单' : 'TA还没有创建歌单') }}</h3>
              <p>{{ searchQuery ? '试试其他关键词' : (isOwner ? '开始收集你喜欢的音乐吧' : '') }}</p>
              <el-button
                v-if="isOwner && !searchQuery"
                type="primary"
                @click="handleOpenCreateDialog"
                :icon="Plus"
                class="empty-create-btn"
              >
                创建歌单
              </el-button>
            </div>
          </div>
        </div>

        <!-- 分页 -->
        <div v-if="!isLoading && total > pageSize" class="pagination-section">
          <el-pagination
            v-model:current-page="page"
            v-model:page-size="pageSize"
            :total="total"
            :page-sizes="[12, 24, 36, 48]"
            layout="total, sizes, prev, pager, next"
            background
            @size-change="handlePageChange"
            @current-change="handlePageChange"
            class="custom-pagination"
          />
        </div>
      </div>

      <PlaylistFormDialog
        v-model:visible="dialogVisible"
        :mode="dialogMode"
        :initial-data="currentEditingPlaylist"
        @success="handleFormSuccess"
      />
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus';
import { Search, Plus, FolderOpened } from '@element-plus/icons-vue';
import { usePlayerStore } from '@/stores/player';
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { playlistApi, userApi } from '@/api';
import type { ModelsPlaylistInfoDTO } from '@/api-client';

// 导入所有子组件
import CommonLayout from '@/layouts/CommonLayout.vue';
import PlaylistGrid from '@/components/PlaylistGrid.vue';
import PlaylistFormDialog from '@/components/PlaylistFormDialog.vue';

// --- Hooks ---
const route = useRoute();
const router = useRouter();
const userStore = useUserStore();
const playerStore = usePlayerStore();
const { userId: currentUserId } = storeToRefs(userStore);

// --- Page State ---

const isLoading = ref(true);
const targetUserId = ref<number | null>(null);

const pageTitle = ref('TA的歌单');


// --- Data & Pagination State ---
const playlists = ref<ModelsPlaylistInfoDTO[]>([]);
const page = ref(1);
const pageSize = ref(12);
const total = ref(0);
const searchQuery = ref('');

// --- Dialog State ---
const dialogVisible = ref(false);
const dialogMode = ref<'create' | 'edit'>('create');
const currentEditingPlaylist = ref<ModelsPlaylistInfoDTO | null>(null);

// --- Computed ---
const isOwner = computed(() => {
  if (targetUserId.value == null || currentUserId.value == null) {
    return false;
  }
  return Number(targetUserId.value) === Number(currentUserId.value);
});

// --- Methods ---

// 获取页面数据
const loadPlaylists = async () => {
  if (!targetUserId.value) return;
  isLoading.value = true;

  try {
    const apiCall = isOwner.value
      ? playlistApi.mePlaylistsGet(page.value, pageSize.value, searchQuery.value)
      : playlistApi.usersIdPlaylistsGet(targetUserId.value, page.value, pageSize.value, searchQuery.value);
    // 意思就是, 如果使用 usersIdPlaylistsGet, 返回的只能是公开的歌单
    // 已经修改了后端, 如果使用usersIdPlaylistsGet, 只能获取公开歌单
    const response = await apiCall;
    playlists.value = (response.data.list as ModelsPlaylistInfoDTO[]) || [];
    total.value = response.data.total || 0;
  } catch (error) {
    console.error('获取歌单列表失败:', error);
    ElMessage.error('获取歌单列表失败');
  } finally {
    isLoading.value = false;
  }
};

// 设置页面标题
const setPageTitleAndId = async () => {
  const id = Number(route.params.id);
  targetUserId.value = id;
  if (isOwner.value) {
    pageTitle.value = '我的歌单';
  } else {
    try {
      const res = await userApi.usersIdGet(id);
      pageTitle.value = `${res.data.name || '用户'}的歌单`;
    } catch {
      pageTitle.value = 'TA的歌单';
    }
  }
};

// 事件处理
const handleSearch = () => {
  page.value = 1;
  loadPlaylists();
};

const handlePageChange = () => {
  loadPlaylists();
};

const navigateToDetail = (id: number) => {
  router.push(`/playlist/${id}`);
};

const handleSongPlay =  async (id: number) => {
  const resp = await playlistApi.playlistsPlaylistIdSongsGet(id);
  const songs = resp.data.List;
  if (songs.length > 0) {
    playerStore.setPlaylist(songs, 0);
  } else {
    ElMessage.warning('该歌单没有歌曲');
  }
}


const handleOpenCreateDialog = () => {
  dialogMode.value = 'create';
  currentEditingPlaylist.value = null;
  dialogVisible.value = true;
};

const handleOpenEditDialog = (playlist: ModelsPlaylistInfoDTO) => {
  dialogMode.value = 'edit';
  currentEditingPlaylist.value = playlist;
  dialogVisible.value = true;
};

const handleDelete = (playlist: ModelsPlaylistInfoDTO) => {
  ElMessageBox.confirm(`确定要删除歌单 "${playlist.name}" 吗？此操作不可撤销。`, '删除确认', {
    confirmButtonText: '确定删除',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    try {
      await playlistApi.mePlaylistsPlaylistIdDelete(playlist.id!);
      ElMessage.success('删除成功');
      loadPlaylists();
    } catch (error) {
      console.error('删除失败:', error);
      ElMessage.error('删除失败');
    }
  }).catch(() => { /* 用户取消 */
  });
};

const handleFormSuccess = () => {
  page.value = 1;
  loadPlaylists();
};

// --- Watchers & Lifecycle ---
onMounted(async () => {
  await setPageTitleAndId();
  await loadPlaylists();
});

watch(() => route.params.id, async (newId) => {
  if (newId) {
    await setPageTitleAndId();
    await loadPlaylists();
  }
});
</script>

<style scoped>
.playlists-page-container {
  padding: 32px;
  height: 100%;
  display: flex;
  flex-direction: column;
  background: transparent;
  position: relative;
}

/* 页面头部 */
.page-header {
  margin-bottom: 32px;
  position: relative;
  z-index: 1;
}

.header-content {
  margin-bottom: 24px;
}

.page-title {
  font-size: 40px;
  font-weight: 900;
  color: #fff;
  margin: 0 0 8px 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
  background: linear-gradient(135deg, #fff, #b3b3b3);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.page-subtitle {
  font-size: 16px;
  color: #b3b3b3;
  margin: 0;
  font-weight: 400;
}

/* 过滤操作栏 */
.filter-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 16px;
  padding: 20px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.left-section {
  display: flex;
  gap: 12px;
  align-items: center;
  flex: 1;
  min-width: 0;
}

.search-input {
  max-width: 400px;
  flex: 1;
}

.search-input :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  transition: all 0.2s ease;
  height: 40px;
}

.search-input :deep(.el-input__wrapper:hover) {
  border-color: rgba(255, 255, 255, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.search-input :deep(.el-input__wrapper.is-focus) {
  border-color: #1db954;
  box-shadow: 0 0 0 2px rgba(29, 185, 84, 0.2);
}

.search-input :deep(.el-input__inner) {
  color: #fff;
}

.search-input :deep(.el-input__inner::placeholder) {
  color: #b3b3b3;
}

.search-input :deep(.el-input__prefix) {
  color: #b3b3b3;
}

.search-btn {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
  border-radius: 20px;
  height: 40px;
  padding: 0 20px;
}

.search-btn:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
  transform: scale(1.02);
}

.create-btn {
  background: #1db954;
  border-color: #1db954;
  color: #fff;
  border-radius: 20px;
  height: 40px;
  padding: 0 24px;
  font-weight: 600;
  transition: all 0.2s ease;
}

.create-btn:hover {
  background: #1ed760;
  border-color: #1ed760;
  transform: scale(1.02);
  box-shadow: 0 4px 12px rgba(29, 185, 84, 0.4);
}

/* 内容区域 */
.content-area {
  flex-grow: 1;
  min-height: 400px;
  position: relative;
  padding: 24px 0;
}

/* 空状态 */
.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  min-height: 400px;
}

.empty-content {
  text-align: center;
  max-width: 400px;
  padding: 40px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.empty-content .el-icon {
  font-size: 64px;
  color: #5e5e5e;
  margin-bottom: 20px;
}

.empty-content h3 {
  font-size: 20px;
  color: #fff;
  margin: 0 0 8px 0;
  font-weight: 600;
}

.empty-content p {
  font-size: 14px;
  color: #b3b3b3;
  margin: 0 0 24px 0;
  line-height: 1.5;
}

.empty-create-btn {
  background: #1db954;
  border-color: #1db954;
  color: #fff;
  border-radius: 20px;
  height: 44px;
  padding: 0 32px;
  font-weight: 600;
}

.empty-create-btn:hover {
  background: #1ed760;
  border-color: #1ed760;
  transform: scale(1.05);
}

/* 分页 */
.pagination-section {
  padding-top: 32px;
  display: flex;
  justify-content: center;
  flex-shrink: 0;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  margin-top: 24px;
}

.custom-pagination :deep(.el-pagination) {
  color: #fff;
}

.custom-pagination :deep(.el-pagination .el-pagination__total) {
  color: #b3b3b3;
}

.custom-pagination :deep(.el-pagination .el-select .el-input .el-input__inner) {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.2);
  color: #fff;
}

.custom-pagination :deep(.el-pagination .btn-prev),
.custom-pagination :deep(.el-pagination .btn-next),
.custom-pagination :deep(.el-pagination .el-pager li) {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
  margin: 0 2px;
  border-radius: 6px;
}

.custom-pagination :deep(.el-pagination .btn-prev:hover),
.custom-pagination :deep(.el-pagination .btn-next:hover),
.custom-pagination :deep(.el-pagination .el-pager li:hover) {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
}

.custom-pagination :deep(.el-pagination .el-pager li.active) {
  background: #1db954;
  border-color: #1db954;
  color: #fff;
}

/* 加载状态覆盖 */
:deep(.el-loading-mask) {
  background-color: rgba(0, 0, 0, 0.8);
  border-radius: 16px;
}

:deep(.el-loading-spinner) {
  color: #1db954;
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .playlists-page-container {
    padding: 24px;
  }

  .page-title {
    font-size: 40px;
  }
}

@media screen and (max-width: 768px) {
  .playlists-page-container {
    padding: 16px;
  }

  .page-title {
    font-size: 32px;
  }

  .filter-section {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }

  .left-section {
    flex-direction: column;
    gap: 8px;
  }

  .search-input {
    max-width: none;
  }

  .create-btn,
  .search-btn {
    width: 100%;
    justify-content: center;
  }
}

@media screen and (max-width: 480px) {
  .playlists-page-container {
    padding: 12px;
  }

  .page-title {
    font-size: 28px;
  }

  .content-area {
    min-height: 300px;
  }

  .empty-content {
    padding: 24px;
  }

  .empty-content .el-icon {
    font-size: 48px;
  }

  .empty-content h3 {
    font-size: 18px;
  }
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .page-title,
  .empty-content h3 {
    text-shadow: 0 0 1px rgba(0, 0, 0, 0.8);
  }

  .filter-section,
  .pagination-section {
    border-color: rgba(255, 255, 255, 0.3);
  }
}

/* 减少动画的用户偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .search-btn,
  .create-btn,
  .empty-create-btn {
    transition: none !important;
  }

  .search-btn:hover,
  .create-btn:hover,
  .empty-create-btn:hover {
    transform: none !important;
  }
}

/* 深色主题下的文本选择效果 */
::selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}

::-moz-selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}
</style>