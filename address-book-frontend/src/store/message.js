import { defineStore } from "pinia";

export const useMessageStore = defineStore("message", {
  state: () => ({ message: "", hasMessage: false }),
  actions: {
    addMessage(text) {
      this.message = text;
      this.hasMessage = true;
      // close message after some secs
      setTimeout(() => {
        this.removeMessage();
      }, 3000);
    },
    removeMessage() {
      this.message = "";
      this.hasMessage = false;
    },
  },
});
