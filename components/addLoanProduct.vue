<template>
  <!-- 添加产品 -->
    <el-form :model="form" ref="ruleFormRef" :rules="rules" label-width="120px">
      <!-- 产品基础信息 -->
      <el-divider content-position="left">产品基础信息</el-divider>

      <el-form-item label="产品名称" prop="productName">
        <el-input
          v-model="form.productName"
          placeholder="请输入产品名称"
          autocomplete="off"
        />
      </el-form-item>

      <el-form-item label="产品头像" prop="productAvatar">
        <el-input
          v-model="form.productAvatar"
          placeholder="请输入产品头像URL"
          autocomplete="off"
        />
      </el-form-item>

      <!-- 机构信息 -->
      <el-divider content-position="left">机构信息</el-divider>

      <el-form-item label="机构ID" prop="financialInstitution.id">
        <el-input
          v-model="form.financialInstitution.id"
          placeholder="请输入机构ID"
          autocomplete="off"
        />
      </el-form-item>

      <el-form-item label="机构全称" prop="financialInstitution.name">
        <el-input
          v-model="form.financialInstitution.name"
          placeholder="请输入机构全称"
          autocomplete="off"
        />
      </el-form-item>

      <el-form-item
        label="客服电话"
        prop="financialInstitution.customerService"
      >
        <el-input
          v-model="form.financialInstitution.customerService"
          placeholder="请输入客服电话"
          autocomplete="off"
        />
      </el-form-item>

      <!-- 融资条款 -->
      <el-divider content-position="left">融资条款</el-divider>

      <el-form-item label="贷款金额范围">
        <el-col :span="11">
          <el-form-item prop="loanAmountRange.min">
            <el-input-number
              v-model="form.loanAmountRange.min"
              :min="0"
              placeholder="最低贷款额"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="2" class="text-center">-</el-col>
        <el-col :span="11">
          <el-form-item prop="loanAmountRange.max">
            <el-input-number
              v-model="form.loanAmountRange.max"
              :min="0"
              placeholder="最高贷款额"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
      </el-form-item>

      <el-form-item label="利率类型" prop="interestRate.type">
        <el-radio-group v-model="form.interestRate.type">
          <el-radio :label="0">固定利率</el-radio>
          <el-radio :label="1">浮动利率</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="执行利率(%)" prop="interestRate.finalRate">
        <el-input-number
          v-model="form.interestRate.finalRate"
          :min="0"
          :max="100"
          :precision="2"
          placeholder="请输入利率"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item
        label="利率优惠说明"
        prop="interestRate.discountDescription"
      >
        <el-input
          v-model="form.interestRate.discountDescription"
          type="textarea"
          placeholder="请输入利率优惠说明"
          autocomplete="off"
        />
      </el-form-item>

      <!-- 贷款期限 -->
      <el-divider content-position="left">贷款期限</el-divider>

      <el-form-item label="贷款期限(月)">
        <el-col :span="11">
          <el-form-item prop="loanTerm.minMonths">
            <el-input-number
              v-model="form.loanTerm.minMonths"
              :min="1"
              placeholder="最短期限"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="2" class="text-center">-</el-col>
        <el-col :span="11">
          <el-form-item prop="loanTerm.maxMonths">
            <el-input-number
              v-model="form.loanTerm.maxMonths"
              :min="1"
              placeholder="最长期限"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
      </el-form-item>

      <!-- 适用条件 -->
      <el-divider content-position="left">适用条件</el-divider>

      <el-form-item label="最低经营年限" prop="eligibility.minOperatingYears">
        <el-input-number
          v-model="form.eligibility.minOperatingYears"
          :min="0"
          placeholder="请输入最低经营年限"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="征信要求" prop="eligibility.creditRequirement">
        <el-select
          v-model="form.eligibility.creditRequirement"
          placeholder="请选择征信要求"
          style="width: 100%"
        >
          <el-option
            label="近2年内无不良信用记录，当前无逾期"
            value="近2年内无不良信用记录，当前无逾期"
          />
          <el-option
            label="近1年无30天以上逾期，累计逾期不超过3次"
            value="近1年无30天以上逾期，累计逾期不超过3次"
          />
          <el-option
            label="当前无重大不良记录，轻微逾期已结清可接受"
            value="当前无重大不良记录，轻微逾期已结清可接受"
          />
          <el-option
            label="无恶意逃废债记录，政府增信项目可适当放宽"
            value="无恶意逃废债记录，政府增信项目可适当放宽"
          />
          <el-option
            label="近5年内无任何不良信用记录，征信完美"
            value="近5年内无任何不良信用记录，征信完美"
          />
          <el-option
            label="接受近期有轻微逾期，重点关注经营状况和还款能力"
            value="接受近期有轻微逾期，重点关注经营状况和还款能力"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="担保要求" prop="eligibility.collateralRequirements">
        <el-select
          v-model="form.eligibility.collateralRequirements"
          placeholder="请选择担保要求"
          style="width: 100%"
        >
          <el-option label="无需抵押" value="无需抵押" />
          <el-option label="农村宅基地使用权" value="农村宅基地使用权" />
          <el-option label="农业设施" value="农业设施" />
          <el-option label="机械设备" value="机械设备" />
          <el-option label="温室大棚" value="温室大棚" />
          <el-option label="定期存单" value="定期存单" />
          <el-option label="保险保单" value="保险保单" />
          <el-option label="应收账款" value="应收账款" />
          <el-option label="政府风险补偿基金" value="政府风险补偿基金" />
          <el-option label="融资担保公司" value="融资担保公司" />
          <el-option label="龙头企业担保" value="龙头企业担保" />
          <el-option label="合作社联保" value="合作社联保" />
          <el-option label="土地经营权" value="土地经营权" />
          <el-option label="养殖水面使用权" value="养殖水面使用权" />
          <el-option label="林权" value="林权" />
          <el-option label="农产品期货仓单" value="农产品期货仓单" />
          <el-option
            label="活体抵押（牲畜、水产等）"
            value="活体抵押（牲畜、水产等）"
          />
          <el-option
            label="知识产权（农产品品牌、专利等）"
            value="知识产权（农产品品牌、专利等）"
          />
        </el-select>
      </el-form-item>

      <!-- 支持的业务范围 -->
      <el-divider content-position="left">支持的业务范围</el-divider>

      <el-form-item label="贷款用途" prop="supportedPurposes">
        <el-checkbox-group v-model="supportedPurposesList">
          <el-checkbox label="production">农业生产</el-checkbox>
          <el-checkbox label="equipment">设备购置</el-checkbox>
          <el-checkbox label="land">土地流转/租赁</el-checkbox>
          <el-checkbox label="operating">经营周转</el-checkbox>
          <el-checkbox label="infrastructure">设施建设</el-checkbox>
        </el-checkbox-group>
      </el-form-item>

      <!-- 状态信息 -->
      <el-divider content-position="left">状态信息</el-divider>

      <el-form-item label="预计审批时间" prop="estimatedTime">
        <el-input
          v-model="form.estimatedTime"
          placeholder="请输入预计审批时间"
          autocomplete="off"
        />
      </el-form-item>

      <el-form-item label="生效日期" prop="effectiveDate">
        <el-date-picker
          v-model="form.effectiveDate"
          type="date"
          placeholder="选择生效日期"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="失效日期" prop="expiryDate">
        <el-date-picker
          v-model="form.expiryDate"
          type="date"
          placeholder="选择失效日期（可选）"
          style="width: 100%"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="resetForm(ruleFormRef)">
          重置
        </el-button>
        <el-button type="primary" @click="formSubmit(ruleFormRef)">
          保存
        </el-button>
      </el-form-item>
    </el-form>
</template>

<script setup lang="ts">
import {addLoanProduct} from "~/composables/useAddLoanproduct"
import type { FormInstance, FormRules } from 'element-plus'
import type { AgriculturalLoanProduct } from "~/types/loanProduct";

const formLabelWidth = "140px";
const addPro = ref(false);
// 表单数据
const form = reactive<AgriculturalLoanProduct>({
  // 产品基础信息
  productId: "",
  productName: "",
  productAvatar: "",

  // 机构信息
  financialInstitution: {
    id: "",
    name: "",
    customerService: "",
  },

  // 融资条款
  loanAmountRange: {
    min: 0,
    max: 0,
  },

  interestRate: {
    type: 0,
    finalRate: 0,
    discountDescription: "",
  },

  // 贷款期限
  loanTerm: {
    minMonths: 0,
    maxMonths: 0,
  },

  // 适用条件
  eligibility: {
    minOperatingYears: 0,
    creditRequirement: "",
    collateralRequirements: "",
  },

  // 支持的业务范围
  supportedPurposes: {
    production: false,
    equipment: false,
    land: false,
    operating: false,
    infrastructure: false,
  },

  // 状态信息
  estimatedTime: "",
  updateTime: new Date(),
  effectiveDate: new Date(),
  expiryDate: undefined,
});
const ruleFormRef = ref<FormInstance>()

// 将支持的用途从布尔值转换为数组，以便在表单中使用
const supportedPurposesList = computed({
  get() {
    const purposes: string[] = [];
    Object.keys(form.supportedPurposes).forEach((key) => {
      if ((form.supportedPurposes as any)[key]) {
        purposes.push(key);
      }
    });
    return purposes;
  },
  set(newValues) {
    // 重置所有值为 false
    Object.keys(form.supportedPurposes).forEach((key) => {
      (form.supportedPurposes as any)[key] = false;
    });

    // 设置选中的值为 true
    newValues.forEach((value) => {
      if (form.supportedPurposes.hasOwnProperty(value)) {
        (form.supportedPurposes as any)[value] = true;
      }
    });
  },
});

const resetForm = (formEl: FormInstance | undefined) => {
  addPro.value = false
  if (!formEl) return
  formEl.resetFields()
}
const formSubmit = async (formEl: FormInstance | undefined) => {
  addPro.value = false
  if (!formEl) return
   await formEl.validate(async (valid, fields) => {
    if (valid) {
      await addLoanProduct(form)
      console.log('submit!')
    } else {
      console.log('error submit!', fields)
    }
  })
}

const rules = reactive<FormRules<AgriculturalLoanProduct>>({
  // 产品基础信息
  productName: [
    { required: true, message: '请输入产品名称', trigger: 'blur' }
  ],
  productAvatar: [
    { required: true, message: '请输入产品头像URL', trigger: 'blur' }
  ],
  
  // 机构信息
  'financialInstitution.id': [
    { required: true, message: '请输入机构ID', trigger: 'blur' }
  ],
  'financialInstitution.name': [
    { required: true, message: '请输入机构全称', trigger: 'blur' }
  ],
  'financialInstitution.customerService': [
    { required: true, message: '请输入客服电话', trigger: 'blur' }
  ],
  
  // 融资条款
  'loanAmountRange.min': [
    { required: true, message: '请输入最低贷款额', trigger: 'blur' },
    { type: 'number', min: 0, message: '贷款额不能小于0', trigger: 'blur' }
  ],
  'loanAmountRange.max': [
    { required: true, message: '请输入最高贷款额', trigger: 'blur' },
    { type: 'number', min: 0, message: '贷款额不能小于0', trigger: 'blur' },
    { 
      validator: (rule: any, value : any, callback : any) => {
        if (value <= form.loanAmountRange.min) {
          callback(new Error('最高贷款额必须大于最低贷款额'))
        } else {
          callback()
        }
      }, 
      trigger: 'blur' 
    }
  ],
  
  'interestRate.type': [
    { required: true, message: '请选择利率类型', trigger: 'change' }
  ],
  'interestRate.finalRate': [
    { required: true, message: '请输入执行利率', trigger: 'blur' },
    { type: 'number', min: 0, max: 100, message: '利率必须在0-100之间', trigger: 'blur' }
  ],
  
  // 贷款期限
  'loanTerm.minMonths': [
    { required: true, message: '请输入最短期限', trigger: 'blur' },
    { type: 'number', min: 1, message: '期限必须大于0', trigger: 'blur' }
  ],
  'loanTerm.maxMonths': [
    { required: true, message: '请输入最长期限', trigger: 'blur' },
    { type: 'number', min: 1, message: '期限必须大于0', trigger: 'blur' },
    { 
      validator: (rule : any, value : any, callback : any) => {
        if (value < form.loanTerm.minMonths) {
          callback(new Error('最长期限必须大于最短期限'))
        } else {
          callback()
        }
      }, 
      trigger: 'blur' 
    }
  ],
  
  // 适用条件
  'eligibility.minOperatingYears': [
    { required: true, message: '请输入最低经营年限', trigger: 'blur' },
    { type: 'number', min: 0, message: '年限不能小于0', trigger: 'blur' }
  ],
  'eligibility.creditRequirement': [
    { required: true, message: '请选择征信要求', trigger: 'change' }
  ],
  'eligibility.collateralRequirements': [
    { required: true, message: '请选择担保要求', trigger: 'change' }
  ],
  
  // 状态信息
  estimatedTime: [
    { required: true, message: '请输入预计审批时间', trigger: 'blur' }
  ],
  effectiveDate: [
    { required: true, message: '请选择生效日期', trigger: 'change' },
  ],
  expiryDate: [
    // 失效日期是可选的，所以没有 required: true
    { 
      validator: (rule : any, value : any, callback : any) => {
        // 如果填写了失效日期，必须大于生效日期
        if (value && form.effectiveDate && value <= form.effectiveDate) {
          callback(new Error('失效日期必须大于生效日期'))
        } else {
          callback()
        }
      }, 
      trigger: 'change' 
    }
  ],
  
  // 支持的业务范围 - 至少选择一个
  supportedPurposes: [
    { 
      validator: (rule : any, value : any, callback : any) => {
        const hasSelected = Object.values(value).some(v => v === true)
        if (!hasSelected) {
          callback(new Error('请至少选择一个业务范围'))
        } else {
          callback()
        }
      }, 
      trigger: 'change' 
    }
  ]
})
</script>

<style scoped>
</style>
