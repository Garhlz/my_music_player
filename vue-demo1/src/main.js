import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'; // <--- 1. 导入插件
import ElementPlus from 'element-plus';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import 'element-plus/dist/index.css';
import App from './App.vue';
import router from './router';

// 创建应用实例
const app = createApp(App);

// 注册 Pinia
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate); // <--- 2. 注册插件
app.use(pinia);

// 注册 Element Plus
app.use(ElementPlus);

// 注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

// 注册路由
app.use(router);

// 挂载应用
app.mount('#app');
