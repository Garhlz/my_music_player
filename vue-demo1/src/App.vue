<template>
  <div id="app">
    <div class="app-content">
      <router-view />
    </div>

    <transition name="slide-up">
      <div v-if="showPlayer" class="global-player-bar">
        <MusicPlayer />
      </div>
    </transition>

    <div class="app-background"></div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import MusicPlayer from '@/components/MusicPlayer.vue';

const route = useRoute();

// 根据路由判断是否显示播放器
// 在登录、注册等认证页面不显示播放器
const showPlayer = computed(() => {
  const authRoutes = ['/login', '/register', '/auth'];
  const currentPath = route.path.toLowerCase();

  // 检查是否在认证相关页面
  return !authRoutes.some(authRoute => currentPath.startsWith(authRoute));
});
</script>

<style>
/* --- 1. 全局样式重置和基础设置 (保持不变) --- */
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html, body {
  width: 100%;
  height: 100%;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Helvetica Neue', 'Arial', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color-scheme: dark; /* 告知浏览器使用深色方案，优化原生控件 */
}

/* --- 2. Element Plus 主题变量覆盖 (核心改动) --- */
/* 我们在这里定义 Spotify 的颜色体系，并应用到 Element Plus 的变量上 */
:root {
  --spotify-green: #1db954;
  --spotify-green-light: #1ed760;
  --spotify-dark-bg: #121212;
  --spotify-dark-component: #282828;
  --spotify-dark-subtle: #181818;
  --spotify-text-primary: #ffffff;
  --spotify-text-secondary: #b3b3b3;
  --spotify-border-color: rgba(255, 255, 255, 0.1);
}

html.dark {
  /* 基础颜色 */
  --el-color-primary: var(--spotify-green);
  --el-bg-color: var(--spotify-dark-bg);
  --el-bg-color-overlay: var(--spotify-dark-component);
  --el-text-color-primary: var(--spotify-text-primary);
  --el-text-color-regular: var(--spotify-text-secondary);
  --el-border-color: var(--spotify-border-color);
  --el-border-color-light: rgba(255, 255, 255, 0.1);
  --el-border-color-lighter: rgba(255, 255, 255, 0.05);
  --el-fill-color-light: rgba(255, 255, 255, 0.1);

  /* 成功/警告/错误状态颜色 (确保对比度) */
  --el-color-success: var(--spotify-green);
  --el-color-success-light: rgba(29, 185, 84, 0.2);
  --el-color-warning: #ffa422;
  --el-color-danger: #f15e6c;
  --el-color-error: #f15e6c;

  /* 组件特定颜色 */
  --el-fill-color-blank: var(--spotify-dark-bg);
}

/* --- 3. 应用布局 (保持不变) --- */
#app {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  position: relative;
  background: #000;
  display: flex;
  flex-direction: column;
}

.app-background {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(29, 185, 84, 0.1) 0%, rgba(0, 0, 0, 0.9) 30%, #000 70%, rgba(29, 185, 84, 0.05) 100%);
  pointer-events: none;
  z-index: -1;
}

.app-content {
  flex-grow: 1;
  min-height: 0;
  position: relative;
  z-index: 1;
}

.global-player-bar {
  position: relative;
  flex-shrink: 0;
  height: 80px;
  z-index: 2000;
  /* 播放器背景使用更深的颜色 */
  background: var(--spotify-dark-subtle);
  border-top: 1px solid var(--el-border-color);
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-up-enter-from,
.slide-up-leave-to {
  transform: translateY(100%);
  opacity: 0;
}

/* --- 4. 精细化样式微调 (移除 !important) --- */

/* 滚动条全局样式 (保持不变) */
::-webkit-scrollbar {
  width: 12px;
}

::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 6px;
}

::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 6px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}

/* 文本选择样式 (保持不变) */
::selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}

/* [THE FIX] 消息提示 (ElMessage) 的优雅覆盖 */
.el-message {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  /* 我们只定义基础容器样式，不强制定义文字颜色 */
  background-color: var(--el-bg-color-overlay);
  border: 1px solid var(--el-border-color);
}

.el-message .el-message__content {
  color: var(--el-text-color-primary); /* 默认文字为白色 */
}

/* 针对不同类型的提示，分别定义背景和文字颜色，确保对比度 */
.el-message--success {
  background-color: var(--el-color-success);
  border-color: var(--el-color-success);
}

.el-message--success .el-message__content {
  color: #000; /* 绿色背景配黑色文字，像 Spotify 一样 */
}

.el-message--warning {
  background-color: var(--el-color-warning);
  border-color: var(--el-color-warning);
}

.el-message--warning .el-message__content {
  color: #000; /* 黄色背景配黑色文字 */
}

.el-message--error, .el-message--danger {
  background-color: var(--el-color-error);
  border-color: var(--el-color-error);
}

.el-message--error .el-message__content, .el-message--danger .el-message__content {
  color: var(--el-text-color-primary); /* 红色背景配白色文字 */
}

/* 下拉菜单选中项 */
.el-select-dropdown__item.selected {
  color: var(--el-color-primary) !important;
  font-weight: 700;
}

/* 输入框聚焦效果 */
.el-input__wrapper.is-focus {
  box-shadow: 0 0 0 1px var(--el-color-primary) inset !important;
}

.el-textarea__inner:focus {
  box-shadow: 0 0 0 1px var(--el-color-primary) inset !important;
}

/* 响应式调整播放器高度 (保持不变) */
@media screen and (max-width: 768px) {
  .global-player-bar {
    height: 72px;
  }
}

@media screen and (max-width: 480px) {
  .global-player-bar {
    height: 64px;
  }
}
</style>