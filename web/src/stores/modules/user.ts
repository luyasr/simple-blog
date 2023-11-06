import { defineStore } from 'pinia'
import { Login, UserInfo, Logout } from '@/api/user'
import type { LoginRequest, LogoutRequest } from '@/types/user'
import type { UserState } from '../types/user'

export const useUserStore = defineStore({
  id: 'user',
  // 开启数据持久化
  persist: true,
  // 存储数据
  state: (): UserState => ({
    access_token: localStorage.getItem('access_token'),
    refresh_token: localStorage.getItem('refresh_token'),
    isAuthenticated: false,
    user_id: 0,
    username: '',
    avatar: ''
  }),
  // 异步逻辑
  actions: {
    async login(data: LoginRequest) {
      const resp = await Login(data)
      if (resp.code == 200) {
        this.user_id = resp.data.user_id
        this.username = resp.data.username
        this.access_token = resp.data.access_token
        this.refresh_token = resp.data.refresh_token
        this.isAuthenticated = true
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async logout() {
      const req: LogoutRequest = {
        access_token: this.access_token,
        refresh_token: this.refresh_token
      }
      const resp = await Logout(this.user_id, req)
      if (resp.code == 200) {
        this.$reset()
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async userInfo() {
      const resp = await UserInfo(this.user_id)
      if (resp.code == 200) {
        this.avatar = resp.data.avatar
        return Promise.resolve(resp.data)
      } else {
        return Promise.reject(new Error(resp.message))
      }
    }
  },
  getters: {}
})
