import { defineStore } from "pinia";
import type { AgriculturalLoanProduct } from "~/types/loanProduct";
export const useLoanStore = defineStore("loanProductStore", {
    state: () => ({
        orderList: [] as AgriculturalLoanProduct[],
    }),
    actions: {
        setOrder(order: AgriculturalLoanProduct[]) {
            this.orderList = order
        },
        //后续这里要写更新列表的方法
    },
})