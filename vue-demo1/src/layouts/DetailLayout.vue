<template>
  <div class="detail-layout-container" v-loading="loading">
    <aside class="detail-aside">
      <slot name="info"></slot>
    </aside>

    <main class="detail-main">
      <slot name="list"></slot>
    </main>
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
  box-sizing: border-box; /* 确保 padding 不会撑大容器 */
}

.detail-aside {
  flex: 0 0 320px; /* 左侧固定宽度，可以根据喜好调整 */
  height: 100%;
}

.detail-main {
  flex: 1; /* 右侧占据所有剩余空间 */
  min-width: 0; /* 防止内容过长时撑破布局 */
  height: 100%;
  overflow-y: auto;
}

/* 针对 .detail-main 的滚动条的美化 */
.detail-main::-webkit-scrollbar {
  width: 6px;
}

.detail-main::-webkit-scrollbar-thumb {
  background-color: #dcdfe6;
  border-radius: 3px;
}

.detail-main::-webkit-scrollbar-track {
  background-color: #f5f7fa;
}
</style>