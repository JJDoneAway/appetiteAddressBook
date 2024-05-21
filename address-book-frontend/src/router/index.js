import { createRouter, createWebHistory } from "vue-router";
import AddressBook from "../pages/AddressBook.vue";

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "Editor",
      component: AddressBook,
    },
  ],
});
