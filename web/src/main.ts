import './assets/main.css'

import { createApp } from 'vue'
// 导入Arco UI库
import ArcoVue from '@arco-design/web-vue';
// 导入Arco 图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import '@arco-design/web-vue/dist/arco.css';
import App from './App.vue'
import router from './router'
import pinia from './stores'

const app = createApp(App)

app.use(pinia)
app.use(ArcoVue)
app.use(ArcoVueIcon)
app.use(router)

app.mount('#app')
