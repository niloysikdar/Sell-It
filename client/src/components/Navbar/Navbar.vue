<template>
  <header>
    <router-link to="/" class="logo">Sell It</router-link>
    <router-link to="/new" class="button">Add New</router-link>
  </header>
</template>

<script>
export default {};
</script>

<style scoped>
header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  padding: 18px 30px;
  background: #e6f6fc;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 2.2rem;
  letter-spacing: 0.5rem;
  text-transform: uppercase;
  color: #1c4e97;
  text-decoration: none;
  font-weight: 500;
}

.button {
  outline: none;
  border: none;
  border-radius: 10px;
  background: #347ce6;
  color: #fff;
  padding: 8px 18px;
  font-size: 1.3rem;
  margin-right: 30px;
  cursor: pointer;
  text-decoration: none;
}
</style>
