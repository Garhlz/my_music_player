<template>
  <div class="common-layout">
    <!-- 顶部固定头部 -->
    <div class="fixed-header">
      <Banner :page-name="pageName" />
    </div>

    <!-- 主要内容区域 -->
    <div class="middle-content">
      <!-- 左侧导航栏 -->
      <div class="fixed-aside">
        <div class="sidebar-container">
          <Navbar />
        </div>
      </div>
      <!-- 主内容区 -->
      <div class="main-content-wrapper">
        <div class="main-content" :class="{ 'is-scrollable': mainContentScrollable }">
          <slot name="main"></slot>
        </div>
      </div>
    </div>

    <!-- 背景装饰效果 -->
    <div class="background-gradient"></div>
  </div>
</template>

<script setup lang="ts">
import Banner from '@/components/Banner.vue';
import Navbar from '@/components/Navbar.vue';

withDefaults(defineProps<{
  pageName: string;
  mainContentScrollable?: boolean;
}>(), {
  mainContentScrollable: true,
});
</script>

<style scoped>
/* Spotify风格的主布局容器 */
.common-layout {
  display: flex;
  flex-direction: column;
  /* [THE FIX] 高度设为 100%，以填满其父容器 (.app-content) */
  height: 100%;
  width: 100%;
  overflow: hidden;
  background: transparent; /* 背景由 App.vue 控制 */
  position: relative;
}

/* 背景渐变装饰 */
.background-gradient {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 300px;
  background: linear-gradient(
    180deg,
    rgba(29, 185, 84, 0.15) 0%,
    rgba(29, 185, 84, 0.08) 50%,
    transparent 100%
  );
  pointer-events: none;
  z-index: 0;
}

/* 顶部固定头部 */
.fixed-header {
  flex-shrink: 0;
  height: 64px;
  background: rgba(18, 18, 18, 0.95);
  backdrop-filter: blur(20px);
  z-index: 1000;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  position: relative;
}

/* 中间内容区域 */
.middle-content {
  flex-grow: 1;
  display: flex;
  overflow: hidden;
  min-height: 0;
  gap: 8px;
  padding: 8px;
  position: relative;
  z-index: 1;
}


/* 左侧导航栏 */
.fixed-aside {
  flex-shrink: 0;
  width: 280px;
  position: relative;
  z-index: 900;
}

.sidebar-container {
  height: 100%;
  background: #121212;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3),
  inset 0 1px 0 rgba(255, 255, 255, 0.05);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.sidebar-container:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.4),
  inset 0 1px 0 rgba(255, 255, 255, 0.08);
}

/* 主内容区包装器 */
.main-content-wrapper {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  min-width: 0;
  position: relative;
}

/* 主内容区 */
.main-content {
  flex-grow: 1;
  background: #121212;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3),
  inset 0 1px 0 rgba(255, 255, 255, 0.05);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

.main-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(29, 185, 84, 0.5),
    transparent
  );
  opacity: 0;
  transition: opacity 0.3s ease;
}

.main-content:hover::before {
  opacity: 1;
}

.main-content.is-scrollable {
  overflow-y: auto;
  overflow-x: hidden;
}

/* 自定义滚动条样式 */
.main-content.is-scrollable::-webkit-scrollbar {
  width: 12px;
}

.main-content.is-scrollable::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 6px;
}

.main-content.is-scrollable::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 6px;
  transition: background 0.2s ease;
}

.main-content.is-scrollable::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}

.main-content.is-scrollable::-webkit-scrollbar-corner {
  background: transparent;
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .fixed-aside {
    width: 240px;
  }
}

@media screen and (max-width: 992px) {
  .fixed-aside {
    width: 200px;
  }

  .middle-content {
    gap: 6px;
    padding: 6px;
  }
}

@media screen and (max-width: 768px) {
  .middle-content {
    gap: 4px;
    padding: 4px;
  }

  .fixed-aside {
    width: 72px;
  }

  .sidebar-container {
    border-radius: 6px;
  }

  .main-content {
    border-radius: 6px;
  }

  .fixed-header {
    height: 56px;
  }

  .player-spacer {
    height: 76px; /* 移动端播放器稍矮一些 */
  }
}

@media screen and (max-width: 480px) {
  .middle-content {
    gap: 2px;
    padding: 2px;
  }

  .fixed-aside {
    width: 64px;
  }

  .player-spacer {
    height: 70px;
  }
}

/* 加载动画效果 */
@keyframes fadeInScale {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

.sidebar-container,
.main-content {
  animation: fadeInScale 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.sidebar-container {
  animation-delay: 0.1s;
}

.main-content {
  animation-delay: 0.2s;
}

/* 暗色模式下的聚焦效果 */
.main-content:focus-within {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.4),
  0 0 0 1px rgba(29, 185, 84, 0.3),
  inset 0 1px 0 rgba(255, 255, 255, 0.08);
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .sidebar-container,
  .main-content {
    border: 1px solid rgba(255, 255, 255, 0.2);
  }

  .fixed-header {
    border-bottom: 2px solid rgba(255, 255, 255, 0.3);
  }
}

/* 减少动画的用户偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .sidebar-container,
  .main-content {
    animation: none;
    transition: none;
  }

  .background-gradient {
    display: none;
  }
}

/* 深色主题下的文本选择效果 */
::selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}

::-moz-selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}
</style>