<template>
  <div class="common-layout">
    <div class="fixed-header">
      <Banner :page-name="pageName" />
    </div>

    <div class="middle-content">
      <div class="fixed-aside">
        <Navbar />
      </div>
      <div class="main-content" :class="{ 'is-scrollable': mainContentScrollable }">
        <slot name="main"></slot>
      </div>
    </div>

    <div class="fixed-footer">
      <MusicPlayer />
    </div>
  </div>
</template>

<script setup lang="ts">
import Banner from '@/components/Banner.vue';
import Navbar from '@/components/Navbar.vue';
import MusicPlayer from '@/components/MusicPlayer.vue';

// 只用调用一次即可, 否则重复
withDefaults(defineProps<{
  pageName: string;
  mainContentScrollable?: boolean;
}>(), {
  mainContentScrollable: true, // 默认值为 true
});
// 潜在新功能（未实现）:
// - 动态主题切换：通过 CSS 变量切换主题颜色。
// - 响应式侧边栏：移动端支持折叠侧边栏，释放主内容空间。
</script>

<style scoped>
/* 替换 CommonLayout.vue 的 <style scoped> */

.common-layout {
  display: flex;
  flex-direction: column; /* 垂直排列：header, 中间区域, footer */
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  background-color: #f5f7fa;
}

/* 头部和底部样式保持基本不变，但不再需要 fixed 定位 */
.fixed-header,
.fixed-footer {
  flex-shrink: 0; /* 防止被压缩 */
  background-color: #fff;
  z-index: 1000;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.fixed-header {
  height: 64px;
  border-radius: 0 0 12px 12px;
}

.fixed-footer {
  height: 80px;
  border-radius: 12px 12px 0 0;
}

/* 中间内容区域，使用 Flexbox 实现侧边栏+主内容的布局 */
.middle-content {
  flex-grow: 1; /* 占据剩余的所有垂直空间 */
  display: flex;
  overflow: hidden; /* 防止内容溢出布局 */
}

.fixed-aside {
  flex-shrink: 0; /* 防止被压缩 */
  width: 240px;
  background-color: #fff;
  z-index: 900;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  border-radius: 0 12px 12px 0;
  /* 这里可以添加 overflow-y: auto; 如果侧边栏内容很多需要滚动 */
}

/* 主内容区现在只负责滚动自己的内容 */
.main-content {
  flex-grow: 1;
  height: 100%;
  overflow: hidden; /* 默认状态是不可滚动的 */
}

/* 当 .main-content 同时拥有 .is-scrollable 这个 class 时，它才会有滚动条 */
.main-content.is-scrollable {
  overflow-y: auto;
}


/* 响应式设计 */
@media screen and (max-width: 768px) {
  .fixed-aside {
    width: 80px;
  }

  .fixed-header,
  .fixed-footer {
    border-radius: 0;
  }
}
</style>
