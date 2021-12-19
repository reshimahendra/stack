<template>
    <div class="container">
        <div class="mt-6 pt-5 is-hidden-touch"></div>
        <div class="container mb-6">
        <h2 class="title has-text-centered mt-6">Fresh Dishes Ready</h2>
        <p class="has-text-centered mb-6">Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.</p>
        </div>
        <hr />
        <div class="container mt-6 has-background-light p-5">
            <div class="columns is-mobile p-5">
                <div class="column">
                    <h2 class="subtitle"><strong>Food Menu</strong></h2>
                </div>
                <div class="column has-text-right">
                    <router-link class="button is-primary" :to="{name: 'Food'}">Show All</router-link>
                </div>
            </div>
            <div class="columns is-multiline">
                <div class="column is-one-third" v-for="product of products" :key="product.id">
                    <CardProduct :product="product"></CardProduct>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import CardProduct from '@/components/CardProduct'
import axios from 'axios'

export default {
    name: 'Food',
    components: {
        CardProduct,
    }, 
    data() {
        return {
            products: [],
        }
    },
    mounted() {
        this.loadProducts()
    },
    methods: {
        async loadProducts() {
            await axios
                .get('/v1/product/')
                .then(res=>{
                    console.log(res.data.data)
                    this.products = res.data.data
                })
                .catch(err=>{
                    console.log(err, 'Error')
                })
        }
    },
}
</script>
