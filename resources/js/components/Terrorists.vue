<template>
    <div>
        <search title="Пошук терористів" @search="search" @clearSearch="clearSearch()"></search>

        <div class="mt-5">
            <h1>
                Список терористів
                <span v-if="hasEntities()" class="float-right">знайдено: {{ metadata.total.toLocaleString() }}</span>
            </h1>
        </div>
        <div class="mb-2 text-center m-5">
            <span v-if="! entities || loading">Завантаження...</span>
            <span v-if="entities && entities.length === 0">Нiчого не знайдено</span>
        </div>

        <div v-if="hasEntities()">
            <table class="table table-striped">
                <thead>
                <tr>
                    <th class="py-4">
                        Номер
                    </th>
                    <th class="py-4">
                        Ім'я
                    </th>
                    <th class="py-4">
                        Статус
                    </th>
                    <th class="py-4">
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="e in entities">
                    <td>{{ e.number_in_list }}</td>
                    <td>
                        <span v-if="e.known_names.length > 0">
                        {{ e.known_names[0].last_name }}
                        {{ e.known_names[0].first_name }}
                        {{ e.known_names[0].middle_name }}
                        {{ e.known_names[0].additional_name }}
                        </span>
                    </td>
                    <td>
                        <fop-status :status="e.status"></fop-status>
                    </td>
                    <td>
                        <router-link class="btn btn_default blue m-0" :to="{name: 'legal_entities_details', params: {id: e.code, q: q}}">Детальнiше</router-link>
                    </td>
                </tr>
                </tbody>
            </table>

            <hr>

            <div>
                Джерело: <a target="_blank" href="https://data.gov.ua/dataset/1c7f3815-3259-45e0-bdf1-64dca07ddc10">https://data.gov.ua/dataset/1c7f3815-3259-45e0-bdf1-64dca07ddc10</a>
            </div>
        </div>

        <hr>

        <pagination :limit="2" :data="metadata" @pagination-change-page="entitiesList" align="right"></pagination>
    </div>
</template>

<script>
    export default {
        data: () => {
            return {

            }
        },

        beforeMount() {
            this.entitiesList()
        },

        methods: {
            entitiesList(page = 1) {
                this.startLoading();
                this.entities = [];
                this.$route.params.page = page;
                this.terrorists(this.$route.params).then(response => {
                    this.entities = response.data.data;
                    this.metadata = response.data.metadata;
                    this.stopLoading()
                }, this.onError)
            },

            search(q) {
                this.q = q;
                this.page = 1;
                this.$router.push({name: 'terrorists', params: {q: this.q}});
                this.entitiesList();
            },

            clearSearch() {
                this.$router.push({name: 'terrorists'});
                this.entitiesList();
            },
        }
    }
</script>

<style>

</style>
