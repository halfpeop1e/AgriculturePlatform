<template>
    <div class="flex h-screen">
        <SizeBar />
        <div class="flex-1 overflow-auto h-full">
            <slot />
        </div>
    </div>
    <div class="absolute top-0 left-4">
        <el-button :icon="Expand" @click="toggleSidebar" />
    </div>
</template>

<script setup lang="ts">
import SizeBar from '~/components/sizeBar.vue'
import buttonStore from "~/utils/buttonStore"
import { Expand } from '@element-plus/icons-vue'
const buttonstore = buttonStore();
function toggleSidebar() {
    buttonstore.isvisible = !buttonstore.isvisible;
}
</script>

<style scoped>

</style>