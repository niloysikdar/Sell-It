<template>
  <main id="app">
    <ProductPage />
  </main>
</template>

<script>
import ProductPage from "./components/ProductPage.vue";
export default {
  components: {
    ProductPage,
  },
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
