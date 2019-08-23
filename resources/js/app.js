window.Vue = require('vue');

import axios from 'axios'
import App from "./App"
import Navigation from "./components/Navigation"
import router from './router';
import api from './minxins/api';
import FopStatus from './components/Fop/Status'
import Search from './components/Common/Search'

Vue.http = Vue.prototype.$http = axios;

Vue.component('fop-status', FopStatus);
Vue.component('search', Search);
Vue.component('pagination', require('laravel-vue-pagination'));

Vue.mixin(api);

const app = new Vue({
    el: '#app',
    components: { App, Navigation },
    router,
});