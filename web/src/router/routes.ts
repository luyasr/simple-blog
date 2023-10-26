// 对外暴露的路由
export const constantRoute = [
    {
        path: '/',
        name: 'layout',
        component: () => import('../layout/IndexView.vue'),
        meta: {
            locale: 'menu.layout',
            requiresAuth: true,
            icon: 'icon-dashboard'
        },
        children: [
            {
                path: '/home',
                name: 'home',
                component: () => import('../views/home/IndexView.vue'),
                meta: {
                    locale: 'menu.home',
                    requiresAuth: true,
                    hideInMenu: false
                }
            }
        ]
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/login/IndexView.vue')
    },
    {
        path: '/404',
        name: '404',
        component: () => import('../views/404/404View.vue')
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'any',
        redirect: '/404'
    }
]
