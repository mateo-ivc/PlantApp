<script>

export default {
  props: {
    svg: {type: Object, required: true},
    to: {type: String, required: true}
  },

}
</script>
<template>

  <a :href="to">
  <component  :is="svg" />
  </a>
</template>

<style scoped>

</style>