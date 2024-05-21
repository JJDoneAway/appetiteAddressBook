import { defineStore } from "pinia";
import axios from "axios";
import { useMessageStore } from "./message.js";

function getFilter(value) {
  if (value == "") {
    return "";
  } else {
    return `&filters=[["firstName","like","${value}"],["OR"],["lastName","like","${value}"],["OR"],["email","like","${value}"],["OR"],["phone","like","${value}"]]`;
  }
}

export const useAddressStore = defineStore("address", {
  state: () => ({
    addresses: [{}],
    selectedAddress: {},
    newAddress: {},
    currentPage: 0,
    lastPage: 0,
    sort: "lastName",
    asc: true,
    filter: "",
    loading: false,
  }),
  actions: {
    setSelected(address) {
      this.selectedAddress = Object.assign({}, address);
    },
    clearNewValue() {
      this.newAddress = {};
    },
    doSearch(value) {
      this.currentPage = 0;
      this.filter = value;
      this.addresses = [{}];
      this.loadAddresses();
    },
    doSort(attribute, asc) {
      this.asc = asc;
      this.sort = attribute;
      this.addresses = [{}];
      this.loadAddresses();
    },
    getSortString() {
      if (this.asc) {
        return this.sort;
      } else {
        return `-${this.sort}`;
      }
    },
    doPaging(dif) {
      if (this.currentPage == 0 && dif < 0) {
        return;
      }
      if (this.currentPage + 1 == this.lastPage && dif > 0) {
        return;
      }
      this.addresses = [{}];
      this.currentPage = this.currentPage + dif;
      this.loadAddresses();
    },
    doClear() {
      this.filter = "";
      this.loadAddresses();
    },
    async loadAddresses() {
      this.loading = true;
      axios
        .get(
          `${import.meta.env.VITE_URL}/addresses?page=${
            this.currentPage
          }&size=16&sort=${this.getSortString()}${getFilter(this.filter)}`
        )
        .then((response) => {
          this.addresses = [...response.data.items];
          this.currentPage = response.data.page;
          this.lastPage = response.data.totalPages;
          this.loading = false;
        })
        .catch((err) => {
          console.log(err);
          useMessageStore().addMessage(
            "Authentication or authorization failed. Please check if you have the required SIAM roles. See Environment Tab"
          );
          this.loading = false;
        });
    },
    async updateAddress() {
      axios
        .put(
          `${import.meta.env.VITE_URL}/addresses/${this.selectedAddress.id}`,
          this.selectedAddress
        )
        .then(() => {
          this.loadAddresses();
        })
        .catch((err) => {
          console.log(err);
          useMessageStore().addMessage(
            "Not able to update. Missing authorization"
          );
        });
    },
    createAddress() {
      axios
        .post(`${import.meta.env.VITE_URL}/addresses`, this.newAddress)
        .then((response) => {
          this.loadAddresses();
          this.newAddress = {};
          this.selectedAddress = { ...response.data };
        });
    },
    deleteAddress() {
      axios
        .delete(
          `${import.meta.env.VITE_URL}/addresses/${this.selectedAddress.id}`
        )
        .then(() => this.loadAddresses())
        .then(() => (this.selectedAddress = {}));
    },
    resetAddresses() {
      axios.delete(`${import.meta.env.VITE_URL}/addresses`).then(() => {
        axios
          .post(`${import.meta.env.VITE_URL}/addresses/default`, null)
          .then(() => {
            this.currentPage = 0;
            this.filter = "";
            this.sort = "lastName";
            this.asc = true;
            this.loadAddresses();
            this.selectedAddress = {};
            this.newAddress = {};
          });
      });
    },
  },
});
