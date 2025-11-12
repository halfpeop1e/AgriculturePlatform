<!--
    sizeBar.vue
    说明：应用的侧边导航（抽屉式）组件
    Props:
        - nav: 路由分组数组（类型 RouteType[]），每个分组包含 label 与 items
    功能：显示分组导航、根据当前路由高亮当前项并支持点击跳转（navigateTo）
    依赖：使用了全局的 buttonStore 用于控制 Drawer 的可见性。
-->
<template>
        <Drawer v-model:visible="buttonstore.isvisible">
        <aside
        class="w-64 min-h-full bg-white/80 dark:bg-slate-900/80 border-r border-slate-200/70 dark:border-slate-700/60 backdrop-blur-lg"
    >   
    <div class="h-16 flex items-center px-4">
            <img class="h-10 w-10 mr-1" src="/logo.png" alt="Logo" />
            <button class="text-lg font-bold mr-4" @click="navigateTo('/homePage')">农产品融销一体平台</button>
            </div>
        <nav class="p-6 space-y-8">
            <div v-for="section in porps.nav" :key="section.label" class="space-y-3">
                <h2 class="text-xl font-semibold uppercase tracking-wide text-slate-400 dark:text-slate-500">
                    {{ section.label }}
                </h2>
                <ul class="space-y-2">
                    <li v-for="item in section.items" :key="item.routername">
                        <button
                            type="button"
                            :icon="item.icon"
                            class="w-full text-left px-4 py-2.5 rounded-xl transition-all duration-200 flex items-center gap-2 text-lg font-medium"
                            :class="[
                                isActive(item.routername)
                                    ? 'bg-emerald-500 text-white shadow-lg shadow-emerald-500/30'
                                    : 'text-slate-600 dark:text-slate-300 bg-white/60 dark:bg-slate-800/60 hover:bg-emerald-50 dark:hover:bg-emerald-500/10 hover:text-emerald-600 hover:shadow-lg'
                            ]"
                            @click="handleNavigate(item.routername)"
                        >   
                            <span :class="item.icon"></span>
                            <span class="flex-1">{{ item.label }}</span>
                            <el-icon v-if="isActive(item.routername)" class="text-white">
                                <ArrowRightBold />
                            </el-icon>
                        </button>
                    </li>
                </ul>
            </div>
        </nav>
    </aside>
    </Drawer>
    
</template>

<script setup lang="ts">
import { ArrowRightBold } from '@element-plus/icons-vue'
import type{ RouteType } from '~/types/routes'
import { nav } from '~/config/topbarRoute'
import buttonStore from "~/utils/buttonStore"
const buttonstore = buttonStore();
const visible = ref(true);
const route = useRoute()
const porps=defineProps({
    nav:{
      type:Array as () => RouteType[],
      required:true,
    }
})
function isActive(path: string) {
    return route.path === path
}
function handleNavigate(path: string) {
    if (route.path !== path) {
        navigateTo(path)
    }
}
</script>
