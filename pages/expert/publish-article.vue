<template>
  <div class="min-h-screen bg-slate-50 py-10">
    <div class="mx-auto w-full max-w-4xl px-4">
      <Card class="shadow-lg">
        <template #title>
          <div class="flex items-center justify-between">
            <div>
              <h1 class="text-2xl font-semibold text-slate-900">发布农业知识文章</h1>
              <p class="mt-1 text-sm text-slate-500">
                分享您的专业知识，帮助更多农户
              </p>
            </div>
          </div>
        </template>

        <div class="space-y-6">
          <div>
            <label for="title" class="block text-sm font-medium text-gray-700 mb-1">文章标题</label>
            <InputText id="title" v-model="form.title" placeholder="请输入文章标题" class="w-full" />
            <p v-if="errors.title" class="mt-1 text-xs text-red-500">{{ errors.title }}</p>
          </div>

          <div>
            <label for="tags" class="block text-sm font-medium text-gray-700 mb-1">文章标签</label>
            <MultiSelect
              id="tags"
              v-model="form.tags"
              :options="defaultTags"
              display="chip"
              optionLabel="label"
              optionValue="value"
              placeholder="请选择或输入标签"
              :maxSelectedLabels="5"
              filter
              class="w-full"
              @create="handleTagCreate"
              :allowCreate="true"
            />
            <div class="text-xs text-gray-500 mt-1">
              提示：选择合适的标签有助于文章被更多人发现，最多选择5个标签。
            </div>
            <p v-if="errors.tags" class="mt-1 text-xs text-red-500">{{ errors.tags }}</p>
          </div>

          <div>
            <label for="content" class="block text-sm font-medium text-gray-700 mb-1">文章内容</label>
            <Textarea
              id="content"
              v-model="form.content"
              rows="15"
              auto-resize
              placeholder="请在此处撰写您的文章内容..."
              class="w-full"
            />
            <p v-if="errors.content" class="mt-1 text-xs text-red-500">{{ errors.content }}</p>
          </div>

          <div class="flex justify-end gap-3">
            <Button label="取消" class="p-button-text" @click="onCancel" />
            <Button label="发布文章" icon="pi pi-send" :loading="submitting" @click="onSubmit" />
          </div>
        </div>
      </Card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import MultiSelect from 'primevue/multiselect'
import Button from 'primevue/button'
import { useUserStore } from '~/utils/userStore'
import { ElMessage } from 'element-plus'
import type { KnowledgeArticle } from '~/types/knowledgeArticle'

definePageMeta({
  layout: 'home-page-layout'
})

const router = useRouter()
const userStore = useUserStore()

const form = reactive<Omit<KnowledgeArticle, 'id' | 'date' | 'author' | 'excerpt'>>({
  title: '',
  content: '',
  tags: [],
})

const rawTags = ref([
  '种植技术',
  '养殖技术',
  '病虫害防治',
  '土壤改良',
  '施肥技巧',
  '灌溉技术',
  '品种选择',
  '气象灾害',
  '农业政策',
  '市场分析',
  '农产品加工',
  '智慧农业'
])

const defaultTags = computed(() => rawTags.value.map((tag) => ({ label: tag, value: tag })))

const errors = reactive({ title: '', tags: '', content: '' })
const submitting = ref(false)

const handleTagCreate = (event: { value: string }) => {
  const newTag = event.value.trim()
  if (!newTag) return
  if (!rawTags.value.includes(newTag)) {
    rawTags.value.push(newTag)
  }
  form.tags = [...new Set([...(form.tags || []), newTag])].slice(0, 5)
}

const resetForm = () => {
  form.title = '',
  form.content = '',
  form.tags = []
  errors.title = '',
  errors.tags = '',
  errors.content = ''
}

const validate = () => {
  errors.title = '',
  errors.tags = '',
  errors.content = ''
  let valid = true
  if (!form.title || form.title.length < 5 || form.title.length > 100) {
    errors.title = '文章标题长度需在 5 到 100 个字符之间'
    valid = false
  }
  if (!form.tags || !form.tags.length) {
    errors.tags = '请至少选择一个标签'
    valid = false
  } else if (form.tags.length > 5) {
    errors.tags = '最多选择 5 个标签'
    valid = false
  }
  if (!form.content || form.content.length < 50) {
    errors.content = '文章内容请至少 50 个字符'
    valid = false
  }
  return valid
}

const onSubmit = async () => {
  if (!validate()) return

  submitting.value = true
  try {
    const response = await $fetch('/api/knowledge/publish', {
      method: 'POST',
      body: {
      title: form.title,
      content: form.content,
      tags: form.tags,
      authorId: userStore.userId, // 假设文章作者为当前登录用户
    })

    if ((response as { status: number }).status === 200) {
      ElMessage.success('文章发布成功！')
      resetForm()
      router.push('/knowledge') // 发布成功后跳转到知识库列表页
    } else {
      ElMessage.error('文章发布失败，请稍后再试。')
    }

  } catch (error) {
    console.error('文章发布失败', error)
    ElMessage.error('文章发布失败，请稍后再试。')
  } finally {
    submitting.value = false
  }
}

const onCancel = () => {
  resetForm()
  router.back()
}
</script>

<style scoped>
/* 可以添加一些 scoped 样式 */
</style>
