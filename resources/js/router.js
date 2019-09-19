import Vue from 'vue';
import Router from 'vue-router'
import Fop from "./components/Fop"
import FopDetails from "./components/FopDetails"
import LegalEntities from "./components/LegalEntities"
import LegalEntityDetails from "./components/LegalEntityDetails"
import Terrorists from "./components/Terrorists"
import TerroristDetails from "./components/TerroristDetails"

Vue.use(Router);

const router = new Router({
    mode: 'history',
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
        {
            path: '/terrorists',
            name: 'terrorists',
            component: Terrorists,
        },
        {
            path: '/terrorists/search/:q',
            name: 'terrorists_search',
            component: Terrorists,
        },
        {
            path: '/terrorists/details/:id',
            name: 'terrorist_details',
            component: TerroristDetails,
        },
    ]
});

export default router
