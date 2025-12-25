<template>
	<div class="preorder-card">
		<div class="preorder-card__header">
			<div>
				<p class="label">专家</p>
				<p class="value">{{ preorder.expertname }}</p>
			</div>
			<Tag :value="statusLabel" :severity="statusSeverity" />
		</div>

		<div class="preorder-card__body">
			<div>
				<p class="label">预约人</p>
				<p class="value">{{ preorder.author }}</p>
			</div>

			<div>
				<p class="label">预约时间</p>
				<p class="value">{{ formattedTime }}</p>
			</div>

			<div>
				<p class="label">订单号</p>
				<p class="value monospace">{{ preorder.orderId }}</p>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import Tag from 'primevue/tag'
import type { PreOrder } from '~/composables/usepreorder'

const props = defineProps<{ preorder: PreOrder }>()

const formattedTime = computed(() => {
	const parsed = new Date(props.preorder.time)
	return isNaN(parsed.getTime())
		? props.preorder.time
		: parsed.toLocaleString('zh-CN', {
				year: 'numeric',
				month: '2-digit',
				day: '2-digit',
				hour: '2-digit',
				minute: '2-digit'
			})
})

const isExpired = computed(() => {
	const parsed = new Date(props.preorder.time)
	return isNaN(parsed.getTime()) ? false : parsed.getTime() < Date.now()
})

const statusLabel = computed(() => (isExpired.value ? '已过期' : '待咨询'))
const statusSeverity = computed(() => (isExpired.value ? 'secondary' : 'success'))
</script>

<style scoped>
.preorder-card {
	border: 1px solid #e2e8f0;
	border-radius: 0.75rem;
	padding: 1.25rem;
	background: #fff;
	box-shadow: 0 10px 25px rgba(15, 23, 42, 0.04);
	display: flex;
	flex-direction: column;
	gap: 1rem;
}

.preorder-card__header {
	display: flex;
	align-items: center;
	justify-content: space-between;
}

.preorder-card__body {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
	gap: 1rem;
}

.label {
	font-size: 0.75rem;
	color: #94a3b8;
	text-transform: uppercase;
	letter-spacing: 0.06em;
	margin-bottom: 0.25rem;
}

.value {
	font-size: 0.95rem;
	color: #0f172a;
	font-weight: 600;
}

.monospace {
	font-family: 'SFMono-Regular', Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace;
	font-size: 0.85rem;
}
</style>
