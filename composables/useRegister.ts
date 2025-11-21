/*
  useRegister.ts
  说明：与用户注册与邮箱验证码相关的请求封装
  导出函数：
    - onSendCode(email)：发送邮箱验证码，会更新 loadingStore 并展示消息
    - onVerifyCode(email, code)：校验验证码，成功时返回状态码 200
    - onRegister(data)：提交注册信息到后端
  注意：该模块使用全局的 ElMessage 与 loadingStore 显示消息与加载状态；调用方无需重复展示提示。
*/
import {useAxios} from '~/composables/useAxios'
import { useEmailAxios } from './useEmailAxios'
import type { RegisterRequest } from '~/types/register'
import { useLoadingStore } from '~/utils/loadingStore'

const useAxiosInstance=useAxios()
const useEmailAxiosInstance=useEmailAxios()
export async function onSendCode(email:string,action:string){
    const LodingStore=useLoadingStore()
    LodingStore.setLoading(true)
    try{
      const respone=await useEmailAxiosInstance.post('/email/sendCode',{email:email,action:action})
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
        const respone=await useEmailAxiosInstance.post('/email/verify',{email:email,code:code})
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
            console.log('注册成功，返回数据：',respone.data)
            return respone.data
        }
        else{
            throw new Error('注册失败')
        }
    }
    catch(err){
        ElMessage.error('注册失败')
    }
}
