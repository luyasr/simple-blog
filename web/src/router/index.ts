import { createRouter, createWebHistory } from 'vue-router'
import { constantRoute } from './routes'
import { beforeEach } from './permission'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoute
})

router.beforeEach(beforeEach)

export default router

