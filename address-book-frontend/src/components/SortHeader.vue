<template>
  <span
    class="w3-button my-button"
    :class="{ active: isActive() }"
    @click="sort"
    >{{ name }}
    <img v-if="!isDesc()" src="@/assets/sortAsc.png" alt="" class="icons" /><img
      v-if="isDesc()"
      src="@/assets/sortDesc.png"
      alt=""
      class="icons"
  /></span>
</template>
<script setup>
import { useAddressStore } from "@/store/address";

const addressStore = useAddressStore();
const props = defineProps(["name", "sort"]);

const isDesc = () => {
  return !addressStore.asc && addressStore.sort == props.sort;
};

const isActive = () => {
  return addressStore.sort == props.sort;
};

const sort = () => {
  addressStore.doSort(props.sort, !addressStore.asc);
};
</script>
<style>
.my-button {
  padding-left: 0;
}
.icons {
  height: 1.3em;
  width: auto;
  margin: -1em 0.1em -0.8em 0em;
}
.active {
  font-weight: bold;
}
</style>
