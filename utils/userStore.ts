/*
  userStore.ts
  说明：全局用户信息状态管理（登录、资料、token 等）
  State:
    - userinfo: 包含 nickname, email, bio, tags, location, joinedAt, phone, address
    - lastActive: 最后活跃时间
    - userId: 用户 ID
    - avatar: 用户头像 URL
    - islogin: 登录状态
    - tokens: JWT 或 session token
  Actions:
    - setUserProfile(profile): 从后端返回的 profileResponse 更新本地用户资料
    - LoginSet(): 设置 islogin = true
    - LogoutSet(): 清空 userId、avatar、tokens 并设置 islogin = false
  使用场景：登录成功后调用 LoginSet 与 setUserProfile；退出时调用 LogoutSet；其他组件从 store 读取用户信息
*/
import { defineStore } from 'pinia'
import type { profileResponse } from '~/types/profile'
export const useUserStore = defineStore('userStore', {
    state:()=>({
    userinfo:{
    nickname:'nickname',
    email:'' ,
    bio:  '',
    tags: [''],
    location:'' ,
    joinedAt:'' ,
    phone: '',
    address:'' ,
    
    },
    lastActive: '',
    userId: '',
    avatar: '',
    islogin:false,
    tokens:'',
    }),
    actions:{
        setUserProfile(profile:profileResponse){
            this.userinfo.nickname=profile.nickname
            this.userinfo.email=profile.email
            this.avatar=profile.avatar
            this.userinfo.bio=profile.bio
            this.userinfo.tags=profile.tags
            this.userinfo.location=profile.location
            this.userinfo.joinedAt=profile.joinedAt
            this.userinfo.phone=profile.phone
            this.userinfo.address=profile.address
            this.lastActive=profile.lastActive
        },
        LoginSet(){
            this.islogin=true
            console.log('用户登录状态已设置为 true',this.islogin)
        },
        LogoutSet(){
            this.userId=''
            this.avatar=''
            this.islogin=false
            this.tokens=''
            this.userinfo={
                nickname:'nickname',
                email:'' ,
                bio:  '',
                tags: [''],
                location:'' ,
                joinedAt:'' ,
                phone: '',
                address:'' ,  
            }
        }
      }
})
