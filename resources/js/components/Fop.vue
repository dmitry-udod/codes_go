<template>
    <div>
        <div class="form active m-0">
            <div class="row">
                <div class="col-xl-9">
                    <label>Пошук ФОП:</label>
                    <input type="text" name="key" class="text" autofocus>
                    <div id="search-error" class="search-error"></div>
                </div>
                <div class="col-xl-3 d-flex align-items-end">
                    <button class="btn_default yellow btn_search">Знайти</button>
                </div>
            </div>
        </div>

        <div class="mt-5">
            <h1>
                Список ФОП
                <span v-if="hasEntities()" class="float-right">знайдено: {{ metadata.total.toLocaleString() }}</span>
            </h1>
        </div>
        <div class="mb-2 text-center">
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
                        <router-link class="btn btn_default blue m-0" :to="{name: 'fop_details', params: {id: e.id}}">Детальнiше</router-link>
                    </td>
                </tr>
                </tbody>
            </table>
        </div>

        <hr>

        <button @click="next()" class="btn btn_default yellow float-right ml-2" v-if="hasEntities()">></button>
        <button @click="prev()" class="btn btn_default yellow float-right" v-if="page > 1"><</button>
    </div>
</template>

<script>
    export default {
        data: () => {
            return {
                metadata: {},
                entities: null,
                page: 1
            }
        },

        beforeMount() {
            this.fops();
        },

        methods: {
            next() {
                this.startLoading();
                this.page++;
                this.fops();
            },

            prev() {
                this.startLoading();
                this.page--;
                this.fops();
            },

            fops() {
                this.startLoading();
                this.fopLatest(this.page).then(response => {
                    this.entities = [];
                    this.entities = response.data.data;
                    this.metadata = response.data.metadata;
                    this.stopLoading()
                }, this.onError)
            },

            hasEntities() {
                return this.entities && this.entities.length > 0;
            }
        }
    }
</script>

<style>

</style>
