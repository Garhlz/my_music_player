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
			:src="currentSong.cover"
			alt="cover"
			class="song-cover"
			@click="goToPlayer"
		/>
		<div class="song-info">
		  <span class="song-name">{{ currentSong.name }}</span>
		  <span class="artist-name">{{ currentSong.artist }}</span>
		</div>
	  </div>

	  <div class="center-section">
		<div class="control-buttons">
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
			<CaretBottom v-if="volume === 0"/>
			<CaretTop v-else/>
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

<script setup>
import {usePlayerStore} from "src/stores/player";
import {
  ArrowLeft,
  ArrowRight,
  CaretBottom,
  CaretTop,
  VideoPause,
  VideoPlay,
} from "@element-plus/icons-vue";
import {computed, onMounted, ref, watch} from "vue";

const playerStore = usePlayerStore();
const audio = ref(null);
const currentProgress = ref(0);
const maxDuration = ref(0);
const volume = ref(100);
const previousVolume = ref(100);

// 计算属性
const currentSong = computed(() => playerStore.getCurrentSong);
const isPlaying = computed(() => playerStore.getIsPlaying);

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
	  console.error("设置播放进度失败:", error);
	}
  }
};

// 格式化时间
const formatTime = (seconds) => {
  if (!isFinite(seconds)) return "0:00";
  const mins = Math.floor(seconds / 60);
  const secs = Math.floor(seconds % 60);
  return `${mins}:${secs.toString().padStart(2, "0")}`;
};

// 播放控制方法
const togglePlay = () => {
  if (!audio.value) return;

  if (isPlaying.value) {
	audio.value.pause();
	playerStore.pause();
  } else if (currentSong.value) {
	audio.value.play().catch((error) => {
	  console.error("播放失败:", error);
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

const setVolume = (value) => {
  if (audio.value) {
	audio.value.volume = value / 100;
  }
};

const handleSongEnd = () => {
  playerStore.next();
};

const onCanPlay = () => {
  if (isPlaying.value) {
	audio.value?.play().catch((error) => {
	  console.error("播放失败:", error);
	});
  }
};

// 监听歌曲变化
watch(
	() => currentSong.value,
	(newSong) => {
	  if (newSong && isPlaying.value && audio.value) {
		audio.value.play().catch((error) => {
		  console.error("播放失败:", error);
		});
	  }
	},
	{immediate: true}
);

// 组件挂载
onMounted(() => {
  audio.value = document.querySelector("audio");
  if (audio.value) {
	audio.value.volume = volume.value / 100;
  }
});
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
