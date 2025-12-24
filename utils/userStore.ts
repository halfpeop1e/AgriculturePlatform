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
import type { profileResponse, ExpertProfile } from '~/types/profile'
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
    role:'',
    expertProfileCompleted:false,
    expertProfile: null as ExpertProfile | null
    }),
    actions:{
        setUserProfile(profile:profileResponse){
            this.userinfo.nickname=profile.data.Nickname
            console.log('设置用户昵称为',profile.data.Nickname)
            this.userinfo.email=profile.data.Email
            this.avatar=profile.data.Avatar
            this.userinfo.bio=profile.data.Bio
            this.userinfo.tags=profile.data.Tags
            this.userinfo.location=profile.data.Location
            this.userinfo.joinedAt=profile.data.JoinedAt
            this.userinfo.phone=profile.data.Phone
            this.userinfo.address=profile.data.Address
            this.lastActive=profile.data.LastActive
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
        this.role=''
        this.expertProfileCompleted=false
        this.expertProfile=null
      },
      setExpertProfile(profile: ExpertProfile, completed = true){
        this.expertProfile = profile
        this.expertProfileCompleted = completed
        }
      }
})
