<template>
  <h1>PlantDetailView</h1>
  <p>Plantname: {{ plant.name }}</p>
</template>

<script>


export default {
  data() {

    return {
      plant: {}
    }
  }, mounted() {
    this.fetchPlant()
  }, methods: {
    fetchPlant() {
      const id = this.$route.params.id
      fetch('http://localhost:8080/plants/'.concat(id)).then(res => res.json()).then(data => this.plant = data).catch(err => console.log(err.message))
    }
  }
}
</script>

<style scoped>

</style>