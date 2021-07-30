<template>
  <section class="products-main">
    <h2>Products for sale</h2>
    <div class="products">
      <Product
        v-for="product in products"
        :key="product.ID"
        :product="product"
      />
    </div>
  </section>
</template>

<script>
import Product from "./Product.vue";
import axios from "axios";

export default {
  components: {
    Product,
  },
  async mounted() {
    const URL = process.env.VUE_APP_BASEURL;
    const result = await axios.get(URL);
    this.products = result.data;
  },
  data() {
    return {
      products: [],
    };
  },
};
</script>

<style scoped>
.products-main {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  padding: 30px;
  padding-top: 110px;
  background: #c0e0e0;
}

.products-main h2 {
  margin-bottom: 30px;
  color: #1e57ad;
  font-size: 2rem;
}

.products {
  min-height: 100%;
  min-width: 100%;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 50px;
}

.new {
  position: fixed;
  top: 1.5rem;
  right: 5rem;
  height: 100px;
  width: 100px;
  font-size: 2rem;
  border-radius: 50%;
  border: none;
  outline: none;
}
</style>
