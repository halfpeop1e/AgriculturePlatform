import { useAxios } from "./useAxios"
import type{ Order } from "~/types/myOrder";
export async function getOrderList(){
    try{
        const useAxiosInstance=useAxios()
        const response=await useAxiosInstance.get<Order[]>(`/order/list`)
        if(response.status===200){
            console.log('获取订单列表成功',response.data)
            return response.data
        }
        else{
            throw new Error('获取订单列表失败')
        }
    }
    catch(err){
        console.error('获取订单列表失败',err)
    }   
}