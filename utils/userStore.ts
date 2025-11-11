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
        },
        LogoutSet(){
            this.userId=''
            this.avatar=''
            this.islogin=false
            this.tokens=''
        }
    }
})
