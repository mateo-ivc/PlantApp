<template xmlns:div="http://www.w3.org/1999/html">
    <Sidebar/>
  <div id="main">
    <router-view/> <!-- Display all routes -->
  </div>

</template>

<script>


import Sidebar from "@/components/sidebar/SidebarComponent.vue";

export default {
  name: 'App',
  components: {Sidebar},
}
</script>

<style>
body {
  margin: 0;

}


#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
  width: 100%;
  max-width: 80rem;
  margin: 0 auto;
  padding: 0 2rem;
  box-sizing: border-box;


}

</style>
