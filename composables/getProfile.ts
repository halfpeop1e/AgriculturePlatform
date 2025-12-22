/*
  getProfile.ts
  说明：向后端请求当前用户资料并写入本地 userStore 的 helper
  行为：
    - 读取 userStore.userId 并请求 `/user/profile/:id` 接口
    - 成功时调用 userStore.setUserProfile 更新本地状态
  返回：无明确返回值；出现错误时在控制台打印错误并吞掉异常（可根据需要改为抛出）
  注意：调用方可以在请求完成后直接从 userStore 中读取最新的用户信息
*/
import {useAxios} from '~/composables/useAxios'
import type { profileResponse, ExpertProfile, ExpertProfileResponse } from '~/types/profile'
import { useUserStore } from '~/utils/userStore'
const useAxiosInstance=useAxios()
export async function getUserProfile(){
    const userStore=useUserStore()
    try{
        const response=await useAxiosInstance.get<profileResponse>(`/user/profile/${userStore.userId}`)
        if(response.status===200){
            console.log('获取用户资料成功',response.data.data)
            try{
                userStore.setUserProfile(response.data)
                console.log('设置用户资料成功',userStore.userinfo)
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

export async function getExpertProfile(){
    const userStore = useUserStore()
    if (!userStore.userId) {
        console.warn('无法获取专家资料：缺少用户 ID')
        return null
    }
    try{
        const response = await useAxiosInstance.get<ExpertProfileResponse | ExpertProfile>(`/expert/profile/${userStore.userId}`)
        if(response.status===200 && response.data){
            const payload = response.data as ExpertProfileResponse | ExpertProfile
            const profile = (typeof payload === 'object' && 'data' in payload ? (payload as ExpertProfileResponse).data : payload) as ExpertProfile
            if(profile){
                const completed = typeof profile.completed === 'boolean' ? profile.completed : Boolean(profile.name && profile.introduction)
                userStore.setExpertProfile(profile, completed)
                return profile
            }
        }
        throw new Error('获取专家资料失败')
    }
    catch(err){
        console.error('获取专家资料失败',err)
        userStore.expertProfile = null
        userStore.expertProfileCompleted = false
        return null
    }
}
