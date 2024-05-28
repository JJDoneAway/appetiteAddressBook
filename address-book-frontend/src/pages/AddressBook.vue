<template>
  <div class="w3-row" style="margin-top: 5em">
    <div class="w3-twothird w3-container">
      <loading v-model:active="isLoading" loader="bars" />
      <div class="w3-bar w3-card-4 w3-white" style="padding-left: 1em">
        <header>&nbsp;</header>
        <SearchBar />
        <footer><h3>&nbsp;</h3></footer>
      </div>
      <p />
      <div class="w3-card-4">
        <table
          class="w3-table-all w3-hoverable w3-container"
          aria-label="A List of all my contacts"
        >
          <tr class="w3-white">
            <th>
              <SortHeader name="Last Name" sort="lastName" />
            </th>
            <th>
              <SortHeader name="First Name" sort="firstName" />
            </th>
            <th>
              <SortHeader name="Email" sort="email" />
            </th>
            <th>
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
      <p />
      <div class="w3-bar w3-card-4 w3-white">
        <header>&nbsp;</header>
        <div style="margin-bottom: 4em">
          <button
            class="w3-bar-item w3-button w3-ripple w3-right w3-border"
            @click="doPaging(+1)"
            style="margin-right: 1em;"
          >
            Next Page &gt;&gt;
          </button>
          <span class="w3-bar-item w3-right"
            >Page {{ currentPage }} of {{ lastPage }}</span
          >
          <button
            class="w3-bar-item w3-button w3-ripple w3-right w3-border"
            @click="doPaging(-1)"
          >
            &lt;&lt; Last Page
          </button>
        </div>
      </div>
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
</style>
