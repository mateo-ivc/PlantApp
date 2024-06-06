<script>
import DoubleAngleBack from "@/components/icons/DoubleAngleBack.vue"
import HomeIcon from "@/components/icons/HomeIcon.vue";
import PlantIcon from "@/components/icons/PlantIcon.vue";
import DashboardIcon from "@/components/icons/DashboardIcon.vue";
import SidebarLinksComponent from "@/components/sidebar/SidebarLinksComponent.vue";
import AccountIcon from "@/components/icons/AccountIcon.vue";

export default {
  computed: {
    AccountIcon() {
      return AccountIcon
    },
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
  components: {SidebarLinksComponent},
  props: {},

}
</script>

<template>
  <div class="sidebar">

    <div class="nav-item" id="logo">
      <a> Home</a>

    </div>
    <div class="nav-item">

    </div>
    <div class="nav-item" id="icons">
      <SidebarLinksComponent to="/dashboard" :svg=DashboardIcon>Dashboard</SidebarLinksComponent>
      <SidebarLinksComponent to="/account" :svg="AccountIcon">Account</SidebarLinksComponent>
    </div>
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
  display: flex;
  gap: 2%;
  font-size: 2rem;

}

.nav-item {
  margin: 10px 0;
  width: 100%;

}
#icons {
  display: flex;
  flex-direction: row;
  justify-content: right;
  gap: 10px;

}

#icons > * {
  width: 30px;
  height: 30px;
  align-self: center;

}

#logo {
  text-align: left;
}


</style>