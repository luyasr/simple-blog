// 对外暴露的路由
export const constantRoute = [
  {
    path: '/',
    name: 'layout',
    component: () => import('../layout/IndexView.vue'),
    meta: {
      title: 'layout',
      requiresAuth: true,
      icon: 'icon-apps',
      hideInMenu: false
    },
    children: [
      {
        path: '/home',
        name: 'home',
        component: () => import('../views/home/IndexView.vue'),
        meta: {
          title: '首页',
          requiresAuth: true,
          icon: 'icon-home',
          hideInMenu: false
        }
      },
      {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('../views/home/IndexView.vue'),
        meta: {
          title: '控制台',
          requiresAuth: true,
          icon: 'icon-dashboard',
          hideInMenu: false
        }
      }
    ]
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/login/IndexView.vue'),
    meta: {
      title: '登录',
      hideInMenu: true
    }
  },
  {
    path: '/404',
    name: '404',
    component: () => import('../views/404/404View.vue'),
    meta: {
      title: '404',
      hideInMenu: true
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'any',
    redirect: '/404',
    meta: {
      title: '任意路由',
      hideInMenu: true
    }
  }
]
