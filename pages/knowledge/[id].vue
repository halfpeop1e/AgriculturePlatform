<template>
   <div class="p-6">
    <el-card shadow="hover" class="rounded-2xl border border-[#E5E7EB] bg-white hover:shadow-md transition">
    <div class="max-w-3xl mx-auto bg-white rounded-lg shadow-sm p-6">
      <div v-if="articelContent">
        <h1 class="text-2xl font-bold mb-2">{{ articelContent.title }}</h1>
        <div class="text-sm text-gray-500 mb-4">{{ articelContent.author }} · {{ articelContent.date }}</div>
        <el-divider />
        <div class="prose max-w-none" v-html="articelContent.content"/>
      </div>
      <div v-else class="text-center text-gray-500">未找到相关文章</div>
      </div>
  </el-card>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import type { KnowledgeArticleContent } from '~/types/knowledgeArticle'
import { ref, onMounted } from 'vue'
definePageMeta({ layout: 'home-page-layout' })
const route = useRoute()
const articelContent = ref<KnowledgeArticleContent>({
  id: '',
  title: '',
  author: '',
  date: '',
  content: ''
})
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
