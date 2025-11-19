import { useAxios } from "./useAxios"
import type { AgriculturalLoanProduct } from '~/types/loanProduct'

const useAxiosInstance = useAxios()

export async function addLoanProduct(AgriculturalLoanProduct:AgriculturalLoanProduct){
    try{
        const response=await useAxiosInstance.post(`/loan/add`,AgriculturalLoanProduct)
        if(response.status===200){ 
            console.log('添加贷款产品成功')
            return true
        }
        else{
            throw new Error('添加贷款产品失败')
        }
    }
    catch(err){
        console.error('添加贷款产品失败',err)
        return false
    }
} 