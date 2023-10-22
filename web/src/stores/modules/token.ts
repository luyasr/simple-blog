import { defineStore } from 'pinia'
import { LoginReq } from '@/api/token'
import type { LoginForm } from '@/types/token'

export const useTokenStore = defineStore({
  id: 'token',
  // 开启数据持久化
  persist: true,
  // 存储数据
  state: () => ({
    access_token: localStorage.getItem('access_token') || '',
    refresh_token: localStorage.getItem('refresh_token') || ''
  }),
  // 异步逻辑
  actions: {
    async login(data: LoginForm) {
      const resp = await LoginReq(data)
      if (resp.code == 0) {
        this.access_token = resp.data.access_token
        this.refresh_token = resp.data.refresh_token
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    }
  },
  getters: {}
})
