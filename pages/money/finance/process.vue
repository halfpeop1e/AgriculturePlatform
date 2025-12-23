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
            <div class="pt-6">
              <Button label="审批产品" icon="pi pi-plus" class="w-full" severity="help" @click="check = true" />
            </div>
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

    <Dialog v-model:visible="check" header="贷款产品审批" :style="{ width: '35rem' }" modal class="p-fluid">
      <div class="flex flex-col gap-6">
        
        <!-- 1. 选择待审批的单据 -->
        <div class="flex flex-col gap-2">
          <label for="audit-select" class="font-medium text-slate-700">选择待审批单据</label>
          <Dropdown 
            id="audit-select"
            v-model="selectedAuditId" 
            :options="pendingApplications" 
            optionLabel="displayLabel" 
            optionValue="id" 
            placeholder="请选择一条申请记录"
            class="w-full"
            emptyMessage="当前没有待审批的申请"
          />
        </div>

        <!-- 2. 单据详情 (选中后显示) -->
        <div v-if="selectedApplication" class="rounded-lg bg-slate-50 p-4 border border-slate-200 space-y-3">
          <div class="flex items-center gap-3 border-b border-slate-200 pb-3">
            <Avatar :image="selectedApplication.avatar || defaultAvatar" shape="circle" />
            <div>
              <p class="font-bold text-slate-800">{{ selectedApplication.product_name }}</p>
              <p class="text-xs text-slate-500">申请编号: {{ selectedApplication.uuid }}</p>
            </div>
          </div>
          <div class="grid grid-cols-2 gap-y-2 text-sm">
            <span class="text-slate-500">当前状态:</span>
            <span class="font-medium text-amber-600">{{ selectedApplication.status }}</span>
          </div>
        </div>

        <!-- 3. 审批操作表单 -->
        <div v-if="selectedApplication" class="flex flex-col gap-4 border-t border-slate-100 pt-4">
          <div class="flex flex-col gap-2">
            <label class="font-medium text-slate-700">审批结果</label>
            <div class="flex gap-4">
              <div class="flex items-center">
                <RadioButton v-model="auditForm.result" inputId="pass" name="audit" value="通过" />
                <label for="pass" class="ml-2 cursor-pointer text-green-600 font-medium">通过</label>
              </div>
              <div class="flex items-center">
                <RadioButton v-model="auditForm.result" inputId="reject" name="audit" value="拒绝" />
                <label for="reject" class="ml-2 cursor-pointer text-red-600 font-medium">驳回</label>
              </div>
            </div>
          </div>

          <div class="flex flex-col gap-2">
            <label for="remark" class="font-medium text-slate-700">审批意见</label>
            <Textarea 
              id="remark" 
              v-model="auditForm.remark" 
              rows="3" 
              placeholder="请输入审批备注..." 
              class="w-full" 
            />
          </div>
        </div>

        <!-- 空状态提示 -->
        <div v-else-if="!pendingApplications.length" class="text-center py-6 text-slate-400 text-sm">
          <i class="pi pi-check-circle text-4xl mb-2 text-slate-300"></i>
          <p>当前无需审批的单据</p>
        </div>

      </div>

      <template #footer>
        <Button label="取消" icon="pi pi-times" severity="secondary" @click="check = false" />
        <Button 
          label="提交审批" 
          icon="pi pi-check" 
          severity="primary" 
          :disabled="!selectedApplication || submitting" 
          :loading="submitting"
          @click="submitAudit" 
        />
      </template>
    </Dialog>
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
import Dropdown from 'primevue/dropdown'
import RadioButton from 'primevue/radiobutton'
import Textarea from 'primevue/textarea'

definePageMeta({ layout: 'home-page-layout' })

const router = useRouter()
const route = useRoute()
const toast = useToast()
const userStore = useUserStore()

const loading = ref(false)
const applications = ref<ApplyListRespond[]>([])
const check = ref(false)

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



// --- 审批相关逻辑 ---
const selectedAuditId = ref(null)
const submitting = ref(false)
const auditForm = ref({
  result: '通过',
  remark: ''
})

// 过滤出待审批的申请（为了演示，这里假设包含'待'字样的状态都可以审批，或者你可以列出所有单据方便测试）
const pendingApplications = computed(() => {
  if (!applications.value) return []
  return applications.value.map(app => ({
    ...app,
    displayLabel: `${app.product_name} - ${app.status} (${app.uuid.substring(0, 8)}...)`
  }))
})

// 获取当前选中的单据详情
const selectedApplication = computed(() => {
  return applications.value.find(app => app.id === selectedAuditId.value)
})

// 提交审批函数
const submitAudit = async () => {
  if (!selectedApplication.value) return

  submitting.value = true
  
  // 模拟API请求延迟
  await useApplyLoanResult(selectedApplication.value.uuid, auditForm.value.result === '通过')

  // 更新本地数据状态（模拟后端更新）
  const appIndex = applications.value.findIndex(app => app.id === selectedAuditId.value)
  if (appIndex !== -1) {
    // 根据选择更新状态
    const newStatus = auditForm.value.result === '通过' ? '审批通过' : '已拒绝'
    applications.value[appIndex]!.status = newStatus
    
    toast.add({ 
      severity: auditForm.value.result === '通过' ? 'success' : 'error', 
      summary: '审批完成', 
      detail: `单据 ${applications.value[appIndex]!.product_id} 已${auditForm.value.result}`, 
      life: 3000 
    })
  }

  // 重置表单并关闭
  submitting.value = false
  selectedAuditId.value = null
  auditForm.value.remark = ''
  auditForm.value.result = '通过'
  check.value = false
}
</script>

<style scoped></style>
