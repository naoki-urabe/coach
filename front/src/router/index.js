import { createWebHistory, createRouter } from "vue-router";
import HelloWorld from "../components/HelloWorld";
import Register from "../components/Register";

const routes = [
    {
        path: "/hello",
        component: HelloWorld,
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