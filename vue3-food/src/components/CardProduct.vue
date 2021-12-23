<template>
    <div class="card p-5">
        <div class="card-image">
            <span class="tag is-pulled-right is-info has-text-weight-medium" style="position:absolute;z-index:2;top:10px;right:10px;">{{ format(product.price, 'idr') }}</span>
            <figure class="image is-4by3">
                <img :src="require('@/assets/img/' + product.picture)">
            </figure>
        </div>
        <div class="content mt-3">
            <h2 class="subtitle is-5">{{ product.name }}</h2>
            <router-link class="button is-primary" :to="{name:'FoodDetail', params:{id:id}}">Add to Cart</router-link>
        </div>
    </div>
</template>

<script>
export default {
    props: ['product','id'],
    name: 'CardProduct',
    methods: {
        format(n, currency) {
            return new Intl.NumberFormat('id', {
                style: 'currency',
                currency: currency
            }).format(n)
        }
    },
}
</script>
