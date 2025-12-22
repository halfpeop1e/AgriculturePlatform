<!--
  comfirmbuy.vue
  说明：购买确认弹窗组件（使用 PrimeVue Dialog）
  Props:
    - product: 产品对象（id, name, price, stock, saler 等）
    - showDialog: 是否显示弹窗（由父组件控制）
  Emits:
    - closeDialog: 关闭弹窗事件（父组件需监听以隐藏弹窗）
  行为：校验购买数量、调用 BuyProduct 下单并显示消息提示；出现错误时显示失败提示。
  注意：BuyProduct 为外部可用的 API/函数（请确保在父作用域或全局导入）。
-->
<template>
  <Dialog
    v-model:visible="visible"
    header="确认购买"
    modal
    class="w-full max-w-md"
    :draggable="false"
  >
    <div class="space-y-5">
      <div class="space-y-2 rounded-lg bg-slate-50 p-4">
        <div class="text-xl font-semibold text-slate-900">{{ product.name }}</div>
        <div class="flex items-center justify-between text-sm text-slate-600">
          <span>单价</span>
          <span class="font-semibold text-emerald-600">￥{{ product.price }}</span>
        </div>
        <div class="flex items-center justify-between text-sm text-slate-600">
          <span>商户</span>
          <span class="font-medium text-slate-800">{{ product.saler }}</span>
        </div>
      </div>

      <div class="space-y-2">
        <label class="text-sm font-medium text-slate-700">购买数量</label>
        <InputNumber
          v-model="orderQuantity"
          :min="1"
          :max="maxQuantity"
          :showButtons="true"
          buttonLayout="horizontal"
          decrementButtonClass="p-button-outlined"
          incrementButtonClass="p-button-outlined"
          inputClass="w-full"
          class="w-full"
        />
        <p class="text-xs text-slate-500">库存：{{ product.stock }} 件</p>
      </div>

      <div class="flex items-center justify-between border-t border-slate-200 pt-4 text-base">
        <span class="text-slate-600">合计</span>
        <span class="font-semibold text-slate-900">￥{{ totalPrice }}</span>
      </div>
    </div>

    <template #footer>
      <div class="flex w-full justify-end gap-3">
        <Button label="取消" class="p-button-outlined" @click="handleCancel" />
        <Button label="确认订单" icon="pi pi-shopping-cart" :loading="submitting" @click="confirmOrder" />
      </div>
    </template>
  </Dialog>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import Dialog from 'primevue/dialog'
import Button from 'primevue/button'
import InputNumber from 'primevue/inputnumber'
import { useToast } from 'primevue/usetoast'
import type { comfirmOrderRequest } from '~/types/comfirmOrder'
import { v4 as uuidv4 } from 'uuid'
import { useUserStore } from '~/utils/userStore'

const toast = useToast()

const emit = defineEmits<{ (e: 'closeDialog'): void }>()

const props = defineProps<{
  product: {
    id: string
    name: string
    image: { url: string }[]
    description: string
    price: number
    stock: number
    saler: string
  }
  showDialog: boolean
}>()

const visible = computed({
  get: () => props.showDialog,
  set: (value: boolean) => {
    if (!value) emit('closeDialog')
  }
})

const orderQuantity = ref<number>(1)
const submitting = ref(false)
const maxQuantity = computed(() => Math.max(props.product.stock, 1))

const totalPrice = computed(() => {
  const qty = Number(orderQuantity.value || 0)
  return (props.product.price * qty).toFixed(2)
})

watch(
  () => props.showDialog,
  (open) => {
    if (open) {
      orderQuantity.value = Math.min(Math.max(1, orderQuantity.value), maxQuantity.value)
    }
  }
)

const order = ref<comfirmOrderRequest>({
  orderId: '',
  productId: '',
  quantity: 0,
  totalprice: 0,
  buyer: '',
  saler: ''
})

const handleCancel = () => {
  visible.value = false
}


async function confirmOrder() {
  const userStore = useUserStore()
  const stock = props.product.stock
  const qty = Number(orderQuantity.value || 0)

  if (!qty || qty <= 0) {
    toast.add({ severity: 'warn', summary: '数量错误', detail: '购买数量必须至少为 1', life: 3000 })
    return
  }

  if (qty > stock) {
    toast.add({ severity: 'warn', summary: '库存不足', detail: '购买数量不能超过库存', life: 3000 })
    return
  }

  submitting.value = true
  try {
    order.value.orderId = uuidv4()
    order.value.productId = props.product.id
    order.value.quantity = qty
    order.value.totalprice = props.product.price * qty
    order.value.buyer = userStore.userinfo.nickname
    order.value.saler = props.product.saler

    await BuyProduct(order.value)
    toast.add({ severity: 'success', summary: '下单成功', detail: `${props.product.name} × ${qty}`, life: 3000 })
    visible.value = false
  } catch (error) {
    console.error(error)
    toast.add({ severity: 'error', summary: '下单失败', detail: '请稍后重试', life: 3000 })
  } finally {
    submitting.value = false
  }
}
</script>
