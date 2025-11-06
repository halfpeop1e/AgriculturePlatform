<template>
   <div class="fixed top-50 left-0 w-[300px] h-screen overflow-y-auto flex flex-col items-center  gap-8">
       <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
         style="width: 250px; height: 300px;"
      >
        <el-menu-item index="1">
          <el-icon><icon-menu /></el-icon>
          <span class="text-lg">个人中心</span>
        </el-menu-item>
        <el-menu-item index="2">
          <el-icon><icon-menu /></el-icon>
          <span class="text-lg">金融产品</span>
        </el-menu-item>
        <el-menu-item index="3">
          <el-icon><document /></el-icon>
          <span class="text-lg">贷款申请</span>
        </el-menu-item>
        <el-menu-item index="4">
          <el-icon><setting /></el-icon>
          <span class="text-lg">贷款进度</span>
        </el-menu-item>
      </el-menu>
    </div>
  <div class="min-h-screen grid grid-cols-[300px_1fr]">
      
    <div class="overflow-y-auto col-start-2">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ layout: "home-page-layout" });

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
</style>
