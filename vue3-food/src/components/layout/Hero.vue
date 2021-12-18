<template>
    <div class="container">
        <div class="columns is-vcentered custom-columns">
            <div class="column is-two-fifth">
                <div class="is-flex is-justify-content-center">
                    <div class="has-text-right">
                        <h1 class="title is-uppercase"><strong>Delicious Food Menu</strong></h1>
                        <h2 class="subtitle">in your gadget</h2>
                        <p>Order your favorite food a click away</p>
                    </div>
                </div>
            </div>
            <div class="column">
                <img src="@/assets/img/hero2.png" width="100%" />
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name:'Hero',
}
</script>

<style lang="scss" scoped>
@media(max-width: 767px) {
  .custom-columns {
    flex-direction: column-reverse;
    display: flex;
  }
}
</style>
