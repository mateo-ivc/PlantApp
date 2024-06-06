<template>
  <div class="plants-container">
    <div v-for="plant in plants" :key="plant.name" class="plant">
      <div class="plant-text-container">
        <div>
          <router-link class="plant-link" :to="{name: 'plant', params: {id: plant.id}}">
            <div class="plant-image-container" style="height: 300px">
              <!-- Click event is removed from the image -->
              <img :src="plant.imageUrl" :id="'plant-image-'+ plant.id" width="100%" height="100%"
                   style="object-fit: cover;" alt=""/>
            </div>
            <h1 class="plant-name">{{ plant.name }}</h1>
          </router-link>
        </div>
        <div class="plant-information-container">
          <div class="plant-information-data">
            <a>Temperature:</a>
            <a>{{ plant.temperature }}</a>
          </div>
          <div class="plant-information-data">
            <a>Humidity:</a>
            <a>{{ plant.humidity }}</a>
          </div>
          <div class="plant-information-data">
            <a>Moisture:</a>
            <a>{{ plant.moisture }}</a>
          </div>
          <div class="plant-information-data">
            <a>Lighting:</a>
            <a>{{ plant.lighting }}</a>
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
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(275px, 1fr));
  gap: 2%;

}


.plants-container > div {
  box-shadow: 0 0 1px 1px #2c3e50;
  border-radius: 10px;
  margin-bottom: 10px;
  width: 100%; /* Ensures the box takes up the full grid cell width */
  min-width: 275px;

}


.plant-information-container {
  display: flex;
  flex-direction: column;
  align-items: start;
}

.plant-name {
  margin:  10px;
  color: black;

}

.plant-link {
  text-decoration: none;
}

.plant-image-container {

  border-radius: 10px 10px 0 0;
  overflow: hidden;
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