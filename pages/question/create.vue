<template>
  <div class="p-6">
    <el-card shadow="hover" class="rounded-2xl border border-[#E5E7EB] bg-white max-w-3xl mx-auto">
      <h1 class="text-2xl font-bold mb-6">提出您的问题</h1>
      
      <el-form ref="questionForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="问题标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入问题标题" />
        </el-form-item>
        
        <el-form-item label="问题标签" prop="tags">
          <el-select
            v-model="form.tags"
            multiple
            placeholder="请选择或输入标签"
            filterable
            allow-create
          >
            <el-option v-for="tag in defaultTags" :key="tag" :label="tag" :value="tag" />
          </el-select>
          <div class="text-xs text-gray-500 mt-1">
            提示：输入标签可帮助专家更快找到您的问题，最多选择5个标签
          </div>
        </el-form-item>
        
        <el-form-item label="问题详情" prop="content">
          <el-input
            type="textarea"
            :rows="8"
            v-model="form.content"
            placeholder="请详细描述您遇到的问题，包括背景、现状和期望解决的问题等"
          />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="onSubmit">提交问题</el-button>
          <el-button @click="onCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { submitQuestion } from '~/composables/useQuestionAnswer'
import { useUserStore } from '~/utils/userStore' // 假设存在用户状态管理

definePageMeta({ layout: 'home-page-layout' })
const router = useRouter()
const userStore = useUserStore() // 获取用户信息

// 表单数据 - 补充authorId字段
const form = ref({
  title: '',
  content: '',
  tags: [] as string[],
  authorId: userStore.userId // 从用户存储中获取当前用户ID
})

// 默认标签
const defaultTags = [
  '种植技术', '养殖技术', '病虫害防治', '土壤改良', 
  '施肥技巧', '灌溉技术', '品种选择', '气象灾害'
]

// 表单验证规则
const rules = ref({
  title: [
    { required: true, message: '请输入问题标题', trigger: 'blur' },
    { min: 5, max: 100, message: '标题长度在5到100个字符之间', trigger: 'blur' }
  ],
  tags: [
    { 
      validator: (rule: any, value: any, callback: any) => {
        if (value.length === 0) {
          callback(new Error('请至少选择一个标签'))
        } else if (value.length > 5) {
          callback(new Error('最多选择5个标签'))
        } else {
          callback()
        }
      }, 
      trigger: 'change' 
    }
  ],
  content: [
    { required: true, message: '请输入问题详情', trigger: 'blur' },
    { min: 20, message: '问题描述请至少20个字符', trigger: 'blur' }
  ]
})

// 表单引用
const questionForm = ref<any>(null)

// 提交表单
async function onSubmit() {
  questionForm.value.validate(async (valid: boolean) => {
    if (valid) {
      // 确保authorId已正确设置
      if (!form.value.authorId) {
        ElMessage.error('用户信息获取失败，请重新登录')
        return
      }
      
      const result = await submitQuestion(form.value)
      if (result) {
        router.push('/question')
      }
    }
  })
}

// 取消
function onCancel() {
  router.back()
}
</script>