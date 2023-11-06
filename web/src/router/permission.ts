import { useUserStore } from '@/stores/modules/user'

// 全局前置守卫
export const beforeEach = async (to: any) => {
  const userStore = useUserStore()

  const isAuthenticated = userStore.isAuthenticated
  if (!isAuthenticated && to.name !== 'login') {
    return {name: 'login'}
  }
}
