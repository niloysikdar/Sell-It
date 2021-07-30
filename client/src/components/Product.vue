<template>
  <div class="product">
    <div class="product-image">
      <img :src="product.imageurl" height="250" width="250" />
    </div>
    <div class="product-text">
      <h3>{{ product.productName }}</h3>
      <h4>{{ product.price }}</h4>
      <p>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Repudiandae
        unde
      </p>
    </div>
    <div class="see-more"></div>
  </div>
</template>

<script>
export default {
  props: ["product"],
};
</script>

<style scoped>
.product {
  max-width: 250px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 15px;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5);
  text-align: center;
  overflow: hidden;
}
</style>
