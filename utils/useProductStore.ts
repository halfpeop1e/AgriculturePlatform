/*
  comfirmBuyStore.ts
  说明：管理购买确认弹窗的状态和被选中产品的信息
  State:
    - showComfirmBuyDialog: boolean — 是否显示购买确认弹窗
    - product: productResponse — 当前选中待购买的产品对象
  Actions:
    - openComfirmBuyDialog(): 打开弹窗
    - closeComfirmBuyDialog(): 关闭弹窗
    - setProduct(product): 设置当前产品
    - resetProduct(): 重置产品为空对象
  使用场景：在产品列表或详情页点击购买时，调用 setProduct + openComfirmBuyDialog；购买成功或取消时调用 closeComfirmBuyDialog
*/
import { defineStore } from 'pinia'
import type {productResponse} from '~/types/product'
export const useComfirmBuyStore = defineStore('comfirmBuyStore', {
    state:()=>({
        showComfirmBuyDialog:false,
        product:{} as productResponse
    }),
    actions:{
        openComfirmBuyDialog(){
            this.showComfirmBuyDialog=true
            console.log('打开购买确认弹窗')
        },
        closeComfirmBuyDialog(){
            this.showComfirmBuyDialog=false
        },
        setProduct(product:productResponse){
            this.product=product
        },
        resetProduct(){
            this.product={} as productResponse
        }
    }
})  
export const useEditProductStore = defineStore('editProductStore', {
    state:()=>({
        showEditProductDialog:false,
        product:{} as productResponse
    }),
    actions:{
        openEditProductDialog(){
            this.showEditProductDialog=true
        },
        closeEditProductDialog(){
            this.showEditProductDialog=false
        },
        setProduct(product:productResponse){
            this.product=product
        },
        resetProduct(){
            this.product={} as productResponse
        }
    }
})
