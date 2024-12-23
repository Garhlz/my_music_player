<template>
  <el-menu
    class="sidebar-menu"
    :default-active="activeRoute"
    background-color="#2b2b2b"
    text-color="#909399"
    active-text-color="#409EFF"
  >
    <el-menu-item index="/player" @click="navigateTo('/player')">
      <el-icon><VideoPlay /></el-icon>
      <span>音乐播放</span>
    </el-menu-item>

    <el-menu-item index="/public-playlist" @click="navigateTo('/public-playlist')">
      <el-icon><VideoPlay /></el-icon>
      <span>音乐库</span>
    </el-menu-item>

    <el-menu-item index="/my-love" @click="navigateTo('/my-love')">
      <el-icon><Star /></el-icon>
      <span>我喜欢的音乐</span>
    </el-menu-item>

    <el-menu-item index="/my-playlist" @click="navigateTo('/my-playlist')">
      <el-icon><Menu /></el-icon>
      <span>我的歌单</span>
    </el-menu-item>

    <el-menu-item index="/my-album" @click="navigateTo('/my-album')">
      <el-icon><Files /></el-icon>
      <span>收藏的专辑</span>
    </el-menu-item>

    <el-menu-item index="/uploaded" @click="navigateTo('/uploaded')">
      <el-icon><Upload /></el-icon>
      <span>上传的音乐</span>
    </el-menu-item>

    <!-- <el-menu-divider /> -->

    <el-menu-item index="/comments" @click="navigateTo('/comments')">
      <el-icon><ChatDotRound /></el-icon>
      <span>所有评论</span>
    </el-menu-item>

    <el-menu-item index="/manage-users" @click="navigateTo('/manage-users')">
      <el-icon><User /></el-icon>
      <span>管理用户</span>
    </el-menu-item>

    <el-menu-item index="/manage-songs" @click="navigateTo('/manage-songs')">
      <el-icon><Setting /></el-icon>
      <span>管理音乐</span>
    </el-menu-item>
  </el-menu>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  HomeFilled,
  VideoPlay,
  Star,
  Menu,
  Files,
  Upload,
  ChatDotRound,
  User,
  Setting
} from '@element-plus/icons-vue'

export default {
  name: 'SidebarNavbar',
  components: {
    HomeFilled,
    VideoPlay,
    Star,
    Menu,
    Files,
    Upload,
    ChatDotRound,
    User,
    Setting
  },
  setup() {
    const router = useRouter()
    const route = useRoute()
    const activeRoute = ref(route.path)

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

    return {
      activeRoute,
      navigateTo
    }
  }
}
</script>

<style scoped>
.sidebar-menu {
  height: 100%;
  border-right: none;
}

.el-menu-item {
  display: flex;
  align-items: center;
  padding: 0 20px;
  height: 56px;
  margin: 4px 0;
}

.el-menu-item:hover {
  background-color: #363636 !important;
}

.el-menu-item.is-active {
  background-color: #363636;
  border-left: 4px solid #409EFF;
}

.el-menu-item .el-icon {
  margin-right: 12px;
  font-size: 18px;
}

.el-menu-item span {
  font-size: 14px;
}

.el-menu-divider {
  margin: 20px 0;
  border-top: 1px solid #404040;
}

/* 自定义滚动条 */
.sidebar-menu::-webkit-scrollbar {
  width: 6px;
}

.sidebar-menu::-webkit-scrollbar-thumb {
  background-color: #404040;
  border-radius: 3px;
}

.sidebar-menu::-webkit-scrollbar-track {
  background-color: #2b2b2b;
}
</style>
