import swal from 'sweetalert2';

export default {
    data () {
        return {
            congif: {},
            loading: false,
        }
    },

    beforeMount() {
        this.config = require(`../config/api.${process.env.NODE_ENV}`);
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

        fopLatest(page = 1) {
            return this.sendRequest('get', `fop/latest?page=${page}`)
        },

        fopDetails(id) {
            return this.sendRequest('get', `fop/view/${id}`)
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
    }
}
