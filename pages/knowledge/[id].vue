<template>
   <div class="p-6">
    <div class="max-w-3xl mx-auto bg-white rounded-lg shadow-sm p-6">
      <div v-if="articelContent">
        <h1 class="text-2xl font-bold mb-2">{{ articelContent.title }}</h1>
        <div class="text-sm text-gray-500 mb-4">{{ articelContent.author }} · {{ articelContent.date }}</div>
        <div class="prose max-w-none" v-html="articelContent.content"></div>
      </div>
      <div v-else class="text-center text-gray-500">未找到相关文章</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import type { KnowledgeArticleContent } from '~/types/knowledgeArticle'
import { ref, onMounted } from 'vue'

const route = useRoute()
const articelContent = ref<KnowledgeArticleContent>()
const id = route.params.id as string

onMounted(() => {
  if (import.meta.client) {
    const stateData = history.state
    if (stateData && stateData.articleData) {
      articelContent.value = stateData.articleData
      console.log('接收到的数据:', articelContent.value)
    }
  }
})
</script>

<style scoped>
.prose img { border-radius: 6px }
</style>
