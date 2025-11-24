<!--
  comfirmbuy.vue
  说明：购买确认弹窗组件（使用 Element Plus 的 el-dialog）
  Props:
    - product: 产品对象（id, name, price, stock, saler 等）
    - showDialog: 是否显示弹窗（由父组件控制）
  Emits:
    - closeDialog: 关闭弹窗事件（父组件需监听以隐藏弹窗）
  行为：校验购买数量、调用 BuyProduct 下单并显示消息提示；出现错误时显示失败提示。
  注意：BuyProduct 为外部可用的 API/函数（请确保在父作用域或全局导入）。
-->
<template>
    <!-- 购买弹窗 -->
  <el-dialog :model-value="showDialog" width="420px" title="确认购买">
    <div class="space-y-4">
      <div class="text-lg font-medium">{{ product.name }}</div>
      <div class="flex items-center justify-between text-sm text-gray-600">
        <div>单价：</div>
        <div class="font-semibold text-green-600">￥{{ product.price }}</div>
      </div>
      <div class="flex items-center justify-between text-sm text-gray-600">
        <div>商户：</div>
        <div class="font-medium">{{ product.saler }}</div>
      </div>

      <div>
        <label class="block text-sm text-gray-700 mb-1">购买数量</label>
        <el-input-number v-model="orderQuantity" :min="0" :max="product.stock" class="w-full" />
        <div class="text-xs text-gray-400 mt-1">库存：{{ product.stock }} 件</div>
      </div>

      <div class="flex items-center justify-between pt-2 border-t border-gray-100">
        <div class="text-sm text-gray-600">合计：</div>
        <div class="text-lg font-bold">￥{{ (product.price * orderQuantity).toFixed(2) }}</div>
      </div>
    </div>

    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
      <el-button type="primary" @click="confirmOrder">确认订单</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import type { comfirmOrderRequest } from '~/types/comfirmOrder'
import {v4 as uuidv4} from 'uuid'
const emit = defineEmits<{
    (e: 'closeDialog'): void
  }>()
const orderQuantity = ref<number>(1)
const props=defineProps<{
  product: {
    id: string
    name: string
    image:{ url: string}[]
    description: string
    price: number
    stock: number
    saler: string
  },
  showDialog: boolean
}>()
function closeDialog() {
  emit('closeDialog')
}
const order = ref<comfirmOrderRequest>({
    orderId: '',
    productId: '',
    quantity: 0,
    totalprice: 0,
    buyer: '',
    saler: ''
})
function confirmOrder() {
  const userStore=useUserStore()
  const stock = props.product.stock
  const qty = Number(orderQuantity.value || 0)
  const totalprice = props.product.price * qty
  if (qty <= 0) {
    ElMessage.error('购买数量必须至少为 1')
    return
  }
  if (qty > stock) {
    ElMessage.error('购买数量不能超过库存')
    return
  }

  try {
    //const pid = Number(props.product.id)
    order.value.orderId=uuidv4()
    order.value.productId = props.product.id
    order.value.quantity = qty
    order.value.totalprice = totalprice
    order.value.buyer=userStore.userinfo.nickname
    order.value.saler=props.product.saler
    BuyProduct(order.value).then(() => {
      ElMessage.success(`已下单：${props.product.name} × ${qty}`)
      emit('closeDialog')
    }).catch((err) => {
      ElMessage.error('下单失败')
      console.error(err)
    })
  } catch (err) {
    ElMessage.error('下单失败')
  }
}
</script>

<style scoped>

</style>
