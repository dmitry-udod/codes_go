import swal from 'sweetalert2';

export default {
    data () {
        return {
            congif: {},
            loading: false,
            metadata: {},
            entities: null,
            q: '',
        }
    },

    beforeMount() {
        this.config = require(`../config/api.${process.env.NODE_ENV}`);
        if (this.$route.params.q) {
            this.q = this.$route.params.q;
        }
    },

    methods: {
        startLoading() {
            this.loading = true;
        },

        stopLoading() {
            this.loading = false;
        },

        onError(err) {
            this.stopLoading();
            if (err.response && err.response.data && err.response.data.error) {
                swal.fire(err.response.data.error, '', 'error').then(() => {
                    if(this.afterError) {
                        this.afterError();
                    }
                });
            } else {
                let message = '';
                if(err.message) {
                    message = err.message;
                }
                if(typeof err === "string") {
                    message = err;
                }
                swal.fire('Some error happened', message, 'error');
            }
        },

        fopLatest(request) {
            request = this.formatRequest(request);
            return this.sendRequest('get', `fop/latest?page=${request.page}&q=${request.q}`);
        },

        legalEntitiesLatest(request) {
            request = this.formatRequest(request);
            return this.sendRequest('get', `legal-entities/latest?page=${request.page}&q=${request.q}`);
        },

        fopDetails(id) {
            return this.sendRequest('get', `fop/view/${id}`)
        },

        legalEntityDetails(id) {
            return this.sendRequest('get', `legal-entities/view/${id}`)
        },

        terrorists(request) {
            request = this.formatRequest(request);
            return this.sendRequest('get', `terrorists?page=${request.page}&q=${request.q}`);
        },

        sendRequest(method, url, data, onlyUrl = false, onlyHost = false) {
            if (this.config.port && this.config.port !=='' && this.config.port.search(/:/) === -1) {
                this.config.port =  ':' + this.config.port
            }

            let fullUrl = '';
            if (onlyHost) {
                fullUrl = `${this.config.host}${this.config.port}/${url}`;
            } else {
                fullUrl = `${this.config.host}${this.config.port}/${this.config.api_prefix}/${this.config.api_version}/${url}`;
            }

            return this.$http[method](fullUrl, data)
        },

        hasEntities() {
            return this.entities && this.entities.length > 0;
        },

        formatRequest(request) {
            request.q = request.q ? request.q : '';
            return request
        }
    }
}
