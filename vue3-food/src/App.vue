<template>
  <div class="container">
    <Navbar />
  </div>
  <router-view/>
</template>

<script>
import Navbar from '@/components/layout/Navbar'
export default {
    name:'App',
    components: {
        Navbar,
    }
}
</script>

<style lang="scss">
@import '../node_modules/bulma';
/* @import '../assets/scss/main.scss'; */
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;700&display=swap');

body{
    font-family: 'Montserrat', sans-serif;
    font-weight: 300;
}

.router-link-exact-active{
    font-weight: bold;
}
</style>
