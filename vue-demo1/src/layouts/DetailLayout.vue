<template>
  <div class="detail-layout-container" v-loading="loading">
    <aside class="detail-aside">
      <div class="aside-content">
        <slot name="info"></slot>
      </div>
    </aside>

    <main class="detail-main">
      <div class="main-content">
        <slot name="list"></slot>
      </div>
    </main>

    <!-- 背景装饰 -->
    <div class="detail-background"></div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  loading?: boolean;
}>();
</script>

<style scoped>
.detail-layout-container {
  display: flex;
  gap: 24px;
  height: 100%;
  padding: 24px;
  width: 100%;
  box-sizing: border-box;
  position: relative;
  background: transparent;
  overflow: hidden;
}

/* 背景装饰 */
.detail-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 400px;
  background: linear-gradient(
    180deg,
    rgba(29, 185, 84, 0.2) 0%,
    rgba(29, 185, 84, 0.1) 40%,
    transparent 100%
  );
  pointer-events: none;
  z-index: 0;
}

.detail-aside {
  flex: 0 0 360px;
  height: 100%;
  position: relative;
  z-index: 1;
}

.aside-content {
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  position: relative;
}

/* 左侧内容区域滚动条样式 */
.aside-content::-webkit-scrollbar {
  width: 8px;
}

.aside-content::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 4px;
}

.aside-content::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 4px;
  transition: background 0.2s ease;
}

.aside-content::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}

.detail-main {
  flex: 1;
  min-width: 0;
  height: 100%;
  position: relative;
  z-index: 1;
}

.main-content {
  height: 100%;
  background: rgba(18, 18, 18, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4),
  inset 0 1px 0 rgba(255, 255, 255, 0.05);
  overflow-y: auto;
  overflow-x: hidden;
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
  opacity: 0.8;
}

/* 主内容区域滚动条样式 */
.main-content::-webkit-scrollbar {
  width: 12px;
}

.main-content::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 6px;
}

.main-content::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 6px;
  transition: background 0.2s ease;
}

.main-content::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}

.main-content::-webkit-scrollbar-corner {
  background: transparent;
}

/* 为子组件提供滚动条样式 */
.main-content :deep(*::-webkit-scrollbar) {
  width: 12px;
}

.main-content :deep(*::-webkit-scrollbar-track) {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 6px;
}

.main-content :deep(*::-webkit-scrollbar-thumb) {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 6px;
  transition: background 0.2s ease;
}

.main-content :deep(*::-webkit-scrollbar-thumb:hover) {
  background: rgba(255, 255, 255, 0.3);
}

.main-content :deep(*::-webkit-scrollbar-corner) {
  background: transparent;
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .detail-aside {
    flex: 0 0 320px;
  }

  .detail-layout-container {
    gap: 20px;
  }
}

@media screen and (max-width: 992px) {
  .detail-layout-container {
    flex-direction: column;
    gap: 16px;
    padding: 16px;
  }

  .detail-aside {
    flex: none;
    height: auto;
  }

  .detail-main {
    flex: 1;
    height: calc(100vh - 400px);
    min-height: 400px;
  }
}

@media screen and (max-width: 768px) {
  .detail-layout-container {
    gap: 12px;
    padding: 12px;
  }

  .main-content {
    border-radius: 8px;
  }

  .detail-main {
    height: calc(100vh - 350px);
    min-height: 350px;
  }
}

@media screen and (max-width: 480px) {
  .detail-layout-container {
    gap: 8px;
    padding: 8px;
  }

  .detail-aside {
    flex: 0 0 auto;
  }
}

/* 加载状态样式覆盖 */
:deep(.el-loading-mask) {
  background-color: rgba(0, 0, 0, 0.8);
  border-radius: 12px;
}

:deep(.el-loading-spinner) {
  color: #1db954;
}

:deep(.el-loading-text) {
  color: #fff;
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .main-content {
    border: 1px solid rgba(255, 255, 255, 0.2);
  }

  .detail-background {
    display: none;
  }
}

/* 减少动画的用户偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .main-content {
    backdrop-filter: none;
  }

  .detail-background {
    display: none;
  }

  * {
    transition: none !important;
  }
}

/* 深色主题下的文本选择效果 */
.detail-layout-container ::selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}

.detail-layout-container ::-moz-selection {
  background: rgba(29, 185, 84, 0.3);
  color: #fff;
}
</style>