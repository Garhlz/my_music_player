export const player = {
  namespaced: true,
  state: () => ({
    currentSong: null,
    isPlaying: false,
  }),
  mutations: {
    setCurrentSong(state, song) {
      state.currentSong = song;
    },
    setPlayingState(state, isPlaying) {
      state.isPlaying = isPlaying;
    },
  },
  actions: {
    playSong({ commit }, song) {
      commit('setCurrentSong', song);
      commit('setPlayingState', true);
    },
    pauseSong({ commit }) {
      commit('setPlayingState', false);
    },
    addToPlaylist({ commit, state }, song) {
      const updatedPlaylist = [...state.playlist, song];
      commit('setPlaylist', updatedPlaylist);
    },
    nextSong({ commit, state }) {
      const currentIndex = state.playlist.findIndex(song => song.id === state.currentSong.id);
      const nextIndex = (currentIndex + 1) % state.playlist.length;
      const nextSong = state.playlist[nextIndex];
      commit('setCurrentSong', nextSong);
      commit('setPlayingState', true);
    },
    previousSong({ commit, state }) {
      const currentIndex = state.playlist.findIndex(song => song.id === state.currentSong.id);
      const previousIndex = (currentIndex - 1 + state.playlist.length) % state.playlist.length;
      const previousSong = state.playlist[previousIndex];
      commit('setCurrentSong', previousSong);
      commit('setPlayingState', true);
    },
  },
  getters: {
    currentSong: (state) => state.currentSong,
    isPlaying: (state) => state.isPlaying,
  },
}; 