<template>
  <div class="min-h-screen grid grid-cols-[300px_1fr]">
    <div
      class="w-[300px] h-screen overflow-y-auto flex flex-col items-center gap-8"
    >
      <el-menu
        default-active="process"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        style="width: 250px; height: 300px"
      >
        <el-menu-item index="product">
          <el-icon><icon-menu /></el-icon>
          <span class="text-lg">金融产品</span>
        </el-menu-item>
        <el-menu-item index="process">
          <el-icon><setting /></el-icon>
          <span class="text-lg">进度查询</span>
        </el-menu-item>
      </el-menu>
    </div>

    <div class="overflow-y-auto col-start-2">
      <div class="flex flex-col justify-between h-screen py-20">
        <div class="flex gap-5 items-center">
          <img
            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
            alt="pic"
            class="w-20"
          />
          <div class="grid grid-cols-3 w-full">
            <span>产品名称</span>
            <span>产品状态</span>
            <span>查看详情</span>
          </div>
        </div>
        <div class="demo-collapse">
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
definePageMeta({ layout: "home-page-layout" });
const router = useRouter();
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  router.push(`/money/finance/${key}`);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const activeName = ref("1");
</script>

<style scoped></style>
