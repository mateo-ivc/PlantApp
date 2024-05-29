<script>

export default {
  props: {
    svg: {type: Object, required: true}
  },

}
</script>
<template>

  <component :is="svg" />

</template>

<style scoped>

</style>