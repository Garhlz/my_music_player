<template>
  <el-menu
    class="sidebar-menu"
    :default-active="activeRoute"
    background-color="#1a1a1a"
    text-color="#8c8c8c"
    active-text-color="#ffffff"
  >
    <div class="menu-header">
      <div class="logo">
        <el-icon :size="32"><Headset /></el-icon>
      </div>
      <h1 class="site-title">享受音乐</h1>
    </div>

    <div class="menu-section">
      <h3 class="section-title">发现音乐</h3>
      <el-menu-item index="/player" @click="navigateTo('/player')" class="menu-item">
        <el-icon><VideoPlay /></el-icon>
        <span>音乐播放</span>
      </el-menu-item>

      <el-menu-item index="/public-playlist" @click="navigateTo('/public-playlist')" class="menu-item">
        <el-icon><List /></el-icon>
        <span>音乐库</span>
      </el-menu-item>
    </div>

    <div class="menu-section">
      <h3 class="section-title">我的音乐</h3>
      <el-menu-item index="/my-love" @click="navigateTo(`/my-love/${getUserId()}`)" class="menu-item">
        <el-icon><Star /></el-icon>
        <span>我喜欢的音乐</span>
      </el-menu-item>

      <el-menu-item 
        index="/my-playlist" 
        @click="navigateTo(`/my-playlist/${getUserId()}`)"
        class="menu-item"
      >
        <el-icon><Menu /></el-icon>
        <span>我的歌单</span>
      </el-menu-item>

      <el-menu-item index="/uploaded" @click="navigateTo('/uploaded')" class="menu-item">
        <el-icon><Upload /></el-icon>
        <span>上传的音乐</span>
      </el-menu-item>
    </div>

    <div class="menu-section">
      <h3 class="section-title">管理</h3>
      <el-menu-item index="/manage-users" @click="navigateTo('/manage-users')" class="menu-item">
        <el-icon><User /></el-icon>
        <span>管理用户</span>
      </el-menu-item>

      <el-menu-item index="/manage-songs" @click="navigateTo('/manage-songs')" class="menu-item">
        <el-icon><Setting /></el-icon>
        <span>管理音乐</span>
      </el-menu-item>
    </div>
  </el-menu>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  Headset,
  VideoPlay,
  Star,
  Menu,
  List,
  Upload,
  User,
  Setting
} from '@element-plus/icons-vue'

export default {
  name: 'SidebarNavbar',
  components: {
    Headset,
    VideoPlay,
    Star,
    Menu,
    List,
    Upload,
    User,
    Setting
  },
  setup() {
    const router = useRouter()
    const route = useRoute()
    const activeRoute = ref(route.path)
    const userId = ref(localStorage.getItem('userId'))
    // 监听路由变化
    router.afterEach((to) => {
      activeRoute.value = to.path
    })

    // 导航方法
    const navigateTo = (path) => {
      router.push(path)
    }

    onMounted(() => {
      activeRoute.value = route.path
    })

    const getUserId = () => {
      try {
        const userStr = localStorage.getItem('userId')
        console.log(userStr)
        return userStr
      } catch (error) {
        console.error('获取用户ID失败:', error)
        return ''
      }
    }

    return {
      activeRoute,
      navigateTo,
      getUserId
    }
  }
}
</script>

<style scoped>
.sidebar-menu {
  height: 100%;
  border-right: none;
  padding: 16px 0;
  background: linear-gradient(180deg, #1a1a1a 0%, #2d2d2d 100%);
}

.menu-header {
  padding: 20px;
  text-align: center;
  margin-bottom: 24px;
}

.logo {
  width: 48px;
  height: 48px;
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--el-color-primary);
  transition: transform 0.3s ease;
}

.logo:hover {
  transform: scale(1.1);
}

.site-title {
  color: #ffffff;
  font-size: 1.2rem;
  margin: 0;
  font-weight: 500;
}

.menu-section {
  margin-bottom: 24px;
  padding: 0 12px;
}

.section-title {
  color: #6b7280;
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  padding: 0 12px;
  margin: 8px 0;
}

.menu-item {
  margin: 4px 0;
  border-radius: 8px;
  height: 48px;
  transition: all 0.3s ease;
}

.menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1) !important;
  transform: translateX(4px);
}

.menu-item.is-active {
  background: linear-gradient(90deg, var(--el-color-primary) 0%, var(--el-color-primary-light-3) 100%) !important;
  box-shadow: 0 4px 12px rgba(var(--el-color-primary-rgb), 0.3);
}

.el-menu-item .el-icon {
  margin-right: 12px;
  font-size: 18px;
  transition: transform 0.3s ease;
}

.menu-item:hover .el-icon {
  transform: scale(1.1);
}

.el-menu-item span {
  font-size: 14px;
  font-weight: 500;
}

/* 自定义滚动条 */
.sidebar-menu::-webkit-scrollbar {
  width: 6px;
}

.sidebar-menu::-webkit-scrollbar-thumb {
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 3px;
}

.sidebar-menu::-webkit-scrollbar-track {
  background-color: transparent;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .sidebar-menu {
    padding: 12px 0;
  }

  .menu-header {
    padding: 12px;
  }

  .logo {
    width: 36px;
    height: 36px;
  }

  .site-title {
    font-size: 1rem;
  }

  .menu-item {
    height: 40px;
  }
}
</style>
