<template>
  <div class="entity-info-card">
    <div class="entity-cover">
      <el-image :src="imageUrl || '/assets/default-cover.jpg'" fit="cover" lazy />
    </div>

    <h2 class="entity-title">{{ title || '加载中...' }}</h2>

    <div class="entity-meta">
      <span v-for="(item, index) in metaItems" :key="index">{{ item }}</span>
    </div>

    <p class="entity-description">{{ description || '暂无描述' }}</p>

    <div class="actions-container">
      <slot name="actions"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  imageUrl?: string;
  title?: string;
  description?: string;
  metaItems?: string[]; // 用于显示 "xx首歌曲", "xx张专辑" 等信息
}>();
</script>

<style scoped>
.entity-info-card {
  padding: 24px;
  background: var(--el-bg-color);
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  height: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
}

.entity-info-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
}

.entity-cover {
  width: 100%;
  padding-bottom: 100%; /* 实现1:1的宽高比 */
  position: relative;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.entity-cover .el-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s ease;
}

.entity-cover:hover .el-image {
  transform: scale(1.1);
}

.entity-title {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  text-align: center;
  word-break: break-all;
}

.entity-meta {
  margin: 12px 0;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  text-align: center;
  display: flex;
  justify-content: center;
  gap: 12px;
}

.entity-description {
  flex-grow: 1;
  color: var(--el-text-color-regular);
  font-size: 14px;
  line-height: 1.7;
  margin: 16px 0 0;
  padding-top: 16px;
  border-top: 1px solid var(--el-border-color-lighter);
  white-space: pre-wrap; /* 保留换行符 */
  overflow-y: auto; /* 当描述过长时滚动 */
}

.actions-container {
  margin-top: 20px;
  text-align: center;
}
</style>