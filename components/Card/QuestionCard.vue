<template>
  <div
    class="cursor-pointer overflow-hidden rounded-lg border border-gray-100 bg-white p-4 shadow-sm transition-all duration-200 hover:-translate-y-0.5 hover:shadow-lg"
  >
    <div class="flex items-start justify-between gap-4">
      <div class="flex-1">
        <div class="mb-2 flex flex-wrap items-center gap-2">
          <Tag
            v-for="tag in question.tags || []"
            :key="tag"
            :value="tag"
            severity="info"
            rounded
          />
        </div>

        <h3 class="mb-2 line-clamp-2 text-lg font-semibold text-gray-800">{{ question.title }}</h3>

        <p class="mb-3 line-clamp-2 text-sm text-gray-500">{{ question.content }}</p>

        <div class="flex flex-wrap items-center justify-between gap-2 text-sm text-gray-500">
          <span>提问者: {{ question.author }}</span>
          <span>{{ question.date }}</span>
          <div class="flex items-center gap-2">
            <i class="pi pi-comments text-gray-400"></i>
            <Badge :value="question.answerCount" severity="info" />
            <span>回答</span>
          </div>
        </div>
      </div>

      <Tag v-if="question.isAnswered" value="已解决" severity="success" rounded class="shrink-0" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Tag from 'primevue/tag'
import Badge from 'primevue/badge'
import type { Question } from '~/types/questionAnswer'

const props = defineProps<{
  question: Question
}>()
</script>
