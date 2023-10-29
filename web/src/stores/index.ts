import { createPinia } from 'pinia'
// 导入让数据持久化存储的 Pinia 插件
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

export default pinia
