<template>
    <div class="p-4">
        <div class="flex mb-2">
            <el-button type="primary" size="default" @click="navigateTo('/product/myorder/my_buy')">买入订单</el-button>
            <el-button size="default" class="ml-0" @click="navigateTo('/product/myorder/my_sales')">卖出订单</el-button>
        </div>
      <div class="flex items-center justify-between gap-4 mb-4">
        <div class="text-lg font-semibold">我的买入订单</div>
        <div class="text-sm text-gray-500">共 {{ buyOrders.length }} 条</div>
      </div>

      <el-alert title="买入订单" type="info" show-icon />
      <ul class="mt-3 space-y-3">
        <li v-for="o in buyOrders" :key="o.orderId">
          <el-card>
            <div class="flex justify-between items-center">
              <div class="font-medium">{{ o.name }} × {{ o.quantity }}</div>
              <div class="text-sm text-gray-500">状态：{{ o.status }}</div>
            </div>
            <div class="flex items-center mt-2">
              <div class="text-sm mr-2">订单价格：</div>
              <div class="text-green-600 text-xl font-bold mr-1">{{ o.totalprice }}</div>
              <div class="text-sm">元</div>
            </div>
          </el-card>
        </li>
        <li v-if="buyOrders.length === 0" class="text-gray-500">暂无买入订单</li>
      </ul>
    </div>
</template>

<script setup lang="ts">
definePageMeta({ layout: 'home-page-layout' })
import { useRoute } from 'nuxt/app'
import { ref, computed } from 'vue'
import type { Order } from '@/types/myOrder'

const orders = ref<Order[]>([
  { orderId: 'b1', name: '绿色有机苹果', quantity: 2, totalprice: 1000, status: '已支付', type: 'buy' },
  { orderId: 'b2', name: '香甜玉米', quantity: 5, totalprice: 1000, status: '待收货', type: 'buy' },
  { orderId: 's1', name: '土鸡蛋', quantity: 10, totalprice: 1000, status: '待发货', type: 'sell' }
])

const buyOrders = computed(() => orders.value.filter((o) => o.type === 'buy'))

function testRoute(){
    const route = useRoute()
    return route.path
}
</script>

<style scoped>
.font-medium { font-weight: 500; }
</style>