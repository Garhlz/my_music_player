<template>
  <div class="player-container" v-if="currentSong">
	<div class="player-info">
	  <el-image
		  :src="currentSong.cover || '/assets/default-cover.jpg'"
		  class="song-cover"
		  fit="cover"
	  />
	  <div class="song-info">
		<div class="song-name">{{ currentSong.name }}</div>
		<div class="song-artist">{{ currentSong.artist }}</div>
	  </div>
	</div>

	<div class="player-controls">
	  <el-button
		  circle
		  :icon="ArrowLeft"
		  @click="previousSong"
	  />
	  <el-button
		  circle
		  :icon="isPlaying ? VideoPause : VideoPlay"
		  @click="togglePlay"
	  />
	  <el-button
		  circle
		  :icon="ArrowRight"
		  @click="nextSong"
	  />
	</div>

	<div class="player-volume">
	  <el-slider v-model="volume" :max="100" :min="0"/>
	</div>
  </div>
</template>

<script setup>
import {ref, computed} from 'vue'
import {ArrowLeft, ArrowRight, VideoPlay, VideoPause} from '@element-plus/icons-vue'
import {usePlayerStore} from 'src/stores/player'

const playerStore = usePlayerStore()

// 状态
const volume = ref(100)

// 计算属性
const currentSong = computed(() => playerStore.getCurrentSong)
const isPlaying = computed(() => playerStore.getIsPlaying)

// 方法
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
</script>

<style scoped>
.player-container {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 72px;
  background: var(--el-bg-color);
  border-top: 1px solid var(--el-border-color-lighter);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  z-index: 1000;
}

.player-info {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 300px;
}

.song-cover {
  width: 48px;
  height: 48px;
  border-radius: 4px;
}

.song-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.song-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--el-text-color-primary);
}

.song-artist {
  font-size: 12px;
  color: var(--el-text-color-secondary);
}

.player-controls {
  display: flex;
  gap: 16px;
  align-items: center;
}

.player-volume {
  width: 200px;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .player-container {
	padding: 0 16px;
  }

  .player-info {
	width: auto;
  }

  .song-info {
	display: none;
  }

  .player-volume {
	display: none;
  }
}
</style> 