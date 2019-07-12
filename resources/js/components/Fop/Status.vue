<template>
    <span :style="`color: ${color()}`">
        {{ status }}
    </span>
</template>

<script>
    export default {
        props: ['status'],
        methods: {
            color() {
                if (this.status === "зареєстровано") {
                    return 'green';
                }

                if (this.status === "припинено") {
                    return 'red';
                }

                return '#1D1D1B';
            }
        }
    }
</script>

<style>

</style>
