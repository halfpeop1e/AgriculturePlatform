<template>
  <div class="min-h-screen grid grid-cols-[300px_1fr]">
    <div
      class="w-[300px] h-screen overflow-y-auto flex flex-col items-center gap-8"
    >
      <el-menu
        default-active="product"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        style="width: 250px; height: 300px"
      >
        <el-menu-item index="person">
          <el-icon><icon-menu /></el-icon>
          <span class="text-lg">个人中心</span>
        </el-menu-item>
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
      <div class="grid grid-cols-4 gap-5">
        <el-card style="max-width: 480px">
          <template #header>
            <div class="grid grid-cols-2 items-center gap-2">
              <div>产品1</div>
              <img
                class="w-20"
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              />
            </div>
          </template>
          <span>描述</span>
        </el-card>
        <el-card style="max-width: 480px">
          <template #header>
            <div class="grid grid-cols-2 items-center gap-2">
              <div>产品1</div>
              <img
                class="w-20"
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              />
            </div>
          </template>
          <span>描述</span>
        </el-card>
        <el-card style="max-width: 480px">
          <template #header>
            <div class="grid grid-cols-2 items-center gap-2">
              <div>产品1</div>
              <img
                class="w-20"
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              />
            </div>
          </template>
          <span>描述</span>
        </el-card>
        <el-card style="max-width: 480px">
          <template #header>
            <div class="grid grid-cols-2 items-center gap-2">
              <div>产品1</div>
              <img
                class="w-20"
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              />
            </div>
          </template>
          <span>描述</span>
        </el-card>
        <el-card style="max-width: 480px">
          <template #header>
            <div class="grid grid-cols-2 items-center gap-2">
              <div>产品1</div>
              <img
                class="w-20"
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              />
            </div>
          </template>
          <span>描述</span>
        </el-card>
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
</script>

<style scoped></style>
