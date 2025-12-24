<template>
	<div class="min-h-screen bg-slate-50 py-10">
		<div class="mx-auto w-full max-w-4xl px-4 space-y-4">
      <el-card>
					<div class="flex items-center justify-between">
						<div>
							<h1 class="text-2xl font-semibold text-slate-900">完善专家资料</h1>
							<p class="mt-1 text-sm text-slate-500">
								请补充以下信息，以便用户更好地了解您的专业背景
							</p>
						</div>
						<Tag value="专家专享" severity="info" />
					</div>
				</el-card>
			<el-card class="shadow-lg">
		
				<div class="space-y-6 h-full">
					<section class="grid gap-6 md:grid-cols-1">
						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">姓名</label>
							<InputText v-model.trim="form.name" placeholder="请输入真实姓名" class="w-full" />
							<p v-if="errors.name" class="text-xs text-red-500">{{ errors.name }}</p>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">职业领域</label>
							<MultiSelect
								v-model="selectedField"
								:options="fieldOptions"
								display="chip"
								optionLabel="label"
								optionValue="value"
								placeholder="请选择领域"
								class="w-full"
								:maxSelectedLabels="1"
								:selectionLimit="1"
								@change="handleFieldChange"
							/>
							<p v-if="errors.field" class="text-xs text-red-500">{{ errors.field }}</p>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">头像链接</label>
							<InputText v-model.trim="form.avatar" placeholder="请输入头像图片链接" class="w-full" />
							<p v-if="errors.avatar" class="text-xs text-red-500">{{ errors.avatar }}</p>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">咨询单价（元）</label>
							<InputNumber v-model="form.price" inputClass="w-full mt-2" mode="decimal" :min="0" />
						</div>
					</section>

					<div class="space-y-2">
						<label class="text-sm font-medium text-slate-700">个人简介</label>
						<Textarea
							v-model.trim="form.introduction"
							rows="4"
							auto-resize
							placeholder="请介绍您的经历、擅长领域及可提供的帮助"
							class="w-full"
						/>
						<p v-if="errors.introduction" class="text-xs text-red-500">{{ errors.introduction }}</p>
					</div>

					<section class="grid gap-6 md:grid-cols-2">
						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">专业技能</label>
							<MultiSelect
								v-model="form.skills"
								:options="skillOptions"
								display="chip"
								placeholder="请选择或输入技能"
								:maxSelectedLabels="5"
								filter
								class="w-full"
								optionLabel="label"
								optionValue="value"
								:allowCreate="true"
								@create="handleSkillCreate"
							/>
							<p v-if="errors.skills" class="text-xs text-red-500">{{ errors.skills }}</p>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">教育背景</label>
							<InputText v-model.trim="form.education" placeholder="例如：华中农大 农学硕士" class="w-full" />
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">从业经验</label>
							<Textarea
								v-model.trim="form.experience"
								rows="3"
								auto-resize
								placeholder="请描述您的从业或研究经历"
								class="w-full"
							/>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">专业认证</label>
							<MultiSelect
								v-model="form.certification"
								:options="certificationOptions"
								display="chip"
								optionLabel="label"
								optionValue="value"
								placeholder="请选择或输入认证"
								:allowCreate="true"
								class="w-full"
								@create="handleCertificationCreate"
							/>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">可服务时段</label>
							<MultiSelect
								v-model="form.availableTime"
								:options="timeOptions"
								display="chip"
								optionLabel="label"
								optionValue="value"
								placeholder="请选择可服务时间"
								class="w-full"
							/>
						</div>

						<div class="space-y-2">
							<label class="text-sm font-medium text-slate-700">服务范围</label>
							<MultiSelect
								v-model="form.serviceScope"
								:options="serviceOptions"
								display="chip"
								optionLabel="label"
								optionValue="value"
								placeholder="请选择服务范围"
								class="w-full"
							/>
						</div>
					</section>

					<div class="flex justify-end gap-3 pt-4">
						<Button label="保存草稿" class="p-button-secondary" @click="handleSaveDraft" />
						<Button label="提交审批" icon="pi pi-check" :loading="submitting" @click="handleSubmit" />
					</div>
				</div>
			</el-card>
		</div>
	</div>
</template>

<script setup lang="ts">
import { reactive, ref, computed } from 'vue'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import MultiSelect from 'primevue/multiselect'
import InputNumber from 'primevue/inputnumber'
import Button from 'primevue/button'
import Tag from 'primevue/tag'
import { useUserStore } from '~/utils/userStore'
import { submitExpertProfile } from '~/composables/useChange'
import type { ExpertProfile } from '~/types/profile'

type SelectOption = { label: string; value: string }

const userStore = useUserStore()
const router = useRouter()

const presetProfile = computed(() => userStore.expertProfile)

const form = reactive({
	name: presetProfile.value?.name ?? '',
	avatar: presetProfile.value?.avatar ?? '',
	field: presetProfile.value?.field ?? '',
	introduction: presetProfile.value?.introduction ?? '',
	skills: [...(presetProfile.value?.skills ?? [])],
	education: presetProfile.value?.education ?? '',
	experience: presetProfile.value?.experience ?? '',
	certification: [...(presetProfile.value?.certification ?? [])],
	availableTime: [...(presetProfile.value?.availableTime ?? [])],
	serviceScope: [...(presetProfile.value?.serviceScope ?? [])],
	price: presetProfile.value?.price ?? null
})

const initialField = form.field ? [{ label: form.field, value: form.field }] : []
const selectedField = ref<SelectOption['value'][]>(initialField.map((option) => option.value))

const errors = reactive({
	name: '',
	avatar: '',
	field: '',
	introduction: '',
	skills: ''
})

const submitting = ref(false)

const fieldOptions: SelectOption[] = [
	{ label: '种植技术', value: '种植技术' },
	{ label: '农机服务', value: '农机服务' },
	{ label: '畜牧养殖', value: '畜牧养殖' },
	{ label: '农业金融', value: '农业金融' },
	{ label: '农村电商', value: '农村电商' },
	{ label: '病虫害防治', value: '病虫害防治' }
]

const rawSkills = ref([
	'大棚管理',
	'病虫害诊断',
	'精准施肥',
	'无人机植保',
	'乳牛饲养',
	'农产品销售'
])

const skillOptions = computed(() => rawSkills.value.map((item) => ({ label: item, value: item })))

const certificationOptions = ref<SelectOption[]>([
	{ label: '高级农艺师', value: '高级农艺师' },
	{ label: '植保工程师', value: '植保工程师' },
	{ label: '农村电商讲师', value: '农村电商讲师' }
])

const timeOptions: SelectOption[] = [
	{ label: '工作日（9:00-12:00）', value: '工作日（9:00-12:00）' },
	{ label: '工作日（14:00-18:00）', value: '工作日（14:00-18:00）' },
	{ label: '周末全天', value: '周末全天' }
]

const serviceOptions: SelectOption[] = [
	{ label: '线上视频', value: '线上视频' },
	{ label: '在线问答', value: '在线问答' },
	{ label: '线下驻场', value: '线下驻场' },
	{ label: '团队培训', value: '团队培训' }
]

const buildProfilePayload = (): ExpertProfile => ({
	name: form.name,
	avatar: form.avatar,
	field: form.field,
	introduction: form.introduction,
	skills: [...form.skills],
	education: form.education || undefined,
	experience: form.experience || undefined,
	certification: form.certification?.length ? [...form.certification] : undefined,
	availableTime: form.availableTime?.length ? [...form.availableTime] : undefined,
	serviceScope: form.serviceScope?.length ? [...form.serviceScope] : undefined,
	price: form.price ?? undefined
})

const resetErrors = () => {
	errors.name = ''
	errors.avatar = ''
	errors.field = ''
	errors.introduction = ''
	errors.skills = ''
}

const validate = () => {
	resetErrors()
	let valid = true
	if (!form.name) {
		errors.name = '请填写真实姓名'
		valid = false
	}
	if (!selectedField.value.length) {
		errors.field = '请至少选择一个职业领域'
		valid = false
	}
	if (!form.avatar) {
		errors.avatar = '请提供头像链接，可使用企业微信或网盘链接'
		valid = false
	}
	if (!form.introduction || form.introduction.length < 20) {
		errors.introduction = '简介不少于 20 个字符，以便用户了解您的专业能力'
		valid = false
	}
	if (!form.skills.length) {
		errors.skills = '请选择至少一项专业技能'
		valid = false
	}
	return valid
}

const handleFieldChange = () => {
	form.field = selectedField.value[0] ?? ''
}

const handleSkillCreate = (event: { value: string }) => {
	const value = event.value.trim()
	if (!value) return
	if (!rawSkills.value.includes(value)) {
		rawSkills.value = [...rawSkills.value, value]
	}
	if (!form.skills.includes(value)) {
		form.skills.push(value)
	}
}

const handleCertificationCreate = (event: { value: string }) => {
	const value = event.value.trim()
	if (!value) return
	if (!certificationOptions.value.some((item) => item.value === value)) {
		certificationOptions.value = [...certificationOptions.value, { label: value, value }]
	}
	if (!form.certification?.includes(value)) {
		form.certification = [...(form.certification ?? []), value]
	}
}

const persistDraft = () => {
	if (confirm('暂存草稿后可稍后继续填写，是否确认？')) {
		const payload = buildProfilePayload()
		userStore.setExpertProfile(payload, false)
		ElMessage.success('草稿已保存')
	}
}

const handleSaveDraft = () => {
	persistDraft()
}

const handleSubmit = async () => {
	if (!validate()) return

	submitting.value = true
	try {
		const payload = buildProfilePayload()
		const success = await submitExpertProfile(payload)
		if (!success) {
			ElMessage.error('资料提交失败，请稍后再试')
			return
		}
		userStore.setExpertProfile(payload)
		ElMessage.success('资料已提交，正在跳转专家工作台')
		await router.push('/expert/dashboard')
	} finally {
		submitting.value = false
	}
}

if (!selectedField.value.length && form.field) {
	selectedField.value = [form.field]
}

definePageMeta({
	layout: 'expert-backend-layout'
})
</script>

<style scoped>
.p-card-title {
	margin-bottom: 0.5rem;
}
</style>

