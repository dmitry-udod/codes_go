<template>
    <div>
        <div class="content" v-if="entity">
            <h1>
                {{ entity.full_name }}
            </h1>

            <ul>
                <li>Повна назва: {{ entity.full_name }}</li>
                <li>Коротка назва: {{ entity.short_name }}</li>
                <li>ЄДРПОУ: {{ entity.code }}</li>
                <li>Адреса: {{ entity.address }}</li>
                <li>Вид діяльності (КВЕД): {{ entity.activity }}</li>
                <li>Статус: <fop-status :status="entity.status"></fop-status></li>
                <li>
                    Засновники:
                    <ul v-if="entity.founders && entity.founders.length > 0">
                        <li v-for="f in entity.founders">{{ f.name }}</li>
                    </ul>
                </li>
            </ul>

            <div class="text-center">
                <router-link class="btn btn_default blue d-inline-block" :to="{name: $route.params.q ? 'legal_entities_search' : 'legal_entities', params: $route.params}">До списку</router-link>
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
            this.legalEntity();
        },

        methods: {
            legalEntity() {
                this.startLoading();
                this.legalEntityDetails(this.$route.params.id).then(response => {
                    this.stopLoading();
                    this.entity = response.data.data;
                }, this.onError)
            },
        }
    }
</script>

<style>

</style>
