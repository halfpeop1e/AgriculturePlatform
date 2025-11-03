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