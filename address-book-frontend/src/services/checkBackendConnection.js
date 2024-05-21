import axios from "axios";

export function checkAPI(connected) {
  axios
    .get(`${import.meta.env.VITE_URL}/status/health`)
    .then(() => {
      connected.value = true;
    })
    .catch((error) => {
      console.log(`Backend is broken. Response was: ${JSON.stringify(error)}`);
      connected.value = false;
    });
}

export function getBackend() {
  return {
    url: import.meta.env.VITE_URL,
  };
}
