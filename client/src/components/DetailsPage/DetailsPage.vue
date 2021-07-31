<template>
  <div class="main">
    <h2>Product Details</h2>
    <div class="details">
      <div class="image-wrapper">
        <img :src="product.imageurl" :alt="product.productName" />
      </div>
      <div class="details-text">
        <h2 class="productName">{{ product.productName }}</h2>
        <h3 class="price">{{ product.price }}</h3>
        <h4 class="description">{{ product.description }}</h4>
        <h3 class="sellerInfo">Seller: {{ product.sellerName }}</h3>
        <h3 class="sellerInfo">Address: {{ product.address }}</h3>
        <h3 class="sellerInfo">Contact: {{ product.contactInfo }}</h3>
      </div>
    </div>
    <button>Buy Now</button>
  </div>
</template>

<script>
import { useRoute } from "vue-router";
import axios from "axios";

export default {
  async mounted() {
    const id = useRoute().params.id;
    const productURL = `${process.env.VUE_APP_BASEURL}/${id}`;
    const result = await axios.get(productURL);
    if (result.data) {
      this.product = result.data;
    }
  },
  data() {
    return {
      product: {},
    };
  },
};
</script>

<style scoped>
.main {
  min-height: 100vh;
  width: 100%;
  padding: 50px;
  padding-top: 110px;
  background: #c0e0e0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
}
.main h2 {
  margin-bottom: 3rem;
  color: #1e57ad;
  font-size: 2rem;
}

.details {
  display: flex;
  gap: 50px;
}

.image-wrapper {
  height: 250px;
  width: 200px;
  border-radius: 10px;
  overflow: hidden;
}

.image-wrapper img {
  height: 100%;
  width: 100%;
}

.details-text {
  max-width: 550px;
  margin-bottom: 5rem;
}

.details-text .productName {
  font-size: 2.2rem;
  margin-bottom: 0.5rem;
}

.details-text .price {
  font-size: 1.7rem;
  color: #4e89e5;
  margin-bottom: 1.2rem;
}

.details-text .description {
  font-size: 1.5rem;
  color: #2051a0;
  font-weight: 400;
  margin-bottom: 2rem;
}

.details-text .sellerInfo {
  color: #0e2d5c;
  font-size: 1.35rem;
}

button {
  outline: none;
  border: none;
  border-radius: 10px;
  background: #347ce6;
  color: #fff;
  padding: 10px 25px;
  font-size: 1.5rem;
  cursor: pointer;
}
</style>
