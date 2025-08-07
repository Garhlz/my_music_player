// vite.config.js
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { fileURLToPath, URL } from 'node:url';

export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      'src': fileURLToPath(new URL('./src', import.meta.url)),
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  // *** 代理配置应该在这里 ***
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // 指向你的 Go 后端服务
        changeOrigin: true, // 允许跨域
      },
    },
  },
});