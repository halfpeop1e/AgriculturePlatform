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