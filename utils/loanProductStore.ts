import { defineStore } from "pinia";
import type { AgriculturalLoanProduct } from "~/types/loanProduct";
export const useLoanStore = defineStore("loanProductStore", {
    state: () => ({
        orderList: [] as AgriculturalLoanProduct[],
        total: 0,
        page: 1,
        pageSize: 10,
        hasMore: false,
    }),
    actions: {
        setOrder(order: AgriculturalLoanProduct[]) {
            this.orderList = order
        },
         setPaginationInfo(total: number, page: number, pageSize: number, hasMore: boolean) {
            this.total = total;
            this.page = page;
            this.pageSize = pageSize;
            this.hasMore = hasMore;
        },
        //后续这里要写更新列表的方法
    },
})