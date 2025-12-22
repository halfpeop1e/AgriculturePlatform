<template>
  <div class="bg-white rounded-md p-3 shadow-sm flex flex-col gap-2">
    <div class="flex items-center gap-3">
      <Avatar :image="reply.author.avatar || defaultAvatar" shape="circle" class="w-9 h-9" />
      <div class="flex flex-col">
        <div class="font-medium text-sm">{{ reply.author.nickname }}</div>
        <div class="text-xs text-gray-500">{{ formattedTime }}</div>
      </div>
    </div>

    <div>
      <p class="text-sm text-gray-700 whitespace-pre-wrap">{{ reply.content }}</p>
    </div>

    <div class="flex items-center">
      <div class="flex-1"></div>
      <Button class="p-button-text" icon="pi pi-thumbs-up" @click.stop="toggleLike" />
      <span class="ml-2 text-sm text-gray-600">{{ localLikes }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Avatar from 'primevue/avatar'
import Button from 'primevue/button'
import type { Reply } from '~/types/post'

const props = defineProps<{ reply: Reply }>()
const emits = defineEmits<{
  (e: 'like', payload: { id: string; liked: boolean }): void
}>()

const defaultAvatar = '/ioanImage/default-avatar.png'

const localLikes = ref(props.reply.likes ?? 0)
const localLiked = ref(!!props.reply.liked)

const formattedTime = computed(() => {
  if (!props.reply.createdAt) return ''
  try {
    const d = new Date(props.reply.createdAt)
    return d.toLocaleString()
  } catch (e) {
    return props.reply.createdAt
  }
})

const toggleLike = () => {
  localLiked.value = !localLiked.value
  localLikes.value += localLiked.value ? 1 : -1
  emits('like', { id: props.reply.id, liked: localLiked.value })
}
</script>

<style scoped>
.p-avatar { width: 36px; height: 36px }
</style>
