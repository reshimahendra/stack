<template>
  <div class="container">
    <Hero />

    <div class="container mt-6 has-background-light p-5">
        <div class="columns is-mobile p-5">
            <div class="column">
                <h2 class="subtitle"><strong>Best Foods</strong></h2>
            </div>
            <div class="column has-text-right">
                <router-link class="button is-primary" :to="{name: 'Food'}">Show All</router-link>
            </div>
        </div>
        <div class="columns">
            <div class="column" v-for="product of products" :key="product.id">
                <CardProduct :product="product"></CardProduct>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Hero from '@/components/layout/Hero'
import CardProduct from '@/components/CardProduct'
import axios from 'axios'

export default {
    name: 'Home',
    data() {
        return {
            products: []
        }
    },
    components: {
        Hero,
        CardProduct
    }, 
    mounted() {
        this.loadData()
    }, 
    methods: {
        async loadData() {
            await axios
                .get('/v1/product/limit/3')
                .then( res => {
                    /* console.log(res.data, 'success') */
                    this.products = res.data.data ? res.data.data : []
                })
                .catch(err => {
                    console.log(err, 'error')
                })
        }
    },
}
</script>

<style lang="scss">

</style>
