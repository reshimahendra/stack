<template>
    <div class="container mt-6">
        <div class="columns is-centered py-6 px-5">
            <div class="column is-two-thirds">
                <div class="columns is-mobile">
                    <div class="column">
                        <h2 class="subtitle is-4">LBW FOOD</h2>
                        <ul class="about">
                            <li>Jl. Perum Abc</li>
                            <li>Ph. 08112324123</li>
                            <li>mail: abc@lbw.com</li>
                        </ul>
                    </div>
                    <div class="column">
                        <h2 class="subtitle is-4">Links</h2>
                        <ul class="footer-link">
                            <li><router-link :to="{name: 'Home'}">Home</router-link></li>
                            <li><router-link :to="{name: 'Food'}">Foods</router-link></li>
                            <li><router-link :to="{name: 'About'}">About</router-link></li>
                        </ul>
                    </div>
                </div>
                <div class="copyright is-hidden-mobile">
                <p>LBW Food &copy; www.lbwfood.com - 2021</p>
                </div>
            </div>
            <div class="column">
                <h2 class="subtitle is-4">Opening Hours</h2>
                <table class="table is-narrow table-footer">
                    <tr>
                        <td><b>Mon</b></td>
                        <td>09.00 - 20.00</td>
                    </tr>
                    <tr>
                        <td><b>Tue</b></td>
                        <td>09.00 - 20.00</td>
                    </tr>
                    <tr>
                        <td><b>Wed</b></td>
                        <td>09.00 - 20.00</td>
                    </tr>
                    <tr>
                        <td><b>Thu</b></td>
                        <td>09.00 - 20.00</td>
                    </tr>
                    <tr>
                        <td><b>Fri</b></td>
                        <td>09.00 - 20.00</td>
                    </tr>
                    <tr>
                        <td><b>Sat</b></td>
                        <td>09.00 - 22.00</td>
                    </tr>
                    <tr>
                        <td><b>Sun</b></td>
                        <td>Closed</td>
                    </tr>
                </table>
            </div>
                <div class="copyright is-hidden-tablet mt-6">
                <p>LBW Food &copy; www.lbwfood.com - 2021</p>
                </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Footer'
}
</script>

<style lang="scss" scoped>
h2, .copyright{
    color: #dadada;
}
.footer-link, 
.about {
    color: white;
    a{
        color: #dadada;
        &:hover,
        &:active,
        &:visited{
            color: #dedede;
        }
    }

}
.table-footer{
    background-color: #363636;
    color: #dadada;
    td {
        border: none;
    }

}
</style>
