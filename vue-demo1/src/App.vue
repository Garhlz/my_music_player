<template>
  <div class="app-container">
    <router-view></router-view>
    <Player v-if="currentSong" />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import Player from '@/components/Player.vue'
import { usePlayerStore } from '@/stores/player'

const playerStore = usePlayerStore()
const currentSong = computed(() => playerStore.getCurrentSong)
</script>

<style>
.app-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>
