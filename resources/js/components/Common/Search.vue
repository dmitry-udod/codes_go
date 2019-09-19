<template>
    <div class="form active m-0">
        <div class="row">
            <div class="col-xl-9">
                <label for="key">{{ title }}:</label>
                <input type="text" id="key" name="key" class="text" autofocus v-model="q" @keyup.enter="fireSearch()" :placeholder="placeholder">
                <span style="position: absolute;top: 40px;right: 25px;cursor: pointer" v-if="q" @click="clearSearch()">х</span>
                <div id="search-error" class="search-error"></div>
            </div>
            <div class="col-xl-3 d-flex align-items-end">
                <button class="btn_default yellow btn_search" :disabled="q === ''" @click="fireSearch()">Знайти</button>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        props: ['title', 'placeholder'],

        data: () => {
            return {
                q: '',
            }
        },

        methods: {
            fireSearch() {
                this.$emit('search', this.q)
            },

            clearSearch() {
                this.q = '';
                this.$emit('clearSearch')
            }
        }
    }
</script>

<style scoped>

</style>