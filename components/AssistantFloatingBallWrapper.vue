<template>
  <!-- 使用 ClientOnly 确保组件只在客户端运行 -->
  <!-- 流式处理需要浏览器 API（ReadableStream），不能在 SSR 中运行 -->
  <ClientOnly>
    <AssistantFloatingBall />
    <template #fallback>
      <!-- SSR 占位符（可选） -->
      <div style="display: none;"></div>
    </template>
  </ClientOnly>
</template>

<script setup lang="ts">
import AssistantFloatingBall from './AssistantFloatingBall.vue'
</script>


