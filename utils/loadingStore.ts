import { defineStore } from 'pinia'
export const useLoadingStore = defineStore('sendCodeLoading', {
    state:()=>({
        sendLoading:false
    }),
    actions:{
        setLoading(loading:boolean){
            this.sendLoading=loading
        },
        cancelLoading(loding:boolean){
            this.sendLoading=loding
        }
    }
})