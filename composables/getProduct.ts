/*
  getProduct.ts
  说明：获取产品列表的网络请求封装
  返回：成功时返回 productResponse[]，失败时在控制台输出错误并返回 undefined
  依赖：内部使用 useAxios() 创建的 Axios 实例（会自动注入 token 等拦截逻辑）
  注意：如果需要分页或查询参数，可将函数签名扩展为接受选项参数。
*/
import { useAxios } from "./useAxios"
import type{ productResponse } from "~/types/product";
const useAxiosInstance=useAxios()
export async function getProductList(){
    try{
        const response=await useAxiosInstance.get<productResponse[]>('/product/list')
        if(response.status===200){
            console.log('获取产品列表成功',response.data)
            return response.data
        }
        else{
            throw new Error('获取产品列表失败')
        }
    }
    catch(err){
        console.error('获取产品列表失败',err)
    }
}
