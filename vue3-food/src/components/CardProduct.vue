<template>
    <div class="card p-5">
        <div class="card-image">
            <figure class="image is-4by3">
                <img :src="require('@/assets/img/' + product.gambar)" :alt="product.nama">
            </figure>
        </div>
        <div class="content mt-3">
            <h2 class="subtitle is-5">{{ product.nama }}</h2>
            <input class="button is-primary" type="button" value="Order">
            
        </div>
    </div>
</template>

<script>
export default {
    props: ['product'],
    name: 'CardProduct'
}
</script>
