import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../layout/IndexView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/404',
      name: '404',
      component: () => import('../views/404View.vue')
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'any',
      redirect: '/404'
    }
  ]
})

export default router
