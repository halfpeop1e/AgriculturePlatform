<template>
  <div class="p-6">
    <div class="max-w-6xl mx-auto">
      <!-- 页面标题 -->
      <div class="mb-6">
        <h1 class="text-2xl font-bold">专家咨询</h1>
        <p class="text-sm text-gray-500">找到合适的农业专家，解决种植、养殖难题</p>
      </div>

      <!-- 筛选区域 -->
      <div class="bg-white rounded-lg shadow-sm p-4 mb-6">
        <div class="flex flex-wrap gap-4 items-center">
          <el-input 
            v-model="searchKey" 
            placeholder="搜索专家姓名或领域" 
            clearable 
            class="w-64"
          />
          <el-select 
            v-model="selectedField" 
            placeholder="选择专业领域" 
            clearable
            class="w-48"
          >
            <el-option 
              v-for="field in expertFields" 
              :key="field" 
              :label="field" 
              :value="field"
            />
          </el-select>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
        </div>
      </div>

      <!-- 专家列表 -->
      <div class="flex flex-col gap-4">
        <ExpertCard
          v-for="expert in filteredExperts" 
          :key="expert.id" 
          :expert="expert"
          @click="navigateToDetail(expert.id)"
        />
      </div>

      <!-- 分页 -->
      <div class="mt-6 flex justify-center">
        <Paginator
          :first="(currentPage - 1) * pageSize"
          :rows="pageSize"
          :totalRecords="ExpertDataStore.total"
          :rowsPerPageOptions="[5,10,20]"
          @page="handlePageChange"
          class="rounded-lg border border-slate-200 bg-white p-2 shadow-sm"
        />
      </div>

      <!-- 空状态 -->
      <div v-if="filteredExperts.length === 0" class="text-center text-gray-500 py-10">
        <p>未找到匹配的专家，请尝试其他筛选条件</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ExpertCard from '~/components/Card/ExpertCard.vue'
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import type { Expert } from '@/types/expert'
import { useExpertDataStore } from '~/utils/expertDataStore'
import Paginator from 'primevue/paginator'

definePageMeta({ layout: 'home-page-layout' })
const ExpertDataStore = useExpertDataStore()
const router = useRouter()

const searchKey = ref('')
const selectedField = ref('')

const experts = ref<Expert[]>([])
const currentPage = ref(1)
const pageSize = ref(10)

const expertFields = computed(() => {
  const fields = new Set<string>()
  // prefer store data if available
  const list = ExpertDataStore.experts.length ? ExpertDataStore.experts : experts.value
  list.forEach((expert) => fields.add(expert.field))
  return Array.from(fields)
})

const filteredExperts = computed<Expert[]>(() => {
  return experts.value.filter(expert => {
    const matchSearch = searchKey.value.trim() === '' 
      || expert.name.includes(searchKey.value)
      || expert.field.includes(searchKey.value)
    
    const matchField = selectedField.value === '' 
      || expert.field === selectedField.value

    return matchSearch && matchField
  })
})

const handleSearch = () => {
  // 可选：添加防抖逻辑
  currentPage.value = 1
  ExpertDataStore.fetchExperts(currentPage.value, pageSize.value, selectedField.value || undefined, searchKey.value || undefined).then(() => {
    experts.value = ExpertDataStore.experts ?? []
  })
}

const navigateToDetail = (expertId: string) => {
  router.push(
    `/expert/${expertId}`
  )
}
onMounted(async () => {
  // 模拟获取专家数据
  await ExpertDataStore.fetchExperts(currentPage.value, pageSize.value)
  experts.value = ExpertDataStore.experts ?? []
})

const handlePageChange = (e: { first: number; rows: number }) => {
  const newPage = Math.floor(e.first / e.rows) + 1
  currentPage.value = newPage
  pageSize.value = e.rows
  ExpertDataStore.fetchExperts(currentPage.value, pageSize.value, selectedField.value || undefined, searchKey.value || undefined).then(() => {
    experts.value = ExpertDataStore.experts ?? []
  })
}
</script>

<style scoped>
/* 自定义样式 */
</style>
