<template>
  <div class="p-6">
    <div class="max-w-6xl mx-auto">
      <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="text-2xl font-bold">专家问答</h1>
          <p class="text-sm text-gray-500">解决您在农业生产中的疑难问题</p>
        </div>
        <el-button type="primary" @click="navigateToCreate">提问</el-button>
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

      <!-- 搜索框添加 -->
      <div class="mb-6">
        <el-input v-model="q" clearable placeholder="搜索问题标题或内容" class="w-64" />
      </div>

      <div class="grid grid-cols-1 gap-6">
        <QuestionCard
          v-for="qItem in filtered"
          :key="qItem.id"
          :question="qItem"
          @click="openQuestion(qItem.id)"
        />
      </div>

      <div v-if="filtered.length === 0" class="text-center text-gray-500 mt-8">没有匹配的问题</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import QuestionCard from '~/components/Card/QuestionCard.vue'
import { useQuestionDataStore } from '~/utils/questionDataStore'
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

definePageMeta({ layout: 'home-page-layout' })
const router = useRouter()
const questionStore = useQuestionDataStore()
const q = ref('') // 搜索关键词
const selectedTags = ref(new Set<string>())

onMounted(() => {
  questionStore.fetchQuestions()
})

const allTags = computed(() => {
  const s = new Set<string>()
  questionStore.questions.forEach((qItem) => (qItem.tags || []).forEach((t) => s.add(t)))
  return Array.from(s)
})

const toggleTag = (tag: string) => {
  if (selectedTags.value.has(tag)) {
    selectedTags.value.delete(tag)
  } else {
    selectedTags.value.add(tag)
  }
}

const filtered = computed(() => {
  return questionStore.questions.filter((qItem) => {
    const matchQ = q.value.trim() === '' || 
                  qItem.title.includes(q.value) || 
                  qItem.content.includes(q.value) || 
                  (qItem.tags || []).some((t) => t.includes(q.value))
    const matchTag = selectedTags.value.size === 0 || 
                    (qItem.tags || []).some((t) => selectedTags.value.has(t))
    return matchQ && matchTag
  })
})

function openQuestion(id: string) {
  router.push(`/question/${id}`)
}

function navigateToCreate() {
  router.push('/question/create')
}

function clearTags() {
  selectedTags.value = new Set<string>()
}
</script>

<style scoped>
.px-3 { padding-left: .75rem; padding-right: .75rem }
</style>
