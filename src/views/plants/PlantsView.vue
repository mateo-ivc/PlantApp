<template>
  <div class="plants-container">
    <div v-for="plant in plants" :key="plant.name" class="plant">
      <div class="plant-image-container">
        <img src="https://placehold.co/600x400" width="100%" height="300px" style="object-fit: cover" alt=""/>
      </div>
      <div class="plant-text-container">
        <div>
          <router-link :to="{name: 'plant', params: {id: plant.id}}">
            <h1>{{ plant.name }}</h1>
          </router-link>
        </div>
        <div class="plant-information-container">
          <div class="plant-information-data">
            <a>Temperature:</a>
            <a>10</a>
          </div>
          <div class="plant-information-data">
            <a>Humidity:</a>
            <a>12</a>
          </div>
          <div class="plant-information-data">
            <a>Moisture:</a>
            <a>14</a>
          </div>
        </div>
      </div>
    </div>
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

<style>
.plants-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  padding: 10px 10px 0 10px;
  gap: 2%;
}

.plants-container > div {

  box-shadow: 0 0 1px 1px gray;
  margin-bottom: 10px;
}

.plant {

}

.plant-image-container {
  max-width: 300px;
}

.plant-information-container {
  display: flex;
  flex-direction: column;
  align-items: start;

}

.plant-text-container{
  padding: 10px
}

.plant-information-data{
  display: flex;
  justify-content: space-between;
  width: 80%;
}
</style>