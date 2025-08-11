// src/stores/user.ts

import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

// defineStore 的第一个参数是 store 的唯一 ID，Pinia 会用它来连接 devtools
// 这个 ID 必须是独一无二的，通常用文件名即可
export const useUserStore = defineStore('user', () => {
  // --- State ---
  // 就像组件里的 data，用来存放核心数据。我们用 ref() 来创建。
  // 我们从 localStorage 读取初始值，这样用户刷新页面后登录状态不会丢失
  const token = ref < string | null > (localStorage.getItem('token'));
  const userId = ref < number | null > (JSON.parse(localStorage.getItem('userId') || 'null'));

  // --- Getters ---
  // 就像组件里的 computed，用来派生出一些状态。
  const isLoggedIn = computed(() => !!token.value && !!userId.value);

  // --- Actions ---
  // 就像组件里的 methods，用来封装业务逻辑，修改 state。
  // Action 可以是异步的，你可以在里面直接使用 await。

  /**
   * 设置用户凭证信息，通常在登录成功后调用
   * @param newToken 新的 token
   * @param newUserId 新的用户 ID
   */
  function setCredentials(newToken: string, newUserId: number) {
    token.value = newToken;
    userId.value = newUserId;

    // 同时持久化到 localStorage
    localStorage.setItem('token', newToken);
    localStorage.setItem('userId', JSON.stringify(newUserId));
  }

  /**
   * 清除用户凭证，用于退出登录
   */
  function logout() {
    token.value = null;
    userId.value = null;

    // 从 localStorage 中移除
    localStorage.removeItem('token');
    localStorage.removeItem('userId');

    // 你可能还想清除其他数据，比如播放列表等
    // location.reload() // 可以选择强制刷新页面回到初始状态
  }

  // 必须返回所有需要暴露给外部使用的 state, getters, 和 actions
  return {
    // State
    token,
    userId,
    // Getters
    isLoggedIn,
    // Actions
    setCredentials,
    logout,
  };
});