<template>
  <!-- 模板内容保持不变 -->
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import type { Expert } from '~/types/expert' // 导入专家类型

// 接收专家数据props（使用明确类型）
const props = defineProps<{
  expert: Expert
}>()
</script>