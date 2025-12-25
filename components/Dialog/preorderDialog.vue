<template>
	<Dialog
		v-model:visible="visible"
		modal
		header="预约服务"
		class="preorder-dialog"
		:dismissable-mask="false"
		:style="{ width: '420px' }"
		@hide="handleClose"
	>
		<div class="space-y-4">
			<div class="bg-gray-50 rounded-md p-3 text-sm text-gray-600">
				正在预约：<span class="font-semibold text-gray-900">{{ expertName }}</span>
			</div>

			<div>
				<label class="text-sm font-medium text-gray-700 mb-2 block">选择预约时间</label>
				<Calendar
					v-model="selectedTime"
					showTime
					hour-format="24"
					:min-date="minDate"
					inline
					class="w-full"
				/>
			</div>
		</div>

		<template #footer>
			<div class="flex items-center justify-end gap-3 w-full">
				<Button label="取消" severity="secondary" outlined @click="handleClose" />
				<Button label="预约" :loading="submitting" :disabled="!selectedTime" @click="handleSubmit" />
			</div>
		</template>
	</Dialog>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import Dialog from 'primevue/dialog'
import Calendar from 'primevue/calendar'
import Button from 'primevue/button'
import { ElMessage } from 'element-plus'
import { PreOrder } from '~/composables/usepreorder'

const props = defineProps<{ 
	modelValue: boolean
	expertId: string
	expertName?: string
}>()

const emit = defineEmits<{ 
	(e: 'update:modelValue', value: boolean): void
	(e: 'success'): void
}>()

const visible = computed({
	get: () => props.modelValue,
	set: (value: boolean) => emit('update:modelValue', value)
})

const selectedTime = ref<Date | null>(null)
const submitting = ref(false)
const minDate = ref(new Date())

const expertName = computed(() => props.expertName || '未知专家')

watch(
	() => props.modelValue,
	(value) => {
		if (!value) {
			selectedTime.value = null
		} else {
			minDate.value = new Date()
		}
	}
)

const handleClose = () => {
	selectedTime.value = null
	emit('update:modelValue', false)
}

const handleSubmit = async () => {
	if (!selectedTime.value) {
		ElMessage.warning('请选择预约时间')
		return
	}
	if (!props.expertId) {
		ElMessage.error('无法确定专家信息')
		return
	}

	submitting.value = true
	try {
		await PreOrder(props.expertId, selectedTime.value)
		emit('success')
		handleClose()
	} finally {
		submitting.value = false
	}
}
</script>

<style scoped>
.preorder-dialog :deep(.p-dialog-content) {
	padding-bottom: 0.5rem;
}

.preorder-dialog :deep(.p-dialog-footer) {
	border-top: 1px solid #e5e7eb;
}
</style>
