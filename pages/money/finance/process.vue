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
          <el-collapse v-model="activeName" accordion>
            <el-collapse-item title="Consistency" name="1">
              <div>
                Consistent with real life: in line with the process and logic of
                real life, and comply with languages and habits that the users
                are used to;
              </div>
              <div>
                Consistent within interface: all elements should be consistent,
                such as: design style, icons and texts, position of elements,
                etc.
              </div>
            </el-collapse-item>
            <el-collapse-item title="Feedback" name="2">
              <div>
                Operation feedback: enable the users to clearly perceive their
                operations by style updates and interactive effects;
              </div>
              <div>
                Visual feedback: reflect current state by updating or
                rearranging elements of the page.
              </div>
            </el-collapse-item>
            <el-collapse-item title="Efficiency" name="3">
              <div>
                Simplify the process: keep operating process simple and
                intuitive;
              </div>
              <div>
                Definite and clear: enunciate your intentions clearly so that
                the users can quickly understand and make decisions;
              </div>
              <div>
                Easy to identify: the interface should be straightforward, which
                helps the users to identify and frees them from memorizing and
                recalling.
              </div>
            </el-collapse-item>
            <el-collapse-item title="Controllability" name="4">
              <div>
                Decision making: giving advices about operations is acceptable,
                but do not make decisions for the users;
              </div>
              <div>
                Controlled consequences: users should be granted the freedom to
                operate, including canceling, aborting or terminating current
                operation.
              </div>
            </el-collapse-item>
          </el-collapse>
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
