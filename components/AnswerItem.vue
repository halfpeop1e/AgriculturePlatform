<template>
  <div class="bg-gray-50 rounded-lg p-4">
    <div class="flex items-start justify-between mb-3">
      <div>
        <span class="font-medium">{{ props.answer }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Answer } from '~/types/questionAnswer'


const props = defineProps<{
  answer: string
}>()
</script>
