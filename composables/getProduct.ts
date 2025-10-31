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