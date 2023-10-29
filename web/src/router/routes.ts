import Layout from '../layout/LayoutView.vue'

// 对外暴露的路由
export const constantRoute = [
  {
    path: '/',
    name: 'layout',
    component: Layout,
    meta: {
      title: 'layout',
      requiresAuth: true,
      icon: 'icon-apps',
      hideInMenu: true
    },
    children: [
      {
        path: '/home',
        name: 'home',
        component: () => import('../views/home/HomeView.vue'),
        meta: {
          title: '首页',
          requiresAuth: true,
          icon: 'icon-home',
          hideInMenu: false
        }
      }
    ]
  },
  {
    path: '/blog',
    name: 'blog',
    component: Layout,
    meta: {
      title: '博客',
      requiresAuth: true,
      icon: 'icon-book',
      hideInMenu: true
    },
    children: [
      {
        path: '/blog/article',
        name: 'article',
        component: () => import('../views/blog/article/ArticleView.vue'),
        meta: {
          title: '文章管理',
          requiresAuth: true,
          icon: 'icon-bookmark',
          hideInMenu: false
        }
      }
    ]
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/login/LoginView.vue'),
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
