<template>
  <div class="min-h-screen bg-slate-50">
    <div class="grid min-h-screen grid-cols-[280px_1fr]">
      <aside class="border-r border-slate-200 bg-white">
        <div class="flex h-full flex-col items-center px-6 py-10">
          <div class="w-full space-y-3">
            <Button
              label="金融产品"
              icon="pi pi-briefcase"
              class="w-full"
              :severity="route.path === '/money/finance/product' ? 'primary' : 'secondary'"
              :outlined="route.path !== '/money/finance/product'"
              @click="router.push('/money/finance/product')"
            />
            <Button
              label="进度查询"
              icon="pi pi-chart-bar"
              class="w-full"
              :severity="route.path === '/money/finance/process' ? 'primary' : 'secondary'"
              :outlined="route.path !== '/money/finance/process'"
              @click="router.push('/money/finance/process')"
            />
            <div class="pt-6">
              <Button label="添加产品" icon="pi pi-plus" class="w-full" severity="help" @click="openAdd = true" />
            </div>
          </div>
        </div>
      </aside>

      <main class="col-start-2 overflow-y-auto px-10 py-12">
        <div class="flex flex-wrap items-center justify-between gap-4">
          <div>
            <h1 class="text-2xl font-semibold text-slate-900">金融产品</h1>
            <p class="text-sm text-slate-500">根据您的业务需求筛选合适的贷款方案</p>
          </div>
          <Button label="重置筛选" icon="pi pi-refresh" severity="secondary" outlined @click="resetFilters" />
        </div>

        <div class="mt-8 grid gap-4 rounded-xl border border-slate-200 bg-white p-6 shadow-sm md:grid-cols-2 xl:grid-cols-4">
          <div class="space-y-2">
            <span class="text-sm font-medium text-slate-700">贷款金额</span>
            <InputNumber
              v-model="amountFilter"
              mode="currency"
              currency="CNY"
              locale="zh-CN"
              :min="0"
              showButtons
              class="w-full"
              inputClass="w-full"
              placeholder="请输入贷款金额"
            />
          </div>
          <div class="space-y-2">
            <span class="text-sm font-medium text-slate-700">业务范围</span>
            <MultiSelect
              v-model="selectedPurposes"
              :options="options"
              optionLabel="label"
              optionValue="value"
              display="chip"
              placeholder="请选择业务范围"
              class="w-full"
              :maxSelectedLabels="3"
              showClear
            />
          </div>
          <div class="space-y-2">
            <span class="text-sm font-medium text-slate-700">担保要求</span>
            <Dropdown
              v-model="collateralFilter"
              :options="danbaoOptions"
              optionLabel="label"
              optionValue="value"
              placeholder="请选择担保要求"
              class="w-full"
              showClear
            />
          </div>
          <div class="space-y-2">
            <span class="text-sm font-medium text-slate-700">征信要求</span>
            <Dropdown
              v-model="creditFilter"
              :options="trustOptions"
              optionLabel="label"
              optionValue="value"
              placeholder="请选择征信要求"
              class="w-full"
              showClear
            >
              <template #option="{ option }">
                <div class="flex flex-col">
                  <span class="text-sm font-medium text-slate-800">{{ option.label }}</span>
                  <span class="text-xs text-slate-500">{{ option.description }}</span>
                </div>
              </template>
            </Dropdown>
          </div>
        </div>

        <Message severity="info" class="mt-6">
          当前共有 {{ filteredProducts.length }} 款贷款产品符合条件
        </Message>

        <div class="mt-6 grid gap-6 md:grid-cols-2 xl:grid-cols-3">
          <Card
            v-for="product in filteredProducts"
            :key="product.productId"
            class="h-full border border-slate-200 shadow-sm transition hover:-translate-y-1 hover:shadow-md"
          >
            <template #title>
              <div class="flex items-start justify-between gap-3">
                <div>
                  <div class="text-xl font-semibold text-slate-900">{{ product.productName }}</div>
                  <p class="text-xs text-slate-400">产品编号：{{ product.productId }}</p>
                </div>
                <img
                  class="h-16 w-16 rounded object-cover"
                  :src="product.productAvatar || defaultProductImage"
                  alt="产品封面"
                />
              </div>
            </template>

            <template #content>
              <div class="space-y-4 text-sm text-slate-600">
                <div class="flex items-center justify-between">
                  <span>贷款额度 (元)</span>
                  <span class="font-semibold text-slate-900">
                    {{ product.loanAmountRange.min }} - {{ product.loanAmountRange.max }}
                  </span>
                </div>
                <div class="flex flex-wrap items-center gap-2">
                  <span class="text-slate-500">业务范围</span>
                  <Tag v-if="product.supportedPurposes.production" severity="success" value="农业生产" />
                  <Tag v-if="product.supportedPurposes.equipment" severity="success" value="设备购置" />
                  <Tag v-if="product.supportedPurposes.land" severity="success" value="土地流转/租赁" />
                  <Tag v-if="product.supportedPurposes.operating" severity="success" value="经营周转" />
                  <Tag v-if="product.supportedPurposes.infrastructure" severity="success" value="设施建设" />
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-slate-500">年利率</span>
                  <span class="font-semibold text-emerald-600">
                    {{ (product.interestRate.finalRate * 100).toFixed(2) }}%
                  </span>
                </div>
                <div class="text-xs text-slate-400">
                  生效日期：{{ new Date(product.effectiveDate).toLocaleDateString() }}
                </div>
              </div>
            </template>

            <template #footer>
              <div class="flex gap-3">
                <Button
                  label="查看详情"
                  icon="pi pi-eye"
                  severity="secondary"
                  outlined
                  class="flex-1"
                  @click="showProductDetail(product)"
                />
                <Button
                  label="申请"
                  icon="pi pi-send"
                  class="flex-1"
                  @click="openApplyDialog(product)"
                />
              </div>
            </template>
          </Card>
        </div>

        <div v-if="!filteredProducts.length" class="mt-10 rounded-xl border border-dashed border-slate-200 bg-white p-12 text-center text-slate-400">
          暂无符合筛选条件的产品，请调整筛选条件后再试。
        </div>

        <div class="mt-10 flex justify-center">
          <Paginator
            :first="(current - 1) * rows"
            :rows="rows"
            :totalRecords="productStore.total"
            :rowsPerPageOptions="[9, 12, 15]"
            @page="handlePageChange"
            class="rounded-lg border border-slate-200 bg-white shadow-sm"
          />
        </div>
      </main>
    </div>

    <Dialog v-model:visible="centerDialogVisible" header="产品详情" modal class="w-full max-w-2xl">
      <template v-if="selectedProduct">
        <div class="space-y-6">
          <div class="flex items-center gap-4 border-b border-slate-200 pb-4">
            <img
              :src="selectedProduct.productAvatar || defaultProductImage"
              alt="产品图片"
              class="h-20 w-20 rounded object-cover"
            />
            <div>
              <h3 class="text-xl font-semibold text-slate-900">{{ selectedProduct.productName }}</h3>
              <p class="text-sm text-slate-500">产品编号：{{ selectedProduct.productId }}</p>
            </div>
          </div>

          <section>
            <h4 class="text-lg font-semibold text-slate-800">银行信息</h4>
            <div class="mt-3 grid gap-4 rounded-lg bg-slate-50 p-4 md:grid-cols-2">
              <div>银行名称：{{ selectedProduct.financialInstitution.name }}</div>
              <div>客服电话：{{ selectedProduct.financialInstitution.customerService }}</div>
            </div>
          </section>

          <section>
            <h4 class="text-lg font-semibold text-slate-800">融资条款</h4>
            <div class="mt-3 grid gap-4 rounded-lg bg-slate-50 p-4 md:grid-cols-2">
              <div>
                <span class="text-slate-600">贷款金额范围：</span>
                <span class="font-medium text-slate-900">{{ selectedProduct.loanAmountRange.min }} - {{ selectedProduct.loanAmountRange.max }}</span>
              </div>
              <div>
                <span class="text-slate-600">年利率：</span>
                <span class="font-medium text-emerald-600">{{ (selectedProduct.interestRate.finalRate * 100).toFixed(2) }}%</span>
                <span class="text-xs text-slate-400">（{{ selectedProduct.interestRate.type === 0 ? '固定' : '浮动' }}利率）</span>
              </div>
              <div>
                <span class="text-slate-600">贷款期限：</span>
                <span class="font-medium text-slate-900">{{ selectedProduct.loanTerm.minMonths }} - {{ selectedProduct.loanTerm.maxMonths }} 个月</span>
              </div>
              <div>
                <span class="text-slate-600">预计审批时间：</span>
                <span class="font-medium text-slate-900">{{ selectedProduct.estimatedTime }}</span>
              </div>
            </div>
            <div v-if="selectedProduct.interestRate.discountDescription" class="mt-2 rounded-lg bg-blue-50 p-3 text-sm text-blue-700">
              利率优惠：{{ selectedProduct.interestRate.discountDescription }}
            </div>
          </section>

          <section>
            <h4 class="text-lg font-semibold text-slate-800">申请条件</h4>
            <div class="mt-3 space-y-2 rounded-lg bg-slate-50 p-4">
              <div>最低经营年限：{{ selectedProduct.eligibility.minOperatingYears }} 年</div>
              <div>征信要求：{{ selectedProduct.eligibility.creditRequirement }}</div>
              <div v-if="selectedProduct.eligibility.collateralRequirements">
                担保要求：{{ selectedProduct.eligibility.collateralRequirements }}
              </div>
            </div>
          </section>

          <section>
            <h4 class="text-lg font-semibold text-slate-800">支持的业务范围</h4>
            <div class="mt-3 flex flex-wrap gap-2">
              <Tag v-if="selectedProduct.supportedPurposes.production" severity="success" value="农业生产" />
              <Tag v-if="selectedProduct.supportedPurposes.equipment" severity="success" value="设备购置" />
              <Tag v-if="selectedProduct.supportedPurposes.land" severity="success" value="土地流转/租赁" />
              <Tag v-if="selectedProduct.supportedPurposes.operating" severity="success" value="经营周转" />
              <Tag v-if="selectedProduct.supportedPurposes.infrastructure" severity="success" value="设施建设" />
            </div>
          </section>
        </div>
      </template>

      <template #footer>
        <div class="flex justify-end">
          <Button label="关闭" icon="pi pi-times" severity="secondary" outlined @click="centerDialogVisible = false" />
        </div>
      </template>
    </Dialog>

    <Dialog v-model:visible="openAdd" header="新增贷款产品" modal class="w-full max-w-2xl">
      <AddLoanProduct />
    </Dialog>

    <Dialog v-model:visible="applyDialogVisible" header="申请贷款" modal class="w-full max-w-xl">
      <template v-if="selectedProduct">
        <div class="space-y-6">
          <div class="border-b border-slate-200 pb-4">
            <h3 class="text-lg font-semibold text-slate-900">{{ selectedProduct.productName }}</h3>
            <p class="text-sm text-slate-500">产品编号：{{ selectedProduct.productId }}</p>
          </div>
          <div class="grid gap-4 md:grid-cols-2">
            <div class="space-y-2">
              <span class="text-sm font-medium text-slate-700">贷款金额 (元)</span>
              <InputNumber
                v-model="applyAmount"
                mode="currency"
                currency="CNY"
                locale="zh-CN"
                :min="0"
                showButtons
                class="w-full"
                inputClass="w-full"
                placeholder="请输入贷款金额"
              />
            </div>
            <div class="space-y-2">
              <span class="text-sm font-medium text-slate-700">贷款期限 (月)</span>
              <InputNumber
                v-model="applyTerm"
                :min="1"
                :max="60"
                showButtons
                class="w-full"
                inputClass="w-full"
                placeholder="请输入贷款期限"
              />
            </div>
          </div>
        </div>
      </template>

      <template #footer>
        <div class="flex justify-end gap-3">
          <Button label="取消" severity="secondary" outlined @click="applyDialogVisible = false" />
          <Button label="确认申请" icon="pi pi-check" @click="handleApplyLoan()" />
        </div>
      </template>
    </Dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'nuxt/app'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Card from 'primevue/card'
import Tag from 'primevue/tag'
import InputNumber from 'primevue/inputnumber'
import MultiSelect from 'primevue/multiselect'
import Dropdown from 'primevue/dropdown'
import Message from 'primevue/message'
import Paginator from 'primevue/paginator'
import { useToast } from 'primevue/usetoast'
import AddLoanProduct from '~/components/addLoanProduct.vue'
import type { AgriculturalLoanProduct } from '~/types/loanProduct'
import { getLoanProductList, useApplyLoan } from '~/composables/getLoanProduct'
import { useLoanStore } from '~/utils/loanProductStore'
import { useUserStore } from '~/utils/userStore'

definePageMeta({ layout: 'home-page-layout' })

const router = useRouter()
const route = useRoute()
const toast = useToast()

const rows = ref(9)
const current = ref(1)
const applyAmount = ref<number | null>(null)
const applyTerm = ref<number | null>(null)

const openAdd = ref(false)
const centerDialogVisible = ref(false)
const applyDialogVisible = ref(false)
const selectedProduct = ref<AgriculturalLoanProduct | null>(null)

const amountFilter = ref<number | null>(null)
const selectedPurposes = ref<string[]>([])
const collateralFilter = ref<string | null>(null)
const creditFilter = ref<string | null>(null)

const defaultProductImage = 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png'

const productStore = useLoanStore()

const orders = computed(() => (Array.isArray(productStore.orderList) ? productStore.orderList : []))

const filteredProducts = computed(() => {
  return orders.value.filter((product) => {
    if (amountFilter.value !== null) {
      const amount = amountFilter.value
      if (amount < product.loanAmountRange.min || amount > product.loanAmountRange.max) {
        return false
      }
    }

    if (collateralFilter.value && product.eligibility.collateralRequirements !== collateralFilter.value) {
      return false
    }

    if (creditFilter.value && product.eligibility.creditRequirement !== creditFilter.value) {
      return false
    }

    if (selectedPurposes.value.length) {
      const mismatch = selectedPurposes.value.some((key) => (product.supportedPurposes as Record<string, boolean | undefined>)[key] === false)
      if (mismatch) return false
    }

    return true
  })
})

const resetFilters = () => {
  amountFilter.value = null
  selectedPurposes.value = []
  collateralFilter.value = null
  creditFilter.value = null
}

const createDefaultProduct = (): AgriculturalLoanProduct => ({
  productId: 'default-001',
  productName: '农业小额贷款',
  productAvatar: defaultProductImage,
  financialInstitution: {
    id: '001',
    name: '农业银行',
    customerService: '1234567890'
  },
  loanAmountRange: {
    min: 10000,
    max: 500000
  },
  interestRate: {
    type: 0,
    finalRate: 0.05,
    discountDescription: '首年利率优惠0.5%'
  },
  loanTerm: {
    minMonths: 6,
    maxMonths: 36
  },
  eligibility: {
    minOperatingYears: 1,
    creditRequirement: '近2年内无不良信用记录，当前无逾期',
    collateralRequirements: '无需抵押'
  },
  supportedPurposes: {
    production: true,
    equipment: true,
    land: false,
    operating: true,
    infrastructure: false
  },
  estimatedTime: '3-5个工作日',
  updateTime: new Date(),
  effectiveDate: new Date()
})

const loadLoanProducts = async (page = 1, pageSize = rows.value) => {
  const data = await getLoanProductList(page, pageSize)
  if (data) {
    productStore.setPaginationInfo(data.total ?? 0, data.page ?? page, data.pageSize ?? pageSize, data.hasmore ?? false)
    productStore.setOrder(data.loanProductList && data.loanProductList.length ? data.loanProductList : [createDefaultProduct()])
  }
}

onMounted(async () => {
  await loadLoanProducts(current.value, rows.value)
})

const handlePageChange = async (event: { page: number; rows: number }) => {
  current.value = event.page + 1
  rows.value = event.rows
  await loadLoanProducts(current.value, rows.value)
}

const showProductDetail = (product: AgriculturalLoanProduct) => {
  selectedProduct.value = product
  centerDialogVisible.value = true
}

const openApplyDialog = (product: AgriculturalLoanProduct) => {
  selectedProduct.value = product
  applyAmount.value = null
  applyTerm.value = null
  applyDialogVisible.value = true
}

const validateApplyInput = (product: AgriculturalLoanProduct, amount: number, term: number) => {
  if (amount < product.loanAmountRange.min || amount > product.loanAmountRange.max) {
    toast.add({ severity: 'warn', summary: '金额错误', detail: `申请金额需在 ${product.loanAmountRange.min} - ${product.loanAmountRange.max} 元之间`, life: 3000 })
    return false
  }
  if (term < product.loanTerm.minMonths || term > product.loanTerm.maxMonths) {
    toast.add({ severity: 'warn', summary: '期限错误', detail: `贷款期限需在 ${product.loanTerm.minMonths} - ${product.loanTerm.maxMonths} 个月之间`, life: 3000 })
    return false
  }
  return true
}

const handleApplyLoan = async () => {
  if (!selectedProduct.value) return
  const amount = applyAmount.value ?? 0
  const term = applyTerm.value ?? 0
  const roundedTerm = Math.floor(term)

  if (!amount || !roundedTerm) {
    toast.add({ severity: 'warn', summary: '提醒', detail: '请填写完整的贷款金额和期限', life: 3000 })
    return
  }

  if (!validateApplyInput(selectedProduct.value, amount, roundedTerm)) {
    toast.add({ severity: 'error', summary: '申请失败', detail: '请填写正确的贷款金额和期限', life: 3000 })
    console.log('贷款申请输入验证未通过')
    return
  }

  try {
    console.log('提交贷款申请，产品ID=', selectedProduct.value.productId, '金额=', amount, '期限=', roundedTerm)
    const userStore = useUserStore()
    const result = await useApplyLoan(selectedProduct.value.productId, userStore.userId, amount, roundedTerm)
    if (result) {
      toast.add({ severity: 'success', summary: '申请成功', detail: '已提交贷款申请，请耐心等待审核', life: 3000 })
      applyDialogVisible.value = false
    }
  } catch (error) {
    console.error('贷款申请失败:', error)
    toast.add({ severity: 'error', summary: '申请失败', detail: '提交申请时出现问题，请稍后重试', life: 3000 })
  }
}

const options = [
  { value: 'production', label: '农业生产' },
  { value: 'equipment', label: '设备购置' },
  { value: 'land', label: '土地流转/租赁' },
  { value: 'operating', label: '经营周转' },
  { value: 'infrastructure', label: '设施建设' }
]

const danbaoOptions = [
  { value: '无需抵押', label: '无需抵押' },
  { value: '农村宅基地使用权', label: '农村宅基地使用权' },
  { value: '农业设施', label: '农业设施' },
  { value: '机械设备', label: '机械设备' },
  { value: '温室大棚', label: '温室大棚' },
  { value: '定期存单', label: '定期存单' },
  { value: '保险保单', label: '保险保单' },
  { value: '应收账款', label: '应收账款' },
  { value: '政府风险补偿基金', label: '政府风险补偿基金' },
  { value: '融资担保公司', label: '融资担保公司' },
  { value: '龙头企业担保', label: '龙头企业担保' },
  { value: '合作社联保', label: '合作社联保' },
  { value: '土地经营权', label: '土地经营权' },
  { value: '养殖水面使用权', label: '养殖水面使用权' },
  { value: '林权', label: '林权' },
  { value: '农产品期货仓单', label: '农产品期货仓单' },
  { value: '活体抵押（牲畜、水产等）', label: '活体抵押（牲畜、水产等）' },
  { value: '知识产权（农产品品牌、专利等）', label: '知识产权（农产品品牌、专利等）' }
]

const trustOptions = [
  {
    value: '近2年内无不良信用记录，当前无逾期',
    label: '严格',
    description: '近2年内无不良信用记录，当前无逾期'
  },
  {
    value: '近1年无30天以上逾期，累计逾期不超过3次',
    label: '标准',
    description: '近1年无30天以上逾期，累计逾期不超过3次'
  },
  {
    value: '当前无重大不良记录，轻微逾期已结清可接受',
    label: '宽松',
    description: '当前无重大不良记录，轻微逾期已结清可接受'
  },
  {
    value: '无恶意逃废债记录，政府增信项目可适当放宽',
    label: '政府增信',
    description: '无恶意逃废债记录，政府增信项目可适当放宽'
  },
  {
    value: '近5年内无任何不良信用记录，征信完美',
    label: '非常严格',
    description: '近5年内无任何不良信用记录，征信完美'
  },
  {
    value: '接受近期有轻微逾期，重点关注经营状况和还款能力',
    label: '灵活',
    description: '接受近期有轻微逾期，重点关注经营状况和还款能力'
  }
]
</script>

<style scoped></style>
