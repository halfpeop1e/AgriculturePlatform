import { defineStore } from "pinia";
import type { Order } from "~/types/myOrder";
export const useOrderStore = defineStore("orderStore", {
    state: () => ({
        orderList: [] as Order[],
    }),
    actions: {
        setOrder(order: Order[]) {
            this.orderList = order
        },
        //后续这里要写更新列表的方法
    },
})