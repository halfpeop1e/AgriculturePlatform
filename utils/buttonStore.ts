import { defineStore } from "pinia";
const buttonStore = defineStore("buttonStore", {
  state: () => ({
    isvisible: false,
  }),
  actions: {
    setvisble() {
      this.isvisible = true;
    },
    cancelvisible() {
        this.isvisible = false;
    },
  },
});
export default buttonStore;