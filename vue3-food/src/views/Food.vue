<template>
    <div class="container">
        <div class="container">
        <h2 class="title has-text-centered mt-5">Fresh Dishes Ready</h2>
        <p class="has-text-centered">Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.</p>
        </div>
        <hr />
        <div class="field">
            <div class="control is-medium" :class="{'is-loading': loading}">
                <input class="input is-medium" v-model="searchText" @keyup.enter="loadProducts()" type="text" placeholder="Search delicious food on our menu">
            </div>
        </div>
        
        <div class="container has-background-light p-5">
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
                    <CardProduct :id="product.id" :product="product.prod"></CardProduct>
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
            loading: false,
            searchText: '',
        }
    },
    mounted() {
        this.loadProducts(this.endpoint)
    },
    computed: {
        apiEndpoint() {
            return this.searchText !== "" ? "name="+this.searchText : "limit/9"
        } 
    },
    methods: {
        async loadProducts() {
            this.loading = true
            await axios
                .get(`/v1/product/` + this.apiEndpoint)
                .then(res=>{
                    // console.log(res.data.data)
                    this.products = res.data.data ? res.data.data : []
                })
                .catch(err=>{
                    console.log(err, 'Error')
                })

            this.loading = false
        },
    },
}
</script>
