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
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <ExpertCard 
          v-for="expert in filteredExperts" 
          :key="expert.id" 
          :expert="expert" 
          @click="navigateToDetail(expert.id)"
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
import ExpertCard from '~/components/ExpertCard.vue'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
// 修复路径导入
import { EXPERTS } from '@/utils/expertDataStore'
// 确保类型正确导入
import type { Expert } from '@/types/expert'

definePageMeta({ layout: 'home-page-layout' })

const router = useRouter()

const searchKey = ref('')
const selectedField = ref('')

const expertFields = [
  '种植业', '养殖业', '农产品加工', '农业技术', '土壤改良', '病虫害防治'
]

const filteredExperts = computed<Expert[]>(() => {
  return EXPERTS.filter(expert => {
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
}

const navigateToDetail = (expertId: string) => {
  router.push(`/experts/${expertId}`)
}
</script>

<style scoped>
/* 自定义样式 */
</style>