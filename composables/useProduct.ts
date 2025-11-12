/*
  useProduct.ts
  说明：与商品发布/购买相关的请求封装
  导出函数：
    - PostProduct(formData)：发布商品（multipart/form-data），会向后端提交表单并弹出提示
    - BuyProduct(productId, quantity, totalprice)：下单购买，内部会生成订单号并调用后端
  注意：函数内部使用了全局 UserStore、ElMessage 等全局对象；若在测试环境使用请提供替代实现或 mock。
*/
import { useAxios } from "./useAxios";
import type{ postProductRequest } from "~/types/product";
import{v4 as uuidv4} from 'uuid'

const useAxiosInstance=useAxios()
export async function PostProduct(formData:any){
    const UserStore=useUserStore()
    try{
        const response=await useAxiosInstance.post<postProductRequest>('/product/sell',{
            formData,
            salerId:UserStore.userId
        },
        {
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
export async function BuyProduct(productId:number,quantity:number,totalprice:number){
    try{
        const response=await useAxiosInstance.post('/product/buy',{
            orderId:uuidv4(),
            productId:productId,
            quantity:quantity,
            totalprice:totalprice
            // buyer:useUserStore().userinfo.userId,
            // saler:从商品的发布者信息中获取
        })
        if(response.status===200){
            console.log('购买产品成功',response.data)
            ElMessage.success('购买产品成功')
            return response
        }
        else{
            throw new Error('购买产品失败')
        }
    }
    catch(err){
        console.error('购买产品失败',err)
    }
}   
