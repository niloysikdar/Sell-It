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
@import url("https://fonts.googleapis.com/css2?family=Roboto:wght@400;500&display=swap");
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: "Roboto", sans-serif;
}
</style>
