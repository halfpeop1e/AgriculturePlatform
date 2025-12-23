<template>
  <div class="bg-gray-50 rounded-lg p-4">
    <div class="flex items-start justify-between mb-3">
      <div>
        <span class="font-medium">{{ answer.expert }}</span>
        <span class="text-sm text-gray-500 ml-2">专家</span>
      </div>
      <span class="text-sm text-gray-500">{{ answer.date }}</span>
    </div>
    
    <div class="prose max-w-none mb-3" v-html="answer.content" />
    
    <div class="flex justify-end">
      <Button
        v-if="answer.isAccepted"
        label="已采纳"
        icon="pi pi-check"
        severity="success"
        size="small"
        disabled
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Answer } from '~/types/questionAnswer'
import Button from 'primevue/button'

const props = defineProps<{
  answer: Answer
}>()
</script>
