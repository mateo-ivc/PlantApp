<template xmlns:div="http://www.w3.org/1999/html">
  <Sidebar/>
  <div id="main">


    <router-view/> <!-- Display all routes -->
  </div>

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
body {
  margin: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
}
#main{

}

</style>
