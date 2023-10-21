import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
// 导入Arco UI库
import ArcoVue from '@arco-design/web-vue';
// 导入Arco 图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import '@arco-design/web-vue/dist/arco.css';
// 导入让数据持久化存储的 Pinia 插件
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(ArcoVue)
app.use(ArcoVueIcon)
app.use(router)

app.mount('#app')
