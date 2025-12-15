<template>
	<div class="p-6 space-y-6">
		<section class="flex items-center justify-between">
			<div>
				<h1 class="text-2xl font-semibold text-gray-800">金融审批中心</h1>
				<p class="text-gray-500">审批贷款、融资以及 money 模块提交的其它金融申请</p>
			</div>
			<el-select v-model="currentFilter" placeholder="筛选状态" style="width: 180px">
				<el-option v-for="option in filterOptions" :key="option.value" :label="option.label" :value="option.value" />
			</el-select>
		</section>

		<el-card shadow="hover" class="rounded-2xl border border-gray-200">
			<el-table :data="filteredRequests" stripe border :max-height="520">
				<el-table-column prop="referenceNo" label="申请编号" width="140" />
				<el-table-column prop="applicant" label="申请人" width="140" />
				<el-table-column prop="type" label="类型" width="120">
					<template #default="{ row }">
						<el-tag :type="typeTagMap[row.type as FinanceRequest['type']] || 'info'" effect="dark">{{ typeLabel(row.type) }}</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="amount" label="金额 (¥)" width="140">
					<template #default="{ row }">{{ formatCurrency(row.amount) }}</template>
				</el-table-column>
				<el-table-column prop="submittedAt" label="提交时间" width="180" />
				<el-table-column prop="status" label="状态" width="120">
					<template #default="{ row }">
						<el-tag :type="statusTagType(row.status)">{{ statusLabel(row.status) }}</el-tag>
					</template>
				</el-table-column>
				<el-table-column label="操作" width="220" fixed="right">
					<template #default="{ row }">
						<el-button type="primary" plain size="small" @click="openDetail(row)">查看详情</el-button>
						<el-button v-if="row.status === 'pending'" type="success" size="small" @click="approve(row)">通过</el-button>
						<el-button v-if="row.status === 'pending'" type="danger" size="small" @click="openReject(row)">驳回</el-button>
					</template>
				</el-table-column>
			</el-table>

			<div v-if="filteredRequests.length === 0" class="py-12 text-center text-gray-400">
				暂无符合筛选条件的审批请求
			</div>
		</el-card>

		<!-- 详情对话框 -->
		<el-dialog v-model="detailVisible" title="申请详情" width="560px" destroy-on-close>
			<el-descriptions :column="2" border label-class-name="w-32">
				<el-descriptions-item label="申请编号">{{ activeRequest?.referenceNo }}</el-descriptions-item>
				<el-descriptions-item label="当前状态">
					<el-tag :type="statusTagType(activeRequest?.status)" effect="dark">{{ statusLabel(activeRequest?.status) }}</el-tag>
				</el-descriptions-item>
				<el-descriptions-item label="申请人">{{ activeRequest?.applicant }}</el-descriptions-item>
				<el-descriptions-item label="联系方式">{{ activeRequest?.contact }}</el-descriptions-item>
				<el-descriptions-item label="类型">{{ typeLabel(activeRequest?.type) }}</el-descriptions-item>
				<el-descriptions-item label="申请金额">{{ formatCurrency(activeRequest?.amount || 0) }}</el-descriptions-item>
				<el-descriptions-item label="提交时间">{{ activeRequest?.submittedAt }}</el-descriptions-item>
				<el-descriptions-item label="贷款期限">{{ activeRequest?.term }} 月</el-descriptions-item>
			</el-descriptions>
			<el-divider />
			<section>
				<h3 class="text-base font-semibold mb-2">资金用途</h3>
				<p class="text-sm text-gray-600 whitespace-pre-line">{{ activeRequest?.purpose }}</p>
			</section>
			<section v-if="activeRequest?.attachments?.length" class="mt-4">
				<h3 class="text-base font-semibold mb-2">附件</h3>
				<ul class="list-disc list-inside space-y-1 text-sm text-blue-500">
					<li v-for="(file, index) in activeRequest?.attachments" :key="index">
						<a :href="file.url" target="_blank" class="hover:underline">{{ file.name }}</a>
					</li>
				</ul>
			</section>
			<section v-if="activeRequest?.status === 'rejected' && activeRequest?.rejectReason" class="mt-4">
				<h3 class="text-base font-semibold mb-2 text-red-500">驳回说明</h3>
				<p class="text-sm text-gray-600 whitespace-pre-line">{{ activeRequest?.rejectReason }}</p>
			</section>
			<template #footer>
				<div class="flex justify-end gap-2">
					<el-button @click="detailVisible = false">关闭</el-button>
					<el-button v-if="activeRequest?.status === 'pending'" type="success" @click="approve(activeRequest)">通过</el-button>
					<el-button v-if="activeRequest?.status === 'pending'" type="danger" @click="openReject(activeRequest)">驳回</el-button>
				</div>
			</template>
		</el-dialog>

		<!-- 驳回原因对话框 -->
		<el-dialog v-model="rejectVisible" title="填写驳回原因" width="460px" destroy-on-close>
			<el-form label-position="top">
				<el-form-item label="驳回原因">
					<el-input
						v-model="rejectReason"
						type="textarea"
						:rows="4"
						maxlength="200"
						show-word-limit
						placeholder="请说明驳回原因"
					/>
				</el-form-item>
			</el-form>
			<template #footer>
				<div class="flex justify-end gap-2">
					<el-button @click="rejectVisible = false">取消</el-button>
					<el-button type="danger" @click="confirmReject">确认驳回</el-button>
				</div>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

definePageMeta({ layout: 'home-page-layout' })

interface FinanceRequest {
	id: string
	referenceNo: string
	applicant: string
	contact: string
	type: 'loan' | 'finance' | 'subsidy'
	amount: number
	term: number
	purpose: string
	submittedAt: string
	status: 'pending' | 'approved' | 'rejected'
	attachments?: { name: string; url: string }[]
	rejectReason?: string
}

const filterOptions = [
	{ label: '全部', value: 'all' },
	{ label: '待审批', value: 'pending' },
	{ label: '已通过', value: 'approved' },
	{ label: '已驳回', value: 'rejected' },
]

const typeTagMap: Record<FinanceRequest['type'], 'info' | 'success' | 'warning'> = {
	loan: 'success',
	finance: 'info',
	subsidy: 'warning',
}

// Mock dataset for demo; replace with API integration when backend is ready.
const requests = ref<FinanceRequest[]>([
	{
		id: 'req-loan-001',
		referenceNo: 'LN-2025-0001',
		applicant: '李四',
		contact: '138-0000-0001',
		type: 'loan',
		amount: 200000,
		term: 12,
		purpose: '购买自动化灌溉设备，提升灌溉效率并降低人工成本。',
		submittedAt: '2025-12-05 10:21',
		status: 'pending',
		attachments: [
			{ name: '商业计划书.pdf', url: '#' },
			{ name: '资金使用预算.xlsx', url: '#' },
		],
	},
	{
		id: 'req-fin-002',
		referenceNo: 'FN-2025-0008',
		applicant: '王梅',
		contact: '139-1111-2222',
		type: 'finance',
		amount: 50000,
		term: 6,
		purpose: '采购冬季保温温室，扩大草莓种植面积。',
		submittedAt: '2025-12-03 14:12',
		status: 'approved',
	},
	{
		id: 'req-sub-003',
		referenceNo: 'SB-2025-0003',
		applicant: '张强',
		contact: '137-8888-9999',
		type: 'subsidy',
		amount: 30000,
		term: 9,
		purpose: '申请防灾补贴，用于修复被风灾损坏的设施。',
		submittedAt: '2025-11-28 09:45',
		status: 'rejected',
		rejectReason: '材料不全，请补充修复方案与费用明细后重新提交。',
	},
])

const currentFilter = ref<'all' | FinanceRequest['status']>('pending')
const detailVisible = ref(false)
const rejectVisible = ref(false)
const rejectReason = ref('')
const activeRequest = ref<FinanceRequest | null>(null)

const filteredRequests = computed(() => {
	if (currentFilter.value === 'all') return requests.value
	return requests.value.filter((item) => item.status === currentFilter.value)
})

function statusLabel(status?: FinanceRequest['status']) {
	switch (status) {
		case 'approved':
			return '已通过'
		case 'rejected':
			return '已驳回'
		case 'pending':
		default:
			return '待审批'
	}
}

function statusTagType(status?: FinanceRequest['status']) {
	switch (status) {
		case 'approved':
			return 'success'
		case 'rejected':
			return 'danger'
		default:
			return 'warning'
	}
}

function typeLabel(type?: FinanceRequest['type']) {
	switch (type) {
		case 'loan':
			return '贷款审批'
		case 'finance':
			return '金融服务'
		case 'subsidy':
			return '补贴申请'
		default:
			return '未知类型'
	}
}

function formatCurrency(val: number) {
	return new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(val)
}

function openDetail(request: FinanceRequest) {
	activeRequest.value = request
	detailVisible.value = true
}

function openReject(request: FinanceRequest | null) {
	if (!request) return
	activeRequest.value = request
	rejectReason.value = request.rejectReason || ''
	rejectVisible.value = true
}

async function approve(request: FinanceRequest | null) {
	if (!request) return
	try {
		await ElMessageBox.confirm(`确认通过【${request.applicant}】的审批申请？`, '审批确认', {
			confirmButtonText: '确认通过',
			cancelButtonText: '取消',
			type: 'success',
		})
		request.status = 'approved'
		request.rejectReason = undefined
		detailVisible.value = false
		ElMessage.success('已通过该申请')
	} catch (error) {
		if (error !== 'cancel' && error !== 'close') {
			ElMessage.error('操作失败，请稍后重试')
		}
	}
}

function confirmReject() {
	if (!activeRequest.value) return
	if (!rejectReason.value.trim()) {
		ElMessage.warning('请填写驳回原因')
		return
	}
	activeRequest.value.status = 'rejected'
	activeRequest.value.rejectReason = rejectReason.value.trim()
	detailVisible.value = false
	rejectVisible.value = false
	rejectReason.value = ''
	ElMessage.success('已驳回该申请')
}
</script>

<style scoped>
.el-table :deep(.el-table__body-wrapper) {
	max-height: 520px;
}
</style>

