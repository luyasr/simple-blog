import { defineStore } from 'pinia'
import { constantRoute } from '@/router/routes'

export const useRouteStore = defineStore({
  id: 'route',
  state: () => ({
    menuRoutes: constantRoute
  }),
  actions: {},
  getters: {}
})
