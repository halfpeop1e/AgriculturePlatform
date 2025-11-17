/*
  useLogin.ts
  说明：用户登录相关的封装
  导出函数：
    - loginUser(loginData): 请求登录接口并更新本地 UserStore（userId, tokens, islogin）
  返回：成功时返回后端的 LoginResponse 数据（包含 userId、tokens 等），失败时在控制台输出错误并返回 undefined
  注意：该函数会直接修改全局 UserStore（因此调用方无需再次写入 store），若需替换为纯函数请修改实现。
*/
import {useAxios} from '~/composables/useAxios'
import Cookies from 'js-cookie'
import type { LoginRequest,LoginResponse } from '~/types/login'
const useAxiosInstance=useAxios()
export async function loginUser(loginData:LoginRequest){
    try{
        const UserStore=useUserStore()
        const response=await useAxiosInstance.post<LoginResponse>('/user/login',loginData)
        if(response.status===200){
            //console.log('用户登录成功',response)
            UserStore.userId=response.data.data.userId
            //console.log("userId:",UserStore.userId)
            UserStore.tokens=response.data.data.tokens
            UserStore.LoginSet()
            navigateTo('/')
            return response.data
        }
        else{
            throw new Error('用户登录失败')
        }
    }
    catch(err){
        console.error('用户登录失败',err)
    }
}

export function logoutUser(){
    const userStore=useUserStore()
    userStore.LogoutSet
    Cookies.remove('AuthToken')
    console.log('用户已登出，清除本地状态')
}
