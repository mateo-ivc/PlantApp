<script>
import {collapsed, toggleSidebar, sidebarWidth} from "@/components/sidebar/state";
import SidebarLinksComponent from "@/components/sidebar/SidebarLinksComponent.vue";
import DoubleAngleBack from "@/components/icons/DoubleAngleBack.vue"
import HomeIcon from "@/components/icons/HomeIcon.vue";
import PlantIcon from "@/components/icons/PlantIcon.vue";
import DashboardIcon from "@/components/icons/DashboardIcon.vue";

export default {
  computed: {
    DashboardIcon() {
      return DashboardIcon
    },
    PlantIcon() {
      return PlantIcon
    },
    HomeIcon() {
      return HomeIcon
    },
    DoubleAngleBack() {
      return DoubleAngleBack
    }
  },
  components: {DoubleAngleBack, SidebarLinksComponent},
  props: {},
  setup() {
    return {collapsed, toggleSidebar, sidebarWidth}
  }
}
</script>

<template>
  <div class="sidebar" :style="{width: sidebarWidth}">

    <h1>
      <span v-if="collapsed">
        <div>V</div>
        <div>S</div>
      </span>
      <span v-else>
        Vue Sidebar
      </span>
    </h1>

    <SidebarLinksComponent to="/" :svg=HomeIcon>Home</SidebarLinksComponent>
    <SidebarLinksComponent to="/dashboard" :svg=DashboardIcon>Dashboard</SidebarLinksComponent>
    <span class="collapse-icon"  :class="{'rotate-180': collapsed}" @click="toggleSidebar">
    <double-angle-back/>
    </span>
  </div>
</template>

<style>
:root {
  --sidebar-bg-color: #2f855a;
  --sidebar-item-hover: #38a169;
  --sidebar-item-active: #276749;
}
</style>

<style scoped>
.sidebar {
  color: white;
  background-color: var(--sidebar-bg-color);

  float: left;

  z-index: 1;
  top: 0;
  left: 0;
  bottom: 0;
  padding: 0.5em;

  transition: 0.3s ease;

  display: flex;
  flex-direction: column;
  height: -webkit-fill-available;
}
.collapse-icon{
  position: absolute;
  bottom: 0;
  padding: 0.75em;

  color: rgba(255,255,255, 0.7);

  transition: 0.2s linear;
}

.rotate-180{
  transform: rotate(180deg);
  transition: 0.2s linear;
}
</style>