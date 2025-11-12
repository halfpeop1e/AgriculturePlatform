/*
  buttonStore.ts
  说明：管理侧边抽屉（Drawer / Sidebar）可见性的全局状态
  State:
    - isvisible: boolean — 抽屉是否可见
  Actions:
    - setvisble(): 显示抽屉
    - cancelvisible(): 隐藏抽屉
  使用场景：在 profilePageLayout 或 sizeBar 组件中调用以控制侧边导航显示/隐藏
*/
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
