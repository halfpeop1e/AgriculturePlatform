<template>
  <div
    class="bg-white rounded-lg border border-gray-100 shadow-sm hover:shadow-lg transform hover:-translate-y-0.5 transition-all duration-200 overflow-hidden cursor-pointer"
    @click="$emit('open', article.id)"
  >
    <div class="flex flex-col md:flex-row">
      <!-- Image on the left for md+ screens, top for small screens -->
      <div class="flex-shrink-0 bg-gray-100 overflow-hidden" style="width: 300px;">
        <div class="w-full h-44 md:h-full">
          <el-image
            v-if="article.image"
            :src="article.image"
            alt="cover"
            class="w-full h-full object-cover"
            fit="cover"
          >
            <template #error>
              <div class="w-full h-full flex items-center justify-center bg-gray-100">
                <el-icon><IconPicture /></el-icon>
              </div>
            </template>
          </el-image>
        </div>
      </div>
      <!-- Text content -->
      <div class="p-4 flex-1">
        <div class="flex items-center gap-2 mb-3 flex-wrap">
          <span
            v-for="tag in article.tags || []"
            :key="tag"
            class="inline-flex items-center text-xs px-4 py-0.75 rounded-lg border border-blue-600 text-black bg-white hover:bg-green-50 hover:border-green-200"
          >
            {{ tag }}
          </span>
        </div>

        <h3 class="text-lg font-semibold text-gray-800 mb-2 line-clamp-2">{{ article.title }}</h3>
        <p class="text-sm text-gray-500 mb-4 line-clamp-3">{{ article.excerpt }}</p>

        <div class="flex items-center justify-between text-sm text-gray-400">
          <div class="truncate">{{ article.author }}</div>
          <div>{{ article.date }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { KnowledgeArticle } from '~/types/knowledgeArticle'
import { Picture as IconPicture } from '@element-plus/icons-vue'
const props = defineProps<{ article: KnowledgeArticle }>()
const emit = defineEmits<{
  (e: 'open', id: string): void
}>()
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
