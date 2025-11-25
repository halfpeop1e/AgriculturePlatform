<template>
  <Dialog v-model:visible="visibleLocal" :modal="true">
    <Card class="order-info-card">
      <template #title>
        <div class="flex items-center justify-between">
          <div class="flex flex-col">
            <div class="text-lg font-semibold">订单详情</div>
            <small class="text-gray-500"><strong>订单号: </strong>{{ order?.orderId || '-' }}</small>
            <small class="text-gray-500"><strong>卖家: </strong>{{ order?.saler || '-' }}</small>
          </div>
          <Tag :value="order?.status || '未知'" :severity="statusSeverity(order?.status || '')" />
        </div>
      </template>
      <template #content>
          <div class="mt-4 grid grid-cols-2 gap-4">
        <div>
          <div class="text-sm text-gray-600">商品名称</div>
          <div class="font-medium">{{ order?.name || '-' }}</div>
        </div>
        <div>
          <div class="text-sm text-gray-600">订单类型</div>
          <div class="font-medium">{{ order?.type === 'buy' ? '买入' : '卖出' }}</div>
        </div>

        <div>
          <div class="text-sm text-gray-600">数量</div>
          <div class="font-medium">{{ order?.quantity ?? '-' }}</div>
        </div>
        <div>
          <div class="text-sm text-gray-600">总价</div>
          <div class="font-medium">{{ order ? formatCurrency(order.totalprice) : '-' }}</div>
        </div>

        <div class="col-span-2">
          <div class="text-sm text-gray-600">备注</div>
          <div class="font-medium">{{ note }}</div>
        </div>
      </div>
      </template>
      

      <template #footer>
        <div class="flex justify-end gap-2">
          <Button label="关闭" class="p-button-text" @click="onClose()" />
          <Button v-if="canPay" label="去支付" icon="pi pi-credit-card" class="p-button-success" @click="onPay()" />
        </div>
      </template>
    </Card>
  </Dialog>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, ref, watch, computed } from 'vue'
import type { Order } from '~/types/myOrder'

const props = defineProps<{
  modelValue: boolean,
  order: Order | null,
  note?: string
}>()
const emits = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
  (e: 'pay', orderId: string): void
}>()

const visibleLocal = ref(!!props.modelValue)
watch(() => props.modelValue, (v) => (visibleLocal.value = !!v))
watch(visibleLocal, (v) => emits('update:modelValue', v))

const order = computed(() => props.order)
const note = computed(() => props.note || '-')

const onClose = () => {
  visibleLocal.value = false
  console.log('Dialog closed')
}

const onPay = () => {
  if (order.value) emits('pay', order.value.orderId)
}

const canPay = computed(() => {
  const s = (order.value?.status || '').toLowerCase()
  return s.includes('未支付') || s.includes('待支付') || s.includes('pending')
})



</script>

<style scoped>
.order-info-card {
  width: 640px;
}
</style>

