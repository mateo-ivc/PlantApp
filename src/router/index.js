import {createRouter, createWebHistory} from 'vue-router'
import HomeView from "@/views/home/HomeView.vue";
import PlantsView from "@/views/plants/PlantsView.vue";
import PlantDetailView from "@/views/plants/PlantDetailView.vue";


const routes = [
    {
        path: '/',
        component: HomeView
    },
    {
        path: '/plants',
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