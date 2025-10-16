<template>
    <div class="w-full bg-neutral-400 text-white flex items-center justify-between shadow-lg rounded-md border-b border-gray-200">
        <div class="h-16 flex items-center px-4">
            <img class="h-10 w-10 mr-2" src="/logo.png" alt="Logo" />
            <button class="text-2xl font-bold mr-4" @click="navigateTo('/homePage')">农产品融销一体平台</button>
            <!-- 默认插槽（如需在左侧插入更多内容） -->
            <slot />
        </div>

        <div class="h-16 flex items-center px-4">
            <slot name="end">
                <button class="bg-white text-green-900 px-3 py-1 rounded">默认按钮</button>
            </slot>
        </div>
    </div>
</template>

<script setup lang="ts">
// 如果需要对插槽类型进行 TypeScript 声明，可以在这里添加，但不是必须的。
</script>

<style scoped>
.bg-green-900 { background-color: #065f46; }

/* Slight inner highlight to make the bar feel elevated */
.rounded-md {
    overflow: hidden;
}

.shadow-lg {
    box-shadow: 0 10px 30px rgba(2,6,23,0.15);
}
</style>