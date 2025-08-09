<template>
  <div class="app-container">
    <router-view></router-view>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { usePlayerStore } from 'src/stores/player';

const playerStore = usePlayerStore();
</script>

<style>
.app-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden; /* 防止子组件溢出 */
}

html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
}
</style>