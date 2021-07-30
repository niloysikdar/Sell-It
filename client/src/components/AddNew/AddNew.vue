<template>
  <div class="new">
    <h2>Add a new Product</h2>
    <form>
      <input type="text" placeholder="Product Name" required />
      <input type="text" placeholder="Description" required />
      <input type="text" placeholder="Price" required />
      <input type="url" placeholder="Image URL" required />
      <input type="text" placeholder="Address" required />
      <input type="text" placeholder="Seller Name" required />
      <input type="text" placeholder="Contact Info" required />
      <button type="submit">ADD</button>
    </form>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.new {
  min-height: 100vh;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding-top: 110px;
  background: #c0e0e0;
}

.new h2 {
  margin-bottom: 30px;
  color: #1e57ad;
  font-size: 2rem;
}

.new form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-width: 400px;
  background: #e6f6fc;
  padding: 20px;
  border-radius: 15px;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.3);
}

.new form input {
  background: #dbdbdb;
  outline: none;
  border: none;
  color: #000;
  margin-bottom: 17px;
  padding: 10px;
  width: 100%;
  border-radius: 10px;
  font-size: 1.1rem;
}

.new form button {
  outline: none;
  border: none;
  border-radius: 10px;
  background: #347ce6;
  color: #fff;
  margin-top: 10px;
  padding: 8px 25px;
  font-size: 1.3rem;
  cursor: pointer;
  text-decoration: none;
}
</style>
