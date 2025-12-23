<template>
  <div class="p-6">
    <div class="max-w-6xl mx-auto ">
      <div style="margin-bottom:20px ;">
        <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="text-2xl font-bold">农业知识库</h1>
          <p class="text-sm text-gray-500">实用的种植、养殖和田间管理技巧</p>
        </div>

        <div class="flex items-center">
          <el-input v-model="q" clearable placeholder="搜索文章标题或关键字" class="w-64" />
          <el-button type="primary">搜索</el-button>
        </div>
      </div>

      <div class="mb-6">
        <div class="flex items-center justify-between mb-2">
          <div class="text-sm text-gray-500">按标签筛选</div>
          <el-button v-if="selectedTags.size" @click="clearTags" class="text-sm text-gray-400">清除</el-button>
        </div>

        <div class="flex gap-3 flex-wrap">
          <el-button
            v-for="tag in allTags"
            :key="tag"
            @click="toggleTag(tag)"
            :type="selectedTags.has(tag) ? 'primary' : 'default'"
            :plain="!selectedTags.has(tag)"
            round
            size="default"
            class="px-3 text-sm"
          >
            {{ tag }}
          </el-button>
        </div>
      </div>
      </div>

      <div class="grid grid-cols-1 gap-6">
        <KnowledgeCard
          v-for="a in filtered"
          :key="a.id"
          :article="a"
          @click="openArticle(a.id,a.title,a.author,a.date,a.content)"
        />
      </div>

      <div v-if="filtered.length === 0" class="text-center text-gray-500 mt-8">没有匹配的文章</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import KnowledgeCard from '~/components/Card/knowledgeCard.vue'
import { ARTICLES } from '~/utils/knowledgeDataStore'
import { ref, computed } from 'vue'
definePageMeta({ layout: 'home-page-layout' })
const router = useRouter()
const q = ref('')
const selectedTags = ref(new Set<string>())

const allTags = computed(() => {
  const s = new Set<string>()
  ARTICLES.forEach((a) => (a.tags || []).forEach((t) => s.add(t)))
  return Array.from(s)
})

const toggleTag = (tag: string) => {
  if (selectedTags.value.has(tag)) selectedTags.value.delete(tag)
  else selectedTags.value.add(tag)
}

const filtered = computed(() => {
  return ARTICLES.filter((a) => {
    const matchQ = q.value.trim() === '' || a.title.includes(q.value) || a.excerpt.includes(q.value) || (a.tags || []).some((t) => t.includes(q.value))
    const matchTag = selectedTags.value.size === 0 || (a.tags || []).some((t) => selectedTags.value.has(t))
    return matchQ && matchTag
  })
})

function openArticle(id: string,title: string,author:string,date:string, content: string) {
 router.push({
    path: `/knowledge/${id}`,
    state: {
      articleData: {
        id,
        title,
        author,
        date,
        content,
      }
    }
  })
}

function clearTags() {
  selectedTags.value = new Set<string>()
}
</script>

<style scoped>
/* small layout polish */
.px-3 { padding-left: .75rem; padding-right: .75rem }
.py-2 { padding-top: .5rem; padding-bottom: .5rem }
</style>
