<template>
  <section class="products-main">
    <h2>All products for sale</h2>
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
export default {
  components: {
    Product,
  },
  data() {
    return {
      products: [
        {
          ID: 1,
          CreatedAt: "2021-07-29T03:49:58.9381337+05:30",
          UpdatedAt: "2021-07-29T03:49:58.9381337+05:30",
          DeletedAt: null,
          productName: "Math Book",
          description: "I want to sell my Math Book which is only 1 year old",
          price: "Rs. 350",
          imageurl:
            "https://images-na.ssl-images-amazon.com/images/I/81OrIk+OffL.jpg",
          address: "India",
          sellerName: "Niloy Sikdar",
          contactInfo: "9541236547",
          isSold: false,
        },
        {
          ID: 2,
          CreatedAt: "2021-07-29T03:49:58.9381337+05:30",
          UpdatedAt: "2021-07-29T03:49:58.9381337+05:30",
          DeletedAt: null,
          productName: "Science Book",
          description:
            "I want to sell my Science Book which is only 1 year old",
          price: "Rs. 350",
          imageurl:
            "https://images-na.ssl-images-amazon.com/images/I/81OrIk+OffL.jpg",
          address: "India",
          sellerName: "Niloy Sikdar",
          contactInfo: "9541236547",
          isSold: false,
        },
        {
          ID: 2,
          CreatedAt: "2021-07-29T03:49:58.9381337+05:30",
          UpdatedAt: "2021-07-29T03:49:58.9381337+05:30",
          DeletedAt: null,
          productName: "Science Book",
          description:
            "I want to sell my Science Book which is only 1 year old",
          price: "Rs. 350",
          imageurl:
            "https://images-na.ssl-images-amazon.com/images/I/81OrIk+OffL.jpg",
          address: "India",
          sellerName: "Niloy Sikdar",
          contactInfo: "9541236547",
          isSold: false,
        },
      ],
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
}

.products-main h2 {
  margin-bottom: 30px;
}

.products {
  min-height: 100%;
  min-width: 100%;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 50px;
}
</style>
