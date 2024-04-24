import {createRouter, createWebHistory} from 'vue-router'


const routes = [
    {
        path: '/plants',
        name: 'plants-item',
        component: 'plants-item'
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router