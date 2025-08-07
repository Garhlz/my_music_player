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

		  <!-- 播放控制 -->
		  <div class="player-controls">
			<el-button circle @click="previousSong">
			  <el-icon>
				<ArrowLeft/>
			  </el-icon>
			</el-button>
			<el-button circle @click="togglePlay">
			  <el-icon>
				<VideoPlay v-if="!isPlaying"/>
				<VideoPause v-else/>
			  </el-icon>
			</el-button>
			<el-button circle @click="nextSong">
			  <el-icon>
				<ArrowRight/>
			  </el-icon>
			</el-button>
		  </div>

		  <!-- 音频控制器 -->
		  <audio
			  v-if="currentSong"
			  :src="currentSong.url"
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

<script setup>
import {ref, computed} from 'vue'
import {useRouter} from 'vue-router'
import CommonLayout from 'src/layouts/CommonLayout.vue'
import {usePlayerStore} from 'src/stores/player'
import {
  ArrowLeft,
  ArrowRight,
  VideoPlay,
  VideoPause
} from '@element-plus/icons-vue'

const router = useRouter()
const playerStore = usePlayerStore()
const currentName = ref('正在播放')

const currentSong = computed(() => playerStore.getCurrentSong)
const isPlaying = computed(() => playerStore.getIsPlaying)

const togglePlay = () => {
  if (isPlaying.value) {
	playerStore.pause()
  } else {
	playerStore.resume()
  }
}

const previousSong = () => {
  playerStore.previous()
}

const nextSong = () => {
  playerStore.next()
}

const handlePlay = () => {
  playerStore.resume()
}

const handlePause = () => {
  playerStore.pause()
}

const handleEnded = () => {
  playerStore.next()
}

const handleTimeUpdate = (event) => {
  // 可以在这里处理播放进度更新
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
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
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

.player-controls {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
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