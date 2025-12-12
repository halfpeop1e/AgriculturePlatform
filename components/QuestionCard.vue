<template>
  <div
    class="bg-white rounded-lg border border-gray-100 shadow-sm hover:shadow-lg transform hover:-translate-y-0.5 transition-all duration-200 overflow-hidden cursor-pointer p-4"
  >
    <div class="flex items-start justify-between">
      <div class="flex-1">
        <!-- 标签区域 -->
        <div class="flex items-center gap-2 mb-2 flex-wrap">
          <span
            v-for="tag in question.tags || []"
            :key="tag"
            class="inline-flex items-center text-xs px-2 py-0.5 rounded border border-blue-600 text-black bg-white"
          >
            {{ tag }}
          </span>
        </div>

        <!-- 标题 -->
        <h3 class="text-lg font-semibold text-gray-800 mb-2 line-clamp-2">{{ question.title }}</h3>

        <!-- 内容摘要 -->
        <p class="text-sm text-gray-500 mb-3 line-clamp-2">{{ question.content }}</p>

        <!-- 底部信息 -->
        <div class="flex items-center justify-between text-sm text-gray-500">
          <span>提问者: {{ question.author }}</span>
          <span>{{ question.date }}</span>
          <span>
            <el-badge :value="question.answerCount" class="ml-2" />
            回答
          </span>
        </div>
      </div>
      
      <!-- 状态标识 -->
      <div v-if="question.isAnswered" class="ml-4">
        <el-tag type="success">已解决</el-tag>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import type { Question } from '~/types/questionAnswer'

const props = defineProps<{
  question: Question
}>()
</script>