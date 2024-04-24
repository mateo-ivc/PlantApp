<template>
  <h1>Plants</h1>
  <div v-for="plant in plants" :key="plant.id" class="plant">
    <router-link :to="{name: 'PlantDetails', params: {id: plant.id}}">
      <h2>
        {{ plant.name }}
      </h2>
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