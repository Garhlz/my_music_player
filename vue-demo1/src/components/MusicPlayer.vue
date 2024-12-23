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
        <img :src="currentSong.cover" alt="cover" class="song-cover" @click="goToPlayer" />
        <div class="song-info">
          <span class="song-name">{{ currentSong.name }}</span>
          <span class="artist-name">{{ currentSong.artist }}</span>
        </div>
      </div>
      
      <div class="center-section">
        <div class="control-buttons">
          <el-button circle @click="previousSong">
            <el-icon><CaretLeft /></el-icon>
          </el-button>
          <el-button circle @click="togglePlay">
            <el-icon>
              <VideoPlay v-if="!isPlaying" />
              <VideoPause v-else />
            </el-icon>
          </el-button>
          <el-button circle @click="nextSong">
            <el-icon><CaretRight /></el-icon>
          </el-button>
        </div>
        <div class="progress-bar">
          <span class="time">{{ formatTime(currentProgress) }}</span>
          <el-slider
            :model-value="currentProgress"
            :max="maxDuration"
            :show-tooltip="false"
            @change="handleProgressChange"
          />
          <span class="time">{{ formatTime(maxDuration) }}</span>
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

<script>
import { mapState, mapActions } from 'vuex';
import { ref, computed, onMounted, watch } from 'vue';
import { 
  CaretLeft, 
  CaretRight, 
  VideoPlay, 
  VideoPause, 
  CaretTop, 
  CaretBottom 
} from '@element-plus/icons-vue';

export default {
  components: {
    CaretLeft,
    CaretRight,
    VideoPlay,
    VideoPause,
    CaretTop,
    CaretBottom
  },
  
  setup() {
    const audio = ref(null);
    const currentProgress = ref(0);
    const maxDuration = ref(0);
    const volume = ref(100);
    const previousVolume = ref(100);

    // 处理音频进度更新
    const updateProgress = () => {
      if (audio.value) {
        const current = audio.value.currentTime;
        const duration = audio.value.duration;
        
        if (isFinite(current)) {
          currentProgress.value = current;
        }
        if (isFinite(duration)) {
          maxDuration.value = duration;
        }
      }
    };

    // 处理进度条改变
    const handleProgressChange = (value) => {
      if (audio.value && isFinite(value)) {
        try {
          audio.value.currentTime = value;
        } catch (error) {
          console.error('设置播放进度失败:', error);
        }
      }
    };

    // 格式化时间
    const formatTime = (seconds) => {
      if (!isFinite(seconds)) return '0:00';
      const mins = Math.floor(seconds / 60);
      const secs = Math.floor(seconds % 60);
      return `${mins}:${secs.toString().padStart(2, '0')}`;
    };

    onMounted(() => {
      // 初始化音频元素引用
      audio.value = document.querySelector('audio');
    });

    return {
      audio,
      currentProgress,
      maxDuration,
      volume,
      previousVolume,
      updateProgress,
      handleProgressChange,
      formatTime
    };
  },

  computed: {
    ...mapState('player', ['currentSong', 'isPlaying']),
  },

  methods: {
    ...mapActions('player', ['playSong', 'pauseSong', 'nextSong', 'previousSong']),

    togglePlay() {
      if (!this.audio) return;

      if (this.isPlaying) {
        this.audio.pause();
        this.pauseSong();
      } else if (this.currentSong) {
        this.audio.play().catch(error => {
          console.error('播放失败:', error);
        });
        this.playSong(this.currentSong);
      }
    },

    toggleMute() {
      if (this.volume > 0) {
        this.previousVolume = this.volume;
        this.volume = 0;
      } else {
        this.volume = this.previousVolume;
      }
      this.setVolume(this.volume);
    },

    setVolume(value) {
      if (this.audio) {
        this.audio.volume = value / 100;
      }
    },

    goToPlayer() {
      this.$router.push('/player');
    },
  },

  watch: {
    currentSong: {
      handler(newSong) {
        if (newSong && this.isPlaying && this.audio) {
          this.$nextTick(() => {
            this.audio.play().catch(error => {
              console.error('播放失败:', error);
            });
          });
        }
      },
      immediate: true
    }
  },

  mounted() {
    // 初始化音量
    if (this.audio) {
      this.audio.volume = this.volume / 100;
    }

    // 添加进度更新监听
    if (this.audio) {
      this.audio.addEventListener('timeupdate', this.updateProgress);
    }
  },

  beforeUnmount() {
    // 清理事件监听
    if (this.audio) {
      this.audio.removeEventListener('timeupdate', this.updateProgress);
    }
  }
};
</script>

<style scoped>
.music-player {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 64px;
  background-color: #fff;
  border-top: 1px solid #e4e4e4;
  z-index: 1000;
}

.player-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 20px;
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
  border-radius: 4px;
  cursor: pointer;
}

.song-info {
  display: flex;
  flex-direction: column;
}

.song-name {
  font-size: 14px;
  font-weight: 500;
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
  gap: 4px;
}

.control-buttons {
  display: flex;
  gap: 16px;
  align-items: center;
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
}
</style>