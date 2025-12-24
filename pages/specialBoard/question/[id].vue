<template>
  <div class="space-y-8">
    <Card v-if="questionDetail" class="border border-slate-200 shadow-sm">
      <template #content>
        <section class="mb-8 space-y-4">
          <h1 class="text-2xl font-semibold text-slate-900">{{ questionDetail.title }}</h1>
          <div class="flex flex-wrap items-center gap-3 text-sm text-slate-500">
            <span>{{ questionDetail.author }} · {{ questionDetail.date }}</span>
            <Tag
              v-for="tag in questionDetail.tags"
              :key="tag"
              :value="tag"
              severity="info"
              rounded
            />
          </div>
          <div class="prose max-w-none text-slate-700" v-html="questionDetail.content" />
        </section>

        <Divider />

        <section class="space-y-6">
          <div>
            <h2 class="text-xl font-semibold text-slate-900">回答 ({{ questionDetail.answerCount }})</h2>
            <Message v-if="!questionDetail.answers.length" severity="info" class="mt-4 justify-center text-sm">
              暂无回答，快来成为第一个回答的专家吧！
            </Message>
          </div>

          <div v-if="questionDetail.answers.length" class="space-y-6">
            <AnswerItem v-for="(answer, index) in questionDetail.answers" :key="index" :answer="answer" />
          </div>
        </section>

        <section class="mt-8 space-y-4">
          <h3 class="text-lg font-semibold text-slate-900">我来回答</h3>
          <div class="rounded-xl border border-slate-200 bg-slate-50 p-4">
            <Textarea v-model="answerContent" rows="6" autoResize placeholder="请输入您的回答..." class="w-full" />
            <div class="mt-4 flex justify-end">
              <Button label="提交回答" icon="pi pi-send" :loading="submitting" @click="handleSubmitAnswer" />
            </div>
          </div>
        </section>
      </template>
    </Card>

    <Card v-else class="border border-dashed border-slate-200 bg-white text-center text-slate-500">
      <template #content>
        未找到相关问题
      </template>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import Card from 'primevue/card'
import Divider from 'primevue/divider'
import Tag from 'primevue/tag'
import Message from 'primevue/message'
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import AnswerItem from '~/components/AnswerItem.vue'
import { useQuestionDataStore } from '~/utils/questionDataStore'
import { submitAnswer as submitAnswerApi } from '~/composables/useQuestionAnswer'

definePageMeta({ layout: 'expert-backend-layout' })

const route = useRoute()
const toast = useToast()
const questionStore = useQuestionDataStore()
const answerContent = ref('')
const submitting = ref(false)
const id = route.params.id as string

onMounted(() => {
  questionStore.fetchQuestionDetail(id)
  console.log(questionStore.currentQuestion)
})

watch(
  () => questionStore.currentQuestion,
  (newVal) => {
    if (!newVal && id) {
      questionStore.fetchQuestionDetail(id)
    }
  }
)

const questionDetail = computed(() => {
  const detail = questionStore.currentQuestion
  console.log('Computed question detail:', detail)
  if (!detail) return null
  return {
    ...detail,
    tags: detail.tags ?? [],
    answers: detail.answer ?? [],
    answerCount:
      detail.answerCount ??
      (detail.answer ? detail.answer.length : 0)
  }
})

const handleSubmitAnswer = async () => {
  const content = answerContent.value.trim()
  if (!content) {
    toast.add({ severity: 'warn', summary: '提示', detail: '回答内容不能为空', life: 3000 })
    return
  }

  try {
    submitting.value = true
    await submitAnswerApi(id, content)
    toast.add({ severity: 'success', summary: '提交成功', detail: '回答已提交', life: 3000 })
    answerContent.value = ''
    questionStore.fetchQuestionDetail(id)
  } catch (error) {
    console.error('提交回答失败', error)
    toast.add({ severity: 'error', summary: '提交失败', detail: '请稍后重试', life: 4000 })
  } finally {
    submitting.value = false
  }
}

onUnmounted(() => {
  questionStore.clearCurrentQuestion()
})
</script>
