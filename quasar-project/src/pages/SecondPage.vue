<template>
  <q-page class="row">
    <p>This is the <span class="page-text">SECOND</span> page</p>
  </q-page>
</template>

<script lang="ts">
export default {
  name: 'SecondPage',
};
</script>

<script setup lang="ts">
//Most code goes here
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.page-text {
  color: red;
}
</style>
