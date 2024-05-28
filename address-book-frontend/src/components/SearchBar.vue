<template>
  <input
    class="w3-bar-item w3-input w3-border w3-hover-light-grey"
    style="width: 30em"
    placeholder="Search..."
    v-model="value"
    @keyup.enter="search()"
  />
  <img
    src="@/assets/search.png"
    alt=""
    class="w3-bar-item w3-button w3-border w3-ripple my-icon"
  />
  <button
    class="w3-bar-item w3-button w3-border w3-ripple"
    style="margin-left: 0.5em"
    @click="clear()"
  >
    Clear
  </button>
</template>
<script setup>
import { useAddressStore } from "@/store/address";
import { ref } from "vue";
const addressStore = useAddressStore();
const value = ref(addressStore.filter);

const search = () => {
  addressStore.doSearch(value.value);
};
const clear = () => {
  value.value = "";
  addressStore.doClear();
};
</script>
<style>
.my-icon {
  height: 2.7em;
  width: auto;
  margin-left: 0.5em;
}
</style>
