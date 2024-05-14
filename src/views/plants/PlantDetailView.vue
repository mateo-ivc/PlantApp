<template>
  <div class="detailed-plant-view">
    <div class="detailed-plant-image-container">
      <img id="plant-image" src="" style="object-fit: cover" alt="" >
      <div class="upload-image">
        <label for="file-upload" class="upload-image-button">Upload File</label>
        <input id="file-upload" @change="sendImage" type="file" class="upload-image-button" name="filename">
      </div>
    </div>
  </div>
  <p>Plant name: {{ plant.name }}</p>
</template>

<script>


export default {
  data() {
    return {
      plant: {},
    }
  }, mounted() {
    this.fetchPlant()
    this.fetchImage()
  }, methods: {
    fetchPlant() {
      const id = this.$route.params.id
      fetch('http://localhost:8080/plants/'.concat(id)).then(res => res.json()).then(data => this.plant = data).catch(err => console.log(err.message))
    },
    fetchImage() {
      const id = this.$route.params.id
      fetch('http://localhost:8080/plants/'.concat(id).concat("/image"), {
        method: 'GET'
      }).then(response => response.text())
          .then(base64Img => {
            const imgElement = document.getElementById('plant-image');
            imgElement.src = `data:image/png;base64,${base64Img}`;
          })
          .catch(error => console.error('Error:', error));
    },
    sendImage(e) {
      var files = e.target.files || e.dataTransfer.files;

      const upload = (file) => {
        fetch('http://localhost:8080/plants/'.concat(this.plant.id).concat("/image"), { // Your POST endpoint
          method: 'POST',
          headers: {
            "Content-Type": "image/jpg"
          },
          body: file // This is your file object
        }).then(
            response => response.json() // if the response is a JSON object
        ).then(
            success => console.log(success) // Handle the success response object
        ).catch(
            error => console.log(error) // Handle the error response object
        );
      };
      upload(files[0])
    },
  },
}
</script>

<style>
input[type="file"] {
  display: none;
}

.detailed-plant-view {
  width: 100%;
  display: flex;
  justify-content: center;
}

.detailed-plant-image-container {
  position: relative;
}

.upload-image {
  position: absolute;
  bottom: 10px;
  left: 10px;
}

.upload-image-button {
  border: 1px solid #2f855a;
  border-radius: 10px;
  background-color: white;
  display: inline-block;
  padding: 6px 12px;
  cursor: pointer;
}

#plant-image{
  width: 400px;
}
</style>