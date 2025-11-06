import { defineStore } from 'pinia'
import type { profileResponse } from '~/types/profile'
export const useUserStore = defineStore('userStore', {
    state:()=>({
    userinfo:{
    userId: '',
    nickname:'nickname',
    email:'' ,
    avatar: '',
    bio:  '',
    tags: [''],
    location:'' ,
    joinedAt:'' ,
    phone: '',
    address:'' ,
    lastActive: '',
    }
     
    }),
    actions:{
        setUserProfile(profile:profileResponse){
            this.userinfo.userId=profile.userId
            this.userinfo.nickname=profile.nickname
            this.userinfo.email=profile.email
            this.userinfo.avatar=profile.avatar
            this.userinfo.bio=profile.bio
            this.userinfo.tags=profile.tags
            this.userinfo.location=profile.location
            this.userinfo.joinedAt=profile.joinedAt
            this.userinfo.phone=profile.phone
            this.userinfo.address=profile.address
            this.userinfo.lastActive=profile.lastActive
        }
    }
})