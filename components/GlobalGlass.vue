<template>
  <!-- 全屏固定的毛玻璃层（轻薄的蓝白渐变 + 毛玻璃） -->
  <div class="global-glass" aria-hidden="true">
    <div class="glass-overlay"></div>
  </div>
</template>

<script setup lang="ts">
// GlobalGlass 只负责渲染浅蓝→白的线性渐变 + backdrop blur，移除所有交互与烟雾效果。
</script>

<style scoped>
.global-glass {
  position: fixed;
  inset: 0;
  z-index: 0;
  pointer-events: none;
  overflow: hidden;
}

.glass-overlay {
  position: absolute;
  inset: 0;
  /* 浅蓝到白色的线性渐变，保持轻薄视觉 */
  background: linear-gradient(135deg, rgba(5, 182, 241, 0.22) 0%,rgba(111, 195, 208, 0.62) 40%, rgba(255,255,255,0.62) 65%);
  backdrop-filter: blur(14px) saturate(120%);
  -webkit-backdrop-filter: blur(14px) saturate(120%);
  z-index: 1;
  pointer-events: none;
}

/* 响应式：移动端降低模糊以节省性能 */
@media (max-width: 640px) {
  .glass-overlay { backdrop-filter: blur(10px); -webkit-backdrop-filter: blur(10px); }
}
</style>
