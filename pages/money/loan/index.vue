<template>
  <div>
    <div className="min-h-screen flex flex-col items-center pt-10 gap-4">
    <div
      class="bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center"
    >
      <el-image
        style="width: 200px; height: 240px"
        :src="farmer1"
        :fit="'cover'"
      />
      <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
        <span
          class="text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2"
          >融资难、流程繁，制约生产扩张?
        </span>
      </div>
    </div>
    <div
      class="bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center"
    >
      <el-image
        style="width: 200px; height: 240px"
        :src="farmer2"
        :fit="'cover'"
      />
      <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
        <span
          class="text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2"
          >抗风险能力弱，制约农业可持续发展？
        </span>
      </div>
    </div>
    <div
      v-on:click="nextPage"
      class="bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center"
    >
      <el-image style="width: 200px; height: 240px" :src="man" :fit="'cover'" />
      <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
        <span
          class="text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2"
          >兄弟,把你的手机给我😄
        </span>
        <div class="flex justify-center">
          <el-icon :size="45"><ArrowDownBold /></el-icon>
        </div>
      </div>
    </div>
  </div>
  <div
    class=" h-screen flex flex-col gap-10 items-center justify-center"
    ref="section2"
  >
  <div class="text-center text-3xl mb-11">
    Ave木吉卡金融 是本平台推出的金融服务，提供贷款功能，为用户带来简约的贷款服务。
  </div>
    <div
      class="w-1/2 h-1/2 bg-slate-100 rounded-lg grid grid-rows-[100px_1fr] grid-flow-col overflow-hidden gap-1"
    >
      <div
        class="h-full min-h-[166px] bg-gradient-to-t from-transparent via-purple-500 to-pink-500"
      >
        <div
          class="flex w-full pt-5 pl-8 justify-between relative h-100px z-10"
        >
          <div class="flex">
            <el-avatar :size="75" :src="userStore.avatar || man" />
            <div class="pl-6 text-lg flex items-center">
              {{ userStore.userinfo.nickname }}
            </div>
          </div>
          <div class="absolute right-10 top-0">
            <el-image
              style="width: 100px; height: 100px"
              :src="muzimi"
              :fit="'cover'"
            />
          </div>
        </div>
      </div>
      <div class="bg-slate-50/45 mx-4 rounded-t-lg grid grid-cols-2 z-50 -mt-1">
        <div class="flex flex-col items-center gap-2 pt-5">
          <span class="text-sm"> 我的额度(元) </span>
          <span class="text-3xl">200000.00</span>
          <el-button type="primary" size="large"> 贷款 </el-button>
        </div>
        <div class="flex flex-col items-center gap-2 pt-5">
          <span class="text-sm"> 我的待还(元) </span>
          <span class="text-3xl">0.00</span>
          <el-button type="primary" size="large"> 还款 </el-button>
        </div>
        <div class="col-span-2 px-10">
          <div class="flex flex-col gap-3 items-center">
            <span class="text-blue-500 ">{{ nextMonth }}月应还(元)</span>
            <span class="text-3xl">已还清</span>
          </div>
        </div>
        <h1 class=" hover:cursor-pointer flex justify-center col-span-2">您的账单已出 ></h1>
      </div>
    </div>
  </div>
  </div>
  
</template>

<script setup lang="ts">
definePageMeta({ layout: "home-page-layout" });
import man from "~/public/ioanImage/man.png";
import { ArrowDownBold, PictureRounded } from "@element-plus/icons-vue";
import farmer1 from "~/public/ioanImage/1762280990064.png";
import farmer2 from "~/public/ioanImage/1762281252724.png";
import muzimi from "~/public/ioanImage/1762350715242.png";

const userStore = useUserStore();
const section2 = ref<HTMLElement | null>(null);
const nextPage = () => {
  if (section2.value) {
    section2.value.scrollIntoView({
      behavior: "smooth",
    });
  }
};

const now = new Date
const nextMonth = now.getMonth() + 2
</script>

<style scoped>
</style>
