<template>
    <div class="p-4">
         <div class="flex mb-2">
            <el-button size="default" @click="navigateTo('/product/myorder/my_buy')">买入订单</el-button>
            <el-button type="primary" size="default" class="ml-0" @click="navigateTo('/product/myorder/my_sales')">卖出订单</el-button>
        </div>
        <div class="flex items-center justify-between gap-4 mb-4">
            <div class="text-lg font-semibold">我的卖出订单</div>
            <div class="text-sm text-gray-500">共 {{ sellOrders.length }} 条</div>
        </div>

        <el-alert title="卖出订单" type="warning" show-icon />
        <ul class="mt-3 space-y-3">
            <li v-for="o in sellOrders" :key="o.orderId">
              <OrderCard :order="o"/>
            </li>
            <li v-if="sellOrders.length === 0" class="text-gray-500">暂无卖出订单</li>
        </ul>
    </div>
</template>

<script setup lang="ts">
definePageMeta({ layout: 'home-page-layout' })
import { ref, computed } from 'vue'
import type { Order } from '@/types/myOrder'

const orders = ref<Order[]>([
  { orderId: 'b1', name: '绿色有机苹果', quantity: 2, totalprice: 1000, status: '未支付', type: 'buy', buyer:'张三',saler:'水果商户' },
  { orderId: 'b2', name: '香甜玉米', quantity: 5, totalprice: 1000, status: '待收货', type: 'buy' ,buyer:'李四',saler:'玉米商户'},
  { orderId: 's1', name: '土鸡蛋', quantity: 10, totalprice: 1000, status: '待发货', type: 'sell', buyer:'王五',saler:'蛋商户' }
])

const sellOrders = computed(() => orders.value.filter((o) => o.type === 'sell'))
</script>

<style scoped>

</style>
