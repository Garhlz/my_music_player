<template>
  <el-menu
    class="sidebar-menu"
    :default-active="activeRoute"
    background-color="#fff"
    text-color="#606266"
    active-text-color="var(--el-color-primary)"
  >


    <div class="menu-section">
      <!--      <el-menu-item index="/player" @click="navigateTo('/player')" class="menu-item">-->
      <!--        <el-icon>-->
      <!--          <VideoPlay />-->
      <!--        </el-icon>-->
      <!--        <span>音乐播放</span>-->
      <!--      </el-menu-item>-->
      <el-menu-item index="/pub" @click="navigateTo('/pub')" class="menu-item">
        <el-icon>
          <List />
        </el-icon>
        <span>音乐库</span>
      </el-menu-item>
      <el-menu-item index="/love" @click="navigateTo(`/love/${userId}`)" class="menu-item">
        <el-icon>
          <Star />
        </el-icon>
        <span>我喜欢的音乐</span>
      </el-menu-item>
      <el-menu-item index="/playlist-grid" @click="navigateTo(`/playlist-grid/${userId}`)" class="menu-item">
        <el-icon>
          <Menu />
        </el-icon>
        <span>我的歌单</span>
      </el-menu-item>
      <el-menu-item index="/uploaded" @click="navigateTo('/uploaded')" class="menu-item">
        <el-icon>
          <Upload />
        </el-icon>
        <span>上传的音乐</span>
      </el-menu-item>
    </div>

    <!--    <div class="menu-section">-->
    <!--      <h3 class="section-title">管理</h3>-->
    <!--      <el-menu-item index="/manage-users" @click="navigateTo('/manage-users')" class="menu-item">-->
    <!--        <el-icon>-->
    <!--          <User />-->
    <!--        </el-icon>-->
    <!--        <span>管理用户</span>-->
    <!--      </el-menu-item>-->
    <!--      <el-menu-item index="/manage-songs" @click="navigateTo('/manage-songs')" class="menu-item">-->
    <!--        <el-icon>-->
    <!--          <Setting />-->
    <!--        </el-icon>-->
    <!--        <span>管理音乐</span>-->
    <!--      </el-menu-item>-->
    <!--    </div>-->
  </el-menu>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { Headset, VideoPlay, Star, Menu, List, Upload, User, Setting } from '@element-plus/icons-vue';

import { useUserStore } from '@/stores/user';

const router = useRouter();
const userStore = useUserStore();
const route = useRoute();
const activeRoute = ref(route.path);

router.afterEach((to) => {
  activeRoute.value = to.path;
});

const navigateTo = (path: string) => {
  router.push(path);
};

const userId = userStore.userId;

// const getUserId = () => {
//   try {
//     const userStr = localStorage.getItem('userId');
//     return userStr || '';
//   } catch (error) {
//     console.error('获取用户ID失败:', error);
//     return '';
//   }
// };

onMounted(() => {
  activeRoute.value = route.path;
});
</script>

<style scoped>
.sidebar-menu {
  height: 100%;
  border-right: none;
  padding: 16px 0;
  background: #fff;
  animation: fadeIn 0.5s ease-in;
}

.menu-header {
  padding: 20px;
  text-align: center;
  margin-bottom: 24px;
}

.logo {
  width: 48px;
  height: 48px;
  margin: 0 auto 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--el-color-primary), var(--el-color-primary-light-5));
  border-radius: 12px;
  transition: all 0.3s ease;
}

.logo:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.site-title {
  color: var(--el-text-color-primary);
  font-size: 20px;
  margin: 0;
  font-weight: 600;
}

.menu-section {
  margin-bottom: 24px;
  padding: 0 12px;
}

.section-title {
  color: #909399;
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  padding: 0 12px;
  margin: 8px 0;
}

.menu-item {
  margin: 4px 0;
  border-radius: 8px;
  height: 48px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.menu-item:hover {
  background-color: var(--el-fill-color-light) !important;
  transform: translateX(4px);
}

.menu-item.is-active {
  background: var(--el-color-primary-light-9) !important;
  border-left: 4px solid var(--el-color-primary);
}

.el-menu-item .el-icon {
  margin-right: 12px;
  font-size: 18px;
}

.menu-item:hover .el-icon {
  transform: scale(1.1);
}

.el-menu-item span {
  font-size: 14px;
  font-weight: 500;
}

/* 滚动条样式 */
.sidebar-menu::-webkit-scrollbar {
  width: 6px;
}

.sidebar-menu::-webkit-scrollbar-thumb {
  background-color: #dcdfe6;
  border-radius: 3px;
}

.sidebar-menu::-webkit-scrollbar-track {
  background-color: #f5f7fa;
}

/* 响应式调整 */
@media screen and (max-width: 768px) {
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
    font-size: 16px;
  }

  .menu-item {
    height: 40px;
  }

  .menu-item span {
    display: none;
  }

  .el-menu-item .el-icon {
    margin-right: 0;
  }
}
</style>
