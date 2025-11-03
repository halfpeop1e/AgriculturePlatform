<template>
  <div class="p-6">
    <div class="max-w-3xl mx-auto bg-white rounded-lg shadow-sm p-6">
      <div v-if="article">
        <h1 class="text-2xl font-bold mb-2">{{ article.title }}</h1>
        <div class="text-sm text-gray-500 mb-4">{{ article.author }} · {{ article.date }}</div>
        <img v-if="article.image" :src="article.image" alt="cover" class="w-full h-64 object-cover rounded-md mb-4" />
        <div class="prose max-w-none" v-html="article.content"></div>
      </div>
      <div v-else class="text-center text-gray-500">未找到相关文章</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { ARTICLES } from '~/utils/knowledgeDataStore'
import type { KnowledgeArticle } from '~/types/knowledgeArticle'
import { computed } from 'vue'

const route = useRoute()
const id = route.params.id as string

const article = computed<KnowledgeArticle | undefined>(() => ARTICLES.find((a) => a.id === id))
</script>

<style scoped>
.prose img { border-radius: 6px }
</style>
