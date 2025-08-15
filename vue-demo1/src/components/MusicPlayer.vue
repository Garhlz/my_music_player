<template>
  <div class="music-player-bar">
    <div v-if="!currentSong" class="empty-state">
      <el-icon>
        <Headset />
      </el-icon>
      <span>把好听的音乐加进来呗~</span>
    </div>

    <div v-else class="player-controls">
      <audio
        :key="currentSong.id"
        ref="audioEl"
        :src="currentSong.url"
        :loop="playerStore.playbackMode === 'single-loop'"
        @timeupdate="onTimeUpdate"
        @loadedmetadata="onLoadedMetadata"
        @ended="handleSongEnd"
        @canplay="onCanPlay"
      ></audio>

      <div class="section left">
        <img :src="currentSong.cover || '/assets/default-cover.jpg'" alt="歌曲封面" class="song-cover"
             @click="goToPlayerPage" />
        <div class="song-info">
          <span class="song-name">{{ currentSong.name }}</span>
          <span class="artist-name">{{ currentSong.artist_name || '未知艺术家' }}</span>
        </div>

      </div>

      <div class="section center">
        <div class="control-buttons">
          <el-tooltip :content="playbackModeTooltip" placement="top" :hide-after="0">
            <el-button text circle @click="onTogglePlaybackMode" class="mode-button">
              <el-icon size="16">
                <Refresh v-if="playerStore.playbackMode === 'list-loop'" />
                <Switch v-if="playerStore.playbackMode === 'single-loop'" />
                <Sort v-if="playerStore.playbackMode === 'shuffle'" />
              </el-icon>
            </el-button>
          </el-tooltip>

          <el-button text circle @click="playerStore.previous()" class="control-button">
            <el-icon size="20">
              <Back />
            </el-icon>
          </el-button>

          <el-button circle class="play-button" @click="playerStore.togglePlay()">
            <el-icon size="20">
              <VideoPause v-if="isPlaying" />
              <VideoPlay v-else />
            </el-icon>
          </el-button>

          <el-button text circle @click="playerStore.next()" class="control-button">
            <el-icon size="20">
              <Right />
            </el-icon>
          </el-button>
          <el-tooltip :content="isCurrentSongLiked ? '取消喜欢' : '喜欢'" placement="top" :hide-after="0">
            <el-button text circle @click="handleLikeToggle" class="control-button">
              <el-icon size="16"
                       :class="{active: isCurrentSongLiked}">
                <StarFilled v-if="isCurrentSongLiked" />
                <Star v-else />
              </el-icon>
            </el-button>
          </el-tooltip>

        </div>

        <div class="progress-bar">
          <span class="time">{{ formatDuration(currentTime) }}</span>
          <el-slider
            v-model="currentTime"
            :max="duration"
            :show-tooltip="false"
            :disabled="duration === 0"
            @input="onProgressInput"
            @change="onProgressChange"
            class="progress-slider"
          />
          <span class="time">{{ formatDuration(duration) }}</span>
        </div>
      </div>

      <div class="section right">
        <el-tooltip content="静音" placement="top" :hide-after="0">
          <el-button text circle @click="toggleMute" class="volume-button">
            <el-icon size="16">
              <Mute v-if="isMuted || volume === 0" />
              <Operation v-else />
            </el-icon>
          </el-button>
        </el-tooltip>

        <div class="volume-control">
          <el-slider
            v-model="volume"
            :max="100"
            :show-tooltip="false"
            @input="onVolumeChange"
            class="volume-slider"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, computed, nextTick, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { usePlayerStore, PlaybackMode } from '@/stores/player';
import { useUserStore } from '@/stores/user';
import { useLikeStore } from '@/stores/like';
import { storeToRefs } from 'pinia';
import { formatDuration } from '@/utils/format';
import {
  VideoPlay, VideoPause, Back, Right, Mute, Headset,
  Sort, Switch, Operation, Refresh, Star, StarFilled,
} from '@element-plus/icons-vue';

import { likeApi } from '@/api';
import type { ModelsSongDetailDTO } from '@/api-client';
import { ElMessage } from 'element-plus';

const playerStore = usePlayerStore();
const userStore = useUserStore();
const likeStore = useLikeStore();
const { currentSong, isPlaying, currentTime, duration } = storeToRefs(playerStore);
const { isLoggedIn } = storeToRefs(userStore);
const router = useRouter();
const audioEl = ref<HTMLAudioElement | null>(null);
const volume = ref(80);
const isMuted = ref(false);
const isDraggingProgress = ref(false);

// 添加一个标记，防止恢复进度时触发 timeupdate
const isRestoringProgress = ref(false);

const isCurrentSongLiked = computed(() => {
  if (!currentSong.value?.id) return false;
  return likeStore.isLiked(currentSong.value.id);
});

const currentVolume = computed(() => (isMuted.value ? 0 : volume.value / 100));
const playbackModeTooltip = computed(() => {
  const modeMap: Record<PlaybackMode, string> = {
    'list-loop': '列表循环', 'single-loop': '单曲循环', 'shuffle': '随机播放',
  };
  return modeMap[playerStore.playbackMode];
});

// 组件挂载时恢复播放进度
onMounted(() => {
  if (currentSong.value && audioEl.value) {
    restorePlaybackState();
  }
});


// 恢复播放状态的函数
const restorePlaybackState = async () => {
  if (!audioEl.value || !currentSong.value) return;

  isRestoringProgress.value = true;

  // 等待音频元素可以播放
  const restoreProgress = () => {
    if (audioEl.value && currentTime.value > 0) {
      audioEl.value.currentTime = currentTime.value;
      console.log('恢复播放进度:', currentTime.value);
    }

    // 恢复播放状态
    if (isPlaying.value && audioEl.value) {
      audioEl.value.play().catch(e => console.error('恢复播放失败:', e));
    }

    setTimeout(() => {
      isRestoringProgress.value = false;
    }, 100);
  };

  // 如果音频已经加载了元数据，直接恢复
  if (audioEl.value.readyState >= 1) {
    restoreProgress();
  } else {
    // 否则等待元数据加载完成
    audioEl.value.addEventListener('loadedmetadata', restoreProgress, { once: true });
  }
};

// --- Event Handlers ---
// 修改 onTimeUpdate 函数，增加更严格的检查
const onTimeUpdate = () => {
  if (audioEl.value && !isDraggingProgress.value) {
    const time = audioEl.value.currentTime;
    const dur = audioEl.value.duration;

    // 确保时间值有效
    if (time >= 0 && !isNaN(time) && dur > 0 && !isNaN(dur)) {
      playerStore.setAudioCurrentTime(time);
    }
  }
};

const onLoadedMetadata = () => {
  if (audioEl.value) {
    playerStore.setAudioDuration(audioEl.value.duration);

    // 如果是组件刚挂载，恢复播放进度
    if (currentTime.value > 0 && !isRestoringProgress.value) {
      restorePlaybackState();
    }
  }
};

const onCanPlay = () => {
  // 只有在不是恢复进度的情况下才自动播放
  if (isPlaying.value && audioEl.value && !isRestoringProgress.value) {
    audioEl.value.play().catch(e => console.error('自动播放失败:', e));
  }
};

// 修改 handleSongEnd 函数
const handleSongEnd = () => {
  // 播放结束时先重置进度
  playerStore.setAudioCurrentTime(0);

  if (playerStore.playbackMode !== 'single-loop') {
    playerStore.next();
  }
};

const onProgressInput = (newTime: number) => {
  if (duration.value === 0) return;
  isDraggingProgress.value = true;
  playerStore.setAudioCurrentTime(newTime);
};

const onProgressChange = (newTime: number) => {
  if (duration.value === 0) return;
  if (audioEl.value) {
    audioEl.value.currentTime = newTime;
  }
  setTimeout(() => {
    isDraggingProgress.value = false;
  }, 50);
};

const onTogglePlaybackMode = () => {
  playerStore.togglePlaybackMode();
};

const toggleMute = () => {
  isMuted.value = !isMuted.value;
};

const onVolumeChange = (newVolume: number) => {
  volume.value = newVolume;
  isMuted.value = newVolume === 0;
};

const goToPlayerPage = () => {
  if (currentSong.value?.id) {
    router.push(`/player/${currentSong.value.id}`);
  }
};

const handleLikeToggle = () => {
  if (!isLoggedIn.value) {
    ElMessage.warning('请先登录');
    return;
  }
  const songId = currentSong.value?.id;
  if (!songId) return;

  if (isCurrentSongLiked.value) {
    likeStore.unlikeSong(songId);
  } else {
    likeStore.likeSong(songId);
  }
};

// --- Watchers for State Synchronization ---
watch(isPlaying, (newIsPlaying) => {
  if (!audioEl.value) return;
  try {
    if (newIsPlaying) {
      audioEl.value.play();
    } else {
      audioEl.value.pause();
    }
  } catch (e) {
    console.error('播放/暂停操作失败:', e);
  }
});

watch(currentVolume, (newVolume) => {
  if (audioEl.value) {
    audioEl.value.volume = newVolume;
  }
});

// --- Watchers for State Synchronization ---
// 修复：监听歌曲切换，重置进度
watch(() => currentSong.value?.id, (newId, oldId) => {
  if (newId && newId !== oldId) {
    // 歌曲切换时，重置进度条
    playerStore.setAudioCurrentTime(0);
    playerStore.setAudioDuration(0);

    nextTick(() => {
      if (audioEl.value) {
        audioEl.value.volume = currentVolume.value;
        // 重置音频元素的当前时间
        audioEl.value.currentTime = 0;

        if (isPlaying.value) {
          audioEl.value.play().catch(e => console.error('切歌后自动播放失败:', e));
        }
      }
    });

  }
}, { immediate: true });
</script>

<style scoped>
/* Spotify 风格的播放器栏 */
.music-player-bar {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #121212;
  padding: 0 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  position: relative;
}

.music-player-bar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(29, 185, 84, 0.5),
    transparent
  );
}

/* 空状态样式 */
.empty-state {
  display: flex;
  align-items: center;
  gap: 12px;
  color: #b3b3b3;
  font-size: 14px;
}

.empty-state .el-icon {
  color: #1db954;
  font-size: 20px;
}

/* 播放器控件布局 */
.player-controls {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  align-items: center;
  gap: 20px;
  max-width: 1800px;
}

.section {
  display: flex;
  align-items: center;
  height: 100%;
}

/* 左侧区域 - 歌曲信息 */
.section.left {
  justify-content: flex-start;
  gap: 12px;
  min-width: 0;
}

.song-cover {
  width: 56px;
  height: 56px;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
  flex-shrink: 0;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.5);
}

.song-cover:hover {
  transform: scale(1.05);
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.6);
}

.song-info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-width: 0;
  flex-grow: 1;
}

.song-name, .artist-name {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  cursor: pointer;
}

.song-name {
  font-size: 14px;
  font-weight: 400;
  color: #fff;
  margin-bottom: 2px;
  transition: color 0.2s ease;
}

.song-name:hover {
  text-decoration: underline;
}

.artist-name {
  font-size: 11px;
  color: #b3b3b3;
  transition: color 0.2s ease;
}

.artist-name:hover {
  color: #fff;
  text-decoration: underline;
}

.like-button .el-button {
  color: #b3b3b3;
  transition: all 0.2s ease;
}

.like-button .el-button:hover {
  color: #1db954;
  transform: scale(1.1);
}

/* 中间区域 - 播放控制 */
.section.center {
  flex-direction: column;
  justify-content: center;
  gap: 8px;
}

.control-buttons {
  display: flex;
  align-items: center;
  gap: 8px;
  justify-content: center;
}

.mode-button,
.control-button {
  color: #b3b3b3;
  transition: all 0.2s ease;
  width: 32px;
  height: 32px;
}

/* 激活状态的图标变为 Spotify 绿 */
.control-button .el-icon.active {
  color: #1db954;
}

.mode-button:hover,
.control-button:hover {
  color: #fff;
  transform: scale(1.1);
}

.play-button {
  width: 32px;
  height: 32px;
  /* 使用 !important 覆盖全局样式，确保背景是白色 */
  //background: #fff !important;
  /* 确保图标是黑色，以形成对比 */
  color: #fff;
  border: none;
  transition: all 0.2s ease;
}

.play-button:hover {
  background: #fff;
  transform: scale(1.05);
}

.play-button:active {
  transform: scale(0.95);
}

.progress-bar {
  width: 100%;
  max-width: 480px;
  display: flex;
  align-items: center;
  gap: 12px;
}

.time {
  font-size: 11px;
  color: #b3b3b3;
  width: 40px;
  text-align: center;
  user-select: none;
  font-variant-numeric: tabular-nums;
}

.progress-slider {
  flex-grow: 1;
}

/* 进度条样式覆盖 */
.progress-slider :deep(.el-slider__runway) {
  background: #5e5e5e;
  height: 4px;
}

.progress-slider :deep(.el-slider__bar) {
  background: #1db954;
  height: 4px;
}

.progress-slider :deep(.el-slider__button) {
  width: 12px;
  height: 12px;
  background: #fff;
  border: none;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.progress-slider:hover :deep(.el-slider__button) {
  opacity: 1;
}

.progress-slider :deep(.el-slider__button-wrapper) {
  z-index: 100;
}

/* 右侧区域 - 音量控制 */
.section.right {
  justify-content: flex-end;
  gap: 8px;
}

.volume-button {
  color: #b3b3b3;
  transition: all 0.2s ease;
  width: 32px;
  height: 32px;
}

.volume-button:hover {
  color: #fff;
  transform: scale(1.1);
}

.volume-control {
  width: 93px;
  display: flex;
  align-items: center;
}

.volume-slider {
  width: 100%;
}

/* 音量滑块样式覆盖 */
.volume-slider :deep(.el-slider__runway) {
  background: #5e5e5e;
  height: 4px;
}

.volume-slider :deep(.el-slider__bar) {
  background: #1db954;
  height: 4px;
}

.volume-slider :deep(.el-slider__button) {
  width: 12px;
  height: 12px;
  background: #fff;
  border: none;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.volume-slider:hover :deep(.el-slider__button) {
  opacity: 1;
}

/* 响应式设计 */
@media screen and (max-width: 1024px) {
  .player-controls {
    grid-template-columns: 1fr 1.5fr 0.8fr;
    gap: 16px;
  }

  .volume-control {
    width: 70px;
  }
}

@media screen and (max-width: 768px) {
  .music-player-bar {
    padding: 0 12px;
  }

  .player-controls {
    grid-template-columns: 1fr 1fr 1fr;
    gap: 8px;
  }

  .section.center {
    gap: 6px;
  }

  .control-buttons {
    gap: 4px;
  }

  .progress-bar {
    max-width: 200px;
    gap: 8px;
  }

  .time {
    width: 35px;
    font-size: 10px;
  }

  .volume-control {
    width: 60px;
  }

  .like-button {
    display: none;
  }
}

@media screen and (max-width: 480px) {
  .player-controls {
    display: flex;
    justify-content: space-between;
  }

  .section.center {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    width: auto;
  }

  .progress-bar {
    display: none;
  }

  .section.right {
    gap: 4px;
  }

  .volume-control {
    display: none;
  }

  .song-info .artist-name {
    display: none;
  }
}

/* 深色主题的按钮样式覆盖 */
.el-button {
  background: transparent !important;
  border: none !important;
}

.el-button:hover {
  background: rgba(255, 255, 255, 0.1) !important;
}

.el-button.is-circle {
  border-radius: 50% !important;
}

/* 无障碍性和高对比度支持 */
@media (prefers-contrast: high) {
  .music-player-bar {
    border-top: 2px solid rgba(255, 255, 255, 0.3);
  }

  .song-name,
  .artist-name {
    text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
  }
}

/* 减少动画的用户偏好支持 */
@media (prefers-reduced-motion: reduce) {
  * {
    transition: none !important;
    animation: none !important;
  }
}
</style>