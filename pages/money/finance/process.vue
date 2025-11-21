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
          <el-icon><Menu/></el-icon>
          <span class="text-lg">金融产品</span>
        </el-menu-item>
        <el-menu-item index="process">
          <el-icon><Setting /></el-icon>
          <span class="text-lg">进度查询</span>
        </el-menu-item>
      </el-menu>
    </div>

    <div class="overflow-y-auto col-start-2">
      <div class="flex flex-col justify-between h-screen py-20">
        <div class="flex gap-5 items-center bg-teal-200  rounded-sm " v-for="value in datas">
          <img
            :src="value.avatar"
            alt="pic"
            class="w-24"
          />
          <div class="grid grid-cols-3 w-full">
            <span>{{ value.product_name }}</span>
            <span>{{ value.status }}</span>
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
import type { ApplyListRespond } from '~/types/loanApply';

import { Setting } from '@element-plus/icons-vue';
definePageMeta({ layout: "home-page-layout" });
const router = useRouter();
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  router.push(`/money/finance/${key}`);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

onMounted(
  async () => {
    const res = await getLoanApplyList(useUserStore().userId)
    datas.push(...res!)
  }
)
const datas: ApplyListRespond[] = []
const activeName = ref("1");
</script>

<style scoped></style>
