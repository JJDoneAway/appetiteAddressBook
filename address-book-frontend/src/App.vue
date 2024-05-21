<template>
  <div v-if="!backendIsConnected">
    <div class="w3-panel w3-red w3-display-container">
      <span
        onclick="parentElement.style.display='none'"
        class="w3-button w3-large w3-display-topright"
        >&times;</span
      >
      <h3>Problem!</h3>
      <div>
        It is not possible to contact the Addressbook Backend.
        <ul class="w3-ul">
          <li>Backend URL: {{ backend.url }}</li>
        </ul>
      </div>
    </div>
  </div>
  <div v-if="backendIsConnected">
    <navigation-bar />
    <message-panel />
    <router-view />
  </div>
</template>

<script setup>
import NavigationBar from "./components/NavigationBar.vue";
import MessagePanel from "./components/MessagePanel.vue";
import { checkAPI, getBackend } from "./services/checkBackendConnection";
import { ref, onMounted, computed } from "vue";

const backendIsConnected = ref(true);
const backend = computed(() => getBackend());

onMounted(() => {
  checkAPI(backendIsConnected);
});
</script>
