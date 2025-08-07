<template>
  <div class="common-layout">
	<!-- 头部区域 -->
	<div class="fixed-header">
	  <Banner :page-name="pageName"/>
	</div>

	<!-- 侧边栏区域 -->
	<div class="fixed-aside">
	  <Navbar/>
	</div>

	<!-- 主内容区域 -->
	<div class="main-content">
	  <slot name="main"></slot>
	</div>

	<!-- 底部区域 -->
	<div class="fixed-footer">
	  <MusicPlayer/>
	</div>
  </div>
</template>

<script>
import Banner from 'src/components/Banner.vue';  // 条幅信息栏
import Navbar from 'src/components/Navbar.vue';  // 导航栏
import MusicPlayer from 'src/components/MusicPlayer.vue';  // 音乐播放栏

export default {
  name: 'BaseLayout',
  components: {
	Banner,
	Navbar,
	MusicPlayer
  },
  props: {
	pageName: String,
  },
};
</script>

<style scoped>
.common-layout {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  position: relative;
}

/* 头部样式 */
.fixed-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 64px;
  background-color: #fff;
  z-index: 1000;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

/* 侧边栏样式 */
.fixed-aside {
  position: fixed;
  top: 64px;
  left: 0;
  bottom: 80px;
  width: 240px;
  background-color: #f4f4f4;
  z-index: 900;
  overflow: hidden;
}

/* 主内容区域样式 */
.main-content {
  position: fixed;
  top: 64px;
  left: 240px;
  right: 0;
  bottom: 80px;
  overflow: hidden;
  background-color: #fff;
}

/* 底部播放器样式 */
.fixed-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 80px;
  background-color: #fff;
  z-index: 1000;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
}
</style>
