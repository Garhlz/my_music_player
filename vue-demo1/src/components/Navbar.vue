<template>
  <div class="spotify-navbar">
    <!-- 导航菜单 -->
    <el-menu
      class="sidebar-menu"
      :default-active="activeRoute"
      background-color="transparent"
      text-color="#b3b3b3"
      active-text-color="#1db954"
      :unique-opened="true"
    >
      <!-- 主要功能区 -->
      <div class="menu-section main-section">
        <div class="section-header">
          <h3 class="section-title">
            <el-icon>
              <House />
            </el-icon>
            主页
          </h3>
        </div>

        <el-menu-item index="/pub" @click="navigateTo('/pub')" class="menu-item">
          <div class="menu-item-content">
            <el-icon class="menu-icon">
              <List />
            </el-icon>
            <span class="menu-text">音乐库</span>
          </div>
          <div class="item-indicator"></div>
        </el-menu-item>

        <el-menu-item index="/love" @click="navigateTo(`/love/${userId}`)" class="menu-item">
          <div class="menu-item-content">
            <el-icon class="menu-icon">
              <Star />
            </el-icon>
            <span class="menu-text">我喜欢的音乐</span>
          </div>
          <div class="item-indicator"></div>
        </el-menu-item>

        <el-menu-item index="/uploaded" @click="navigateTo('/uploaded')" class="menu-item">
          <div class="menu-item-content">
            <el-icon class="menu-icon">
              <Upload />
            </el-icon>
            <span class="menu-text">上传的音乐</span>
          </div>
          <div class="item-indicator"></div>
        </el-menu-item>
      </div>

      <!-- 歌单区域 -->
      <div class="menu-section playlist-section">
        <div class="section-header">
          <h3 class="section-title">
            <el-icon>
              <Menu />
            </el-icon>
            我的歌单
          </h3>
          <el-tooltip content="创建歌单" placement="right">
            <el-button class="create-playlist-btn" text circle @click="navigateTo(`/playlist-grid/${userId}`)">
              <el-icon>
                <Plus />
              </el-icon>
            </el-button>
          </el-tooltip>
        </div>

        <el-menu-item index="/playlist-grid" @click="navigateTo(`/playlist-grid/${userId}`)" class="menu-item">
          <div class="menu-item-content">
            <el-icon class="menu-icon">
              <Menu />
            </el-icon>
            <span class="menu-text">所有歌单</span>
          </div>
          <div class="item-indicator"></div>
        </el-menu-item>

        <!-- 最近播放的歌单可以在这里动态添加 -->
        <div class="recent-playlists">
          <!-- 这里可以动态渲染最近的歌单 -->
        </div>
      </div>
    </el-menu>

    <!-- 底部装饰 -->
    <div class="navbar-footer">
      <div class="footer-gradient"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import {
  House, List, Star, Menu, Upload, Plus,
} from '@element-plus/icons-vue';
import { useUserStore } from '@/stores/user';

const router = useRouter();
const userStore = useUserStore();
const route = useRoute();

const userId = userStore.userId;

// 【修改点2】: 使用 computed 智能计算 activeRoute
const activeRoute = computed(() => {
  const currentPath = route.path;

  // 优先匹配带参数的路由
  // 如果当前路径以 /love 开头 (例如 /love/123), 我们就认为激活的是 /love
  if (currentPath.startsWith('/love')) {
    return '/love';
  }
  // 同理, 匹配所有歌单页面 (例如 /playlist-grid/123)
  if (currentPath.startsWith('/playlist-grid')) {
    return '/playlist-grid';
  }

  // 对于其他精确匹配的路由，直接返回其路径
  return currentPath;
});


const navigateTo = (path: string) => {
  router.push(path);
};


</script>

<style scoped>
/* Spotify风格的导航栏 */
.spotify-navbar {
  height: 100%;
  display: flex;
  flex-direction: column;
  background: #121212;
  color: #fff;
  position: relative;
  overflow: hidden;
}

/* 侧边栏菜单 */
.sidebar-menu {
  height: 100%;
  border: none;
  background: transparent;
  padding: 24px 0;
  overflow-y: auto;
  overflow-x: hidden;
  flex: 1;
}

/* 自定义滚动条 */
.sidebar-menu::-webkit-scrollbar {
  width: 6px;
}

.sidebar-menu::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 3px;
}

.sidebar-menu::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 3px;
  transition: background 0.2s ease;
}

.sidebar-menu::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}

/* 菜单区块 */
.menu-section {
  margin-bottom: 32px;
  padding: 0 16px;
}

.main-section {
  position: relative;
}

.main-section::after {
  content: '';
  position: absolute;
  bottom: -16px;
  left: 16px;
  right: 16px;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(255, 255, 255, 0.1),
    transparent
  );
}

/* 区块标题 */
.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
  padding: 0 8px;
}

.section-title {
  display: flex;
  align-items: center;
  gap: 12px;
  color: #fff;
  font-size: 14px;
  font-weight: 700;
  margin: 0;
  text-transform: uppercase;
  letter-spacing: 0.1em;
}

.section-title .el-icon {
  font-size: 18px;
  color: #1db954;
}

/* 创建歌单按钮 */
.create-playlist-btn {
  width: 24px;
  height: 24px;
  background: rgba(255, 255, 255, 0.1);
  color: #b3b3b3;
  transition: all 0.3s ease;
}

.create-playlist-btn:hover {
  background: #1db954;
  color: #000;
  transform: scale(1.1);
}

.create-playlist-btn .el-icon {
  font-size: 14px;
}

/* 菜单项样式 */
.menu-item {
  margin: 2px 0;
  border-radius: 8px;
  height: 48px;
  background: transparent;
  border: none;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

/* 菜单项内容 */
.menu-item-content {
  display: flex;
  align-items: center;
  gap: 16px;
  width: 100%;
  height: 100%;
  padding: 0 12px;
  position: relative;
  z-index: 2;
}

.menu-icon {
  font-size: 20px;
  color: #b3b3b3;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

.menu-text {
  font-size: 14px;
  font-weight: 500;
  color: #b3b3b3;
  transition: all 0.3s ease;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 活跃状态指示器 */
.item-indicator {
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 4px;
  height: 0;
  background: #1db954;
  border-radius: 0 2px 2px 0;
  transition: height 0.3s ease;
}

/* 菜单项hover效果 */
.menu-item:hover {
  background: rgba(255, 255, 255, 0.1);
}

.menu-item:hover .menu-icon,
.menu-item:hover .menu-text {
  color: #fff;
  transform: translateX(4px);
}

.menu-item:hover .menu-icon {
  transform: translateX(4px) scale(1.1);
}

/* 活跃状态 */
.menu-item.is-active {
  background: rgba(29, 185, 84, 0.1) !important;
}

.menu-item.is-active .menu-icon,
.menu-item.is-active .menu-text {
  color: #1db954 !important;
}

.menu-item.is-active .item-indicator {
  height: 24px;
}

.menu-item.is-active::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    90deg,
    rgba(29, 185, 84, 0.1),
    transparent
  );
  z-index: 1;
}

/* 最近播放的歌单区域 */
.recent-playlists {
  margin-top: 8px;
}

/* 底部装饰 */
.navbar-footer {
  position: relative;
  height: 20px;
  margin-top: auto;
}

.footer-gradient {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(29, 185, 84, 0.5),
    transparent
  );
}

/* 响应式设计 */
@media screen and (max-width: 1200px) {
  .sidebar-menu {
    padding: 20px 0;
  }

  .menu-section {
    padding: 0 12px;
    margin-bottom: 24px;
  }
}

@media screen and (max-width: 992px) {
  .menu-item {
    height: 44px;
  }

  .menu-text {
    font-size: 13px;
  }

  .menu-icon {
    font-size: 18px;
  }
}

@media screen and (max-width: 768px) {
  .sidebar-menu {
    padding: 16px 0;
  }

  .menu-section {
    padding: 0 8px;
    margin-bottom: 20px;
  }

  .section-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .section-title {
    font-size: 12px;
  }

  .menu-text {
    display: none;
  }

  .menu-item-content {
    justify-content: center;
    padding: 0;
  }

  .menu-icon {
    font-size: 20px;
  }

  .create-playlist-btn {
    width: 20px;
    height: 20px;
  }
}

/* 加载动画 */
@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.menu-item {
  animation: slideInLeft 0.3s ease;
}

.menu-item:nth-child(1) {
  animation-delay: 0.1s;
}

.menu-item:nth-child(2) {
  animation-delay: 0.2s;
}

.menu-item:nth-child(3) {
  animation-delay: 0.3s;
}

.menu-item:nth-child(4) {
  animation-delay: 0.4s;
}

/* 高对比度模式 */
@media (prefers-contrast: high) {
  .menu-item {
    border: 1px solid rgba(255, 255, 255, 0.2);
  }

  .menu-item.is-active {
    border-color: #1db954;
  }
}

/* 减少动画偏好 */
@media (prefers-reduced-motion: reduce) {
  .menu-item {
    animation: none;
    transition: none;
  }

  .menu-icon,
  .menu-text,
  .item-indicator,
  .create-playlist-btn {
    transition: none;
  }
}</style>