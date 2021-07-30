<template>
  <main id="app">
    <Navbar />
    <ProductPage />
  </main>
</template>

<script>
import Navbar from "./components/Navbar/Navbar";
import ProductPage from "./components/ProductPage/ProductPage";
export default {
  components: {
    Navbar,
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
