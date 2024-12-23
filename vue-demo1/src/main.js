import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
//注意引入路由
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'  
// 引入 Element Plus 的 CSS 样式
import store from './store' // 引入 Vuex store

const app = createApp(App)

app.use(router)
app.use(store) // 使用 Vuex store
app.use(ElementPlus)

app.mount('#app')
