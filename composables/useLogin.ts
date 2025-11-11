import { User } from '@element-plus/icons-vue'
import {useAxios} from '~/composables/useAxios'
import type { LoginRequest,LoginResponse } from '~/types/login'
const useAxiosInstance=useAxios()
export async function loginUser(loginData:LoginRequest){
    try{
        const UserStore=useUserStore()
        const response=await useAxiosInstance.post<LoginResponse>('/user/login',loginData)
        if(response.status===200){
            console.log('用户登录成功',response.data)
            UserStore.userId=response.data.userId
            UserStore.tokens=response.data.tokens
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
