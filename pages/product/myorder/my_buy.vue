<template>
    <div class="p-4">
        <div class="flex items-center gap-4 mb-4">
            <el-button-group>
                <el-button :type="selectedTab === 'buy' ? 'primary' : 'default'" @click="selectedTab = 'buy'">买入订单</el-button>
                <el-button :type="selectedTab === 'sell' ? 'primary' : 'default'" @click="selectedTab = 'sell'">卖出订单</el-button>
            </el-button-group>

            <div class="text-sm text-gray-500">共 {{ orders.length }} 条</div>
        </div>

        <div v-if="selectedTab === 'buy'">
            <el-alert title="买入订单" type="info" show-icon />
            <ul class="mt-3 space-y-3">
                <li v-for="o in buyOrders" :key="o.orderId">
                    <el-card>
                        <div class="flex justify-between items-center">
                            <div class="font-medium">{{ o.name }} × {{ o.quantity }}</div>
                            <div class="text-sm text-gray-500">状态：{{ o.status }}</div>
                        </div>
                        <div class="flex items-center">
                            <h1 class="text-black text-m ">
                            订单价格 ：
                            </h1>
                            <div class="text-green-600 text-xl font-bold mr-1">
                            {{ o.totalprice }} 
                            </div>
                            元
                        </div>
                    </el-card>
                </li>
                <li v-if="buyOrders.length === 0" class="text-gray-500">暂无买入订单</li>
            </ul>
        </div>

        <div v-else>
            <el-alert title="卖出订单" type="warning" show-icon />
            <ul class="mt-3 space-y-3">
                <li v-for="o in sellOrders" :key="o.orderId">
                    <el-card>
                        <div class="flex justify-between items-center">
                            <div class="font-medium">{{ o.name }} × {{ o.quantity }}</div>
                            <div class="text-sm text-gray-500">状态：{{ o.status }}</div>
                        </div>
                    </el-card>
                </li>
                <li v-if="sellOrders.length === 0" class="text-gray-500">暂无卖出订单</li>
            </ul>
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({ layout: 'home-page-layout' })
import { ref, computed } from 'vue'
import type { Order } from '@/types/myOrder'
type Tab = 'buy' | 'sell'

const selectedTab = ref<Tab>('buy')


const orders = ref<Order[]>([
    { orderId: 'b1', name: '绿色有机苹果', quantity: 2,totalprice:1000, status: '已支付', type: 'buy' },
    { orderId: 'b2', name: '香甜玉米', quantity: 5, totalprice:1000,status: '待收货', type: 'buy' },
    { orderId: 's1', name: '土鸡蛋', quantity: 10,totalprice:1000, status: '待发货', type: 'sell' }
])

const buyOrders = computed(() => orders.value.filter((o) => o.type === 'buy'))
const sellOrders = computed(() => orders.value.filter((o) => o.type === 'sell'))
</script>

<style scoped>
.font-medium { font-weight: 500; }
</style>