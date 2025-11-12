/*
  loadingStore.ts
  说明：管理发送验证码等操作的加载状态
  State:
    - sendLoading: boolean — 是否正在加载/发送中（如：发送邮箱验证码时显示按钮加载状态）
  Actions:
    - setLoading(loading): 设置加载状态
    - cancelLoading(loading): 取消/关闭加载状态（参数名为 loding 有拼写错误）
  使用场景：在 registerform 等组件中点击"发送验证码"时调用 setLoading(true)；倒计时结束后调用 cancelLoading(false)
*/
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
