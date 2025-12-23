<template>
  <div class="min-h-screen bg-slate-50">
    <div class="grid min-h-screen grid-cols-[280px_1fr]">
      <aside class="border-r border-slate-200 bg-white">
        <div class="flex h-full flex-col items-center px-6 py-10">
          <div class="w-full space-y-3">
            <Button
              label="金融产品"
              icon="pi pi-briefcase"
              class="w-full"
              :severity="route.path === '/money/finance/product' ? 'primary' : 'secondary'"
              :outlined="route.path !== '/money/finance/product'"
              @click="router.push('/money/finance/product')"
            />
            <Button
              label="进度查询"
              icon="pi pi-chart-bar"
              class="w-full"
              :severity="route.path === '/money/finance/process' ? 'primary' : 'secondary'"
              :outlined="route.path !== '/money/finance/process'"
              @click="router.push('/money/finance/process')"
            />
          </div>
        </div>
      </aside>

      <main class="col-start-2 overflow-y-auto px-10 py-12">
        <div class="flex flex-wrap items-center justify-between gap-4">
          <div>
            <h1 class="text-2xl font-semibold text-slate-900">进度查询</h1>
            <p class="text-sm text-slate-500">查看您的贷款申请最新进度</p>
          </div>
          <Button label="刷新" icon="pi pi-refresh" severity="secondary" outlined :loading="loading" @click="fetchApplications" />
        </div>

        <Message v-if="!loading" severity="info" class="mt-6">
          当前共有 {{ totalApplications }} 条申请记录
        </Message>

        <div class="mt-8">
          <div v-if="loading" class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
            <Card v-for="index in 3" :key="index" class="border border-slate-200 shadow-sm">
              <template #content>
                <div class="flex items-center gap-4">
                  <Skeleton shape="circle" size="4rem" />
                  <div class="flex-1 space-y-2">
                    <Skeleton width="50%" height="1.25rem" />
                    <Skeleton width="80%" />
                    <Skeleton width="60%" />
                  </div>
                </div>
              </template>
            </Card>
          </div>

          <div v-else-if="applications.length" class="space-y-4">
            <Card
              v-for="application in applications"
              :key="application.id"
              class="border border-slate-200 shadow-sm transition hover:-translate-y-1 hover:shadow-md"
            >
              <template #content>
                <div class="flex flex-col gap-4 md:flex-row md:items-center">
                  <Avatar :image="application.avatar || defaultAvatar" shape="circle" size="large" />
                  <div class="flex-1 space-y-2">
                    <div class="flex flex-wrap items-center justify-between gap-2">
                      <span class="text-lg font-semibold text-slate-900">{{ application.product_name }}</span>
                      <Tag :value="application.status" :severity="statusSeverity(application.status)" />
                    </div>
                    <div class="flex flex-wrap items-center gap-4 text-sm text-slate-500">
                      <span>申请编号：{{ application.uuid }}</span>
                      <span>产品编号：{{ application.product_id }}</span>
                    </div>
                  </div>
                  <Button
                    label="查看详情"
                    icon="pi pi-eye"
                    severity="secondary"
                    outlined
                    class="w-full md:w-auto"
                    @click="router.push('/money/finance/product')"
                  />
                </div>
              </template>
            </Card>
          </div>

          <Message v-else severity="warn" class="mt-6">
            暂无申请记录，欢迎前往金融产品页面提交贷款申请。
          </Message>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'nuxt/app'
import Button from 'primevue/button'
import Card from 'primevue/card'
import Avatar from 'primevue/avatar'
import Tag from 'primevue/tag'
import Message from 'primevue/message'
import Skeleton from 'primevue/skeleton'
import { useToast } from 'primevue/usetoast'
import type { ApplyListRespond } from '~/types/loanApply'
import { getLoanApplyList } from '~/composables/getLoanProduct'
import { useUserStore } from '~/utils/userStore'

definePageMeta({ layout: 'home-page-layout' })

const router = useRouter()
const route = useRoute()
const toast = useToast()
const userStore = useUserStore()

const loading = ref(false)
const applications = ref<ApplyListRespond[]>([])

const defaultAvatar = 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png'

const totalApplications = computed(() => applications.value.length)

const statusSeverity = (status: string): 'success' | 'info' | 'warning' | 'danger' | 'secondary' => {
  const normalized = status.toLowerCase()
  if (normalized.includes('通过') || normalized.includes('成功')) return 'success'
  if (normalized.includes('拒') || normalized.includes('失败')) return 'danger'
  if (normalized.includes('待') || normalized.includes('审') || normalized.includes('处理中')) return 'warning'
  return 'info'
}

const fetchApplications = async () => {
  try {
    loading.value = true
    const result = await getLoanApplyList(userStore.userId)
    applications.value = Array.isArray(result) ? result : []
  } catch (error) {
    console.error('获取申请列表失败', error)
    toast.add({ severity: 'error', summary: '获取失败', detail: '无法加载申请进度，请稍后重试。', life: 4000 })
  } finally {
    loading.value = false
  }
}

onMounted(fetchApplications)
</script>

<style scoped></style>
