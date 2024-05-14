<template>
  <div class="plants-container">
    <div v-for="plant in plants" :key="plant.name" class="plant">
      <div class="plant-image-container">
        <!-- Click event is removed from the image -->
        <img :src="plant.imageUrl" :id="'plant-image-'+ plant.id" width="100%" height="300px" style="object-fit: cover" alt=""/>
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
  },
  mounted() {
    // Fetch plants data
    fetch('http://localhost:8080/plants')
        .then(res => res.json())
        .then(data => {
          // Set plants data
          this.plants = data;
          // Fetch images for each plant
          this.plants.forEach(plant => this.fetchImage(plant.id));
        })
        .catch(err => console.error(err));
  },
  methods: {
    fetchImage(id) {
      fetch(`http://localhost:8080/plants/${id}/image`, {
        method: 'GET'
      })
          .then(response => response.text())
          .then(base64Img => {
            // Update the imageUrl property of the corresponding plant
            const plant = this.plants.find(p => p.id === id);
            if (plant) {
              plant.imageUrl = `data:image/png;base64,${base64Img}`;
            }
          })
          .catch(error => console.error('Error:', error));
    }
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

.plant {}

.plant-image-container {
  max-width: 300px;
}

.plant-information-container {
  display: flex;
  flex-direction: column;
  align-items: start;
}

.plant-text-container {
  padding: 10px
}

.plant-information-data {
  display: flex;
  justify-content: space-between;
  width: 80%;
}
</style>