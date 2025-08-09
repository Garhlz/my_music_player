<template>
  <div class="music-player">
    <audio
      ref="audio"
      :src="currentSong?.url"
      @timeupdate="updateProgress"
      @ended="nextSong"
      @canplay="onCanPlay"
    ></audio>
    <div v-if="currentSong" class="player-controls">
      <div class="left-section">
        <img
          :src="currentSong.cover || '/assets/default-cover.jpg'"
          alt="cover"
          class="song-cover"
          @click="goToPlayer"
        />
        <div class="song-info">
          <span class="song-name">{{ currentSong.name }}</span>
          <span class="artist-name">{{ currentSong.artist_name || '未知歌手' }}</span>
        </div>
      </div>

      <div class="center-section">
        <div class="control-buttons">
          <el-button circle @click="previousSong">
            <el-icon>
              <ArrowLeft />
            </el-icon>
          </el-button>
          <el-button circle @click="togglePlay" class="play-btn">
            <el-icon>
              <VideoPlay v-if="!isPlaying" />
              <VideoPause v-else />
            </el-icon>
          </el-button>
          <el-button circle @click="nextSong">
            <el-icon>
              <ArrowRight />
            </el-icon>
          </el-button>
        </div>
        <div class="progress-bar">
          <span class="time">{{ formatDuration(currentProgress) }}</span>
          <el-slider
            :model-value="currentProgress"
            :max="maxDuration"
            :show-tooltip="false"
            @change="handleProgressChange"
          />
          <span class="time">{{ formatDuration(maxDuration) }}</span>
        </div>
      </div>

      <div class="right-section">
        <div class="volume-control">
          <el-icon @click="toggleMute">
            <CaretBottom v-if="volume === 0" />
            <CaretTop v-else />
          </el-icon>
          <el-slider
            v-model="volume"
            :max="100"
            :show-tooltip="false"
            @input="setVolume"
            class="volume-slider"
          />
        </div>
      </div>
    </div>
    <div v-else class="empty-player">
      <span>暂无播放歌曲</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import { usePlayerStore } from '@/stores/player';
import { ArrowLeft, ArrowRight, CaretBottom, CaretTop, VideoPause, VideoPlay } from '@element-plus/icons-vue';
import { formatDuration } from '@/utils/format';

const playerStore = usePlayerStore();
const router = useRouter();
const audio = ref<HTMLAudioElement | null>(null);
const currentProgress = ref(0);
const maxDuration = ref(0);
const volume = ref(100);
const previousVolume = ref(100);

const currentSong = computed(() => playerStore.getCurrentSong);
const isPlaying = computed(() => playerStore.getIsPlaying);

const updateProgress = () => {
  if (audio.value) {
    const current = audio.value.currentTime;
    const duration = audio.value.duration;
    if (isFinite(current)) currentProgress.value = current;
    if (isFinite(duration)) maxDuration.value = duration;
  }
};

const handleProgressChange = (value: number) => {
  if (audio.value && isFinite(value)) {
    try {
      audio.value.currentTime = value;
    } catch (error) {
      console.error('设置播放进度失败:', error);
    }
  }
};

const togglePlay = () => {
  if (!audio.value) return;
  if (isPlaying.value) {
    audio.value.pause();
    playerStore.pause();
  } else if (currentSong.value) {
    audio.value.play().catch((error) => {
      console.error('播放失败:', error);
    });
    playerStore.resume();
  }
};

const previousSong = () => {
  playerStore.previous();
};

const nextSong = () => {
  playerStore.next();
};

const toggleMute = () => {
  if (volume.value > 0) {
    previousVolume.value = volume.value;
    volume.value = 0;
  } else {
    volume.value = previousVolume.value;
  }
  setVolume(volume.value);
};

const setVolume = (value: number) => {
  if (audio.value) {
    audio.value.volume = value / 100;
  }
};

const onCanPlay = () => {
  if (isPlaying.value) {
    audio.value?.play().catch((error) => {
      console.error('播放失败:', error);
    });
  }
};

const goToPlayer = () => {
  router.push('/player');
};

watch(
  () => currentSong.value,
  (newSong) => {
    if (newSong && isPlaying.value && audio.value) {
      audio.value.play().catch((error) => {
        console.error('播放失败:', error);
      });
    }
  },
  { immediate: true },
);

onMounted(() => {
  if (audio.value) {
    audio.value.volume = volume.value / 100;
  }
});

// 潜在新功能（未实现）:
// - 迷你模式：将播放器切换为右下角悬浮小窗口。
// - 播放队列：在右边显示当前播放队列（el-drawer）。
// - 歌词显示：点击封面时显示歌词面板（el-dialog）。
</script>

<style scoped>
.music-player {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 80px;
  background: linear-gradient(180deg, #fff, #f5f7fa);
  border-radius: 12px 12px 0 0;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  animation: fadeIn 0.5s ease-in;
}

.player-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 24px;
}

.left-section {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 300px;
}

.song-cover {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.song-cover:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.song-info {
  display: flex;
  flex-direction: column;
  min-width: 0;
}

.song-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--el-text-color-primary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.artist-name {
  font-size: 12px;
  color: #909399;
}

.center-section {
  flex: 1;
  max-width: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.control-buttons {
  display: flex;
  gap: 12px;
  align-items: center;
}

.control-buttons .el-button {
  border-radius: 50%;
  transition: all 0.3s ease;
}

.control-buttons .el-button:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.play-btn {
  width: 48px;
  height: 48px;
}

.progress-bar {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 8px;
}

.time {
  font-size: 12px;
  color: #909399;
  width: 40px;
  text-align: center;
}

.el-slider {
  flex: 1;
  margin: 0;
}

.right-section {
  width: 200px;
  display: flex;
  justify-content: flex-end;
}

.volume-control {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 150px;
}

.volume-slider {
  width: 100px;
}

.empty-player {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #909399;
  font-size: 14px;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .music-player {
    padding: 0 16px;
  }

  .left-section {
    width: 200px;
  }

  .song-cover {
    width: 40px;
    height: 40px;
  }

  .center-section {
    max-width: 300px;
  }

  .right-section {
    width: 120px;
  }

  .volume-slider {
    width: 80px;
  }

  .control-buttons {
    gap: 8px;
  }

  .play-btn {
    width: 40px;
    height: 40px;
  }
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
</style>
