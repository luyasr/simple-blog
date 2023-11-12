import { defineStore } from 'pinia'
import { Login, Logout } from '@/api/user'
import type { LoginRequest, LogoutRequest } from '@/types/user'
import type { TokenState } from './types'
import { useUserStore } from '../user'

export const useTokenStore = defineStore({
  id: 'token',
  // 开启数据持久化
  persist: true,
  // 存储数据
  state: (): TokenState => ({
    access_token: localStorage.getItem('access_token'),
    refresh_token: localStorage.getItem('refresh_token')
  }),
  // 异步逻辑
  actions: {
    async login(data: LoginRequest) {
      const userStore = useUserStore()
      const resp = await Login(data)
      if (resp.code == 200) {
        this.access_token = resp.data.access_token
        this.refresh_token = resp.data.refresh_token

        userStore.user_id = resp.data.user_id
        userStore.username = resp.data.username
        userStore.isAuthenticated = true
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async logout() {
        const userStore = useUserStore()
      const req: LogoutRequest = {
        access_token: this.access_token,
        refresh_token: this.refresh_token
      }
      const resp = await Logout(userStore.user_id, req)
      if (resp.code == 200) {
        this.$reset()
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    }
  },
  getters: {}
})
