// vue-router 配置
import { createRouter, createWebHashHistory, type RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/login/index.vue')
    },
    {
        path: '/home',
        name: 'home',
        component: () => import('../views/home/index.vue')
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
