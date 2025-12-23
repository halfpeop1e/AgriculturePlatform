<!-- 
 这个页面用于展示专家收到的问题 -->
<template>
  <div class="space-y-8">
    <section class="flex flex-wrap items-center justify-between gap-4 rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <div>
        <h1 class="text-2xl font-semibold text-slate-900">专家问答</h1>
        <p class="text-sm text-slate-500">解决您在农业生产中的疑难问题</p>
      </div>
    </section>

    <section class="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <div class="flex flex-wrap items-center justify-between gap-2">
        <span class="text-sm text-slate-500">按标签筛选</span>
        <Button
          v-if="selectedTags.size"
          label="清除"
          severity="secondary"
          text
          size="small"
          @click="clearTags"
        />
      </div>
      <div class="mt-4 flex flex-wrap gap-2">
        <Button
          v-for="tag in allTags"
          :key="tag"
          :label="tag"
          :severity="selectedTags.has(tag) ? 'primary' : 'secondary'"
          :outlined="!selectedTags.has(tag)"
          rounded
          size="small"
          @click="toggleTag(tag)"
        />
      </div>
    </section>

    <section class="flex flex-wrap items-center justify-between gap-4 rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <div class="text-sm text-slate-500">搜索问题标题或内容</div>
      <IconField iconPosition="left">
        <InputIcon class="pi pi-search" />
        <InputText v-model="q" placeholder="输入关键词" class="w-72" />
      </IconField>
    </section>

    <div v-if="loading" class="grid grid-cols-1 gap-6">
      <div v-for="i in 3" :key="i" class="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
        <Skeleton width="50%" height="1.5rem" class="mb-4" />
        <Skeleton width="100%" height="4rem" class="mb-4" />
        <div class="flex justify-between">
          <Skeleton width="20%" height="1rem" />
          <Skeleton width="20%" height="1rem" />
        </div>
      </div>
    </div>

    <div v-else class="grid grid-cols-1 gap-6">
      <QuestionCard
        v-for="qItem in filtered"
        :key="qItem.id"
        :question="qItem"
        @click="openQuestion(qItem.id)"
      />
    </div>

    <Message v-if="!loading && !filtered.length" severity="warn" class="justify-center text-sm">
      没有匹配的问题，试试其他筛选条件
    </Message>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Message from 'primevue/message'
import IconField from 'primevue/iconfield'
import InputIcon from 'primevue/inputicon'
import Skeleton from 'primevue/skeleton'
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
    const title = qItem.title.toLowerCase()
    const content = qItem.content.toLowerCase()
    const tags = (qItem.tags || []).map((t) => t.toLowerCase())

    const matchKeyword =
      !keyword ||
      title.includes(keyword) ||
      content.includes(keyword) ||
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
