<template>
  <div class="entity-info-block">
    <div class="entity-cover">
      <el-image
        :src="imageUrl || '/assets/default-cover.jpg'"
        fit="cover"
        lazy
        class="cover-image"
      />
    </div>

    <div class="entity-content">
      <div class="info-main">
        <p class="entity-type">{{ classification || '未知' }}</p>
        <h1 class="entity-title">{{ title || '加载中...' }}</h1>
        <p v-if="description" class="entity-description">{{ description }}</p>
      </div>

      <div class="info-footer">
        <div v-if="metaItems && metaItems.length > 0" class="entity-meta">
          <span v-for="(item, index) in metaItems" :key="index" class="meta-item">
            {{ item }}
          </span>
        </div>
        <div class="actions-container">
          <slot name="actions"></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  imageUrl?: string;
  classification?: string;
  title?: string;
  description?: string;
  metaItems?: string[];
}>();
</script>

<style scoped>
/* 根容器：信息区块 */
.entity-info-block {
  display: flex;
  flex-direction: column;
  gap: 20px;
  color: #fff;
  height: 100%; /* 确保在需要时能撑满父容器 */
}

/* 封面图片 */
.entity-cover {
  width: 100%;
  padding-bottom: 100%; /* 1:1 宽高比 */
  position: relative;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.6);
  flex-shrink: 0;
}

.cover-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}

.entity-cover:hover .cover-image {
  transform: scale(1.05);
}

/* 内容区域 */
.entity-content {
  display: flex;
  flex-direction: column;
  flex-grow: 1; /* 占据剩余所有空间 */
  min-height: 0;
}

.info-main {
  flex-grow: 1;
}

.entity-type {
  font-size: 12px;
  font-weight: 700;
  text-transform: uppercase;
  color: #fff;
  margin: 0 0 4px 0;
}

.entity-title {
  font-size: 32px;
  font-weight: 900;
  margin: 0 0 16px 0;
  line-height: 1.2;
  letter-spacing: -0.04em;
  word-break: break-word;
  color: #fff;
}

.entity-description {
  font-size: 14px;
  color: #b3b3b3;
  line-height: 1.6;
  margin: 0;
  white-space: pre-wrap; /* 保持文本换行 */
  word-break: break-word;
}

.entity-description:empty {
  display: none;
}

/* 底部区域 */
.info-footer {
  margin-top: 24px;
  flex-shrink: 0;
}

/* 元信息行 */
.entity-meta {
  display: flex;
  align-items: center;
  flex-wrap: wrap; /* 允许换行 */
  gap: 4px;
  font-size: 14px;
  color: #fff;
  font-weight: 500;
}

.meta-item:not(:last-child)::after {
  content: '•'; /* 使用伪元素添加分隔点 */
  margin-left: 8px;
  margin-right: 4px;
  color: #b3b3b3;
}

/* 操作按钮 */
.actions-container {
  margin-top: 20px;
}

.actions-container :deep(.el-button) {
  border-radius: 500px;
  font-weight: 700;
  border: 0;
  padding: 12px 24px;
  transition: transform 0.2s ease, background-color 0.2s ease;
}

.actions-container :deep(.el-button:hover) {
  transform: scale(1.05);
}

.actions-container :deep(.el-button--primary) {
  background-color: #1db954; /* Spotify Green */
  color: #000;
}

.actions-container :deep(.el-button--primary:hover) {
  background-color: #1ed760;
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .entity-title {
    font-size: 28px;
  }
}

@media screen and (max-width: 992px) {
  /* 在 DetailLayout 变为垂直布局时，我们也调整为水平布局以更好地利用空间 */
  .entity-info-block {
    flex-direction: row;
    align-items: center;
    gap: 24px;
  }

  .entity-cover {
    width: 200px; /* 固定宽度 */
    height: 200px;
    padding-bottom: 0;
  }

  .entity-title {
    font-size: 24px;
  }
}

@media screen and (max-width: 768px) {
  .entity-cover {
    width: 180px;
    height: 180px;
  }

  .entity-title {
    font-size: 20px;
  }

  .entity-description {
    font-size: 13px;
  }
}

@media screen and (max-width: 480px) {
  .entity-info-block {
    flex-direction: column; /* 在非常窄的屏幕上再次变回垂直 */
    align-items: flex-start;
  }

  .entity-cover {
    width: 100%;
    height: auto;
    padding-bottom: 100%;
  }

  .entity-title {
    font-size: 18px;
  }
}
</style>