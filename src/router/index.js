import {createRouter, createWebHistory} from 'vue-router'

import PlantsView from "@/views/plants/PlantsView.vue";
import PlantDetailView from "@/views/plants/PlantDetailView.vue";


const routes = [
    {
        path: '/',
        component: PlantsView
    },
    {
        path: '/plants/:id',
        name: 'plant',
        component: PlantDetailView
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router