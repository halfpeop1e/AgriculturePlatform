<!--
    profilePageLayout.vue
    说明：个人资料页布局（侧边导航 + 主内容区）
    - 使用 `SizeBar` 作为侧边导航，接收 nav 配置并展示分组导航项
    - 右侧为主内容区（通过 slot 注入页面内容）
    - 顶部左上角放置折叠侧边栏的按钮（调用 buttonStore 控制 Drawer 可见性）
    使用场景：用户个人中心、设置页等需要侧边导航的页面。
-->
<template>
  <div>
<div class="flex h-screen">
        <SizeBar :nav="nav"/>
        <div class="flex-1 overflow-auto h-full">
            <slot />
        </div>
    </div>
    <div class="absolute top-0 left-4">
        <el-button :icon="Expand" @click="toggleSidebar" />
    </div>
  </div>
        
</template>

<script setup lang="ts">
import SizeBar from '~/components/sizeBar.vue'
import buttonStore from "~/utils/buttonStore"
import { Expand } from '@element-plus/icons-vue'
import {nav} from '~/config/topbarRoute'
const buttonstore = buttonStore();
function toggleSidebar() {
    buttonstore.isvisible = !buttonstore.isvisible;
}
</script>

<style scoped>

</style>
