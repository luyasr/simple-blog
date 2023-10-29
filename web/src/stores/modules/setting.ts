// layout组件相关配置
import { defineStore } from 'pinia'

export const useLayoutSettingStore = defineStore({
  id: 'layoutSetting',
  state: () => ({
    refresh: false // 用于控制刷新效果
  }),
  actions: {},
  getters: {}
})
