<template>
  <CommonLayout :page-name="pageTitle" main-content-scrollable>
    <template #main>
      <div class="playlists-page-container">
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
          </div>
          <el-button v-if="isOwner" type="primary" @click="handleOpenCreateDialog" :icon="Plus">
            新建歌单
          </el-button>
        </div>

        <div class="content-area" v-loading="isLoading">
          <PlaylistGrid
            :playlists="playlists"
            :is-owner="isOwner"
            @view="navigateToDetail"
            @edit="handleOpenEditDialog"
            @delete="handleDelete"
          />
        </div>

        <div v-if="!isLoading && total > 0" class="pagination-section">
          <el-pagination
            v-model:current-page="page"
            v-model:page-size="pageSize"
            :total="total"
            :page-sizes="[12, 24, 36, 48]"
            layout="total, sizes, prev, pager, next"
            background
            @size-change="handlePageChange"
            @current-change="handlePageChange"
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
import { Search, Plus } from '@element-plus/icons-vue';
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { playlistApi, userApi } from '@/api'; // 引入 userApi
import type { ModelsPlaylistInfoDTO } from '@/api-client';

// 导入所有子组件
import CommonLayout from '@/layouts/CommonLayout.vue';
import PlaylistGrid from '@/components/PlaylistGrid.vue';
import PlaylistFormDialog from '@/components/PlaylistFormDialog.vue';

// --- Hooks ---
const route = useRoute();
const router = useRouter();
const userStore = useUserStore();
const { userId: currentUserId } = storeToRefs(userStore);

// --- Page State ---
const pageTitle = ref('TA的歌单');
const isLoading = ref(true);
const targetUserId = ref<number | null>(null);

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
// 在比较时，为防止任何意外的类型不匹配，对两边都进行 Number() 转换
// 这是最健壮的做法，可以免疫所有 string/number 混淆问题
const isOwner = computed(() => {
  // 确保两个值都不是 null 或 undefined
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
      loadPlaylists(); // 刷新列表
    } catch (error) {
      console.error('删除失败:', error);
      ElMessage.error('删除失败');
    }
  }).catch(() => { /* 用户取消 */
  });
};

const handleFormSuccess = () => {
  // 当弹窗操作成功后，重新加载第一页数据
  page.value = 1;
  loadPlaylists();
};

// --- Watchers & Lifecycle ---
onMounted(async () => {
  await setPageTitleAndId();
  await loadPlaylists();
});

// 监听路由参数变化，以便在 /playlists/1 -> /playlists/2 之间导航时能刷新数据
watch(() => route.params.id, async (newId) => {
  if (newId) {
    await setPageTitleAndId();
    await loadPlaylists();
  }
});

</script>

<style scoped>
.playlists-page-container {
  padding: 0 24px 24px 24px;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.filter-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px 0;
  flex-shrink: 0;
}

.left-section {
  display: flex;
  gap: 16px;
  align-items: center;
}

.search-input {
  width: 300px;
}

.content-area {
  flex-grow: 1;
  min-height: 0; /* 关键，让 flex item 在空间不足时可以收缩 */
}

.pagination-section {
  padding-top: 24px;
  display: flex;
  justify-content: center;
  flex-shrink: 0;
}
</style>