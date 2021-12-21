<template>
  <div class="container">
    <Navbar />
  </div>
  <router-view/>
  <div class="has-background-dark">
    <Footer />
  </div>
</template>

<script>
import Navbar from '@/components/layout/Navbar'
import Footer from '@/components/layout/Footer'
export default {
    name:'App',
    components: {
        Navbar,
        Footer,
    },
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
