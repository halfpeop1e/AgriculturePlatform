<template>
  <div class="p-6">
    <div class="max-w-6xl mx-auto">
      <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="text-2xl font-bold">专家问答问题浏览</h1>
          <p class="text-sm text-gray-500">浏览已解决的问题</p>
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
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import QuestionCard from '~/components/Card/QuestionCard.vue'
import { useQuestionDataStore } from '~/utils/questionDataStore'

definePageMeta({ layout: 'expert-backend-layout' })

const router = useRouter()
const questionStore = useQuestionDataStore()
const q = ref('')
const loading = ref(false)
const selectedTags = ref<Set<string>>(new Set())

onMounted(async () => {
  loading.value = true
  try {
    await questionStore.fetchQuestions()
  } finally {
    loading.value = false
  }
})

const allTags = computed(() => {
  const tagSet = new Set<string>()
  questionStore.questions.forEach((qItem) => (qItem.tags || []).forEach((t) => tagSet.add(t)))
  return Array.from(tagSet).sort()
})

const toggleTag = (tag: string) => {
  const updated = new Set(selectedTags.value)
  if (updated.has(tag)) {
    updated.delete(tag)
  } else {
    updated.add(tag)
  }
  selectedTags.value = updated
}

const filtered = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  return questionStore.questions.filter((qItem) => {
    const safeTitle = (qItem.title || '').toLowerCase()
    const safeContent = (qItem.content || '').toLowerCase()
    const tags = (qItem.tags || []).map((t) => t?.toLowerCase?.() || '')

    const matchKeyword =
      !keyword ||
      safeTitle.includes(keyword) ||
      safeContent.includes(keyword) ||
      tags.some((t) => t.includes(keyword))

    const matchTag =
      selectedTags.value.size === 0 ||
      (qItem.tags || []).some((t) => selectedTags.value.has(t))

    return matchKeyword && matchTag
  })
})

const openQuestion = (id: string) => {
  router.push(`/specialBoard/question/${id}`)
}


const clearTags = () => {
  selectedTags.value = new Set()
}
</script>
