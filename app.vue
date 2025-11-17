<template>
  <!-- 全局毛玻璃背景（位于页面下方，不阻塞交互） -->
  <GlobalGlass />

  <!-- 页面主内容，确保 class app-content 使内容位于背景上方 -->
  <NuxtLayout class="app-content">
    <NuxtPage />
  </NuxtLayout>
</template>

<script setup lang="ts">
import GlobalGlass from '~/components/GlobalGlass.vue'
</script>
