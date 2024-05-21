<template>
  <div class="w3-row" style="margin-top: 4em">
    <div class="w3-twothird w3-container">
      <loading v-model:active="isLoading" loader="bars" />
      <div class="w3-bar w3-grey">
        <button
          class="w3-bar-item w3-button w3-ripple w3-border-right"
          @click="doPaging(-1)"
        >
          &lt;&lt; Last Page
        </button>
        <button
          class="w3-bar-item w3-button w3-ripple w3-border-right"
          @click="doPaging(+1)"
        >
          Next Page &gt;&gt;
        </button>
        <SearchBar />
        <span class="w3-bar-item w3-right"
          >Page {{ currentPage }} of {{ lastPage }}</span
        >
      </div>
      <p />
      <table
        class="w3-table-all w3-hoverable"
        aria-label="A List of all my contacts"
      >
        <tr class="w3-grey">
          <th class="my-header">
            <SortHeader name="Last Name" sort="lastName" />
          </th>
          <th class="my-header">
            <SortHeader name="First Name" sort="firstName" />
          </th>
          <th class="my-header">
            <SortHeader name="Email" sort="email" />
          </th>
          <th class="my-header">
            <SortHeader name="Phone" sort="phone" />
          </th>
        </tr>
        <tr
          v-for="address in addresses"
          :key="address.id"
          @click="selectAddress(address)"
          :class="{
            'selected-row': isSelected(address),
          }"
        >
          <td>
            {{ address.lastName }}
          </td>
          <td>{{ address.firstName }}</td>
          <td>{{ address.email }}</td>
          <td>{{ address.phone }}</td>
        </tr>
      </table>
    </div>
    <div class="w3-third w3-container">
      <AddressEditor title="Create new address" create />
      <p />
      <AddressEditor title="Update address" />
      <p />
    </div>
  </div>
</template>

<script setup>
import { onMounted, computed } from "vue";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/css/index.css";
import AddressEditor from "../components/AddressEditor.vue";
import SortHeader from "../components/SortHeader.vue";
import SearchBar from "../components/SearchBar.vue";
import { useAddressStore } from "@/store/address";

const addressStore = useAddressStore();
const addresses = computed(() => addressStore.addresses);
// we add one to have it human readable (start to count with 1)
const currentPage = computed(() => addressStore.currentPage + 1);
const lastPage = computed(() => addressStore.lastPage);

const isLoading = computed(() => addressStore.loading);

const selectAddress = (address) => {
  addressStore.setSelected(address);
};

const isSelected = (address) => {
  return (
    addressStore.selectedAddress != null &&
    address != null &&
    addressStore.selectedAddress.id == address.id
  );
};

const doPaging = (dif) => {
  addressStore.doPaging(dif);
};

onMounted(() => {
  addressStore.loadAddresses();
});
</script>

<style>
tr:hover td {
  background-color: Silver;
}
.selected-row td {
  background-color: Silver !important;
}
.my-header {
  font-weight: normal !important;
  font-size: 18px;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  height: 2.6em;
  padding-top: 0 !important;
  padding-bottom: 0 !important;
}
</style>
