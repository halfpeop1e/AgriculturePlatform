import {useAxios} from '~/composables/useAxios'
const useAxiosInstance=useAxios()
import { useUserStore } from '~/utils/userStore'

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
        const response=await useAxiosInstance.post(`/user/profile/${userStore.userinfo.userId}/update`,profileData)
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
        const response=await useAxiosInstance.post(`/user/security/${userStore.userinfo.userId}/update`,securityData)
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
