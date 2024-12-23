import { defineStore } from 'pinia'

export const usePlayerStore = defineStore('player', {
  state: () => ({
    currentSong: null,
    isPlaying: false,
    playlist: [],
    currentIndex: -1
  }),
  
  getters: {
    getCurrentSong: (state) => state.currentSong,
    getIsPlaying: (state) => state.isPlaying,
    getPlaylist: (state) => state.playlist,
    getCurrentIndex: (state) => state.currentIndex
  },
  
  actions: {
    setPlaylist(songs) {
      this.playlist = songs
      this.currentIndex = 0
      this.currentSong = songs[0]
    },
    
    play(index) {
      if (index >= 0 && index < this.playlist.length) {
        this.currentIndex = index
        this.currentSong = this.playlist[index]
        this.isPlaying = true
      }
    },
    
    pause() {
      this.isPlaying = false
    },
    
    resume() {
      this.isPlaying = true
    },
    
    next() {
      if (this.playlist.length > 0) {
        this.currentIndex = (this.currentIndex + 1) % this.playlist.length
        this.currentSong = this.playlist[this.currentIndex]
      }
    },
    
    previous() {
      if (this.playlist.length > 0) {
        this.currentIndex = this.currentIndex - 1 < 0 ? this.playlist.length - 1 : this.currentIndex - 1
        this.currentSong = this.playlist[this.currentIndex]
      }
    }
  }
}) 