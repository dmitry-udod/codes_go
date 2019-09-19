<template>
    <div>
        <div class="content" v-if="entity">
            <h1>
                {{ entity.known_names[0].last_name }}
                {{ entity.known_names[0].first_name }}
                {{ entity.known_names[0].middle_name }}
                {{ entity.known_names[0].additional_name }}
                ({{entity.number_in_list}})
            </h1>

            <ul>
                <li v-if="entity.known_names && entity.known_names.length > 1">
                    Також вiдомий як:
                    <ul>
                        <li v-for="(e, index) in entity.known_names" v-if="index > 0">
                            {{ e.last_name }}
                            {{ e.first_name }}
                            {{ e.middle_name }}
                            {{ e.additional_name }}
                        </li>
                    </ul>
                </li>
                <li v-if="entity.birth_places">Місце народження: {{ entity.birth_places.join(', ') }}</li>
                <li>Дата народження: {{ entity.birth_day }}</li>
                <li v-if="entity.nationalities">Нацiональнiсть: {{ entity.nationalities.join(', ') }}</li>
                <li>Назва органа видачі: {{ entity.source }}</li>
                <li>Дата внесння: {{ entity.added_at }}</li>
                <li v-if="entity.comments">Додатково: {{ entity.comments }}</li>
            </ul>

            <div class="text-center">
                <router-link class="btn btn_default blue d-inline-block" :to="{name: $route.params.q ? 'terrorists_search' : 'terrorists', params: $route.params}">До списку</router-link>
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
                this.terroristDetails(this.$route.params.id).then(response => {
                    this.stopLoading();
                    this.entity = response.data.data;
                }, this.onError)
            },
        }
    }
</script>

<style>

</style>
