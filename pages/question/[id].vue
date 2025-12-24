<template>
  <div class="p-6">
    <el-card shadow="hover" class="rounded-2xl border border-[#E5E7EB] bg-white hover:shadow-md transition max-w-4xl mx-auto">
      <div v-if="questionDetail">
        <!-- 问题部分 -->
        <div class="mb-8">
          <h1 class="text-2xl font-bold mb-4">{{ questionDetail.title }}</h1>
          <div class="text-sm text-gray-500 mb-4">
            {{ questionDetail.author }} · {{ questionDetail.date }}
            <el-tag v-for="tag in questionDetail.tags" :key="tag" class="ml-2">{{ tag }}</el-tag>
          </div>
          <div class="prose max-w-none mb-6" v-html="questionDetail.content" />
        </div>

        <el-divider />

        <!-- 回答列表 -->
        <div class="mt-6">
          <h2 class="text-xl font-semibold mb-4">回答 ({{ questionDetail.answerCount }})</h2>
          
          <div v-if="questionDetail.answer.length === 0" class="text-center text-gray-500 py-8">
            暂无回答
          </div>
          
          <div class="space-y-6">
            <AnswerItem 
              v-for="(answer,index) in questionDetail.answer" 
              :key="index" 
              :answer="answer" 
            />
          </div>
        </div>
      </div>
      
      <div v-else class="text-center text-gray-500 py-8">
        未找到相关问题
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useQuestionDataStore } from '~/utils/questionDataStore'
import { ref, onMounted, watch, computed } from 'vue'
import AnswerItem from '~/components/AnswerItem.vue'

definePageMeta({ layout: 'home-page-layout' })
const route = useRoute()
const questionStore = useQuestionDataStore()
const id = route.params.id as string  // 确保id类型正确（如果后端需要number可改为Number(route.params.id)）

onMounted(() => {
  questionStore.fetchQuestionDetail(id)
})

watch(
  () => questionStore.currentQuestion,
  (newVal) => {
    if (!newVal && id) {
      questionStore.fetchQuestionDetail(id)
    }
  }
)

const questionDetail = computed(() => {
  const detail = questionStore.currentQuestion
  console.log('Computed question detail:', detail)
  if (!detail) return null
  return {
    ...detail,
    tags: detail.tags ?? [],
    answers: detail.answer ?? [],
    answerCount:
      detail.answerCount ??
      (detail.answer ? detail.answer.length : 0)
  }
})
// 页面离开时清除当前问题数据
onUnmounted(() => {
  questionStore.clearCurrentQuestion()
})
</script>
