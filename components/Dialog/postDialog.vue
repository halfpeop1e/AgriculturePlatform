<template>
  <Dialog v-model:visible="visible" header="发表帖子" :closable="true" class="w-full max-w-xl">
    <div class="space-y-4">
      <div class="flex items-center gap-3">
        <Avatar :image="author.avatar || defaultAvatar" shape="circle" class="w-10 h-10" />
        <div>
          <div class="font-medium">{{ author.nickname }}</div>
          <div class="text-xs text-gray-500">发布者</div>
        </div>
      </div>

      <div>
        <label class="block text-sm mb-1">标题</label>
        <InputText v-model="title" class="w-full" placeholder="输入标题" />
      </div>

      <div>
        <label class="block text-sm mb-1">内容</label>
        <Textarea v-model="content" rows="6" class="w-full" placeholder="输入内容" auto-resize />
      </div>

      <div class="flex justify-end gap-3">
        <Button label="取消" class="p-button-text" @click="onCancel" />
        <Button label="发布" icon="pi pi-send" :disabled="!canSubmit" @click="onSubmit" />
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
import Avatar from 'primevue/avatar'
import type { Author } from '~/types/post'
import type { CreatePostInput } from '~/composables/usePost'

const props = defineProps<{ modelValue: boolean }>()
const emits = defineEmits<{
  (e: 'update:modelValue', v: boolean): void
  (e: 'submit', payload: CreatePostInput): void
}>()

const visible = computed({
  get: () => props.modelValue,
  set: (v: boolean) => emits('update:modelValue', v)
})

const title = ref('')
const content = ref('')

const userStore = useUserStore()

const author = computed<Author>(() => {
  const info: any = userStore.userinfo || {}
  return {
    id: userStore.userId || String(info.id || `user_${Date.now()}`),
    nickname: info.nickname || info.name || '匿名',
    avatar: userStore.avatar || ''
  }
})

const defaultAvatar = '/ioanImage/default-avatar.png'

const canSubmit = computed(() => title.value.trim().length > 0 && content.value.trim().length > 0)

const resetForm = () => {
  title.value = ''
  content.value = ''
}

const onCancel = () => {
  visible.value = false
  resetForm()
}

const onSubmit = () => {
  if (!canSubmit.value) return

  const payload: CreatePostInput = {
    title: title.value.trim(),
    content: content.value.trim(),
    author: author.value
  }

  emits('submit', payload)
  visible.value = false
  resetForm()
}
</script>

<style scoped>
/* Tailwind 做主要样式，保留少量 PrimeVue 兼容性调整 */
.p-avatar { width: 40px; height: 40px }
</style>
