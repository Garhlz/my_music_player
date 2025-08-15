// src/stores/like.ts

import { defineStore } from 'pinia';
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { likeApi } from '@/api';
import type { ModelsSongDetailDTO } from '@/api-client';

export const useLikeStore = defineStore('like', () => {
  // --- State ---
  const likedSongIds = ref(new Set<number>());

  // --- Getters ---
  const isLiked = (songId: number | undefined): boolean => {
    if (!songId) return false;
    return likedSongIds.value.has(songId);
  };

  // --- Actions ---
  async function fetchLikedSongs() {
    try {
      const response = await likeApi.meLikedSongsGet(1, 1000);
      const songs = (response.data.list as ModelsSongDetailDTO[]) || [];
      likedSongIds.value = new Set(songs.map(song => song.id!));
    } catch (error) {
      console.error('获取喜欢列表失败:', error);
    }
  }

  async function likeSong(songId: number) {
    if (likedSongIds.value.has(songId)) return;
    likedSongIds.value.add(songId);
    try {
      await likeApi.meLikedSongsSongIdPost(songId);
    } catch (error) {
      likedSongIds.value.delete(songId);
      ElMessage.error('添加喜欢失败');
    }
  }

  async function unlikeSong(songId: number) {
    if (!likedSongIds.value.has(songId)) return;
    likedSongIds.value.delete(songId);
    try {
      await likeApi.meLikedSongsSongIdDelete(songId);
    } catch (error) {
      likedSongIds.value.add(songId);
      ElMessage.error('取消喜欢失败');
    }
  }

  function clearLikes() {
    likedSongIds.value.clear();
  }

  return {
    likedSongIds,
    isLiked,
    fetchLikedSongs,
    likeSong,
    unlikeSong,
    clearLikes,
  };
}, {
  // [THE FIX] 自定义持久化策略
  persist: {
    // 告诉插件如何序列化和反序列化我们的状态
    serializer: {
      // serialize: 当状态需要被保存时调用
      serialize: (state) => {
        // 将 Set 转换为 Array，然后再转换为 JSON 字符串
        const setToArray = Array.from(state.likedSongIds);
        return JSON.stringify(setToArray);
      },
      // deserialize: 当页面加载，需要恢复状态时调用
      deserialize: (storedState) => {
        // 将 JSON 字符串转换回 Array，然后再转换回 Set
        const arrayToSet = new Set<number>(JSON.parse(storedState));
        return {
          likedSongIds: arrayToSet,
        };
      },
    },
  },
});