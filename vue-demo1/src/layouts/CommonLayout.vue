<template>
  <div class="common-layout">
    <!-- 头部区域 -->
    <div class="fixed-header">
      <Banner :page-name="pageName" />
    </div>

    <!-- 侧边栏区域 -->
    <div class="fixed-aside">
      <Navbar />
    </div>

    <!-- 主内容区域 -->
    <div class="main-content">
      <div class="main-content-inner">
        <slot name="main"></slot>
      </div>
    </div>

    <!-- 底部区域 -->
    <div class="fixed-footer">
      <MusicPlayer />
    </div>
  </div>
</template>

<script setup lang="ts">
import Banner from '@/components/Banner.vue';
import Navbar from '@/components/Navbar.vue';
import MusicPlayer from '@/components/MusicPlayer.vue';

defineProps<{
  pageName: string;
}>();

// 潜在新功能（未实现）:
// - 动态主题切换：通过 CSS 变量切换主题颜色。
// - 响应式侧边栏：移动端支持折叠侧边栏，释放主内容空间。
</script>

<style scoped>
.common-layout {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  position: relative;
  background-color: #f5f7fa;
}

/* 头部样式 */
.fixed-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 64px;
  background-color: #fff;
  z-index: 1000;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border-radius: 0 0 12px 12px;
}

/* 侧边栏样式 */
.fixed-aside {
  position: fixed;
  top: 64px;
  left: 0;
  bottom: 80px;
  width: 240px;
  background-color: #fff;
  z-index: 900;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  border-radius: 0 12px 12px 0;
}

/* 主内容区域样式 */
.main-content {
  position: fixed;
  top: 64px;
  left: 240px;
  right: 0;
  bottom: 80px;
  overflow: hidden;
}

.main-content-inner {
  height: 100%;
  overflow-y: hidden;
  padding-bottom: 80px; /* 为分页器预留空间 */
}

/* 底部播放器样式 */
.fixed-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 80px;
  background-color: #fff;
  z-index: 1000;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.1);
  border-radius: 12px 12px 0 0;
}

/* 滚动条样式 */
.main-content-inner::-webkit-scrollbar {
  width: 6px;
}

.main-content-inner::-webkit-scrollbar-thumb {
  background-color: #dcdfe6;
  border-radius: 3px;
}

.main-content-inner::-webkit-scrollbar-track {
  background-color: #f5f7fa;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .fixed-aside {
    width: 80px;
    border-radius: 0 8px 8px 0;
  }

  .main-content {
    left: 80px;
  }

  .fixed-header,
  .fixed-footer {
    border-radius: 0;
  }
}
</style>
