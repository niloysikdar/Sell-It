<template>
  <div class="product">
    <div class="product-image">
      <img :src="product.imageurl" height="250" width="300" />
    </div>
    <div class="product-text">
      <h3>{{ product.productName }}</h3>
      <h4>{{ product.price }}</h4>
      <p>{{ product.description }}</p>
      <div class="buttons">
        <button>Details</button>
        <button>Edit</button>
      </div>
      <button class="delete">Delete</button>
    </div>
  </div>
</template>

<script>
export default {
  props: ["product"],
};
</script>

<style scoped>
.product {
  max-width: 300px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 15px;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.3);
  text-align: center;
  overflow: hidden;
  background: #e6f6fc;
}

.product-text {
  padding: 10px 10px;
}

.product-text h3 {
  color: #3977ce;
  font-size: 1.4rem;
}

.product-text h4 {
  color: #4383e4;
  font-size: 1.1rem;
}

.product-text p {
  color: #1c4e97;
  font-weight: 500;
  font-size: 1.1rem;
  margin: 12px 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.product-text button {
  outline: none;
  border: none;
  border-radius: 10px;
  background: #347ce6;
  color: #fff;
  padding: 8px 18px;
  font-size: 1.1rem;
  margin: 5px 10px;
  cursor: pointer;
}

.delete {
  background: #e90000 !important;
}
</style>
