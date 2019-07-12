import Vue from 'vue';
import Router from 'vue-router'
import Fop from "./components/Fop"
import FopDetails from "./components/FopDetails"

Vue.use(Router);

const router = new Router({
    routes: [
        {
            path: '/',
            name: 'main',
            component: Fop,
        },
        {
            path: '/fop',
            name: 'fop',
            component: Fop,
        },
        {
            path: '/fop/details/:id',
            name: 'fop_details',
            component: FopDetails,
        },
    ]
});

export default router
