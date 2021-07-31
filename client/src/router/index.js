import { createWebHistory, createRouter } from "vue-router";
import ProductPage from "../components/ProductPage/ProductPage.vue";
import AddNew from "../components/AddNew/AddNew.vue";
import DetailsPage from "../components/DetailsPage/DetailsPage.vue";

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
  {
    path: "/details/:id",
    name: "Details",
    component: DetailsPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
