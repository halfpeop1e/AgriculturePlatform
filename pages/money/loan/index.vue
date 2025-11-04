<template>
   <div className="min-h-screen flex flex-col items-center pt-10 gap-4">
         <div class=" bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center">
            <el-image style="width: 200px; height: 240px" :src="farmer1" :fit="'cover'" />
            <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
               <span class=" text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2 ">融资难、流程繁，制约生产扩张?
            </span>
            </div>
         </div>
         <div class=" bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center">
            <el-image style="width: 200px; height: 240px" :src="farmer2" :fit="'cover'" />
            <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
               <span class=" text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2 ">抗风险能力弱，制约农业可持续发展？
            </span>
            </div>
         </div>
         <div v-on:click="nextPage" class=" bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center">
            <el-image style="width: 200px; height: 240px" :src="man" :fit="'cover'" />
            <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
               <span class=" text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2 ">兄弟,把你的手机给我😄
            </span>
            <div class=" flex justify-center">
               <el-icon :size="45"><ArrowDownBold /></el-icon>
            </div>
            </div>
         </div>
    </div>
    <div class=" bg-red-300 h-screen flex items-center justify-center" ref="section2">
      to be continued
    </div>

</template>

<script setup lang="ts">
definePageMeta({ layout: 'home-page-layout' })
import man from '~/public/ioanImage/man.png'
import { ArrowDownBold } from '@element-plus/icons-vue';
import farmer1 from '~/public/ioanImage/1762280990064.png'
import farmer2 from '~/public/ioanImage/1762281252724.png'
const section2 = ref<HTMLElement | null>(null)
const nextPage = ()=>{
  if (section2.value) {
    section2.value.scrollIntoView({
      behavior: 'smooth'
    });
  }
}
</script>

<style scoped>
</style>