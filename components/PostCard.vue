<template>
  <div class="bg-white rounded-lg p-4 shadow flex flex-col gap-3" @click.self="openPost">
    <div class="flex items-center gap-3">
      <Avatar :image="post.author.avatar || defaultAvatar" shape="circle" class="w-10 h-10" />
      <div class="flex flex-col">
        <div class="font-medium text-sm">{{ post.author.nickname }}</div>
        <div class="text-xs text-gray-500">{{ formattedTime }}</div>
      </div>
    </div>

    <div class="cursor-pointer" @click.stop="openPost">
      <h4 class="text-base font-semibold mb-1">{{ post.title }}</h4>
      <p class="text-sm text-gray-700 whitespace-pre-wrap">{{ post.content }}</p>
    </div>

    <div class="flex items-center">
      <Button label="回贴" class="p-button-text text-blue-600" @click.stop="openPost" />
      <div class="flex-1"></div>
      <Button
        class="p-button-text"
        @click.stop="toggleLike"
        :aria-pressed="localLiked"
        :icon="localLiked ? 'pi pi-thumbs-up-fill' : 'pi pi-thumbs-up'"
      />
      <span class="ml-2 text-sm text-gray-600">{{ localLikes }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Avatar from 'primevue/avatar'
import Button from 'primevue/button'
import type { Post } from '~/types/post'

const props = defineProps<{ post: Post }>()
const emits = defineEmits<{
  (e: 'open', id: string): void
  (e: 'like', payload: { id: string; liked: boolean }): void
}>()

const defaultAvatar = '/ioanImage/default-avatar.png'

const localLikes = ref(props.post.likes ?? 0)
const localLiked = ref(!!props.post.liked)

const formattedTime = computed(() => {
  if (!props.post.createdAt) return ''
  try {
    const d = new Date(props.post.createdAt)
    return d.toLocaleString()
  } catch (e) {
    return props.post.createdAt
  }
})

const openPost = () => {
  emits('open', props.post.id)
}

const toggleLike = () => {
  localLiked.value = !localLiked.value
  localLikes.value += localLiked.value ? 1 : -1
  emits('like', { id: props.post.id, liked: localLiked.value })
}
</script>

<style scoped>
/* 使用 Tailwind 处理大部分样式；保留少量本地样式以兼容 PrimeVue Avatar sizing */
.p-avatar { width: 40px; height: 40px; }
</style>
