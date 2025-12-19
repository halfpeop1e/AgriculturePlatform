<template>
    <div class="w-full py-10 flex items-center justify-center">
        <Card class="w-4/5 rounded-3xl bg-white/90 shadow-2xl backdrop-blur-md border border-white/40">
            <template #content>
                <div class="space-y-10">
                    <div class="rounded-2xl bg-slate-900/5 p-6">
                        <ImageShow size="500px" :images="Imageitem" />
                    </div>

                    <Divider />

                    <div class="flex flex-col items-center justify-center lg:flex-row">
                        <div class="w-full max-w-2xl">
                            <ImageShow size="350px" :images="ImageNavigate" />
                        </div>

                        <Divider layout="vertical" class="hidden h-72 lg:flex" />

                        <Card class="w-full max-w-sm rounded-2xl border border-slate-200 shadow-lg">
                            <template #content>
                                <div class="space-y-6">
                                    <header class="space-y-2">
                                        <h2 class="text-3xl font-semibold text-slate-900">平台宗旨</h2>
                                        <p class="text-base leading-relaxed text-slate-600">
                                            以数据驱动为核心，融合金融服务与农产品营销，实现产业链上下游的高效对接与智能化管理，助力农业数字化转型。
                                        </p>
                                    </header>

                                    <Divider />

                                    <div class="space-y-3">
                                         <div
                                            class="text-3xl mb-1 fancy-underline text-cyan-700"
                                            role="button"
                                            tabindex="0"
                                            @click="onClickExperience"    
                                          >
                                            立刻体验
                                          </div>
                                        <p class="text-sm text-slate-500">注册或登录账号，开启您的农业数字化之旅！</p>
                                    </div>
                                </div>
                            </template>
                        </Card>
                    </div>

                    <Divider />

                    <div class="flex w-9/10 justify-center">
                        <DataInHome />
                    </div>

                    <Divider />

                    <div class="flex w-full justify-center">
                        <Contactus />
                    </div>
                </div>
            </template>
        </Card>
    </div>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import Divider from 'primevue/divider'
import Button from 'primevue/button'
import DataInHome from '~/components/dataInHome.vue'
import ImageShow from '~/components/imageShow.vue'
import Contactus from '~/components/Contactus.vue'
import { Imageitem,ImageNavigate } from '~/types/imageshow'

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
