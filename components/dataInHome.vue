<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mt-6">
    <el-card
      v-for="(stat, index) in stats"
      :key="index"
      shadow="hover"
      style="width: 250px;"
      class="rounded-2xl border border-[#E5E7EB] bg-white hover:shadow-md transition"
    >
      <div class="flex items-center justify-between">
        <div>
          <p class="text-[#6B7280] text-sm mb-1">{{ stat.label }}</p>
          <p class="text-2xl font-semibold text-[#111827]">{{ stat.value }}</p>
        </div>
        <div
          class="w-12 h-12 flex items-center justify-center rounded-xl"
          :class="stat.iconBg"
        >
          <el-icon :size="24" :color="stat.iconColor">
            <component :is="stat.icon" />
          </el-icon>
        </div>
      </div>
      <p
        class="mt-2 text-sm"
        :class="stat.trend > 0 ? 'text-green-600' : 'text-red-600'"
      >
        {{ stat.trend > 0 ? '+' + stat.trend : stat.trend }}% 本周变化
      </p>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ArrowUp, ArrowDown, User, ShoppingCart, Wallet, PieChart } from '@element-plus/icons-vue'

const stats = [
  {
    label: '本月销售额',
    value: '¥256,430',
    trend: 8.4,
    icon: Wallet,
    iconBg: 'bg-[#1E3A8A]/10',
    iconColor: '#1E3A8A',
  },
  {
    label: '新增客户',
    value: '1,203',
    trend: 5.2,
    icon: User,
    iconBg: 'bg-[#C9A227]/10',
    iconColor: '#C9A227',
  },
  {
    label: '转化率',
    value: '12.8%',
    trend: -1.3,
    icon: PieChart,
    iconBg: 'bg-[#3B82F6]/10',
    iconColor: '#3B82F6',
  },
  {
    label: '订单数',
    value: '3,475',
    trend: 3.9,
    icon: ShoppingCart,
    iconBg: 'bg-[#16A34A]/10',
    iconColor: '#16A34A',
  },
]
</script>
