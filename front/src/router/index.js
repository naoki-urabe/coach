import { createWebHistory, createRouter } from "vue-router";
import Kuji from "../components/Kuji";
import Register from "../components/Register";

const routes = [
    {
        path: "/kuji",
        component: Kuji,
    },
    {
        path: "/register",
        component: Register,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;