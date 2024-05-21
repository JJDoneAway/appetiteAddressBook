<template>
  <input
    class="w3-bar-item w3-input w3-round w3-hover-light-grey my-input"
    placeholder="Search..."
    v-model="value"
    @keyup.enter="search()"
  />
  <img
    src="@/assets/search.png"
    alt=""
    class="w3-bar-item w3-button w3-ripple my-icon"
    style="height: 2.1em; margin-top: 0.19em"
  />
  <button
    class="w3-bar-item w3-button w3-ripple w3-border-right w3-border-left"
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
  height: 2.1em;
  width: auto;
  margin-left: -3em;
  margin-top: 0.19em;
}
.my-input {
  height: 2.1em;
  margin-top: 0.2em;
  margin-left: 1em;
  width: 25em !important;
}
</style>
