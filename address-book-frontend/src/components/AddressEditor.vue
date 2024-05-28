<template>
  <div v-if="showEditor" class="w3-card-4">
    <!-- title-->
    <header class="w3-container w3-white">
      <h2>{{ title }}</h2>
    </header>

    <!-- Validation errors -->
    <div v-if="validationErrors.length > 0">
      <div class="w3-container w3-red">
        <ul class="w3-ul">
          <li v-for="error in validationErrors" :key="error">
            {{ error }}
          </li>
        </ul>
      </div>
      <p />
    </div>

    <!-- input fields-->
    <div class="w3-container">
      <label>First Name:</label>
      <input
        class="w3-input w3-border"
        type="text"
        name="firstName"
        v-model="address.firstName"
      />
      <label>Last Name:</label>
      <input
        class="w3-input w3-border"
        type="text"
        name="lastName"
        v-model="address.lastName"
      />
      <label>Email:</label>
      <input
        class="w3-input w3-border"
        type="text"
        name="email"
        v-model="address.email"
      />
      <label>Phone:</label>
      <input
        class="w3-input w3-border"
        type="text"
        name="phone"
        v-model="address.phone"
      />
      <p />

      <!-- Buttons-->
      <div class="w3-center">
        <div class="w3-bar">
          <button
            v-if="create"
            class="w3-button w3-ripple w3-border"
            @click="createAddress()"
            style="margin-right: 5px; width: 8em"
          >
            Save
          </button>
          <button
            v-if="create"
            class="w3-button w3-ripple w3-red"
            @click="clearAddress()"
            style="margin-right: 5px; width: 8em"
          >
            Clear
          </button>
          <button
            v-if="!create"
            class="w3-button w3-ripple w3-border"
            @click="updateAddress()"
            style="margin-right: 5px; width: 8em"
          >
            Update
          </button>
          <button
            v-if="!create"
            class="w3-button w3-ripple w3-red"
            @click="deleteAddress()"
            style="margin-right: 5px; width: 8em"
          >
            Delete
          </button>
        </div>
      </div>
      <p />
    </div>
  </div>
</template>
<script setup>
import { ref, computed } from "vue";
import { useAddressStore } from "@/store/address";
import { validateEmail, validatePhone } from "../services/helper";

const props = defineProps({
  title: String,
  create: Boolean,
});

const addressStore = useAddressStore();

const address = computed(() =>
  props.create ? addressStore.newAddress : addressStore.selectedAddress
);
const validationErrors = ref([]);

const showEditor = computed(() => {
  return (
    props.create | (address.value != null && Object.keys(address).length > 0)
  );
});

const clearAddress = () => {
  validationErrors.value = {};
  addressStore.clearNewValue();
};

const updateAddress = () => {
  validateAddress(address.value);
  if (validationErrors.value.length > 0) return;
  addressStore.updateAddress();
};

const createAddress = () => {
  validateAddress(address.value);
  if (validationErrors.value.length > 0) return;
  addressStore.createAddress();
};

const deleteAddress = () => addressStore.deleteAddress();

const validateAddress = (address) => {
  validationErrors.value = [];
  if (address.firstName == null || address.firstName.trim() == "") {
    validationErrors.value.push("First name must not be empty");
  }

  if (address.lastName == null || address.lastName.trim() == "") {
    validationErrors.value.push("Last name must not be empty");
  }

  if (address.email == null || !validateEmail(address.email)) {
    validationErrors.value.push("Email must be set and valid");
  }

  if (address.phone == null || !validatePhone(address.phone)) {
    validationErrors.value.push(
      "Phone must be set and valid like '+490913132587'"
    );
  }
};
</script>
