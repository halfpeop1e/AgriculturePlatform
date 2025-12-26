<template>
	<div class="preorder-page">
		<div class="preorder-page__header">
			<div>
				<h1>预约管理</h1>
				<p>查看农户提交的预约并提前做好准备。</p>
			</div>
			<Button label="刷新" icon="pi pi-refresh" outlined :loading="loading" @click="fetchPreorders" />
		</div>

		<div v-if="loading" class="preorder-page__state">正在同步预约数据...</div>
		<div v-else-if="!upcomingPreorders.length" class="preorder-page__state preorder-page__state--empty">
			暂无新的预约，去知识社区提升曝光吧。
		</div>
		<div v-else class="preorder-page__grid">
			<PreorderCard v-for="item in upcomingPreorders" :key="item.preorderId" :preorder="item" />
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import Button from 'primevue/button'
import PreorderCard from '~/components/Card/preorderCard.vue'
import { GetPreOrder, type PreOrder } from '~/composables/usepreorder'

definePageMeta({ layout: 'expert-backend-layout'})

const loading = ref(false)
const preorders = ref<PreOrder[]>([])

const fetchPreorders = async () => {
	loading.value = true
	try {
		const data = await GetPreOrder()
		preorders.value = Array.isArray(data) ? data : []
	} finally {
		loading.value = false
	}
}

const upcomingPreorders = computed(() => {
	const now = Date.now()
	return preorders.value
		.filter((item) => {
			const parsed = Date.parse(item.time)
			return !isNaN(parsed) && parsed >= now
		})
		.sort((a, b) => Date.parse(a.time) - Date.parse(b.time))
})

onMounted(fetchPreorders)
</script>

<style scoped>
.preorder-page {
	min-height: 100vh;
	padding: 2.5rem 1.5rem;
	background: #eef2ff;
}

.preorder-page__header {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 2rem;
}

.preorder-page__header h1 {
	font-size: 1.75rem;
	color: #1e1b4b;
	margin-bottom: 0.25rem;
}

.preorder-page__header p {
	color: #6d28d9;
	font-size: 0.95rem;
}

.preorder-page__state {
	padding: 2.5rem;
	text-align: center;
	color: #4c1d95;
	background: #fff;
	border-radius: 0.75rem;
	border: 1px dashed #c4b5fd;
}

.preorder-page__state--empty {
	color: #7c3aed;
}

.preorder-page__grid {
	display: flex;
	flex-direction: column;
	gap: 1rem;
}

@media (min-width: 768px) {
	.preorder-page__grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
		gap: 1.25rem;
	}
}
</style>
