/*
  useChange.ts
  说明：包含修改用户资料与安全信息的请求封装
  导出函数：
    - changeUserProfile(profileData)：更新用户资料（昵称、简介、头像、位置、电话、地址、标签）
      返回 true 表示成功，false 表示失败
    - securityChangeInfo(securityData)：更新安全相关信息（如修改密码）
      返回 true 表示成功，false 表示失败
  副作用：函数会读取当前 userStore.userId 作为请求目标，并不会直接更新 store（除非后端返回并且调用方处理）
  注意：该模块以布尔值表示成功/失败；如果需要错误详情，请改为抛出异常或返回包含错误信息的对象。
*/
import {useAxios} from '~/composables/useAxios'
const useAxiosInstance=useAxios()
import { useUserStore } from '~/utils/userStore'
import type { ExpertProfile } from '~/types/profile'

export async function changeUserProfile(profileData:{
    nickname:string,
    bio:string,
    avatar:string,
    location:string,
    phone:string,
    address:string,
    tags:string[]
}){
    const userStore=useUserStore()
    try{
        const response=await useAxiosInstance.post(`/user/profile/${userStore.userId}/update`,profileData)
        if(response.status===200){
            console.log('修改用户资料成功')
            return true
        }
        else{
            throw new Error('修改用户资料失败')
        }
    }
    catch(err){
        console.error('修改用户资料失败',err)
        return false
    }
} 
export async function securityChangeInfo(securityData:{
    email?:string,
    currentpassword:string,
    newpassword:string,
}){
   try{
        const userStore=useUserStore()
        const response=await useAxiosInstance.post(`/user/security/${userStore.userId}/update`,securityData)
        if(response.status===200){
            console.log('修改用户安全信息成功')
            return true
        }
        else{
            throw new Error('修改用户安全失败')
        }
    }
    catch(err){
        console.error('修改用户安全失败',err)
        return false
    }
}

export async function submitExpertProfile(profileData: ExpertProfile){
    const userStore=useUserStore()
    try{
        const response=await useAxiosInstance.post(`/expert/profile/${userStore.userId}/submit`,profileData)
        if(response.status===200){
            console.log('提交专家资料成功')
            return true
        }
        throw new Error('提交专家资料失败')
    }
    catch(err){
        console.error('提交专家资料失败',err)
        return false
    }
}
