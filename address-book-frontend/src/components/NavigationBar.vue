<template>
  <div class="w3-top">
    <div class="w3-bar w3-grey">
      <router-link
        :to="{ name: 'Editor' }"
        class="w3-bar-item w3-border-right w3-button w3-padding-large"
        style="width: 14em"
        exact
      >
        <img src="@/assets/addressbookIcon.png" alt="" class="icons" />
        Address Book</router-link
      >
      <div class="w3-dropdown-hover">
        <button
          class="w3-button w3-border-right w3-button w3-padding-large"
          style="width: 13em"
        >
          <img src="@/assets/toolsIcon.png" alt="" class="icons" /> Tools
        </button>
        <div class="w3-dropdown-content w3-bar-block w3-card-4">
          <span @click="reset()" class="w3-bar-item w3-button"
            >Reset Content</span
          >
        </div>
      </div>
      <div class="w3-bar-item w3-right" style="padding: 1em 1em">
        Address Book of {{ userName }}
      </div>
    </div>
  </div>
</template>
<script setup>
import { useRouter } from "vue-router";
import { useAddressStore } from "@/store/address";
import { useMessageStore } from "@/store/message";
import { onMounted, computed } from "vue";

const router = useRouter();
const addressStor = useAddressStore();
const messageStor = useMessageStore();
const userName = computed(() => addressStor.userName);

const putMessage = (text) => messageStor.addMessage(text);

const reset = () => {
  addressStor.resetAddresses();
  putMessage("Database was reset to its default values.");
  router.push({ name: "Editor" });
};

onMounted(() => {
  addressStor.getUserName();
});
</script>

<style scoped>
.router-link-active {
  font-weight: bold;
}

.icons {
  height: 1.8em;
  width: auto;
  margin: -1em 0.1em -0.8em 0em;
}
</style>
