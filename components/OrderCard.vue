<template>
  <Card class="order-card shadow-sm">
    <template #title>
      <div class="flex justify-between items-center w-full">
        <div>
          <div class="text-2xl font-semibold">{{ order.name }}</div>
          <small class="text-gray-500">订单号: {{ order.orderId }}</small>
        </div>
        <div class="flex items-center gap-2">
          <Tag :value="order.status" :severity="statusSeverity(order.status)" />
          <Tag :value="order.type === 'buy' ? '买入' : '卖出'" severity="info" />
        </div>
      </div>
    </template>
<template #content>
    <div class="mt-3 space-y-2 text-gray-800 ">
      <div class="flex gap-2 items-center"><strong>数量:</strong> <div class="text-blue-500 text-2xl">{{ order.quantity }}</div></div>
      <div class="flex gap-2 items-center"><strong>总价: </strong><div class="text-green-700 text-2xl">{{ formatCurrency(order.totalprice) }}</div></div>
    </div>
</template>
    <template #footer>
      <div class="flex gap-2 justify-end w-full">
        <Button label="查看" icon="pi pi-external-link" class="p-button-text" @click="onView" />
        <Button
          v-if="canPay&& typeSeverity(order.type)"
          label="去支付"
          icon="pi pi-credit-card"
          class="p-button-sm p-button-success"
          @click="onPay"
        />
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import type { Order } from '~/types/myOrder'
import { defineEmits, defineProps, computed } from 'vue'

const props = defineProps<{ order: Order }>()
const emits = defineEmits<{
  (e: 'view', orderId: string): void
  (e: 'pay', orderId: string): void
}>()

const order = props.order

const onView = () => {
  emits('view', order.orderId)
}

const onPay = () => {
  emits('pay', order.orderId)
}

const canPay = computed(() => order.status === '待支付' || order.status === '未支付')

function formatCurrency(v: number) {
  try {
    return new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(v)
  } catch (e) {
    return `¥ ${v}`
  }
}

function statusSeverity(status: string) {
  const s = (status || '').toLowerCase()
  if (s.includes('已完成') || s.includes('完成') || s.includes('success') || s.includes('已付')) return 'success'
  if (s.includes('待支付') || s.includes('未支付') || s.includes('pending')) return 'warning'
  if (s.includes('已取消') || s.includes('取消') || s.includes('cancel')) return 'danger'
  return 'info'
}
function typeSeverity(type: string) {
  const t = (type || '').toLowerCase()
  if (t.includes('buy')) return true
  if (t.includes('sell')) return false
  return true
}
</script>

<style scoped>
.order-card {
  width: 100%;
  border-radius: 8px;
}

.p-button-text {
  background:rgb(220, 242, 245)
}
</style>
