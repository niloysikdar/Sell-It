import { createWebHistory, createRouter } from "vue-router";
import ProductPage from "../components/ProductPage/ProductPage.vue";
import AddNew from "../components/AddNew/AddNew.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: ProductPage,
  },
  {
    path: "/new",
    name: "New",
    component: AddNew,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
