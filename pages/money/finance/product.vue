<template>
  <div>
    <div class="min-h-screen grid grid-cols-[300px_1fr]">
    <div
      class="w-[300px] h-screen overflow-y-auto flex flex-col items-center gap-8"
    >
      <el-menu
        default-active="product"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        style="width: 250px; height: 300px"
      >
        <el-menu-item index="product">
          <el-icon><Menu /></el-icon>
          <span class="text-lg">金融产品</span>
        </el-menu-item>
        <el-menu-item index="process">
          <el-icon><Setting /></el-icon>
          <span class="text-lg">进度查询</span>
        </el-menu-item>
      </el-menu>
    </div>

    <div class="overflow-y-auto col-start-2">
      <Button label="添加" @click="openAdd = true" severity="help" />
      <div class="flex flex-col gap-6">
        <div class="flex gap-4 mt-20 flex-wrap">
          <div class="w-200 flex gap-1 items-center">
            <span>请输入贷款金额:</span>
            <el-input
              v-model="input"
              style="width: 200px"
              :formatter="(value : any) => `￥ ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
              :parser="(value : any) => value.replace(/[^\d]/g, '')"
              maxlength="15"
            />
          </div>
          <div>
            <span>业务范围:</span>
            <el-select
              multiple
              v-model="value"
              clearable
              placeholder="Select"
              style="width: 250px"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </div>
          <div>
            <span>担保要求：</span>
            <el-select
              v-model="danbao"
              placeholder="Select"
              style="width: 240px"
            >
              <el-option
                v-for="item in danbaoOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </div>
          <div>
            <span>征信要求：</span>
            <el-select
              v-model="trust"
              placeholder="Select"
              style="width: 240px"
            >
              <el-option
                v-for="item in trustOptions"
                :key="item.value"
                :label="item.description"
                :value="item.value"
              />
            </el-select>
          </div>
        </div>
        <div class="grid grid-cols-3 gap-5 grid-rows-3 ">
          <div
            v-for="(product, index) in filteredProducts"
            :key="index"
            class="w-full overflow-auto"
            
          >
            <el-card style="max-width: 480px">
              <template #header>
                <div class="grid grid-cols-2 items-center gap-2">
                  <div class="text-xl font-bold">{{ product.productName }}</div>
                  <img
                    class="w-20"
                    :src="
                      product.productAvatar ||
                      `https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png`
                    "
                  />
                </div>
              </template>
              <span class="block mb-2"
                >贷款额度(元)：{{ product.loanAmountRange.min }} -
                {{ product.loanAmountRange.max }}</span
              >
              <div class="flex gap-2">
                <span>业务范围:</span>
                <el-tag
                  v-if="product.supportedPurposes.production"
                  type="success"
                  >农业生产</el-tag
                >
                <el-tag
                  v-if="product.supportedPurposes.equipment"
                  type="success"
                  >设备购置</el-tag
                >
                <el-tag v-if="product.supportedPurposes.land" type="success"
                  >土地流转/租赁</el-tag
                >
                <el-tag
                  v-if="product.supportedPurposes.operating"
                  type="success"
                  >经营周转</el-tag
                >
                <el-tag
                  v-if="product.supportedPurposes.infrastructure"
                  type="success"
                  >设施建设</el-tag
                >
              </div>

              <span class="block mb-2"
                >生效日期：{{ new Date(product.effectiveDate).toLocaleDateString()}}</span
              >
              <template #footer>
                <div class="grid grid-cols-2 items-center gap-2">
                  <el-button
                    type="info"
                    @click="
                      {
                        centerDialogVisible = true;
                        selectedProduct = product;
                      }
                    "
                    >查看详情</el-button
                  >
                  <el-button type="primary">申请</el-button>
                </div>
              </template>
            </el-card>
          </div>
        </div>
      </div>
    </div>
  </div>
  <el-dialog
    v-model="centerDialogVisible"
    title="产品详情"
    width="600"
    align-center
  >
    <div v-if="selectedProduct" class="space-y-4">
      <!-- 产品基本信息 -->
      <div class="flex items-center gap-4 pb-4 border-b">
        <img
          :src="
            selectedProduct.productAvatar ||
            'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png'
          "
          alt="产品图片"
          class="w-20 h-20 object-cover rounded"
        />
        <div>
          <h3 class="text-xl font-semibold">
            {{ selectedProduct.productName }}
          </h3>
          <p class="text-gray-500">产品编号: {{ selectedProduct.productId }}</p>
        </div>
      </div>

      <!-- 银行信息 -->
      <div>
        <h4 class="font-medium mb-3 text-lg">银行信息</h4>
        <div class="grid grid-cols-2 gap-4 bg-gray-50 p-4 rounded">
          <div>
            <span
              >银行名称:{{ selectedProduct.financialInstitution.name }}</span
            >
          </div>
          <div>
            <span
              >客服电话:{{
                selectedProduct.financialInstitution.customerService
              }}</span
            >
          </div>
        </div>
      </div>
      <!-- 融资条款 -->
      <div>
        <h4 class="font-medium mb-3 text-lg">融资条款</h4>
        <div class="grid grid-cols-2 gap-4 bg-gray-50 p-4 rounded">
          <div>
            <span class="text-gray-600">贷款金额范围：</span>
            <span class="font-medium"
              >{{ selectedProduct.loanAmountRange.min }} -
              {{ selectedProduct.loanAmountRange.max }}</span
            >
          </div>
          <div>
            <span class="text-gray-600">年利率：</span>
            <span class="font-medium"
              >{{
                (selectedProduct.interestRate.finalRate * 100).toFixed(2)
              }}%</span
            >
            <span class="text-xs text-gray-500"
              >({{
                selectedProduct.interestRate.type === 0 ? "固定" : "浮动"
              }}利率)</span
            >
          </div>
          <div>
            <span class="text-gray-600">贷款期限：</span>
            <span class="font-medium"
              >{{ selectedProduct.loanTerm.minMonths }} -
              {{ selectedProduct.loanTerm.maxMonths }} 个月</span
            >
          </div>
          <div>
            <span class="text-gray-600">预计审批时间：</span>
            <span class="font-medium">{{ selectedProduct.estimatedTime }}</span>
          </div>
        </div>

        <div
          v-if="selectedProduct.interestRate.discountDescription"
          class="mt-2 p-3 bg-blue-50 rounded text-sm"
        >
          <span class="text-blue-700"
            >😲 利率优惠：{{
              selectedProduct.interestRate.discountDescription
            }}</span
          >
        </div>
      </div>

      <!-- 申请条件 -->
      <div>
        <h4 class="font-medium mb-3 text-lg">申请条件</h4>
        <div class="bg-gray-50 p-4 rounded space-y-2">
          <div>
            <span class="text-gray-600">最低经营年限：</span>
            <span>{{ selectedProduct.eligibility.minOperatingYears }} 年</span>
          </div>
          <div>
            <span class="text-gray-600">征信要求：</span>
            <span>{{ selectedProduct.eligibility.creditRequirement }}</span>
          </div>
          <div v-if="selectedProduct.eligibility.collateralRequirements">
            <span class="text-gray-600">担保要求：</span>
            <span>{{
              selectedProduct.eligibility.collateralRequirements
            }}</span>
          </div>
        </div>
      </div>

      <!-- 支持的业务范围 -->
      <div>
        <h4 class="font-medium mb-3 text-lg">支持的业务范围</h4>
        <div class="flex flex-wrap gap-2">
          <el-tag
            v-if="selectedProduct.supportedPurposes.production"
            type="success"
            >农业生产</el-tag
          >
          <el-tag
            v-if="selectedProduct.supportedPurposes.equipment"
            type="success"
            >设备购置</el-tag
          >
          <el-tag v-if="selectedProduct.supportedPurposes.land" type="success"
            >土地流转/租赁</el-tag
          >
          <el-tag
            v-if="selectedProduct.supportedPurposes.operating"
            type="success"
            >经营周转</el-tag
          >
          <el-tag
            v-if="selectedProduct.supportedPurposes.infrastructure"
            type="success"
            >设施建设</el-tag
          >
        </div>
      </div>

      <!-- 有效期信息 -->
     <!--  <div class="text-sm text-gray-500 pt-2 border-t">
        生效日期：{{ selectedProduct.effectiveDate.toLocaleDateString() }}
        <span v-if="selectedProduct.expiryDate">
          至 {{new Date(selectedProduct.expiryDate).toLocaleDateString() }} 
        </span>
        <span v-else> (长期有效) </span>
      </div> -->
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="centerDialogVisible = false">关闭</el-button>
      </div>
    </template>
  </el-dialog>
    <el-dialog v-model="openAdd" title="产品详情" width="600" align-center>
      <AddLoanProduct />
    </el-dialog>
<div class="flex justify-center mt-4">
    <el-pagination background :current-page="current" @current-change="handlePageChange"  layout="prev, pager, next" :total="productStore.total" />
  </div>
  </div>
</template>

<script setup lang="ts">
import type { AgriculturalLoanProduct } from "~/types/loanProduct";
import { reactive, ref, computed } from "vue";import { Setting } from "@element-plus/icons-vue";
definePageMeta({ layout: "home-page-layout" });
const current = ref(1)

// 加载数据的函数
const loadLoanProducts = async (page = 1, pageSize = 9) => {
  const data = await getLoanProductList(page, pageSize);
  
  if (data) {
    productStore.setPaginationInfo(
      data.total ?? 0, 
      data.page ?? page, 
      data.pageSize ?? pageSize, 
      data.hasmore ?? false
    );
    
    productStore.setOrder(data.loanProductList ?? [createDefaultProduct()]);
    console.log('加载产品列表:', data?.loanProductList);
  }
};
// 首次加载数据
onMounted(async () => {
  await loadLoanProducts(current.value, 9);
});
// 分页变化处理函数
const handlePageChange = async (newPage: number) => {
  current.value = newPage;
  await loadLoanProducts(newPage, 9);
};

const openAdd = ref(false);
const createDefaultProduct = (): AgriculturalLoanProduct => {
  return {
    productId: "default-001",
    productName: "农业小额贷款",
    productAvatar:
      "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
    financialInstitution: {
      id: "001",
      name: "农业银行",
      customerService: "1234567890",
    },
    loanAmountRange: {
      min: 10000,
      max: 500000,
    },
    interestRate: {
      type: 0,
      finalRate: 0.05,
      discountDescription: "首年利率优惠0.5%",
    },
    loanTerm: {
      minMonths: 6,
      maxMonths: 36,
    },
    eligibility: {
      minOperatingYears: 1,
      creditRequirement: "近2年内无不良信用记录，当前无逾期",
      collateralRequirements: "无需抵押",
    },
    supportedPurposes: {
      production: true,
      equipment: true,
      land: false,
      operating: true,
      infrastructure: false,
    },
    estimatedTime: "3-5个工作日",
    updateTime: new Date(),
    effectiveDate: new Date(),
  };
};
const productStore = useLoanStore();
const router = useRouter();
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  router.push(`/money/finance/${key}`);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const centerDialogVisible = ref(false);
const selectedProduct = ref<AgriculturalLoanProduct | null>(null);
const value = ref<string[]>([]);
const input = ref("");
const danbao = ref("");
const trust = ref("");
const filteredProducts = computed(() => {
  return productStore.orderList.filter(product => {
    // 金额筛选
    if (
      input.value !== "" &&
      (parseFloat(input.value) < product.loanAmountRange.min ||
        parseFloat(input.value) > product.loanAmountRange.max)
    ) {
      return false;
    }
    // 担保要求筛选
    if (
      danbao.value !== "" &&
      product.eligibility.collateralRequirements !== danbao.value
    ) {
      return false;
    }
    // 征信要求筛选
    if (
      trust.value !== "" &&
      product.eligibility.creditRequirement !== trust.value
    ) {
      return false;
    }
    // 业务范围筛选
    if (
      value.value.length > 0 &&
      value.value.some(
        (item) => (product.supportedPurposes as any)[item] === false
      )
    ) {
      return false;
    }
    return true;
  });
});
const options = [
  {
    value: "production",
    label: "农业生产",
  },
  {
    value: "equipment",
    label: "设备购置",
  },
  {
    value: "land",
    label: "土地流转/租赁",
  },
  {
    value: "operating",
    label: "经营周转",
  },
  {
    value: "infrastructure",
    label: "设施建设",
  },
];
const danbaoOptions = [
  {
    value: "无需抵押",
    label: "无需抵押",
  },
  {
    value: "农村宅基地使用权",
    label: "农村宅基地使用权",
  },
  {
    value: "农业设施",
    label: "农业设施",
  },
  {
    value: "机械设备",
    label: "机械设备",
  },
  {
    value: "温室大棚",
    label: "温室大棚",
  },
  {
    value: "定期存单",
    label: "定期存单",
  },
  {
    value: "保险保单",
    label: "保险保单",
  },
  {
    value: "应收账款",
    label: "应收账款",
  },
  {
    value: "政府风险补偿基金",
    label: "政府风险补偿基金",
  },
  {
    value: "融资担保公司",
    label: "融资担保公司",
  },
  {
    value: "龙头企业担保",
    label: "龙头企业担保",
  },
  {
    value: "合作社联保",
    label: "合作社联保",
  },
  {
    value: "土地经营权",
    label: "土地经营权",
  },
  {
    value: "养殖水面使用权",
    label: "养殖水面使用权",
  },
  {
    value: "林权",
    label: "林权",
  },
  {
    value: "农产品期货仓单",
    label: "农产品期货仓单",
  },
  {
    value: "活体抵押（牲畜、水产等）",
    label: "活体抵押（牲畜、水产等）",
  },
  {
    value: "知识产权（农产品品牌、专利等）",
    label: "知识产权（农产品品牌、专利等）",
  },
];
const trustOptions = [
  {
    value: "近2年内无不良信用记录，当前无逾期",
    label: "严格",
    description: "近2年内无不良信用记录，当前无逾期",
  },
  {
    value: "近1年无30天以上逾期，累计逾期不超过3次",
    label: "标准",
    description: "近1年无30天以上逾期，累计逾期不超过3次",
  },
  {
    value: "当前无重大不良记录，轻微逾期已结清可接受",
    label: "宽松",
    description: "当前无重大不良记录，轻微逾期已结清可接受",
  },
  {
    value: "无恶意逃废债记录，政府增信项目可适当放宽",
    label: "政府增信",
    description: "无恶意逃废债记录，政府增信项目可适当放宽",
  },
  {
    value: "近5年内无任何不良信用记录，征信完美",
    label: "非常严格",
    description: "近5年内无任何不良信用记录，征信完美",
  },
  {
    value: "接受近期有轻微逾期，重点关注经营状况和还款能力",
    label: "灵活",
    description: "接受近期有轻微逾期，重点关注经营状况和还款能力",
  },
];

</script>

<style scoped></style>
