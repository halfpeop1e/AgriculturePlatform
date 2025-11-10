<template>
  <div class="p-6">
    <div class="max-w-6xl mx-auto flex gap-6">
      <SizeBar :nav="settingNav" />
        <div class="absolute top-16 left-1">
        <el-button :icon="Expand" @click="toggleSidebar" />
      </div>
      <div class="flex-1 bg-white rounded-lg shadow-sm p-6">
        <!-- 子路由渲染点 -->
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import SizeBar from '~/components/sizeBar.vue'
import { settingNav } from '~/config/settingRoute'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Expand } from '@element-plus/icons-vue'
definePageMeta({ layout: 'home-page-layout' })
const buttonstore = buttonStore();
const router = useRouter()
function toggleSidebar() {
    buttonstore.isvisible = !buttonstore.isvisible;
}
// 当访问 /setting 时，自动导航到 /setting/profile
onMounted(() => {
  if (router.currentRoute.value.path === '/setting') {
    router.replace('/setting/profile')
  }
})
</script>

<style scoped>
.max-w-6xl { max-width: 72rem }
</style>
