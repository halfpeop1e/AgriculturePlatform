<template>
	<el-dialog
		v-model="visibleLocal"
		title="扫码支付"
		width="420px"
		:close-on-click-modal="false"
		@close="handleDialogClose"
	>
		<div class="dialog-body">
			<div class="countdown">请在 {{ formattedCountdown }} 内完成支付</div>
			<img class="qr-image" src="/pay.png" alt="支付二维码" />
			<p class="hint">使用手机银行或第三方支付 App 扫码完成支付</p>
		</div>
		<template #footer>
			<div class="dialog-footer">
				<el-button @click="handleCancel">取消支付</el-button>
			</div>
		</template>
	</el-dialog>
</template>

<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

const props = withDefaults(
	defineProps<{
		modelValue: boolean
		expireSeconds?: number
	}>(),
	{
		expireSeconds: 180,
	}
)

const emit = defineEmits<{
	(e: 'update:modelValue', value: boolean): void
	(e: 'cancel'): void
	(e: 'expired'): void
}>()

const visibleLocal = ref<boolean>(props.modelValue)
const remainingSeconds = ref<number>(props.expireSeconds)
const closeReason = ref<'cancel' | 'expired' | 'external' | null>(null)
let timer: ReturnType<typeof setInterval> | null = null

watch(
	() => props.modelValue,
	(visible) => {
		if (visible) {
			closeReason.value = null
			visibleLocal.value = true
			resetTimer()
		} else {
			if (visibleLocal.value && !closeReason.value) {
				closeReason.value = 'external'
			}
			visibleLocal.value = false
			stopTimer()
		}
	}
)

watch(
	() => props.expireSeconds,
	(seconds) => {
		if (visibleLocal.value) {
			resetTimer(seconds)
		}
	}
)

watch(visibleLocal, (visible) => {
	emit('update:modelValue', visible)
	if (!visible) {
		stopTimer()
	}
})

const formattedCountdown = computed(() => {
	const minutes = Math.floor(remainingSeconds.value / 60)
	const seconds = Math.max(remainingSeconds.value % 60, 0)
	const pad = (val: number) => val.toString().padStart(2, '0')
	return `${pad(minutes)}:${pad(seconds)}`
})

function resetTimer(initialSeconds?: number) {
	const base = typeof initialSeconds === 'number' ? initialSeconds : props.expireSeconds
	remainingSeconds.value = Math.max(0, base)
	stopTimer()
	startTimer()
}

function startTimer() {
	timer = setInterval(() => {
		if (remainingSeconds.value > 0) {
			remainingSeconds.value -= 1
		} else {
			handleExpired()
		}
	}, 1000)
}

function stopTimer() {
	if (timer) {
		clearInterval(timer)
		timer = null
	}
}

function handleExpired() {
	stopTimer()
	closeReason.value = 'expired'
	visibleLocal.value = false
	ElMessage.warning('支付二维码已过期，请重新获取')
}

function handleCancel() {
	closeReason.value = 'cancel'
	visibleLocal.value = false
}

function handleDialogClose() {
	if (closeReason.value === 'expired') {
		emit('expired')
	} else if (closeReason.value === 'cancel' || closeReason.value === null) {
		emit('cancel')
	}
	closeReason.value = null
}

onBeforeUnmount(() => {
	stopTimer()
})

onMounted(() => {
	if (props.modelValue) {
		resetTimer()
	}
})
</script>

<style scoped>
.dialog-body {
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 1rem;
	text-align: center;
}

.countdown {
	font-size: 1.1rem;
	font-weight: 600;
	color: #1f2937;
}

.qr-image {
	width: 240px;
	height: 240px;
	padding: 12px;
	border-radius: 12px;
	border: 1px solid #e5e7eb;
	background-color: #ffffff;
	box-shadow: 0 10px 25px rgba(15, 23, 42, 0.05);
	object-fit: contain;
}

.hint {
	font-size: 0.875rem;
	color: #6b7280;
}

.dialog-footer {
	display: flex;
	justify-content: flex-end;
}
</style>

