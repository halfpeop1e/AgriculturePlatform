<template>
  <Dialog
    v-model:visible="visible"
    header="提出您的问题"
    modal
    class="w-full max-w-3xl"
  >
    <div class="space-y-6">
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">问题标题</label>
          <InputText v-model="form.title" placeholder="请输入问题标题" class="w-full" />
          <p v-if="errors.title" class="mt-1 text-xs text-red-500">{{ errors.title }}</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">问题标签</label>
          <MultiSelect
            v-model="form.tags"
            :options="defaultTags"
            display="chip"
            optionLabel="label"
            optionValue="value"
            placeholder="请选择或输入标签"
            :maxSelectedLabels="3"
            filter
            class="w-full"
            @create="handleTagCreate"
            :allowCreate="true"
          />
          <div class="text-xs text-gray-500 mt-1">
            提示：输入标签可帮助专家更快找到您的问题，最多选择5个标签。
          </div>
          <p v-if="errors.tags" class="mt-1 text-xs text-red-500">{{ errors.tags }}</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">问题详情</label>
          <Textarea
            v-model="form.content"
            rows="6"
            auto-resize
            placeholder="请详细描述您遇到的问题，包括背景、现状和期望解决的问题等"
            class="w-full"
          />
          <p v-if="errors.content" class="mt-1 text-xs text-red-500">{{ errors.content }}</p>
        </div>
      </div>

      <div class="flex justify-end gap-3">
        <Button label="取消" class="p-button-text" @click="onCancel" />
        <Button label="提交问题" icon="pi pi-send" :loading="submitting" @click="onSubmit" />
      </div>
    </div>
  </Dialog>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Dialog from 'primevue/dialog'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import MultiSelect from 'primevue/multiselect'
import { submitQuestion } from '~/composables/useQuestionAnswer'
import { useUserStore } from '~/utils/userStore'

const props = defineProps<{ modelValue: boolean;expertId:string }>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'submitted'): void
}>()

const visible = computed({
  get: () => props.modelValue,
  set: (value: boolean) => emit('update:modelValue', value)
})

const userStore = useUserStore()

const form = ref({
  title: '',
  content: '',
  tags: [] as string[],
  authorId: userStore.userId
})

const rawTags = [
  '种植技术',
  '养殖技术',
  '病虫害防治',
  '土壤改良',
  '施肥技巧',
  '灌溉技术',
  '品种选择',
  '气象灾害'
]

const defaultTags = rawTags.map((tag) => ({ label: tag, value: tag }))

const errors = ref({ title: '', tags: '', content: '' })
const submitting = ref(false)

const handleTagCreate = (event: { value: string }) => {
  const newTag = event.value.trim()
  if (!newTag) return
  if (!rawTags.includes(newTag)) {
    rawTags.push(newTag)
  }
  form.value.tags = [...new Set([...form.value.tags, newTag])].slice(0, 5)
}

const resetForm = () => {
  form.value = {
    title: '',
    content: '',
    tags: [],
    authorId: userStore.userId
  }
  errors.value = { title: '', tags: '', content: '' }
}

const validate = () => {
  errors.value = { title: '', tags: '', content: '' }
  let valid = true
  if (!form.value.title || form.value.title.length < 5 || form.value.title.length > 100) {
    errors.value.title = '标题长度需在 5 到 100 个字符之间'
    valid = false
  }
  if (!form.value.tags.length) {
    errors.value.tags = '请至少选择一个标签'
    valid = false
  } else if (form.value.tags.length > 5) {
    errors.value.tags = '最多选择 5 个标签'
    valid = false
  }
  if (!form.value.content || form.value.content.length < 20) {
    errors.value.content = '问题描述请至少 20 个字符'
    valid = false
  }
  return valid
}

const onSubmit = async () => {
  if (!validate()) return

  if (!form.value.authorId) {
    emit('submitted')
    visible.value = false
    return
  }

  submitting.value = true
  try {
    const result = await submitQuestion({
      title: form.value.title,
      content: form.value.content,
      tags: form.value.tags,
      authorId: form.value.authorId,
      expertId:props.expertId
    })

    if (result) {
      emit('submitted')
      visible.value = false
      resetForm()
    }
  } finally {
    submitting.value = false
  }
}

const onCancel = () => {
  visible.value = false
  resetForm()
}
</script>
