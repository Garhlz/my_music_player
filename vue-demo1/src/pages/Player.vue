<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="playlist-container">
        <div class="player-content">
          <!-- 旋转的专辑封面 -->
          <div class="album-container" :class="{ 'is-playing': isPlaying }">
            <div class="album-cover">
              <el-image 
                :src="currentSong?.cover || '/assets/default-cover.jpg'"
                fit="cover"
              />
            </div>
          </div>

          <!-- 歌曲信息 -->
          <div class="song-info">
            <h2 class="song-name">{{ currentSong?.name || '未播放' }}</h2>
            <p class="song-artist">{{ currentSong?.artist || '未知歌手' }}</p>
          </div>

          <!-- 音频控制器 -->
          <audio 
            v-if="currentSong" 
            :src="currentSong.file" 
            ref="audioPlayer"
            @play="handlePlay"
            @pause="handlePause"
            @ended="handleEnded"
            @timeupdate="handleTimeUpdate"
            controls 
            class="audio-player"
          ></audio>
        </div>
      </div>
    </template>
  </CommonLayout>
</template>

<script>
import CommonLayout from '@/layouts/CommonLayout.vue'
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Player',
  components: {
    CommonLayout
  },
  data() {
    return {
      currentName: '正在播放',
      isPlaying: false,
      currentTime: 0,
      duration: 0
    }
  },
  computed: {
    ...mapState(['currentSong'])
  },
  methods: {
    ...mapActions(['playSong', 'playNext']),

    handlePlay() {
      this.isPlaying = true
    },

    handlePause() {
      this.isPlaying = false
    },

    handleEnded() {
      this.isPlaying = false
      this.playNext()
    },

    handleTimeUpdate(e) {
      this.currentTime = e.target.currentTime
      this.duration = e.target.duration
    }
  }
}
</script>

<style scoped>
@import '../styles/common.css';

.player-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 70vh;
  padding: 20px;
}

.album-container {
  width: 300px;
  height: 300px;
  margin-bottom: 30px;
  perspective: 1000px;
}

.album-cover {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  overflow: hidden;
  box-shadow: 0 0 20px rgba(0,0,0,0.2);
  transition: transform 0.3s ease;
}

.album-cover .el-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* 旋转动画 */
.is-playing .album-cover {
  animation: rotate 20s linear infinite;
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.song-info {
  text-align: center;
  margin-bottom: 30px;
}

.song-name {
  font-size: 24px;
  font-weight: 600;
  margin: 0 0 8px;
}

.song-artist {
  font-size: 16px;
  color: #666;
  margin: 0;
}

.audio-player {
  width: 80%;
  max-width: 500px;
}

/* 自定义音频控制器样式 */
.audio-player::-webkit-media-controls-panel {
  background-color: #f5f7fa;
}

.audio-player::-webkit-media-controls-play-button {
  background-color: #409EFF;
  border-radius: 50%;
}

.audio-player::-webkit-media-controls-current-time-display,
.audio-player::-webkit-media-controls-time-remaining-display {
  color: #606266;
}
</style> 