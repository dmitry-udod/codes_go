<template>
    <div>
        <search title="Пошук ФОП" @search="search" @clearSearch="clearSearch()"></search>

        <div class="mt-5">
            <h1>
                Список ФОП
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
                    <th class="py-4 px-6 bg-gray-200 font-bold uppercase text-sm text-grey-dark border-b border-grey-light">
                        Повне I'мя
                    </th>
                    <th class="py-4 px-6 bg-gray-200 font-bold uppercase text-sm text-grey-dark border-b border-grey-light">
                        Статус
                    </th>
                    <th class="py-4 px-2 bg-gray-200 font-bold uppercase text-sm text-grey-dark border-b border-grey-light">
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="e in entities">
                    <td>{{ e.full_name }}</td>
                    <td>
                        <fop-status :status="e.status"></fop-status>
                    </td>
                    <td>
                        <router-link class="btn btn_default blue m-0" :to="{name: 'fop_details', params: {id: e.id, q: q}}">Детальнiше</router-link>
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

        <pagination :limit="2" :data="metadata" @pagination-change-page="fops" align="right"></pagination>
    </div>
</template>

<script>
    export default {
        data: () => {
            return {

            }
        },

        beforeMount() {
            this.fops();
        },

        methods: {
            fops(page = 1) {
                this.startLoading();
                this.entities = [];
                this.$route.params.page = page;
                this.fopLatest(this.$route.params).then(response => {
                    this.entities = response.data.data;
                    this.metadata = response.data.metadata;
                    this.stopLoading()
                }, this.onError)
            },

            search(q) {
                this.q = q;
                this.page = 1;
                this.$router.push({name: 'fop_search', params: {q: this.q}});
                this.fops();
            },

            clearSearch() {
                this.$router.push({name: 'fop'});
                this.fops();
            },
        }
    }
</script>

<style>

</style>
