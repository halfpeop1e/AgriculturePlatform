<template>
  <div>
    <div className="min-h-screen flex flex-col items-center pt-10 gap-4">
      <div
        class="bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center"
      >
        <el-image
          style="width: 200px; height: 240px"
          :src="farmer1"
          :fit="'cover'"
        />
        <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
          <span
            class="text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2"
            >融资难、流程繁，制约生产扩张?
          </span>
        </div>
      </div>
      <div
        class="bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center"
      >
        <el-image
          style="width: 200px; height: 240px"
          :src="farmer2"
          :fit="'cover'"
        />
        <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
          <span
            class="text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2"
            >抗风险能力弱，制约农业可持续发展？
          </span>
        </div>
      </div>
      <div
        v-on:click="nextPage"
        class="bg-slate-500/45 h-[240px] w-2/3 grid grid-cols-[220px_1fr] rounded-lg overflow-hidden items-center"
      >
        <el-image
          style="width: 200px; height: 240px"
          :src="man"
          :fit="'cover'"
        />
        <div class="grid grid-flow-col grid-rows-[66px_1fr_66px] h-full">
          <span
            class="text-balance font-bold text-4xl whitespace-pre-wrap justify-self-center row-start-2"
            >兄弟,把你的手机给我😄
          </span>
          <div class="flex justify-center">
            <el-icon :size="45"><ArrowDownBold /></el-icon>
          </div>
        </div>
      </div>
    </div>
    <div
      class="h-screen flex flex-col gap-10 items-center justify-center"
      ref="section2"
    >
      <div class="text-center text-3xl mb-11">
        Ave木吉卡金融，提供智能匹配贷款功能，为用户带来简约的贷款服务。
      </div>
      <div
        class="w-1/2 h-1/2 bg-slate-100 rounded-lg grid grid-rows-[100px_1fr] grid-flow-col overflow-hidden gap-1"
      >
        <div
          class="h-full min-h-[166px] bg-gradient-to-t from-transparent via-purple-500 to-pink-500"
        >
          <div
            class="flex w-full pt-5 pl-8 justify-between relative h-100px z-10"
          >
            <div class="flex">
              <el-avatar :size="75" :src="userStore.avatar || man" />
              <div class="pl-6 text-lg flex items-center">
                {{ userStore.userinfo.nickname }}
              </div>
            </div>
            <div class="absolute right-10 top-0">
              <el-image
                style="width: 100px; height: 100px"
                :src="muzimi"
                :fit="'cover'"
              />
            </div>
          </div>
        </div>
        <div
          class="bg-slate-50/45 mx-4 rounded-t-lg grid grid-cols-2 z-50 -mt-1"
        >
          <div class="flex flex-col items-center gap-2 pt-5">
            <span class="text-sm"> 我的贷款(元) </span>
            <span class="text-3xl">{{
              formatCurrency(loanData.loanedSum)
            }}</span>
            <el-button type="primary" size="large" @click="openLoanDialog">
              贷款
            </el-button>
          </div>
          <div class="flex flex-col items-center gap-2 pt-5">
            <span class="text-sm"> 我的待还(元) </span>
            <span class="text-3xl">{{ formatCurrency(loanData.loanSum) }}</span>
            <el-button type="primary" size="large" @click="openRepayDialog">
              还款
            </el-button>
          </div>
          <div class="col-span-2 px-10">
            <div class="flex flex-col gap-3 items-center">
              <span class="text-blue-500">{{ nextMonth % 12 }}月应还(元)</span>
              <span class="text-3xl">{{
                formatCurrency(nextMonthAmount)
              }}</span>
            </div>
          </div>
          <h1
            class="hover:cursor-pointer flex justify-center col-span-2 text-gray-600 hover:text-blue-500 transition-colors"
            @click="openAllBillsDialog"
          >
            您的账单已出 >
          </h1>
        </div>
      </div>
    </div>

    <!-- 弹窗 1：还款选择页面 -->
    <el-dialog v-model="repayDialogVisible" title="选择账单还款" width="600px">
      <el-table
        :data="pendingLoans"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="期数" width="120">
          <template #default="scope">
            {{ scope.row.year }}年{{ scope.row.month }}月
          </template>
        </el-table-column>
        <el-table-column prop="loanName" label="贷款名称" />
        <el-table-column prop="amount" label="金额(元)">
          <template #default="scope">
            {{ formatCurrency(scope.row.amount) }}
          </template>
        </el-table-column>
        <el-table-column prop="loanStatus" label="状态">
          <template #default>
            <el-tag type="danger">待还款</el-tag>
          </template>
        </el-table-column>
      </el-table>

      <div class="mt-4 flex justify-end items-center gap-4">
        <span class="text-sm text-gray-600">
          已选总额:
          <span class="text-xl font-bold text-red-500">{{
            formatCurrency(selectedRepayAmount)
          }}</span>
        </span>
        <el-button
          type="primary"
          :disabled="selectedRepayAmount <= 0"
          @click="handleRepayAction"
        >
          确认还款
        </el-button>
      </div>
    </el-dialog>

    <!-- 弹窗 2：申请贷款页面 -->
    <el-dialog v-model="loanDialogVisible" title="申请新贷款" width="500px">
      <el-form label-position="top">
        <el-form-item label="贷款用途">
          <el-input v-model="loanForm.purpose" placeholder="请输入贷款用途" />
        </el-form-item>
        <el-form-item label="申请金额（元）">
          <el-input-number
            v-model="loanForm.amount"
            :min="1000"
            :step="1000"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="分期月数">
          <el-select
            v-model="loanForm.term"
            placeholder="请选择期数"
            style="width: 100%"
          >
            <el-option label="3个月" :value="3" />
            <el-option label="6个月" :value="6" />
            <el-option label="9个月" :value="9" />
            <el-option label="12个月" :value="12" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="loanDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitLoanApplication"
            >提交申请</el-button
          >
        </span>
      </template>
    </el-dialog>

    <!-- 弹窗 3：AI分析贷款选择页面 -->
    <el-dialog
      v-model="aiAnalysisVisible"
      title="AI智能推荐贷款方案"
      width="600px"
    >
      <div class="space-y-4">
        <!-- AI分析结果 -->
        <div class="bg-blue-50 border-l-4 border-blue-400 p-4 rounded">
          <div class="flex items-center mb-2">
            <el-icon class="text-blue-600 mr-2" size="20"
              ><PictureRounded
            /></el-icon>
            <span class="font-semibold text-blue-800">AI分析结果</span>
          </div>
          <div class="text-gray-700 text-sm leading-relaxed">
            <p class="mb-2">
              根据您的申请信息（用途：<strong>{{
                loanForm.purpose || "未知"
              }}</strong
              >，金额：<strong>{{ loanForm.amount }}元</strong>，期限：<strong
                >{{ loanForm.term }}个月</strong
              >），AI为您推荐以下贷款方案：
            </p>
            <div class="mt-3 p-3 bg-white rounded border border-blue-200">
              <div class="grid grid-cols-2 gap-4 text-sm">
                <div>
                  <div class="font-semibold">方案名称</div>

                  <div>{{ loanPlans[selectedPlan]!.name }}</div>
                </div>

                <div>
                  <div class="font-semibold">利率</div>

                  <div>{{ loanPlans[selectedPlan]!.rate }}</div>
                </div>

                <div>
                  <div class="font-semibold">月供</div>

                  <div>{{ loanPlans[selectedPlan]!.monthlyPayment }}</div>
                </div>

                <div>
                  <div class="font-semibold">总利息</div>

                  <div>{{ loanPlans[selectedPlan]!.totalInterest }}</div>
                </div>
              </div>
            </div>
            <p class="mt-3 text-lg text-gray-600">
              {{ aiSuggestion }}
            </p>
            <p class="mt-3 text-xs text-gray-600">
              *
              此方案基于您的信用评级和申请条件智能匹配，具有审批快速、利率优惠等特点。
            </p>
          </div>
        </div>

        <!-- 可选方案列表 -->
        <div>
          <h4 class="font-semibold mb-3 text-gray-800">可选方案对比</h4>
          <div class="space-y-3">
            <div
              v-for="(plan, index) in loanPlans"
              :key="index"
              class="border rounded-lg p-3 cursor-pointer transition-all"
              :class="
                selectedPlan === index
                  ? 'border-blue-500 bg-blue-50'
                  : 'border-gray-200 hover:border-gray-300'
              "
              @click="selectedPlan = index"
            >
              <div class="flex justify-between items-start">
                <div>
                  <div class="font-medium flex items-center">
                    <el-tag
                      :type="
                        index === 0
                          ? 'success'
                          : index === 1
                          ? 'warning'
                          : 'info'
                      "
                      size="small"
                      class="mr-2"
                    >
                      {{ plan.tag }}
                    </el-tag>
                    {{ plan.name }}
                  </div>
                  <div class="text-sm text-gray-600 mt-1">
                    利率：{{ plan.rate }} | 月供：{{ plan.monthlyPayment }}元 |
                    总利息：{{ plan.totalInterest }}元
                  </div>
                  <div class="text-xs text-gray-500 mt-1">
                    {{ plan.description }}
                  </div>
                </div>
                <el-radio v-model="selectedPlan" :label="index"></el-radio>
              </div>
            </div>
          </div>
        </div>
      </div>

      <template #footer>
        <span class="dialog-footer flex justify-between">
          <el-button @click="aiAnalysisVisible = false">返回修改</el-button>
          <div class="space-x-2">
            <el-button @click="selectedPlan = 0">选择推荐方案</el-button>
            <el-button type="primary" @click="confirmLoanPlan">
              一键确认申请
            </el-button>
          </div>
        </span>
      </template>
    </el-dialog>

    <!-- 2. 新增弹窗：显示所有账单明细 -->
    <el-dialog v-model="allBillsVisible" title="全部账单明细" width="700px">
      <el-table
        :data="loanData.loanList"
        stripe
        style="width: 100%"
        height="400"
      >
        <!-- 时间列 -->
        <el-table-column label="账单周期" width="140">
          <template #default="scope">
            {{ scope.row.year }}年{{ scope.row.month }}月
          </template>
        </el-table-column>

        <!-- 名称列 -->
        <el-table-column prop="loanName" label="贷款项目" width="150" />

        <!-- 金额列 -->
        <el-table-column prop="amount" label="账单金额">
          <template #default="scope">
            <span class="font-bold">{{
              formatCurrency(scope.row.amount)
            }}</span>
          </template>
        </el-table-column>

        <!-- 状态列：根据状态显示不同颜色 -->
        <el-table-column prop="loanStatus" label="状态" align="center">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.loanStatus)">
              {{ getStatusText(scope.row.loanStatus) }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>

      <template #footer>
        <el-button @click="allBillsVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ layout: "home-page-layout" });
import man from "~/public/ioanImage/man.png";
import { ArrowDownBold, PictureRounded } from "@element-plus/icons-vue";
import farmer1 from "~/public/ioanImage/1762280990064.png";
import farmer2 from "~/public/ioanImage/1762281252724.png";
import muzimi from "~/public/ioanImage/1762350715242.png";
import { on } from "events";
import { month } from "@primeuix/themes/aura/datepicker";

const userStore = useUserStore();
const section2 = ref<HTMLElement | null>(null);

const nextPage = () => {
  if (section2.value) {
    section2.value.scrollIntoView({
      behavior: "smooth",
    });
  }
};

const now = new Date();
const nextMonth = now.getMonth() + 2;

const GetLoanData = async () => {
  const data = await addLoanProduct();
  console.log(data, "data");
  loanData.value = data ?? loanData.value;
  if(loanData.value.loanList == null){
    loanData.value.loanList = []
  }
};

onMounted(() => {
  GetLoanData();
});
// --- 模拟数据 (实际开发中请替换为 API 请求) ---
const loanData = ref<CheckMyLoanRespond>({
  loanedSum: 0,
  loanSum: 0,
  loanList: [],
});

// --- 状态控制 ---
const repayDialogVisible = ref(false);
const loanDialogVisible = ref(false);
const aiAnalysisVisible = ref(false);
const selectedLoans = ref<LoanOrder[]>([]);
const selectedPlan = ref(0); // 选中的贷款方案索引

// --- 计算属性：下个月份逻辑 ---
const currentDate = new Date();
// 获取下个月的月份 (JS getMonth 是 0-11，所以下个月是 getMonth()+1+1，取模处理跨年)
// 但为了简单展示，我们通常只展示数字
const nextMonthDisplay = computed(() => {
  let m = currentDate.getMonth() + 2; // 当前月(0-11) + 1变成人类月 + 1变成下个月
  if (m > 12) m = 1;
  return m;
});

// --- 计算属性：下月应还金额 ---
const nextMonthAmount = computed(() => {
  const nextM =
    currentDate.getMonth() + 2 > 12 ? 1 : currentDate.getMonth() + 2;
  const nextY =
    nextM === 1 ? currentDate.getFullYear() + 1 : currentDate.getFullYear();

  // 筛选：年份匹配 && 月份匹配 && 状态未还
  const bills = loanData.value.loanList.filter(
    (item) =>
      item.year === nextY && item.month === nextM && item.loanStatus !== "Paid" // 假设后端 'Paid' 表示已还
  );

  return bills.reduce((sum, item) => sum + item.amount, 0);
});

// --- 计算属性：还款弹窗里的待还列表 ---
const pendingLoans = computed(() => {
  return loanData.value.loanList.filter((item) => item.loanStatus !== "Paid");
});

// --- 计算属性：选中要还款的总额 ---
const selectedRepayAmount = computed(() => {
  return selectedLoans.value.reduce((sum, item) => sum + item.amount, 0);
});

// --- 方法 ---

// 格式化金额
const formatCurrency = (val: number) => {
  return val.toFixed(2);
};
// 打开还款弹窗
const openRepayDialog = () => {
  repayDialogVisible.value = true;
};

// 打开贷款弹窗
const openLoanDialog = () => {
  loanDialogVisible.value = true;
};

// 表格多选变化
const handleSelectionChange = (val: LoanOrder[]) => {
  selectedLoans.value = val;
};

// 执行还款
const handleRepayAction = async () => {
  // 这里调用后端还款接口
  console.log("正在还款订单:", selectedLoans.value);
  selectedLoans.value.map(async (item) => {
    await giveMoney(item.id);
  });

  ElMessage.success(`成功还款 ${formatCurrency(selectedRepayAmount.value)} 元`);
  repayDialogVisible.value = false;

  // 模拟前端更新数据（实际应重新请求API）
  selectedLoans.value.forEach((order) => {
    order.loanStatus = "Paid";
  });
  // 更新总待还金额 (简易模拟)
  loanData.value.loanSum -= selectedRepayAmount.value;
};

// 贷款申请表单数据
const loanForm = ref({
  purpose: "", // 贷款用途
  amount: 1000, // 申请金额（默认1000）
  term: 3, // 分期月数（默认12个月）
});
 const aiSuggestion = ref("Ai真是太好用了");
  const loanPlans = ref([
    {
      name: "AI智能推荐方案",
      tag: "推荐",
      rate: "年化4.8%",
      monthlyPayment: "336",
      totalInterest: "8.01",
      description: "基于您的信用评级智能匹配，审批快速，利率优惠",
      id:"1"
    },
  ]);

// 提交贷款申请
const submitLoanApplication = async () => {
  console.log("贷款申请数据:", loanForm.value);
  loanDialogVisible.value = false;

  // 显示AI分析弹窗
  aiAnalysisVisible.value = true;
  // 贷款方案数据
 
  const respong = await useAiSuggestion({
    purpose: loanForm.value.purpose,
    amount: loanForm.value.amount,
    term: loanForm.value.term,
  });
  console.log(respong, "respong");
  if (respong?.aiSuggestion == "当前未找到合适的产品") {
    // 关闭AI分析弹窗
    aiAnalysisVisible.value = false;

    // 清空表单
    loanForm.value = {
      purpose: "",
      amount: 1000,
      term: 3,
    };
    return
  }
  aiSuggestion.value = respong!.aiSuggestion;
  loanPlans.value = respong!.loanPlans;
};

// 确认贷款方案
const confirmLoanPlan = async () => {
  const selectedPlanData = loanPlans.value[selectedPlan.value];
  console.log("确认的贷款方案:", selectedPlanData);
  console.log("贷款申请信息:", loanForm.value);

  // 这里可以调用实际的贷款申请API
  ElMessage.success(
    `贷款申请已提交！您选择了${selectedPlanData!.name}，请等待审批`
  );
  const result = await useApplyLoan(selectedPlanData!.id, userStore.userId, loanForm.value.amount, loanForm.value.term)

  // 关闭AI分析弹窗
  aiAnalysisVisible.value = false;

  // 清空表单
  loanForm.value = {
    purpose: "",
    amount: 1000,
    term: 3,
  };

  // 重置选中的方案
  selectedPlan.value = 0;
};

const allBillsVisible = ref(false);

// --- 新增方法：打开弹窗 ---
const openAllBillsDialog = () => {
  allBillsVisible.value = true;
};

// --- 辅助方法：状态显示转换 ---
// 根据状态返回 Element Plus 的 Tag 颜色类型
const getStatusType = (status: string) => {
  // 假设 'Paid' 是已还，'Unpaid' 是未还，你可以根据后端实际返回字符串修改
  if (status === "Paid") return "success"; // 绿色
  if (status === "Overdue") return "danger"; // 红色（逾期）
  return "warning"; // 橙色（待还）
};

// 根据状态返回中文文本
const getStatusText = (status: string) => {
  if (status === "Paid") return "已还清";
  if (status === "Overdue") return "已逾期";
  return "待还款";
};
</script>

<style scoped></style>
