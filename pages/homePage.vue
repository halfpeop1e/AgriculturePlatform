<template>
    <div>
       <el-card style="border-radius: 20px;background-color: rgba(255, 255, 255, 0.244);">
      <div>
        <ImageShow size='500px' :images="Imageitem"/>
      </div>
    <el-divider />
    <div class="flex items-center justify-center w-full">
        <div style="width: 70%;height:350px; margin-top: 20px;">
        <ImageShow size="350px" :images="Imageitem"/>  
        </div> 
        <el-divider direction="vertical" class="h-full"/>
        <el-card style="width: 310px;margin-top: 20px;height: 350px;">
        <div>
        <p class="text-3xl mb-1">平台宗旨</p>
        <div class="text-lg">以数据驱动为核心，融合金融服务与农产品营销，实现产业链上下游的高效对接与智能化管理，助力农业数字化转型。</div>
        </div>
        <el-divider/>
        <div>
                        <div
                            class="text-3xl mb-1 fancy-underline text-cyan-700"
                            role="button"
                            tabindex="0"
                            @click="onClickExperience"    
                        >
                            立刻体验
                        </div>
            <div class="text-lg">注册或登录账号，开启您的农业数字化之旅！</div>
        </div>
        </el-card>
        </div>
    
    <el-divider />
    <div style="width: 100%;" class="flex justify-center">
    <DataInHome />      
    </div>
    <el-divider />
    <div style="width: 100%; margin-top: 20px;" class="flex justify-center">
        <Contactus/>
    </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import DataInHome from '~/components/dataInHome.vue'
import ImageShow from '~/components/imageShow.vue'
import Contactus from '~/components/Contactus.vue'
import { Imageitem } from '~/types/imageshow'
definePageMeta({ layout: 'home-page-layout' })

function onClickExperience() {
    navigateTo('/login')
}
</script>

<style scoped>
.fancy-underline {
    position: relative;
    display: inline-block;
    cursor: pointer;
}

.fancy-underline::after {
    content: '';
    position: absolute;
    left: 0;
    bottom: -4px;
    width: 0%;
    height: 3px;
    background: linear-gradient(90deg, rgb(111, 111, 111), rgb(161, 175, 38),rgb(92, 187, 221));
    border-radius: 4px;
    transition: width 850ms cubic-bezier(.2,.8,.2,1), opacity 200ms ease;
    opacity: 0;
}

.fancy-underline:hover::after,
.fancy-underline.clicked::after {
    width: 100%;
    opacity: 1;
}


</style>
