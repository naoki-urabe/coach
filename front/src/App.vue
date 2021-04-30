<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <Kuji />
</template>

<script>
import Kuji from './components/Kuji.vue'

export default {
  name: 'App',
  components: {
    Kuji
  }
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
