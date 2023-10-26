import { defineStore } from 'pinia'
import { LoginReq, UserInfoReq } from '@/api/user'
import type { LoginForm } from '@/types/user'
import type { UserState } from '../types/user'

export const useUserStore = defineStore({
  id: 'token',
  // 开启数据持久化
  persist: true,
  // 存储数据
  state: (): UserState => ({
    access_token: localStorage.getItem('access_token'),
    refresh_token: localStorage.getItem('refresh_token'),
    user_id: 0,
    username: "",
    avatar: "",
  }),
  // 异步逻辑
  actions: {
    async login(data: LoginForm) {
      const resp = await LoginReq(data)
      if (resp.code == 0) {
        this.user_id = resp.data.user_id
        this.access_token = resp.data.access_token
        this.refresh_token = resp.data.refresh_token
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    },
    async userInfo() {
      const resp = await UserInfoReq(this.user_id)
      if (resp.code == 0) {
        this.username = resp.data.username
        this.avatar = resp.data.avatar
        return 'ok'
      } else {
        return Promise.reject(new Error(resp.message))
      }
    }
  },
  getters: {}
})
