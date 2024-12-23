<template>
  <el-header class="header">
    <el-row class="header-container" type="flex" justify="space-between" align="middle">
      <!-- 左侧：网站Logo和名称 -->
      <div class="logo-section">
        <el-icon class="logo-icon"><Headset /></el-icon>
        <h2 class="website-name">音乐之声</h2>
      </div>

      <!-- 中间：当前页面名称 -->
      <div class="page-title">
        <h2>{{ pageName }}</h2>
      </div>

      <!-- 右侧：用户操作区 -->
      <div class="header-actions">
        <el-button class="action-btn" @click="goToProfile">
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
      this.$router.push('/profile');
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
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 0 24px;
  height: 64px;
  display: flex;
  align-items: center;
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
  gap: 12px;
}

.logo-icon {
  font-size: 28px;
  color: #409EFF;
}

.website-name {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: #303133;
  background: linear-gradient(45deg, #409EFF, #36cfc9);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
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
  color: #606266;
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
  gap: 4px;
  height: 36px;
  padding: 0 16px;
  border-radius: 18px;
  border: 1px solid #dcdfe6;
  background-color: transparent;
  transition: all 0.3s ease;
}

.action-btn:hover {
  background-color: #f5f7fa;
  border-color: #c6e2ff;
  color: #409EFF;
}

.action-btn .el-icon {
  font-size: 16px;
}

.action-btn span {
  font-size: 14px;
}

.logout-btn {
  color: #f56c6c;
  border-color: #fbc4c4;
}

.logout-btn:hover {
  background-color: #fef0f0;
  border-color: #f56c6c;
  color: #f56c6c;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .website-name {
    display: none;
  }
  
  .action-btn span {
    display: none;
  }
  
  .action-btn {
    padding: 0 12px;
  }
}
</style>
