import {useAxios} from '~/composables/useAxios'
import type { profileResponse } from '~/types/profile'
import { useUserStore } from '~/utils/userStore'
const useAxiosInstance=useAxios()
export async function getUserProfile(email:string){
    const userStore=useUserStore()
    try{
        const response=await useAxiosInstance.get<profileResponse>(`/user/profile/${email}`)
        if(response.status===200){
            console.log('获取用户资料成功')
            try{
                userStore.setUserProfile(response.data)
            }
            catch(err){
                console.error('设置用户资料失败',err)
            }
        }
        else{
            throw new Error('获取用户资料失败')
        }
    }
    catch(err){
        console.error('获取用户资料失败',err)
    }
}