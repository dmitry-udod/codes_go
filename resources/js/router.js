import Vue from 'vue';
import Router from 'vue-router'
import Fop from "./components/Fop"
import FopDetails from "./components/FopDetails"
import LegalEntities from "./components/LegalEntities"
import LegalEntityDetails from "./components/LegalEntityDetails"

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
            path: '/fop/search/:q',
            name: 'fop_search',
            component: Fop,
        },
        {
            path: '/fop/details/:id',
            name: 'fop_details',
            component: FopDetails,
        },
        {
            path: '/legal-entities',
            name: 'legal_entities',
            component: LegalEntities,
        },
        {
            path: '/legal-entities/search/:q',
            name: 'legal_entities_search',
            component: LegalEntities,
        },
        {
            path: '/legal-entities/details/:id',
            name: 'legal_entities_details',
            component: LegalEntityDetails,
        },
    ]
});

export default router
