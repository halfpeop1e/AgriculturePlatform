<template>
  <div class="min-h-screen grid grid-cols-[300px_1fr]">
    <div
      class="w-[300px] h-screen overflow-y-auto flex flex-col items-center gap-8"
    >
      <el-menu
        default-active="person"
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
      <div class="pl-10 pt-10">
        <el-descriptions
          title="Width vertical list"
          direction="vertical"
          style="margin-top: 20px"
        >
          <el-descriptions-item :rowspan="2" :width="140" label="Photo">
            <el-image
              style="width: 100px; height: 100px"
              src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
            />
          </el-descriptions-item>
          <el-descriptions-item label="Username"
            >kooriookami</el-descriptions-item
          >
          <el-descriptions-item label="Telephone"
            >18100000000</el-descriptions-item
          >
          <el-descriptions-item label="Place">Suzhou</el-descriptions-item>
          <el-descriptions-item label="Remarks">
            <el-tag size="small">School</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="Address">
            No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province
          </el-descriptions-item>
        </el-descriptions>
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
