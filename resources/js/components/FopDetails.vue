<template>
    <div>
        <div class="content" v-if="entity">
            <h1>
                {{ entity.full_name }}
            </h1>

            <ul>
                <li>Адреса: {{ entity.address }}</li>
                <li>Вид діяльності (КВЕД): {{ entity.activity }}</li>
                <li>Статус: <fop-status :status="entity.status"></fop-status></li>
            </ul>

            <div class="text-center">
                <router-link class="btn btn_default blue d-inline-block" :to="{name: 'fop'}">До списку</router-link>
            </div>
        </div>
        <div class="mb-2 text-center">
            <span v-if="loading">Завантаження...</span>
        </div>
    </div>
</template>

<script>
    export default {
        data: () => {
            return {
                entity: null,
            }
        },

        beforeMount() {
            this.fop();
        },

        methods: {
            fop() {
                this.startLoading();
                this.fopDetails(this.$route.params.id).then(response => {
                    this.stopLoading();
                    this.entity = response.data.data;
                }, this.onError)
            },
        }
    }
</script>

<style>

</style>
