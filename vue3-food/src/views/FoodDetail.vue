<template>
    <div class="container">
        <nav class="breadcrumb mt-5" aria-label="breadcrumbs">
            <ul>
                <li><router-link :to="{name:'Home'}">Home</router-link></li>
                <li><router-link :to="{name:'Food'}">Food</router-link></li>
            </ul>
        </nav>
        <div class="tile is-ancestor mt-3 mb-5">
            <div class="tile is-child">
                <figure class="image is-4by3">
                    <img :src="product.picture" class="box-image">
                </figure>
            </div>
            <div class="tile is-child">
                <div class="box-detail">
                    <h2 class="subtitle has-text-small mb-3 is-uppercase"><b>{{ product.prod.name }}</b></h2>
                    <div class="columns is-mobile pl-3 mb-0">
                        <div class="column is-one-third">
                            <ul>
                                <li>Price</li>
                                <li>Type</li>
                            </ul>
                        </div>
                        <div class="column">
                            <ul>
                                <li>{{ format(product.prod.price, 'idr') }}</li>
                                <li>{{ product.type }}</li>
                            </ul>
                        </div>
                    </div>
                    <form @submit.prevent="submitOrder" class="mt-3">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="number" v-model="order.quantity" min="1" max="20">
                            </div>
                        </div>
                        <div class="field">
                            <div class="control">
                                <textarea class="textarea" v-model="order.note" placeholder="Additional note for your order" rows="2"></textarea>
                            </div>
                        </div>

                        <div class="control">
                            <button class="button is-primary" type="submit">Place Order</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios'
    import { toast }from 'bulma-toast'

    export default {
        name: 'FoodDetail',
        components: {
            // CardDetail,
        },
        data() {
            return {
                productDefault: {
                    id: 0,
                    prod: {
                        sku: '',
                        name: '',
                        price: 0,
                        is_ready: false,
                        is_available: false,
                        picture: require('@/assets/img/default.jpg'),
                    },
                },
                product : {
                    id : 0,
                    prod:{}
                },
                order: {
                    quantity    : 1,
                    note   : '',
                },

            }
        },
        mounted() {
            this.loadProduct()
        },
        methods: {
            async loadProduct(){
                const id = this.$route.params.id
                await axios
                    .get(`/v1/product/${id}`)
                    .then( res => {
                        this.product = res.data.data ? res.data.data : this.productDefault
                        let pic = res.data.data ? require('@/assets/img/' + res.data.data.prod.picture) : require('@/assets/img/default.jpg')
                        console.log(this.product)
                        this.product.picture = pic
                    })
                    .catch(err=>{
                        console.log('Error', err)
                    })
                await axios
                    .get(`/v1/product-type/${this.product.prod.type_id}`)
                    .then(res=>{
                        this.product.type = res.data.data.type.name ? res.data.data.type.name : '-'
                    })
                    .catch(err=>{
                        console.log('Error', err)
                    })
            },
            submitOrder() {
                this.order.product = this.product
                console.log('Order submited.\n(' + this.order.quantity+ ')'+ this.product.name+ '\n'+ this.order.note)
                toast({
                    message: `'${this.product.name}' has been added.`,
                    type: 'is-success',
                    dismissible: true,
                    duration: 2000,
                    position: 'bottom-right',
                })
            },
            format(n, currency) {
                return new Intl.NumberFormat('id', {
                    style: 'currency',
                    currency: currency
                }).format(n)
            }
        },
    }
</script>

<style lang="scss" scoped>
.box-detail {
    height: 100%;
    /* margin-left: 1rem; */
    padding: 3rem;
}
.box-image{
    border-radius: 1rem;
}
</style>
