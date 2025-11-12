/*
  loadOrderList.global.ts
  说明：路由中间件，在用户访问订单页时自动加载订单列表并写入 orderStore
  触发条件：当路由 path 为 '/product/myorder/my_buy' 或 '/product/myorder/my_sales' 时触发
  行为：调用 getOrderList() 获取订单数据，若成功则调用 orderStore.setOrder 写入本地 store
  注意：该中间件是客户端/服务端可运行的 Nuxt 中间件（使用 defineNuxtRouteMiddleware）；若 getOrderList 需要 token，请确保全局拦截器已注入 Authorization header
*/
import type { Order } from "~/types/myOrder";
export default defineNuxtRouteMiddleware(async(to, from) => {
    if(to.path=='/product/myorder/my_buy' || to.path==='/product/myorder/my_sales'){ 
        try{
            const orderlist = await getOrderList()
            if(orderlist?.length===0){
                console.log('订单列表为空')
                return
            }
            const orderStore = useOrderStore()
            orderStore.setOrder(orderlist as Order[])
        }
        catch(err){
            console.error('加载订单列表失败',err)
        }
    }
})
