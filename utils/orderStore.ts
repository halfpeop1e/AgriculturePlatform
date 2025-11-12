/*
  orderStore.ts
  说明：全局订单列表状态管理
  State:
    - orderList: Order[] — 当前用户的订单列表
  Actions:
    - setOrder(order): 设置订单列表
    - 待实现：更新单个订单、删除订单等方法
  使用场景：在订单页（my_buy / my_sales）由中间件 loadOrderList 自动调用 setOrder 初始化；组件直接从 store 读取 orderList 渲染
*/
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
