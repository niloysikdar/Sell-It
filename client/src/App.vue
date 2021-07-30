<template>
  <main id="app">
    <Navbar />
    <router-view />
  </main>
</template>

<script>
import Navbar from "./components/Navbar/Navbar";
export default {
  components: {
    Navbar,
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
