<template>
  <el-header class="header">
    <el-row class="header-container" type="flex" justify="space-between" align="middle">
      <!-- 左侧：网站Logo和名称 -->
      <div class="logo-section">
        <div class="logo-wrapper">
          <el-icon class="logo-icon"><Headset /></el-icon>
        </div>
        <h2 class="website-name">音乐之声</h2>
      </div>

      <!-- 中间：当前页面名称 -->
      <div class="page-title">
        <transition name="fade" mode="out-in">
          <h2 :key="pageName">{{ pageName }}</h2>
        </transition>
      </div>

      <!-- 右侧：用户操作区 -->
      <div class="header-actions">
        <el-button class="action-btn profile-btn" @click="goToProfile">
          <el-icon><UserFilled /></el-icon>
          <span>个人主页</span>
        </el-button>
        <el-button class="action-btn logout-btn" @click="logout">
          <el-icon><SwitchButton /></el-icon>
          <span>退出登录</span>
        </el-button>
      </div>
    </el-row>
  </el-header>
</template>

<script>
import { Headset, UserFilled, SwitchButton } from '@element-plus/icons-vue'

export default {
  components: {
    Headset,
    UserFilled,
    SwitchButton
  },
  methods: {
    goToProfile() {
      const userId = localStorage.getItem('userId');
      this.$router.push(`/profile/${userId}`);
    },
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('isAuthenticated');
      this.$router.push('/login');
      console.log("已退出登录");
    }
  },
  props: {
    pageName: String,
  },
};
</script>

<style scoped>
.header {
  background: rgba(255, 255, 255, 0.98);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  padding: 0 32px;
  height: 64px;
  display: flex;
  align-items: center;
  position: sticky;
  top: 0;
  z-index: 1000;
  transition: all 0.3s ease;
}

.header:hover {
  background: rgba(255, 255, 255, 1);
}

.header-container {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

/* Logo区域样式 */
.logo-section {
  display: flex;
  align-items: center;
  gap: 16px;
}

.logo-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, var(--el-color-primary), var(--el-color-primary-light-5));
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.logo-wrapper:hover {
  transform: rotate(15deg) scale(1.1);
}

.logo-icon {
  font-size: 24px;
  color: #ffffff;
}

.website-name {
  margin: 0;
  font-size: 22px;
  font-weight: 600;
  background: linear-gradient(45deg, var(--el-color-primary), var(--el-color-primary-light-3));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: 0.5px;
}

/* 页面标题样式 */
.page-title {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

.page-title h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 500;
  color: var(--el-text-color-primary);
  transition: all 0.3s ease;
}

/* 操作按钮区域样式 */
.header-actions {
  display: flex;
  gap: 16px;
  align-items: center;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  height: 38px;
  padding: 0 20px;
  border-radius: 19px;
  border: 2px solid transparent;
  background: var(--el-fill-color-light);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.profile-btn {
  color: var(--el-color-primary);
}

.profile-btn:hover {
  background: var(--el-color-primary-light-9);
  border-color: var(--el-color-primary-light-5);
  transform: translateY(-2px);
}

.logout-btn {
  color: var(--el-color-danger);
}

.logout-btn:hover {
  background: var(--el-color-danger-light-9);
  border-color: var(--el-color-danger-light-5);
  transform: translateY(-2px);
}

.action-btn .el-icon {
  font-size: 18px;
  transition: transform 0.3s ease;
}

.action-btn:hover .el-icon {
  transform: scale(1.1);
}

.action-btn span {
  font-size: 14px;
  font-weight: 500;
}

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .header {
    padding: 0 16px;
  }

  .website-name {
    display: none;
  }
  
  .action-btn {
    padding: 0 12px;
  }
  
  .action-btn span {
    display: none;
  }
  
  .page-title h2 {
    font-size: 16px;
  }
}

/* 暗色模式适配 */
@media (prefers-color-scheme: dark) {
  .header {
    background: rgba(0, 0, 0, 0.8);
  }
  
  .action-btn {
    background: rgba(255, 255, 255, 0.05);
  }
}
</style>
