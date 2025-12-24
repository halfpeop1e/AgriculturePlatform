<template>
	<div class="min-h-screen bg-slate-100 py-8">
		<div class="mx-auto w-full max-w-5xl px-4">
			<Button label="返回" icon="pi pi-arrow-left" class="p-button-text mb-4" @click="handleBack" />

			<div v-if="loading" class="flex flex-col items-center justify-center rounded-lg bg-white py-16 shadow">
				<ProgressSpinner strokeWidth="4" style="width: 3rem; height: 3rem" />
				<p class="mt-4 text-sm text-slate-500">正在加载专家资料...</p>
			</div>

			<Card v-else-if="expertDetail" class="shadow">
				<template #title>
					<div class="flex flex-col gap-4 md:flex-row md:items-center">
						<Avatar :image="expertDetail.avatar || defaultAvatar" shape="circle" class="h-24 w-24" />
						<div class="space-y-2">
							<div class="flex flex-wrap items-center gap-2">
								<h1 class="text-2xl font-semibold text-slate-900">{{ expertDetail.name }}</h1>
								<Tag :value="expertDetail.field" severity="info" />
								<Tag v-if="expertDetail.consultCount" value="认证专家" severity="success" />
							</div>
							<div class="flex flex-wrap gap-6 text-sm text-slate-500">
								<span class="flex items-center gap-1">
									<i class="pi pi-users text-primary-500"></i>
									{{ expertDetail.consultCount }} 次咨询
								</span>
								<span class="flex items-center gap-1">
									<i class="pi pi-star-fill text-yellow-500"></i>
									{{ expertDetail.rating.toFixed(1) }} 分
								</span>
								<span v-if="expertDetail.responseTime" class="flex items-center gap-1">
									<i class="pi pi-clock text-emerald-500"></i>
									{{ expertDetail.responseTime }}
								</span>
							</div>
						</div>
					</div>
				</template>

				<template #content>
					<section class="space-y-8">
						<div>
							<h2 class="text-lg font-semibold text-slate-800">个人简介</h2>
							<p class="mt-2 text-sm leading-relaxed text-slate-600">
								{{ expertDetail.introduction }}
							</p>
						</div>

						<Divider />

						<div class="grid gap-6 md:grid-cols-2">
							<div>
								<h3 class="text-base font-semibold text-slate-800">专业技能</h3>
								<div class="mt-3 flex flex-wrap gap-2">
									<Tag
										v-for="skill in expertDetail.skills"
										:key="skill"
										:value="skill"
										severity="secondary"
										class="bg-slate-100 text-slate-700"
									/>
								</div>
							</div>

							<div v-if="expertDetail.availableTime?.length">
								<h3 class="text-base font-semibold text-slate-800">可咨询时间</h3>
								<ul class="mt-3 space-y-2 text-sm text-slate-600">
									<li v-for="time in expertDetail.availableTime" :key="time">{{ time }}</li>
								</ul>
							</div>

							<div v-if="expertDetail.education">
								<h3 class="text-base font-semibold text-slate-800">教育背景</h3>
								<p class="mt-3 text-sm text-slate-600">{{ expertDetail.education }}</p>
							</div>

							<div v-if="expertDetail.experience">
								<h3 class="text-base font-semibold text-slate-800">从业经验</h3>
								<p class="mt-3 text-sm text-slate-600">{{ expertDetail.experience }}</p>
							</div>

							<div v-if="expertDetail.certification?.length">
								<h3 class="text-base font-semibold text-slate-800">专业认证</h3>
								<ul class="mt-3 space-y-2 text-sm text-slate-600">
									<li v-for="cert in expertDetail.certification" :key="cert">{{ cert }}</li>
								</ul>
							</div>

							<div v-if="expertDetail.serviceScope?.length">
								<h3 class="text-base font-semibold text-slate-800">服务范围</h3>
								<ul class="mt-3 space-y-2 text-sm text-slate-600">
									<li v-for="scope in expertDetail.serviceScope" :key="scope">{{ scope }}</li>
								</ul>
							</div>
						</div>

						<Divider v-if="expertDetail.recentCases?.length" />

						<div v-if="expertDetail.recentCases?.length">
								<h3 class="text-base font-semibold text-slate-800">近期案例</h3>
							<div
								v-for="item in expertDetail.recentCases"
								:key="`${item.date}-${item.question}`"
								class="mt-3 rounded-lg border border-slate-200 bg-slate-50 p-4"
							>
								<div class="text-xs text-slate-400">{{ item.date }}</div>
								<div class="mt-2 text-sm font-medium text-slate-700">问题：{{ item.question }}</div>
								<div class="mt-1 text-sm text-slate-600">解答：{{ item.answer }}</div>
							</div>
						</div>
					</section>
				</template>

				<template #footer>
					<div class="flex flex-wrap items-center justify-between gap-3">
						<div class="text-lg font-semibold text-primary-600">
							<span v-if="expertDetail.price">￥{{ expertDetail.price }}/次</span>
							<span v-else>咨询价格面议</span>
						</div>
						<div class="flex gap-3">
							<Button label="编辑资料" icon="pi pi-pencil" severity="secondary" @click="handleEdit" />
						</div>
					</div>
				</template>
			</Card>

			<div
				v-else
				class="rounded-lg bg-white p-10 text-center text-sm text-slate-500 shadow"
			>
				{{ errorMessage || '未找到对应的专家资料，请返回重试。' }}
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Card from 'primevue/card'
import Avatar from 'primevue/avatar'
import Tag from 'primevue/tag'
import Button from 'primevue/button'
import Divider from 'primevue/divider'
import ProgressSpinner from 'primevue/progressspinner'
import type { ExpertDetail } from '~/types/expert'
import { getExpertDetail } from '~/composables/getExpert'
import { getExpertProfile } from '~/composables/getProfile'
import { getMockExpertById } from '~/types/experttest'
import { useUserStore } from '~/utils/userStore'

definePageMeta({
	layout: 'expert-backend-layout'
})

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const loading = ref(true)
const expertDetail = ref<ExpertDetail | null>(null)
const errorMessage = ref('')

const defaultAvatar = '/ioanImage/default-avatar.png'

const buildDetailFromProfile = () => {
	const profile = userStore.expertProfile
	if (!profile) return null
	return {
		id: userStore.userId || 'current-expert',
		name: profile.name,
		avatar: userStore.avatar,
		field: profile.field,
		introduction: profile.introduction,
		skills: [...profile.skills],
		consultCount: 0,
		rating: 5,
		responseTime: profile.availableTime?.[0] ? '当日响应' : undefined,
		education: profile.education,
		experience: profile.experience,
		certification: profile.certification,
		availableTime: profile.availableTime,
		serviceScope: profile.serviceScope,
		price: profile.price,
		recentCases: []
	} satisfies ExpertDetail
}

const loadExpertDetail = async (userId: string) => {
	loading.value = true
	errorMessage.value = ''
	expertDetail.value = null

	try {
		// 如果查看的是当前登录用户，优先调用 getExpertProfile 同步并基于 profile 构建详情
		if (userStore.userId && userId === userStore.userId) {
			const profile = await getExpertProfile()
			if (profile) {
				const localProfile = buildDetailFromProfile()
				if (localProfile) {
					expertDetail.value = localProfile
					return
				}
			}
		} else {
			// 查看他人，调用按 id 查询的接口
			const other = await getExpertDetail(userId)
			if (other) {
				expertDetail.value = other
				return
			}
		}

		const mock = getMockExpertById(userId)
		if (mock) {
			expertDetail.value = mock
			return
		}

		errorMessage.value = '暂未获取到该专家的详细资料'
	} catch (error) {
		console.error('加载专家详情失败', error)
		errorMessage.value = '加载专家资料时出现问题，请稍后再试'
	} finally {
		loading.value = false
	}
}

watch(
	() => route.params.userId,
	(id) => {
		const parsedId = typeof id === 'string' ? id : Array.isArray(id) ? id[0] : ''
		if (!parsedId) {
			errorMessage.value = '缺少专家标识'
			loading.value = false
			return
		}
		loadExpertDetail(parsedId)
	},
	{ immediate: true }
)

const handleBack = () => {
	navigateTo('/specialBoard')
}

const handleEdit = () => {
	if (userStore.userId && route.params.userId === userStore.userId) {
		router.push('/infomationcomplete')
		return
	}
	ElMessage.info('仅限专家本人编辑资料')
}
</script>

<style scoped>
.p-card-title {
	margin-bottom: 0.5rem;
}
</style>
