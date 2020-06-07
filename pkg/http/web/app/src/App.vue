<template>
  <v-app>
    <router-view></router-view>
    <Footer></Footer>
  </v-app>
</template>

<script>
import Footer from './components/Footer'

export default {
  name: 'App',
  components: { Footer },
  data: () => ({})
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
