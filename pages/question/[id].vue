<template>
  <div class="p-6">
    <el-card shadow="hover" class="rounded-2xl border border-[#E5E7EB] bg-white hover:shadow-md transition max-w-4xl mx-auto">
      <div v-if="questionDetail">
        <!-- 问题部分 -->
        <div class="mb-8">
          <h1 class="text-2xl font-bold mb-4">{{ questionDetail.title }}</h1>
          <div class="text-sm text-gray-500 mb-4">
            {{ questionDetail.author }} · {{ questionDetail.date }}
            <el-tag v-for="tag in questionDetail.tags" :key="tag" class="ml-2">{{ tag }}</el-tag>
          </div>
          <div class="prose max-w-none mb-6" v-html="questionDetail.content" />
        </div>

        <el-divider />

        <!-- 回答列表 -->
        <div class="mt-6">
          <h2 class="text-xl font-semibold mb-4">回答 ({{ questionDetail.answerCount }})</h2>
          
          <div v-if="questionDetail.answers.length === 0" class="text-center text-gray-500 py-8">
            暂无回答，快来成为第一个回答的专家吧！
          </div>
          
          <div class="space-y-6">
            <AnswerItem 
              v-for="answer in questionDetail.answers" 
              :key="answer.id" 
              :answer="answer" 
            />
          </div>
        </div>

        <!-- 回答表单 -->
        <div class="mt-8">
          <h3 class="text-lg font-semibold mb-4">我来回答</h3>
          <el-form>
            <el-form-item>
              <el-input
                type="textarea"
                :rows="6"
                placeholder="请输入您的回答..."
                v-model="answerContent"
              />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSubmitAnswer">提交回答</el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>
      
      <div v-else class="text-center text-gray-500 py-8">
        未找到相关问题
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { useQuestionDataStore } from '~/utils/questionDataStore'
import { ref, onMounted, watch, computed } from 'vue'
import AnswerItem from '~/components/AnswerItem.vue'
import { submitAnswer as submitAnswerApi } from '~/composables/useQuestionAnswer'
import { ElMessage } from 'element-plus'  // 补充导入ElMessage

definePageMeta({ layout: 'home-page-layout' })
const route = useRoute()
const router = useRouter()
const questionStore = useQuestionDataStore()
const answerContent = ref('')
const id = route.params.id as string  // 确保id类型正确（如果后端需要number可改为Number(route.params.id)）

onMounted(() => {
  questionStore.fetchQuestionDetail(id)
})

watch(
  () => questionStore.currentQuestion,
  (newVal) => {
    if (!newVal && id) {
      questionStore.fetchQuestionDetail(id)
    }
  }
)

const questionDetail = computed(() => questionStore.currentQuestion)

async function handleSubmitAnswer() {
  if (!answerContent.value.trim()) {
    ElMessage.warning('回答内容不能为空')
    return
  }
  
  // 调用API提交回答（不判断返回值是否为true，因为原函数返回void）
  await submitAnswerApi(id, answerContent.value)
  answerContent.value = ''
  // 重新获取问题详情以更新回答列表
  questionStore.fetchQuestionDetail(id)
}

// 页面离开时清除当前问题数据
onUnmounted(() => {
  questionStore.clearCurrentQuestion()
})
</script>