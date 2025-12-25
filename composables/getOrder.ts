/*
  getOrder.ts
  说明：获取当前用户订单列表的 helper 函数（基于 Axios 封装）
  返回：成功时返回 Order[]，失败时在控制台输出错误并返回 undefined
  副作用：无持久化动作，仅发起网络请求；依赖 useAxios 返回的 Axios 实例
  注意：调用方应处理返回为 undefined 的情形。
*/
import type { orderDate } from "~/types/product";
import { useAxios } from "./useAxios"
import type{ Order } from "~/types/myOrder";
export async function getOrderList(){
    try{
        const useAxiosInstance=useAxios()
        const response=await useAxiosInstance.get(`/order/list`)
        if(response.status===200){
            console.log('获取订单列表成功',response.data)
            return response.data.data.orders as Order[]
        }
        else{
            throw new Error('获取订单列表失败')
        }
    }
    catch(err){
        console.error('获取订单列表失败',err)
    }   
}

export async function getOrderDate(){
    try{
        const useAxiosInstance=useAxios()
        const response=await useAxiosInstance.get(`/product/date`)
        if(response.status===200){
            console.log('获取订单列表成功',response.data)
            return response.data.data as orderDate
        }
        else{
            throw new Error('获取订单列表失败')
        }
    }
    catch(err){
        console.error('获取订单列表失败',err)
    }   
}
