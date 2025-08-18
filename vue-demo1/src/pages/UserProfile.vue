<template>
  <CommonLayout :page-name="isOwner ? '我的主页' : userInfo?.name" :main-content-scrollable="true">
    <template #main>
      <div class="profile-page" v-loading="isLoading">
        <div class="profile-banner">
          <el-avatar :size="120" :src="userInfo?.avatar" class="profile-avatar" />
          <div class="profile-meta">
            <span class="profile-type">个人主页</span>
            <h1 class="profile-name">{{ userInfo?.name || userInfo?.username }}</h1>
            <p class="profile-bio">{{ userInfo?.bio || '一个热爱音乐的人' }}</p>

            <div class="profile-additional-info">
              <span v-if="userInfo?.location">
                <el-icon><Location /></el-icon> {{ userInfo.location }}
              </span>
              <span>
                <el-icon><User /></el-icon> {{ formatGender(userInfo?.gender) }}
              </span>
            </div>

            <div class="profile-stats">
              <span><b>{{ userInfo?.followers || 0 }}</b> 粉丝数 </span>
              <span><b>{{ userInfo?.followings || 0 }}</b> 关注数</span>
            </div>
          </div>


          <div class="actions-bar">
            <div class="main-actions">
              <el-button v-if="isOwner" type="primary" round @click="showEditDialog = true">
                <el-icon>
                  <Edit />
                </el-icon>
                编辑个人资料
              </el-button>
              <el-button v-else type="primary" round @click="followCurrentUser">
                <!--TODO 需要根据关注情况确定显示关注还是取关... -->
                <el-icon>
                  <Plus />
                </el-icon>
                关注
              </el-button>
            </div>
            <div class="quick-links">
              <el-button text @click="navigate('love')">
                <el-icon>
                  <Star />
                </el-icon>
                喜欢的音乐
              </el-button>
              <el-button text @click="navigate('playlist-grid')">
                <el-icon>
                  <Menu />
                </el-icon>
                公开的歌单
              </el-button>
            </div>
          </div>
        </div>


        <div class="profile-content">
          <!--          <el-empty description="这里空空如也..." />-->
          <div class="playlists-page-container">
            <!-- 页面头部 -->
            <div class="page-header">
              <!--              <div class="header-content">-->
              <!--                <h1 class="page-title">{{ playlistPageTitle }}</h1>-->
              <!--                <p class="page-subtitle" v-if="!isLoading">-->
              <!--                  {{ total }} 个歌单-->
              <!--                </p>-->
              <!--              </div>-->

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
              />

              <!-- 空状态 -->
              <div v-if="!isLoading && playlists.length === 0" class="empty-state">
                <div class="empty-content">
                  <el-icon>
                    <FolderOpened />
                  </el-icon>
                  <h3>{{ searchQuery ? '没有找到匹配的歌单' : (isOwner ? '创建你的第一个歌单' : 'TA还没有创建歌单')
                    }}</h3>
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
        </div>
      </div>

      <EditProfileDialog
        v-if="isOwner"
        v-model:visible="showEditDialog"
        :initial-data="editForm"
        @success="handleUpdateProfile"
      />
    </template>
  </CommonLayout>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router'; // 引入 useRouter
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { userApi, playlistApi } from '@/api';
import { ElMessage } from 'element-plus';
import type { ModelsUserProfile, ModelsUpdateUserRequest, ModelsPlaylistInfoDTO } from '@/api-client';
import { Edit, Plus, Star, Menu, Location, User, Search, FolderOpened } from '@element-plus/icons-vue';
import CommonLayout from '@/layouts/CommonLayout.vue';
import EditProfileDialog from '@/components/EditProfileDialog.vue';
import PlaylistGrid from '@/components/PlaylistGrid.vue';
import PlaylistFormDialog from '@/components/PlaylistFormDialog.vue';

const route = useRoute();
const router = useRouter(); // 初始化 router
const userStore = useUserStore();
const { userId: currentUserId } = storeToRefs(userStore);

const isLoading = ref(true);


const userInfo = ref<ModelsUserProfile | null>(null);
const targetUserId = ref<number | null>(null);

const showEditDialog = ref(false);
const editForm = ref<ModelsUpdateUserRequest | null>(null);


const isOwner = computed(() => {
  if (targetUserId.value == null || currentUserId.value == null) {
    return false;
  }
  return Number(targetUserId.value) === Number(currentUserId.value);
});


// --------------------------------------------------------------------------------

const playlistPageTitle = ref('某人的歌单');

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
    playlistPageTitle.value = '我的歌单';
  } else {
    try {
      const res = await userApi.usersIdGet(id);
      playlistPageTitle.value = `${res.data.name || '用户'}的歌单`;
    } catch {
      playlistPageTitle.value = 'TA的歌单';
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


// --------------------------------------------------------------------------------


const formatGender = (gender?: number) => {
  const map = { 0: '保密', 1: '男', 2: '女' };
  return map[gender || 0];
};

const navigate = (path: 'love' | 'playlist-grid') => {
  router.push(`/${path}/${targetUserId.value}`);
};

const fetchUserInfo = async () => {
  isLoading.value = true;
  try {
    const res = await userApi.usersIdGet(targetUserId.value);
    userInfo.value = res.data;
    if (isOwner.value) {
      // [THE FIX] 填充所有可编辑字段
      editForm.value = {
        name: res.data.name,
        bio: res.data.bio,
        avatar: res.data.avatar,
        gender: res.data.gender,
        birthday: res.data.birthday,
        location: res.data.location,
        email: res.data.email,
        phone: res.data.phone,
      };
    }
  } catch (error) {
    ElMessage.error('获取用户信息失败');
  } finally {
    isLoading.value = false;
  }
};

const handleUpdateProfile = async (updatedData: ModelsUpdateUserRequest) => {
  try {
    await userApi.meProfilePut(updatedData);
    ElMessage.success('个人资料更新成功');
    showEditDialog.value = false;
    await fetchUserInfo();
  } catch (error) {
    ElMessage.error('更新失败，请稍后重试');
  }
};

const followCurrentUser = async() => {
  try {
    await userApi.follow(targetUserId.value);
    ElMessage.success("关注成功");
    
  }
}
onMounted(async () => {
    await setPageTitleAndId();
    await loadPlaylists();
    await fetchUserInfo();
    console.log(currentUserId.value);
    console.log(targetUserId.value);
  },
);

watch(() => route.params.id, async (newId) => {
  if (newId) {
    await setPageTitleAndId();
    await loadPlaylists();
  }
});

</script>

<style scoped>
.profile-page {
  padding: 16px;
}

/* 顶部横幅 */
.profile-banner {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 24px;
  border-radius: 8px;
  background: linear-gradient(135deg, rgba(29, 185, 84, 0.2), transparent 60%);
  justify-content: space-between; /* 确保内容和按钮两端对齐 */
  flex-wrap: wrap; /* 允许换行 */
}

.profile-avatar {
  width: 100px;
  height: 100px;
  border: 3px solid rgba(0, 0, 0, 0.2);
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.4);
  flex-shrink: 0;
}

.profile-meta {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1; /* 占用剩余空间 */
  min-width: 0; /* 防止溢出 */
}

.profile-type {
  font-size: 10px;
  font-weight: 700;
  color: #fff;
  text-transform: uppercase;
}

.profile-name {
  font-size: 36px;
  font-weight: 900;
  color: #fff;
  margin: 0;
  letter-spacing: -0.03em;
}

.profile-bio {
  font-size: 13px;
  color: #b3b3b3;
  margin: 0;
  max-width: 500px;
}

.profile-additional-info {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 13px;
  color: #b3b3b3;
  margin-top: 2px;
}

.profile-additional-info span {
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.profile-stats {
  display: flex;
  gap: 12px;
  font-size: 13px;
  color: #b3b3b3;
  margin-top: 6px;
}

.profile-stats b {
  color: #fff;
}

/* 操作栏（现位于 profile-banner 右侧） */
.actions-bar {
  display: flex;
  align-items: center;
  gap: 12px; /* 减小按钮组间距 */
  flex-wrap: wrap; /* 允许按钮换行 */
}

.main-actions .el-button {
  border-radius: 500px;
  font-weight: 700;
  border: 0;
}

.main-actions .el-button--primary {
  background-color: #1db954;
  color: #000;
  padding: 8px 16px; /* 进一步减小内边距 */
  font-size: 13px; /* 减小字体 */
}

.quick-links {
  display: flex;
  gap: 6px;
}

.quick-links .el-button {
  color: #b3b3b3;
  font-weight: 500;
  padding: 6px 12px; /* 减小内边距 */
  font-size: 12px; /* 减小字体 */
}

.quick-links .el-button:hover {
  color: #fff;
  background-color: rgba(255, 255, 255, 0.1);
}

/* 内容区 */
.profile-content {
  padding: 0 24px;
}

.playlists-page-container {
  padding: 16px;
  height: 100%;
  display: flex;
  flex-direction: column;
  background: transparent;
  position: relative;
}

/* 页面头部 */
.page-header {
  margin-bottom: 16px;
  position: relative;
  z-index: 1;
}

.header-content {
  margin-bottom: 16px;
}

.page-title {
  font-size: 18px;
  font-weight: 900;
  color: #fff;
  margin: 0 0 6px 0;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
  background: linear-gradient(135deg, #fff, #b3b3b3);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.page-subtitle {
  font-size: 12px;
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
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.left-section {
  display: flex;
  gap: 8px;
  align-items: center;
  flex: 1;
  min-width: 0;
}

.search-input {
  max-width: 300px;
  flex: 1;
}

.search-input :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  transition: all 0.2s ease;
  height: 32px;
}

.search-input :deep(.el-input__wrapper:hover) {
  border-color: rgba(255, 255, 255, 0.3);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
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
  border-radius: 16px;
  height: 32px;
  padding: 0 16px;
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
  border-radius: 16px;
  height: 32px;
  padding: 0 20px;
  font-weight: 600;
  transition: all 0.2s ease;
}

.create-btn:hover {
  background: #1ed760;
  border-color: #1ed760;
  transform: scale(1.02);
  box-shadow: 0 3px 10px rgba(29, 185, 84, 0.3);
}

/* 内容区域 */
.content-area {
  flex-grow: 1;
  min-height: 350px;
  position: relative;
  padding: 16px 0;
}

/* 空状态 */
.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  min-height: 350px;
}

.empty-content {
  text-align: center;
  max-width: 350px;
  padding: 32px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.empty-content .el-icon {
  font-size: 56px;
  color: #5e5e5e;
  margin-bottom: 16px;
}

.empty-content h3 {
  font-size: 18px;
  color: #fff;
  margin: 0 0 6px 0;
  font-weight: 600;
}

.empty-content p {
  font-size: 13px;
  color: #b3b3b3;
  margin: 0 0 20px 0;
  line-height: 1.4;
}

.empty-create-btn {
  background: #1db954;
  border-color: #1db954;
  color: #fff;
  border-radius: 16px;
  height: 36px;
  padding: 0 28px;
  font-weight: 600;
}

.empty-create-btn:hover {
  background: #1ed760;
  border-color: #1ed760;
  transform: scale(1.05);
}

/* 分页 */
.pagination-section {
  padding-top: 24px;
  display: flex;
  justify-content: center;
  flex-shrink: 0;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  margin-top: 16px;
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
  border-radius: 12px;
}

:deep(.el-loading-spinner) {
  color: #1db954;
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .playlists-page-container {
    padding: 12px;
  }

  .page-title {
    font-size: 16px;
  }

  .profile-banner {
    flex-direction: column; /* 垂直排列 */
    align-items: flex-start;
    gap: 12px;
  }

  .actions-bar {
    justify-content: flex-start; /* 按钮左对齐 */
    width: 100%;
  }
}

@media screen and (max-width: 768px) {
  .playlists-page-container {
    padding: 8px;
  }

  .page-title {
    font-size: 14px;
  }

  .filter-section {
    flex-direction: column;
    align-items: stretch;
    gap: 8px;
  }

  .left-section {
    flex-direction: column;
    gap: 6px;
  }

  .search-input {
    max-width: none;
  }

  .create-btn,
  .search-btn {
    width: 100%;
    justify-content: center;
    height: 30px;
  }

  .profile-banner {
    text-align: center;
    padding: 16px;
  }

  .profile-name {
    font-size: 28px;
  }

  .profile-stats {
    justify-content: center;
  }

  .profile-additional-info {
    justify-content: center;
  }

  .actions-bar {
    flex-direction: column;
    align-items: center;
    gap: 8px;
  }

  .main-actions .el-button,
  .quick-links .el-button {
    width: 100%; /* 按钮占满宽度 */
    justify-content: center;
  }
}

@media screen and (max-width: 480px) {
  .playlists-page-container {
    padding: 8px;
  }

  .page-title {
    font-size: 13px;
  }

  .content-area {
    min-height: 250px;
  }

  .empty-content {
    padding: 20px;
  }

  .empty-content .el-icon {
    font-size: 40px;
  }

  .empty-content h3 {
    font-size: 16px;
  }

  .profile-avatar {
    width: 80px;
    height: 80px;
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
  .empty-create-btn,
  .main-actions .el-button,
  .quick-links .el-button {
    transition: none !important;
  }

  .search-btn:hover,
  .create-btn:hover,
  .empty-create-btn:hover,
  .main-actions .el-button:hover,
  .quick-links .el-button:hover {
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