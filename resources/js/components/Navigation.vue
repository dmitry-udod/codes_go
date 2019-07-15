<template>
    <section class="header-bg">
        <a class="main-logo" href="/">
            <div class="icon"></div>
            <span>da.org.ua</span>
            <span class="light">Збірка відкритих даних</span>
        </a>
        <div class="row header_content justify-content-center">
            <div class="col-md-1"></div>
            <div class="col-md-4 col-lg-3 d-lg-none menu-toggle_box">
                <div id="menuToggle" class="menu-toggle">
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div>
        </div>
        <div class="header_bottom">
            <nav class="menu" id="menu">
                <ul class="menu__list" id="menu-desktop">
                    <li class="menu__list-item" v-bind:class="{'active' : $route.name === 'fop' ||  $route.name === 'main'}">
                        <a @click.stop.prevent="$router.push({'name': 'fop'})">ФОПи</a>
                    </li>
                    <li class="menu__list-item" v-bind:class="{'active' : $route.name === 'legal_entities'}">
                        <a @click.stop.prevent="$router.push({'name': 'legal_entities'})">Юридичні особи</a>
                    </li>
                </ul>
            </nav>
        </div>
    </section>
</template>

<script>
    export default {}
</script>

<style>

</style>
