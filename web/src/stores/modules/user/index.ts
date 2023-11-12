import { defineStore } from 'pinia'
import { UserInfo } from '@/api/user'
import type { UserState } from './types'

export const useUserStore = defineStore({
  id: 'user',
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
