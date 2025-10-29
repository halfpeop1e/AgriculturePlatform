import { useAxios } from "./useAxios";
import type{ postProductRequest } from "~/types/product";
const useAxiosInstance=useAxios()
export async function PostProduct(formData:postProductRequest){
    try{
        const response=await useAxiosInstance.post('/product/sell',formData,{
            headers:{
                'Content-Type':'multipart/form-data'
            }
        })
        if(response.status===200){
            console.log('发布产品成功',response.data)
            ElMessage.success('发布产品成功')
            return response
        }
        else{
            throw new Error('发布产品失败')
        }
    }
    catch(err){
        console.error('发布产品失败',err)
    }
}   