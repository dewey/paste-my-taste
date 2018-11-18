<template>
  <div id="app">
    <PopularArtists msg="Paste Your Taste for Last.FM"/>
  </div>
</template>

<script>
import PopularArtists from './components/PopularArtists.vue'

export default {
  name: 'app',
  components: {
    PopularArtists 
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
