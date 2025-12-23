<template>
	<div class="p-6 min-h-screen bg-gray-50">
		<div class="max-w-3xl mx-auto space-y-4">
			<div class="flex items-center justify-between">
				<h2 class="text-2xl font-semibold text-gray-800">全部帖子</h2>
				<Button icon="pi pi-refresh" class="p-button-text" @click="reload" :loading="loading" />
			</div>

			<div v-if="error" class="p-3 rounded-md bg-red-50 text-red-600 text-sm">
				{{ error }}
			</div>

			<div v-if="loading" class="flex justify-center py-10">
				<ProgressSpinner style="width:50px;height:50px" strokeWidth="4" />
			</div>

			<div v-else-if="posts.length === 0" class="text-center py-12 text-gray-500">
				暂无帖子，点击右下角按钮发布第一条吧。
			</div>

			<div v-else class="grid gap-4">
				<PostCard
					v-for="post in posts"
					:key="post.id"
					:post="post"
					@open="openPost"
					@like="handleLike"
				/>
			</div>
		</div>

		<postDialog v-model="dialogVisible" @submit="submitPost" />

		<Button
			icon="pi pi-pencil"
			class="p-button-rounded p-button-lg p-button-raised shadow-lg !w-14 !h-14 fixed bottom-6 left-12"
			@click="dialogVisible = true"
		/>
	</div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import ProgressSpinner from 'primevue/progressspinner'
import type { CreatePostInput } from '~/composables/usePost'
import { usePost } from '~/composables/usePost'
import postDialog from '~/components/Dialog/postDialog.vue'
definePageMeta({
  layout:'home-page-layout'
})
const router = useRouter()

const { posts, loading, error, fetchPosts, createPost, togglePostLike } = usePost()

const dialogVisible = ref(false)

onMounted(() => {
	if (!posts.value.length) {
		fetchPosts().catch(() => {})
	}
})

const reload = () => {
	fetchPosts().catch(() => {})
}

const openPost = (id: string) => {
	router.push(`/allPost/${id}`)
}

const submitPost = async (payload: CreatePostInput) => {
	await createPost(payload)
}

const handleLike = async ({ id, liked }: { id: string; liked: boolean }) => {
	await togglePostLike(id, liked)
}

</script>

<style scoped>
.p-progress-spinner { display: flex; }
</style>
