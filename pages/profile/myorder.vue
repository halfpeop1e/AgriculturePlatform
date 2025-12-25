<template>
	<div class="preorder-page">
		<div class="preorder-page__header">
			<div>
				<h1>我的预约</h1>
				<p>查看即将到来的专家预约，及时准备沟通内容。</p>
			</div>
			<Button label="刷新" icon="pi pi-refresh" outlined :loading="loading" @click="fetchPreorders" />
		</div>

		<div v-if="loading" class="preorder-page__state">正在加载预约信息...</div>
		<div v-else-if="!upcomingPreorders.length" class="preorder-page__state preorder-page__state--empty">
			暂无新的预约，去专家列表看看吧。
		</div>
		<div v-else class="preorder-page__grid">
			<PreorderCard v-for="item in upcomingPreorders" :key="item.orderId" :preorder="item" />
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import Button from 'primevue/button'
import PreorderCard from '~/components/Card/preorderCard.vue'
import { GetPreOrder, type PreOrder } from '~/composables/usepreorder'

definePageMeta({ layout: 'home-page-layout' })

const loading = ref(false)
const allPreorders = ref<PreOrder[]>([])

const fetchPreorders = async () => {
	loading.value = true
	try {
		const data = await GetPreOrder()
		allPreorders.value = Array.isArray(data) ? data : []
	} finally {
		loading.value = false
	}
}

const upcomingPreorders = computed(() => {
	const now = Date.now()
	return allPreorders.value
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
	background: #f8fafc;
}

.preorder-page__header {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 2rem;
}

.preorder-page__header h1 {
	font-size: 1.75rem;
	color: #0f172a;
	margin-bottom: 0.25rem;
}

.preorder-page__header p {
	color: #64748b;
	font-size: 0.95rem;
}

.preorder-page__state {
	padding: 2.5rem;
	text-align: center;
	color: #475569;
	background: #fff;
	border-radius: 0.75rem;
	border: 1px dashed #cbd5f5;
}

.preorder-page__state--empty {
	color: #94a3b8;
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
