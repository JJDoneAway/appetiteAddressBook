<template>
  <div v-if="hasMessage" class="fixed">
    <div class="w3-panel w3-dark-gray w3-display-container">
      <span
        onclick="this.parentElement.style.display='none'"
        @click="clear()"
        class="w3-button w3-large w3-display-topright"
        >&times;</span
      >
      <h3>Message</h3>
      <p>{{ message }}</p>
    </div>
  </div>
</template>
<script setup>
import { computed } from "vue";
import { useMessageStore } from "../store/message";

const messageStore = useMessageStore();
const message = computed(() => {
  return messageStore.message;
});
const hasMessage = computed(() => {
  return messageStore.hasMessage;
});
const clear = () => {
  messageStore.removeMessage();
};
</script>
<style scoped>
.fixed {
  position: fixed;
  left: 20%;
  width: 60%;
  z-index: 99;
}
</style>
