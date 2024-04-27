<template>
  <h1>Plants</h1>
  <div v-for="plant in plants" :key="plant.name" class="plant">
      <router-link :to="{name: 'plant', params: {id: plant.id}}">
        {{ plant.name }}
      </router-link>
  </div>
</template>
<script>


export default {
  data() {
    return {
      plants: []
    }
  }, mounted() {
    fetch('http://localhost:8080/plants').then((res) => res.json()).then(data => this.plants = data).catch(err => console.log(err.message))
  }
}

</script>