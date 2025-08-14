<template>
  <el-header class="spotify-header">
    <div class="header-container">
      <!-- 左侧：网站Logo和名称 -->
      <div class="logo-section">
        <div class="logo-wrapper">
          <img src="/icons/music.png" alt="网站Logo" class="site-logo-image" />
        </div>
        <h2 class="website-name">Elaine's Music</h2>
      </div>

      <!-- 中间：当前页面名称 -->
      <div class="page-title-section">
        <transition name="slide-fade" mode="out-in">
          <h2 :key="pageName" class="page-title">{{ pageName }}</h2>
        </transition>
      </div>

      <!-- 右侧：用户操作区 -->
      <div class="header-actions">
        <el-button class="action-btn profile-btn" @click="goToProfile">
          <el-icon>
            <UserFilled />
          </el-icon>
          <span class="btn-text">个人主页</span>
        </el-button>
        <el-button class="action-btn logout-btn" @click="logout">
          <el-icon>
            <SwitchButton />
          </el-icon>
          <span class="btn-text">退出登录</span>
        </el-button>
      </div>
    </div>

    <!-- 底部渐变装饰 -->
    <div class="header-gradient"></div>
  </el-header>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { UserFilled, SwitchButton } from '@element-plus/icons-vue';
import { storeToRefs } from 'pinia';

defineProps<{
  pageName: string;
}>();

const router = useRouter();
const userStore = useUserStore();
const { userId: currentUserId } = storeToRefs(userStore);
const goToProfile = () => {
  router.push(`/profile/${Number(currentUserId.value)}`);
};

const logout = () => {
  userStore.logout();
  router.push('/auth');
  console.log('已退出登录');
};
</script>

<style scoped>
/* Spotify风格的头部 */
.spotify-header {
  background: rgba(18, 18, 18, 0.95);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: none;
  box-shadow: none;
  padding: 0;
  height: 64px;
  position: relative;
  overflow: hidden;
}

.header-container {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
  position: relative;
  z-index: 2;
}

/* 底部装饰渐变 */
.header-gradient {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(29, 185, 84, 0.6),
    transparent
  );
  opacity: 0.8;
}

/* Logo区域样式 */
.logo-section {
  display: flex;
  align-items: center;
  gap: 16px;
  min-width: 0;
  flex-shrink: 0;
}

.logo-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #1db954, #1ed760);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 4px 12px rgba(29, 185, 84, 0.3);
  position: relative;
}

.logo-wrapper::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.logo-wrapper:hover::before {
  opacity: 1;
}

.logo-wrapper:hover {
  transform: scale(1.1) rotate(5deg);
  box-shadow: 0 8px 24px rgba(29, 185, 84, 0.5);
}

.site-logo-image {
  width: 24px;
  height: 24px;
  object-fit: contain;
  filter: brightness(0) invert(1);
  transition: transform 0.3s ease;
}

.logo-wrapper:hover .site-logo-image {
  transform: scale(1.1);
}

.website-name {
  margin: 0;
  font-size: 22px;
  font-weight: 700;
  background: #fff;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  white-space: nowrap;
  letter-spacing: -0.5px;
}

/* 页面标题样式 */
.page-title-section {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  text-align: center;
}

.page-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #fff;
  white-space: nowrap;
  letter-spacing: 0.2px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

/* 操作按钮区域样式 */
.header-actions {
  display: flex;
  gap: 12px;
  align-items: center;
  flex-shrink: 0;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  height: 36px;
  padding: 0 16px;
  border-radius: 18px;
  border: none;
  font-weight: 600;
  font-size: 14px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.action-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  transition: left 0.5s ease;
}

.action-btn:hover::before {
  left: 100%;
}

.profile-btn {
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  backdrop-filter: blur(10px);
}

.profile-btn:hover {
  background: rgba(29, 185, 84, 0.2);
  color: #1ed760;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(29, 185, 84, 0.3);
}

.logout-btn {
  background: rgba(255, 255, 255, 0.05);
  color: #b3b3b3;
  backdrop-filter: blur(10px);
}

.logout-btn:hover {
  background: rgba(239, 68, 68, 0.2);
  color: #f87171;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(239, 68, 68, 0.3);
}

.action-btn .el-icon {
  font-size: 16px;
  transition: transform 0.3s ease;
}

.action-btn:hover .el-icon {
  transform: scale(1.1);
}

.btn-text {
  white-space: nowrap;
}

/* 过渡动画 */
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-fade-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.slide-fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .header-container {
    padding: 0 24px;
  }
}

@media screen and (max-width: 992px) {
  .header-container {
    padding: 0 20px;
  }

  .website-name {
    font-size: 20px;
  }

  .page-title {
    font-size: 16px;
  }
}

@media screen and (max-width: 768px) {
  .header-container {
    padding: 0 16px;
  }

  .website-name {
    display: none;
  }

  .logo-section {
    gap: 8px;
  }

  .action-btn {
    padding: 0 12px;
    height: 32px;
  }

  .btn-text {
    display: none;
  }

  .action-btn .el-icon {
    font-size: 18px;
  }

  .page-title {
    font-size: 16px;
  }
}

@media screen and (max-width: 480px) {
  .header-container {
    padding: 0 12px;
  }

  .header-actions {
    gap: 8px;
  }

  .action-btn {
    width: 32px;
    height: 32px;
    padding: 0;
    border-radius: 16px;
  }

  .page-title-section {
    max-width: 200px;
  }

  .page-title {
    font-size: 14px;
    overflow: hidden;
    text-overflow: ellipsis;
  }
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .spotify-header {
    background: rgba(0, 0, 0, 1);
    border-bottom: 2px solid rgba(29, 185, 84, 0.8);
  }

  .action-btn {
    border: 1px solid rgba(255, 255, 255, 0.3);
  }
}

/* 减少动画偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .logo-wrapper,
  .action-btn,
  .slide-fade-enter-active,
  .slide-fade-leave-active {
    transition: none;
  }

  .logo-wrapper:hover {
    transform: none;
  }

  .action-btn:hover {
    transform: none;
  }
}</style>