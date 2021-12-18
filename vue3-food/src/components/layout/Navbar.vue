<template>
    <nav class="navbar p-3" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <router-link class="navbar-item" :to="{name: 'Home'}">LBW</router-link>

        <a @click="makeBurger" :class="{'is-active': activator}" role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="nav-food">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="nav-food" class="navbar-menu" :class="{'is-active': activator}">
        <div class="navbar-item">
            <router-link class="navbar-item" :to="{name: 'Home'}">Home</router-link>
            <router-link class="navbar-item" :to="{name: 'Foods'}">Foods</router-link>
            <router-link class="navbar-item" :to="{name: 'About'}">About</router-link>
        </div>

        <div class="navbar-end">
          <router-link class="navbar-item " :to="{name: 'About'}">Cart <span class="tag is-primary">0</span></router-link>
          <div class="navbar-item">
            <div class="buttons">
              <a class="button is-primary">
                <strong>Sign up</strong>
              </a>
              <a class="button is-light">
                Log in
              </a>
            </div>
          </div>
        </div>
      </div>
    </nav>
</template>

<script>
export default {
    name: 'Navbar',
    data() {
        return {
            activator: false,
        }
    },
    methods: {
        makeBurger () {
            this.activator = !this.activator
            return this.activator
        }
    }
}
</script>
