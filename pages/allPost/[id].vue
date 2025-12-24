<template>
	<div class="p-6 min-h-screen bg-gray-50">
		<div class="max-w-3xl mx-auto space-y-4">
			<div class="flex items-center justify-between">
				<Button label="返回" icon="pi pi-arrow-left" class="p-button-text" @click="goBack" />
				<Button label="刷新" icon="pi pi-refresh" class="p-button-text" @click="reload" :loading="loading" />
			</div>

			<div v-if="!post && loading" class="flex justify-center py-12">
				<ProgressSpinner style="width:50px;height:50px" strokeWidth="4" />
			</div>

			<div v-else-if="!post && error" class="p-3 rounded-md bg-red-50 text-red-600 text-sm">
				{{ error }}
			</div>

			<PostCard
				v-else-if="post"
				:post="post"
				@open="noop"
				@like="handlePostLike"
			/>

			<div v-if="post" class="bg-white rounded-lg shadow p-4 space-y-4">
				<h3 class="text-lg font-semibold text-gray-800">全部回复</h3>

				<div v-if="post.replies && post.replies.length" class="space-y-3">
					<AnswerCard
						v-for="reply in post.replies"
						:key="reply.id"
						:reply="reply"
						@like="handleReplyLike"
					/>
				</div>
				<div v-else class="text-sm text-gray-500">暂时还没有回复。</div>
			</div>

			<div class="bg-white rounded-lg shadow p-4 space-y-3">
				<h3 class="text-lg font-semibold text-gray-800">我要回复</h3>
				<Textarea
					v-model="replyContent"
					rows="5"
					class="w-full"
					auto-resize
					placeholder="输入回复内容..."
				/>
				<div class="flex justify-end gap-3">
					<Button label="清空" class="p-button-text" @click="replyContent = ''" />
					<Button
						label="发布回复"
						icon="pi pi-send"
						:disabled="!canReply || submitting"
						:loading="submitting"
						@click="submitReply"
					/>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import ProgressSpinner from 'primevue/progressspinner'
import type { Author } from '~/types/post'
import { usePost } from '~/composables/usePost'
import PostCard from '~/components/Card/PostCard.vue'
import AnswerCard from '~/components/Card/answerCard.vue'
definePageMeta({
  layout:'home-page-layout'
})
const route = useRoute()
const router = useRouter()

const { posts, loading, error, fetchPostById, getPostById, addReply, togglePostLike, toggleReplyLike } = usePost()

const postId = computed(() => String(route.params.id))
const post = computed(() => getPostById(postId.value))

const replyContent = ref('')
const submitting = ref(false)

const userStore = useUserStore()

const author = computed<Author>(() => {
	const info: any = userStore.userinfo || {}
	return {
		id: userStore.userId || String(info.id || `user_${Date.now()}`),
		nickname: info.nickname || info.name || '匿名',
		avatar: userStore.avatar || ''
	}
})

const canReply = computed(() => replyContent.value.trim().length > 0)

const ensurePost = async () => {
	if (!post.value) {
		try {
			await fetchPostById(postId.value)
      console.log('获取帖子成功')
		} catch (e) {
			console.error('获取帖子失败', e)
		}
	}
}
const refreshData = async () => {
  await fetchPostById(postId.value)
}
watch(() => postId.value, () => {
	replyContent.value = ''
	ensurePost()
})

const reload = async () => {
	await refreshData()
}

const goBack = () => {
	router.back()
}

const noop = () => {}

const handlePostLike = async ({ id, liked }: { id: string; liked: boolean }) => {
	await togglePostLike(id, liked)
}

const handleReplyLike = async ({ id, liked }: { id: string; liked: boolean }) => {
	await toggleReplyLike(postId.value, id, liked)
}

const submitReply = async () => {
	if (!canReply.value || !post.value) return
	submitting.value = true
	try {
		await addReply(postId.value, {
			content: replyContent.value.trim(),
			author: author.value
		})
		replyContent.value = ''
    submitting.value = false
    console.log('状态回复',submitting.value)
    await refreshData()
	}
  catch(err){
    console.log(err)
  } finally {
		submitting.value = false
    await refreshData()
	}
}

onMounted(async () => {
	await ensurePost()
})

</script>
