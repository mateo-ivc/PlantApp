<template xmlns:div="http://www.w3.org/1999/html">
  <Sidebar/>
  <div :style="{'margin-left': sidebarWidth}">
    <router-view/> <!-- Display all routes -->
  </div>
<!--  <div id="nav">-->

<!--    <router-link to="/plants">Plant</router-link>-->
<!--  </div>-->

</template>

<script>


import Sidebar from "@/components/sidebar/SidebarComponent.vue";
import {sidebarWidth} from "@/components/sidebar/state";

export default {
  name: 'App',
  computed: {
    sidebarWidth() {
      return sidebarWidth
    }
  },
  components: {Sidebar}
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
