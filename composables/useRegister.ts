import {useAxios} from '~/composables/useAxios'
import type { RegisterRequest } from '~/types/register'
import { useLoadingStore } from '~/utils/loadingStore'

const useAxiosInstance=useAxios()
export async function onSendCode(email:string){
    const LodingStore=useLoadingStore()
    LodingStore.setLoading(true)
    try{
      const respone=await useAxiosInstance.post('/email/sendCode',{email:email})
        console.log('服务器响应：',respone.data)
      if(respone.status===200){
        ElMessage.success('发送成功')
        setTimeout(()=>{LodingStore.cancelLoading(false)},60000)
      }
      else{
        throw new Error('发送失败')
      }
    }
    catch(err){
        LodingStore.cancelLoading(false)
        ElMessage.error('发送失败')
}
}
export async function onVerifyCode(email:string,code:string){
    try{
        const respone=await useAxiosInstance.post('/email/verify',{email:email,code:code})
        console.log('服务器响应：',respone.status)
        if(respone.status===200){
          ElMessage.success('验证成功')
          console.log('验证成功，状态码：',respone.status)
          return respone.status
        }
        else{
          throw new Error('验证失败')
        }
      }
    catch(err){
        ElMessage.error('验证失败')
    }
}
export async function onRegister(data:RegisterRequest){
    try{
        const respone=await useAxiosInstance.post('/user/register',data)
        if(respone.status===200){
            ElMessage.success('注册成功')
        }
        else{
            throw new Error('注册失败')
        }
    }
    catch(err){
        ElMessage.error('注册失败')
    }
}