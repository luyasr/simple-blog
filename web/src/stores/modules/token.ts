import { defineStore } from 'pinia'
import { LoginReq } from '../../api/token/index'
import type { LoginForm } from '../../api/token/type'

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
      const resp: any = await LoginReq(data)
      if (resp.data.code == 0) {
        this.access_token = resp.data.data.access_token
        this.refresh_token = resp.data.data.refresh_token
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.data.message))
      }
    }
  },
  getters: {}
})
