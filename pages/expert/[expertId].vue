<template>
  <div class="p-6 min-h-screen bg-gray-50">
    <div class="max-w-4xl mx-auto">
      <Button label="返回" icon="pi pi-arrow-left" class="p-button-text mb-4" @click="goBack" />

      <Card v-if="expertDetail" class="shadow-sm">
        <template #title>
          <div class="flex gap-4 items-center">
            <Avatar :image="expertDetail.avatar || defaultAvatar" shape="circle" class="w-20 h-20" />
            <div>
              <div class="flex items-center gap-2">
                <h2 class="text-2xl font-semibold text-gray-800">{{ expertDetail.name }}</h2>
                <Tag :value="expertDetail.field" severity="info" />
              </div>
              <div class="mt-2 flex flex-wrap gap-4 text-sm text-gray-500">
                <span class="flex items-center gap-1">
                  <i class="pi pi-users text-primary-500"></i>
                  {{ expertDetail.consultCount }} 次咨询
                </span>
                <span class="flex items-center gap-1">
                  <i class="pi pi-star-fill text-yellow-500"></i>
                  {{ expertDetail.rating.toFixed(1) }} 分
                </span>
                <span v-if="expertDetail.responseTime" class="flex items-center gap-1">
                  <i class="pi pi-clock text-emerald-500"></i>
                  {{ expertDetail.responseTime }}
                </span>
              </div>
            </div>
          </div>
        </template>

        <template #content>
          <p class="text-gray-600 leading-relaxed">
            {{ expertDetail.introduction }}
          </p>

          <Divider />

          <section class="grid md:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold text-gray-800 mb-3">擅长技能</h3>
              <div class="flex flex-wrap gap-2">
                <Tag
                  v-for="skill in expertDetail.skills"
                  :key="skill"
                  :value="skill"
                  severity="secondary"
                  class="bg-slate-100 text-slate-700"
                />
              </div>
            </div>

            <div v-if="expertDetail.availableTime?.length">
              <h3 class="text-lg font-semibold text-gray-800 mb-3">可咨询时间</h3>
              <ul class="space-y-2 text-sm text-gray-600">
                <li v-for="time in expertDetail.availableTime" :key="time">{{ time }}</li>
              </ul>
            </div>

            <div v-if="expertDetail.education">
              <h3 class="text-lg font-semibold text-gray-800 mb-3">教育背景</h3>
              <p class="text-sm text-gray-600">{{ expertDetail.education }}</p>
            </div>

            <div v-if="expertDetail.experience">
              <h3 class="text-lg font-semibold text-gray-800 mb-3">工作经历</h3>
              <p class="text-sm text-gray-600">{{ expertDetail.experience }}</p>
            </div>

            <div v-if="expertDetail.certification?.length">
              <h3 class="text-lg font-semibold text-gray-800 mb-3">资质证书</h3>
              <ul class="space-y-2 text-sm text-gray-600">
                <li v-for="cert in expertDetail.certification" :key="cert">{{ cert }}</li>
              </ul>
            </div>

            <div v-if="expertDetail.serviceScope?.length">
              <h3 class="text-lg font-semibold text-gray-800 mb-3">服务范围</h3>
              <ul class="space-y-2 text-sm text-gray-600">
                <li v-for="scope in expertDetail.serviceScope" :key="scope">{{ scope }}</li>
              </ul>
            </div>
          </section>

          <Divider />

          <section v-if="expertDetail.recentCases?.length" class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-800">近期案例</h3>
            <div
              v-for="caseItem in expertDetail.recentCases"
              :key="`${caseItem.date}-${caseItem.question}`"
              class="rounded-md border border-slate-200 bg-slate-50 p-4"
            >
              <div class="text-xs text-gray-400">{{ caseItem.date }}</div>
              <div class="mt-2 text-sm text-gray-700 font-medium">问题：{{ caseItem.question }}</div>
              <div class="mt-1 text-sm text-gray-600">解答：{{ caseItem.answer }}</div>
            </div>
          </section>
        </template>

        <template #footer>
          <div class="flex items-center justify-between">
            <div class="flex gap-3">
              <Button label="在线咨询" icon="pi pi-comments" severity="primary" @click="showDialog()"/>
              <Button label="预约服务" icon="pi pi-calendar" severity="secondary" />
            </div>
          </div>
        </template>
      </Card>

      <div v-else class="bg-white rounded-md shadow p-6 text-center text-gray-500">
        未找到专家信息，请返回列表重试。
      </div>
      <ExpertQuestionDialog v-model:model-value="showdialog" :expert-id="expertId" :expert="expertDetail?.name || ''" @close="showdialog = false" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import ExpertQuestionDialog from '~/components/Dialog/expertQuestionDialog.vue'
import { useRoute, useRouter } from 'vue-router'
import Card from 'primevue/card'
import Avatar from 'primevue/avatar'
import Tag from 'primevue/tag'
import Button from 'primevue/button'
import Divider from 'primevue/divider'
import type { ExpertDetail } from '~/types/expert'
import { getMockExpertById } from '~/types/experttest'
import { useExpertDataStore } from '~/utils/expertDataStore'
const expertDataStore = useExpertDataStore()
definePageMeta({ layout: 'home-page-layout' })

const props = defineProps<{
  expert?: ExpertDetail
}>()

const route = useRoute()
const router = useRouter()
const showdialog = ref(false)
const defaultAvatar = '/ioanImage/default-avatar.png'
const expertId = computed(() => String(route.params.expertId || ''))
const expertDetail = ref<ExpertDetail | null>(null)
const showDialog = () => {
  showdialog.value = true
}
const goBack = () => {
  router.back()
}
onMounted(async () => {
  await expertDataStore.fetchExpertDetail(expertId.value)
  expertDetail.value = expertDataStore.currentExpert
  console.log(expertDetail.value)
})
</script>
<style scoped>

</style>
