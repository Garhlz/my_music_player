<template>
  <el-dialog
    :model-value="modelValue"
    @update:modelValue="emit('update:modelValue', $event)"
    title="添加到歌单"
    :width="dialogWidth"
    class="add-to-playlist-dialog"
    @open="fetchUserPlaylists"
    @close="resetState"
    center
    :lock-scroll="true"
  >
    <div v-loading="isLoading" class="dialog-content-wrapper">
      <div class="content-layout">
        <!-- 左侧：现有歌单列表 -->
        <div class="existing-playlists">
          <el-input
            v-model="playlistSearch"
            placeholder="搜索你的歌单"
            clearable
            class="playlist-search-input"
            prefix-icon="Search"
            @input="filterPlaylists"
          />
          <el-scrollbar height="300px">
            <div v-if="filteredPlaylists.length > 0" class="playlist-list">
              <div
                v-for="playlist in filteredPlaylists"
                :key="playlist.id"
                class="playlist-item"
                :class="{ 'playlist-item-selected': selectedPlaylistId === playlist.id }"
                @click="selectedPlaylistId = playlist.id"
              >
                <div class="playlist-info-row">
                  <el-image
                    :src="playlist.cover || '/assets/default-cover.jpg'"
                    class="playlist-cover-img"
                    fit="cover"
                    lazy
                  />
                  <div class="playlist-details">
                    <div class="playlist-name">{{ playlist.name }}</div>
                    <div class="playlist-meta">
                      <span class="playlist-count">{{ playlist.song_count || 0 }}首</span>
                      <span class="playlist-visibility" :class="{ 'public': playlist.is_public }">
                        {{ playlist.is_public ? '公开' : '私密' }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <el-empty v-else description="没有找到匹配的歌单" />
          </el-scrollbar>
        </div>

        <!-- 右侧：创建新歌单 -->
        <div class="create-playlist">
          <el-divider class="create-divider">创建新歌单</el-divider>
          <div class="create-playlist-form">
            <el-input
              v-model="newPlaylistForm.name"
              placeholder="请输入歌单名称"
              clearable
              maxlength="50"
              show-word-limit
            />
            <el-input
              v-model="newPlaylistForm.description"
              type="textarea"
              :rows="3"
              placeholder="添加歌单描述（选填）"
              class="mt-2"
              maxlength="200"
              show-word-limit
            />
            <el-switch
              v-model="newPlaylistForm.is_public"
              active-text="公开"
              inactive-text="私密"
              class="mt-2"
            />
            <el-button
              type="primary"
              @click="createNewPlaylist"
              class="mt-3 create-playlist-btn"
              :loading="isCreating"
              plain
            >
              创建歌单
            </el-button>
          </div>
        </div>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="emit('update:modelValue', false)">取消</el-button>
        <el-button
          type="primary"
          @click="confirmAddToPlaylist"
          :loading="isConfirming"
          :disabled="!selectedPlaylistId"
        >
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
import { ElMessage } from 'element-plus';
import { playlistApi } from '@/api';
import type { ModelsSongDetailDTO, ModelsPlaylistInfoDTO, ModelsCreatePlaylistRequestDTO } from '@/api-client';
import { AxiosError } from 'axios';
import { Search } from '@element-plus/icons-vue';

// 定义 Props 和 Emits
interface Props {
  modelValue: boolean;
  songToAdd: ModelsSongDetailDTO | null;
}

const props = defineProps<Props>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void;
  (e: 'confirm', data: { playlistId: number; songName: string }): void;
}>();

// 状态管理
const isLoading = ref(false);
const isCreating = ref(false);
const isConfirming = ref(false);
const userPlaylists = ref<ModelsPlaylistInfoDTO[]>([]);
const filteredPlaylists = ref<ModelsPlaylistInfoDTO[]>([]);
const selectedPlaylistId = ref<number | null>(null);
const playlistSearch = ref('');
const newPlaylistForm = reactive<ModelsCreatePlaylistRequestDTO>({
  name: '',
  description: '',
  is_public: false,
});

// 响应式对话框宽度
const dialogWidth = computed(() => {
  return window.innerWidth <= 768 ? '90%' : '600px';
});

// 搜索歌单
const filterPlaylists = () => {
  const query = playlistSearch.value.toLowerCase().trim();
  filteredPlaylists.value = query
    ? userPlaylists.value.filter(
      p => p.name.toLowerCase().includes(query) || p.description?.toLowerCase().includes(query),
    )
    : userPlaylists.value;
};

// 获取用户歌单
const fetchUserPlaylists = async () => {
  isLoading.value = true;
  try {
    const response = await playlistApi.mePlaylistsGet(1, 200);
    userPlaylists.value = (response.data.list as ModelsPlaylistInfoDTO[]) || [];
    filteredPlaylists.value = userPlaylists.value;
  } catch (error) {
    console.error('获取歌单失败:', error);
    ElMessage.error('获取歌单列表失败');
  } finally {
    isLoading.value = false;
  }
};

// 创建新歌单
const createNewPlaylist = async () => {
  if (!newPlaylistForm.name.trim()) {
    ElMessage.warning('请输入歌单名称');
    return;
  }
  isCreating.value = true;
  try {
    const response = await playlistApi.mePlaylistsPost(newPlaylistForm);
    await fetchUserPlaylists();
    selectedPlaylistId.value = response.data.id!;
    newPlaylistForm.name = '';
    newPlaylistForm.description = '';
    newPlaylistForm.is_public = false;
    ElMessage.success('歌单创建成功');
  } catch (error) {
    console.error('创建歌单失败:', error);
    ElMessage.error('创建歌单失败');
  } finally {
    isCreating.value = false;
  }
};

// 确认添加到歌单
const confirmAddToPlaylist = async () => {
  if (!selectedPlaylistId.value || !props.songToAdd) {
    ElMessage.warning('请选择一个歌单');
    return;
  }
  isConfirming.value = true;
  try {
    await playlistApi.mePlaylistsPlaylistIdSongsPost(selectedPlaylistId.value, {
      song_id: props.songToAdd.id!,
    });
    emit('confirm', {
      playlistId: selectedPlaylistId.value,
      songName: props.songToAdd.name,
    });
    emit('update:modelValue', false);
    // ElMessage.success(`已将《${props.songToAdd.name}》添加到歌单`);
  } catch (error) {
    let errorMessage = '添加失败，请稍后重试';
    if (error instanceof AxiosError && error.response?.status === 409) {
      errorMessage = '歌曲已在该歌单中';
    }
    console.error('添加失败:', error);
    ElMessage.error(errorMessage);
  } finally {
    isConfirming.value = false;
  }
};

// 重置状态
const resetState = () => {
  selectedPlaylistId.value = null;
  playlistSearch.value = '';
  newPlaylistForm.name = '';
  newPlaylistForm.description = '';
  newPlaylistForm.is_public = false;
  filteredPlaylists.value = userPlaylists.value;
};
</script>

<style scoped>
/* Spotify 深色主题样式覆盖 */
.add-to-playlist-dialog :deep(.el-dialog) {
  background: #282828;
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #fff;
}

.add-to-playlist-dialog :deep(.el-dialog__header) {
  background: #282828;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.add-to-playlist-dialog :deep(.el-dialog__title) {
  color: #fff;
}

.add-to-playlist-dialog :deep(.el-dialog__headerbtn .el-dialog__close) {
  color: #b3b3b3;
}

.add-to-playlist-dialog :deep(.el-dialog__headerbtn .el-dialog__close:hover) {
  color: #fff;
}

.add-to-playlist-dialog :deep(.el-dialog__body) {
  background: #282828;
  color: #fff;
}

.add-to-playlist-dialog :deep(.el-dialog__footer) {
  background: #282828;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

/* 内容布局 */
.dialog-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.content-layout {
  display: flex;
  gap: 16px;
}

/* 左侧：现有歌单列表 */
.existing-playlists {
  flex: 1;
  min-width: 0;
  max-width: 280px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.05);
}

.playlist-search-input {
  margin: 12px;
  max-width: 240px;
}

.playlist-search-input :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  transition: all 0.3s;
}

.playlist-search-input :deep(.el-input__wrapper:hover) {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
}

.playlist-search-input :deep(.el-input__wrapper.is-focus) {
  border-color: #1db954;
  box-shadow: 0 0 0 1px rgba(29, 185, 84, 0.2);
}

.playlist-search-input :deep(.el-input__inner) {
  color: #fff;
}

.playlist-search-input :deep(.el-input__inner::placeholder) {
  color: #b3b3b3;
}

.playlist-search-input :deep(.el-input__prefix) {
  color: #b3b3b3;
}

.playlist-list {
  display: flex;
  flex-direction: column;
}

.playlist-item {
  padding: 12px 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  color: #fff;
}

.playlist-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
  transform: translateX(4px);
}

.playlist-item-selected {
  background-color: rgba(29, 185, 84, 0.2);
  border-left: 4px solid #1db954;
}

.playlist-info-row {
  display: flex;
  gap: 12px;
  align-items: center;
}

.playlist-cover-img {
  width: 48px;
  height: 48px;
  border-radius: 6px;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
  transition: transform 0.3s;
}

.playlist-item:hover .playlist-cover-img {
  transform: scale(1.05);
}

.playlist-details {
  flex: 1;
  min-width: 0;
}

.playlist-name {
  font-size: 14px;
  font-weight: 500;
  color: #fff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.playlist-meta {
  display: flex;
  gap: 12px;
  font-size: 12px;
  color: #b3b3b3;
  margin-top: 4px;
}

.playlist-visibility.public {
  color: #1db954;
}

.playlist-item-selected .playlist-name {
  color: #1db954;
  font-weight: 600;
}

.playlist-item-selected .playlist-meta {
  color: rgba(29, 185, 84, 0.8);
}

/* 空状态样式覆盖 */
.existing-playlists :deep(.el-empty) {
  background: transparent;
}

.existing-playlists :deep(.el-empty__description) {
  color: #b3b3b3;
}

/* 右侧：创建新歌单 */
.create-playlist {
  flex: 1;
  min-width: 200px;
  padding: 12px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.05);
}

.create-divider {
  margin: 0 0 12px 0;
}

.create-divider :deep(.el-divider__text) {
  white-space: nowrap;
  font-size: 14px;
  font-weight: 500;
  color: #fff;
  background: rgba(255, 255, 255, 0.05);
}

.create-divider :deep(.el-divider__line) {
  background-color: rgba(255, 255, 255, 0.1);
}

.create-playlist-form {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.create-playlist-form :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
}

.create-playlist-form :deep(.el-input__wrapper:hover) {
  border-color: rgba(255, 255, 255, 0.3);
}

.create-playlist-form :deep(.el-input__wrapper.is-focus) {
  border-color: #1db954;
  box-shadow: 0 0 0 1px rgba(29, 185, 84, 0.2);
}

.create-playlist-form :deep(.el-input__inner) {
  color: #fff;
}

.create-playlist-form :deep(.el-input__inner::placeholder) {
  color: #b3b3b3;
}

.create-playlist-form :deep(.el-textarea__inner) {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
}

.create-playlist-form :deep(.el-textarea__inner:hover) {
  border-color: rgba(255, 255, 255, 0.3);
}

.create-playlist-form :deep(.el-textarea__inner:focus) {
  border-color: #1db954;
  box-shadow: 0 0 0 1px rgba(29, 185, 84, 0.2);
}

.create-playlist-form :deep(.el-textarea__inner::placeholder) {
  color: #b3b3b3;
}

.create-playlist-form :deep(.el-switch__label) {
  color: #fff;
}

.create-playlist-form :deep(.el-switch__core) {
  background-color: #5e5e5e;
  border-color: #5e5e5e;
}

.create-playlist-form :deep(.el-switch.is-checked .el-switch__core) {
  background-color: #1db954;
  border-color: #1db954;
}

.mt-2 {
  margin-top: 8px;
}

.mt-3 {
  margin-top: 12px;
}

.create-playlist-btn {
  border-radius: 20px;
  transition: all 0.3s;
  background: #1db954;
  border-color: #1db954;
  color: #fff;
}

.create-playlist-btn:hover {
  transform: scale(1.02);
  box-shadow: 0 4px 12px rgba(29, 185, 84, 0.4);
  background: #1ed760;
  border-color: #1ed760;
}

/* 滚动条样式 */
.existing-playlists :deep(.el-scrollbar__wrap) {
  overflow-x: hidden;
  padding-right: 8px;
}

.existing-playlists :deep(.el-scrollbar__thumb) {
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 3px;
}

.existing-playlists :deep(.el-scrollbar__thumb:hover) {
  background-color: rgba(255, 255, 255, 0.3);
}

/* 对话框底部 */
.dialog-footer {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.dialog-footer :deep(.el-button) {
  color: #b3b3b3;
  background: transparent;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.dialog-footer :deep(.el-button:hover) {
  color: #fff;
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.3);
}

.dialog-footer :deep(.el-button--primary) {
  background: #1db954;
  border-color: #1db954;
  color: #fff;
}

.dialog-footer :deep(.el-button--primary:hover) {
  background: #1ed760;
  border-color: #1ed760;
}

.dialog-footer :deep(.el-button--primary:disabled) {
  background: #5e5e5e;
  border-color: #5e5e5e;
  color: #999;
}

/* 字数统计样式 */
:deep(.el-input__count) {
  color: #b3b3b3;
  background: transparent;
}

:deep(.el-textarea__count) {
  color: #b3b3b3;
  background: rgba(40, 40, 40, 0.9);
}

/* 加载状态覆盖 */
:deep(.el-loading-mask) {
  background-color: rgba(0, 0, 0, 0.8);
}

:deep(.el-loading-spinner) {
  color: #1db954;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .content-layout {
    flex-direction: column;
    gap: 16px;
  }

  .existing-playlists,
  .create-playlist {
    border: none;
    padding: 0;
    max-width: none;
  }

  .playlist-search-input {
    max-width: none;
  }

  .playlist-item {
    padding: 10px 12px;
  }

  .playlist-cover-img {
    width: 40px;
    height: 40px;
  }

  .create-playlist-form {
    gap: 8px;
  }
}
</style>