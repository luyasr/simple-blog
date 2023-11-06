import Layout from '../layout/LayoutView.vue'

// 对外暴露的路由
export const constantRoute = [
  {
    path: '/',
    name: 'layout',
    component: Layout,
    meta: {
      title: '',
      requiresAuth: true,
      icon: '',
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
        name: 'blogArticle',
        component: () => import('../views/blog/article/ArticleView.vue'),
        meta: {
          title: '文章管理',
          requiresAuth: true,
          icon: 'icon-bookmark',
          hideInMenu: false
        }
      },
      {
        path: '/blog/edit',
        name: 'blogEdit',
        component: () => import('../views/blog/article/EditView.vue'),
        meta: {
          title: '文章编辑',
          requiresAuth: true,
          icon: 'icon-bookmark',
          hideInMenu: true
        }
      },
      {
        path: '/blog/avatar',
        name: 'blogAvatar',
        component: () => import('../views/upload/AvatarView.vue'),
        meta: {
          title: '头像上传',
          requiresAuth: true,
          icon: 'icon-upload',
          hideInMenu: false
        }
      }
    ]
  },
  {
    path: '/system',
    name: 'system',
    component: Layout,
    meta: {
      title: '系统设置',
      requiresAuth: true,
      icon: 'icon-settings',
      hideInMenu: false
    },
    children: [
      {
        path: '/system/user',
        name: 'user',
        component: () => import('@/views/system/UserView.vue'),
        meta: {
          title: '个人中心',
          requiresAuth: true,
          icon: 'icon-user',
          hideInMenu: true
        }
      },
      {
        path: '/system/permissions',
        name: 'permissions',
        component: () => import('@/views/system/PermView.vue'),
        meta: {
          title: '权限管理',
          requiresAuth: true,
          icon: 'icon-lock',
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
