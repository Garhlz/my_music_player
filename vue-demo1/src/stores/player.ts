// @/stores/player.ts

import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import type { ModelsSongDetailDTO } from '@/api-client';
// import { songApi } from '@/api'; // 暂时注释，如果需要动态获取URL再打开

// 定义播放模式的类型
export type PlaybackMode = 'list-loop' | 'single-loop' | 'shuffle';

export const usePlayerStore = defineStore('player', () => {
// --- State ---
  const playlist = ref<ModelsSongDetailDTO[]>([]);
  const originalPlaylist = ref<ModelsSongDetailDTO[]>([]);
  const currentIndex = ref(-1);
  const isPlaying = ref(false);
  const playbackMode = ref<PlaybackMode>('list-loop');

  const currentTime = ref(0);
  const duration = ref(0);

// --- Getters ---
  const currentSong = computed(() => playlist.value[currentIndex.value] || null);
  const getIsPlaying = computed(() => isPlaying.value); // 为Vue组件提供一个标准的getter

// --- Actions ---

  const setPlaylist = (songs: ModelsSongDetailDTO[], startIndex: number = 0) => {
    // 重置播放状态
    currentTime.value = 0;
    duration.value = 0;
    originalPlaylist.value = [...songs];
    if (playbackMode.value === 'shuffle') {
// 如果当前是随机模式，则直接打乱新列表
      playlist.value = getShuffledList([...songs]);
// 找到开始播放的歌曲在新列表中的位置
      const startSongId = songs[startIndex]?.id;
      const newIndex = playlist.value.findIndex(s => s.id === startSongId);
      currentIndex.value = newIndex !== -1 ? newIndex : 0;
    } else {
      playlist.value = [...songs];
      currentIndex.value = startIndex;
    }
    isPlaying.value = true;
  };

  const play = () => {
    if (currentSong.value) isPlaying.value = true;
  };
  const pause = () => {
    isPlaying.value = false;
  };
// **[FIXED] 增加 togglePlay action**
  const togglePlay = () => {
    if (currentSong.value) {
      isPlaying.value = !isPlaying.value;
    }
  };

  const next = () => {
    if (playlist.value.length < 2) return;

    // 切歌前重置进度
    currentTime.value = 0;
    duration.value = 0;

    let nextIndex = currentIndex.value + 1;
    if (nextIndex >= playlist.value.length) {
      nextIndex = 0;
    }
    currentIndex.value = nextIndex;
    isPlaying.value = true;
  };

  const previous = () => {
    if (playlist.value.length < 2) return;

    // 切歌前重置进度
    currentTime.value = 0;
    duration.value = 0;

    let prevIndex = currentIndex.value - 1;
    if (prevIndex < 0) {
      prevIndex = playlist.value.length - 1;
    }
    currentIndex.value = prevIndex;
    isPlaying.value = true;
  };

// **[FIXED] 增加切换播放模式的 action**
  const togglePlaybackMode = () => {
    const modes: PlaybackMode[] = ['list-loop', 'single-loop', 'shuffle'];
    const currentModeIndex = modes.indexOf(playbackMode.value);
    const nextModeIndex = (currentModeIndex + 1) % modes.length;
    const newMode = modes[nextModeIndex];
    playbackMode.value = newMode;

    if (newMode === 'shuffle') {
      const currentId = currentSong.value?.id;
      playlist.value = getShuffledList([...originalPlaylist.value]);
      // 保持当前歌曲不变
      const newIndex = playlist.value.findIndex(s => s.id === currentId);
      if (newIndex !== -1) currentIndex.value = newIndex;
    } else {
      // 从随机切回时，恢复原始列表顺序
      const currentId = currentSong.value?.id;
      playlist.value = [...originalPlaylist.value];
      const newIndex = playlist.value.findIndex(s => s.id === currentId);
      if (newIndex !== -1) currentIndex.value = newIndex;
    }

  };

// 内部辅助函数，用于打乱数组
  const getShuffledList = (list: ModelsSongDetailDTO[]) => {
    const shuffled = [...list];
    for (let i = shuffled.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]];
    }
    return shuffled;
  };

  /**

   * 动态获取歌曲的播放 URL
   * 这是连接前端和后端（OSS）的关键
   */
  const fetchAndSetSongUrl = async (song: ModelsSongDetailDTO): Promise<string | null> => {
    if (!song || !song.id) return null;
    // 如果歌曲对象上已经有URL，直接返回，避免重复请求
    if (song.url) return song.url;

    try {
      // 假设你有一个 songApi.songsIdPlayInfoGet 的接口
      const response = await songApi.songsIdPlayInfoGet(song.id);
      const url = response.data.url;
      if (url) {
        // 将获取到的 URL 动态添加到歌曲对象上，作为缓存
        song.url = url;
        return url;
      }
      return null;
    } catch (error) {
      console.error('获取歌曲URL失败:', error);
      return null;
    }
  };

  const setAudioCurrentTime = (time: number) => {
    currentTime.value = time;
  };
  const setAudioDuration = (d: number) => {
    duration.value = d;
  };

  return {
    playlist,
    currentIndex,
    isPlaying,
    getIsPlaying,
    playbackMode,
    currentSong,
    currentTime,
    duration,
    setPlaylist,
    play,
    pause,
    togglePlay, // **[FIXED] 导出 togglePlay**
    next,
    previous,
    togglePlaybackMode, // **[FIXED] 导出新 action**
    setAudioCurrentTime,
    setAudioDuration,
  };
});