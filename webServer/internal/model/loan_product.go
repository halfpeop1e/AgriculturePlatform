package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type LoanProduct struct {
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement"`

	// 产品基础信息
	ProductID     string `json:"productId" gorm:"column:product_id;type:varchar(100);not null;uniqueIndex"` // 产品唯一标识
	ProductName   string `json:"productName" gorm:"column:product_name;type:varchar(255);not null"`         // 产品名称
	ProductAvatar string `json:"productAvatar" gorm:"column:product_avatar;type:varchar(500)"`              // 产品头像URL

	// 机构信息 - 使用JSON类型存储
	FinancialInstitution FinancialInstitution `json:"financialInstitution" gorm:"type:jsonb;not null"`

	// 融资条款 - 使用JSON类型存储
	LoanAmountRange LoanAmountRange `json:"loanAmountRange" gorm:"type:jsonb;not null"`
	InterestRate    InterestRate    `json:"interestRate" gorm:"type:jsonb;not null"`
	LoanTerm        LoanTerm        `json:"loanTerm" gorm:"type:jsonb;not null"`

	// 适用条件 - 使用JSON类型存储
	Eligibility Eligibility `json:"eligibility" gorm:"type:jsonb;not null"`

	// 支持的业务范围 - 使用JSON类型存储
	SupportedPurposes SupportedPurposes `json:"supportedPurposes" gorm:"type:jsonb;not null"`

	// 状态信息
	EstimatedTime string     `json:"estimatedTime" gorm:"column:estimated_time;type:varchar(100);not null"` // 预计审批时间
	UpdateTime    time.Time  `json:"updateTime" gorm:"column:update_time;not null"`                         // 最后更新时间
	EffectiveDate time.Time  `json:"effectiveDate" gorm:"column:effective_date;not null"`                   // 生效日期
	ExpiryDate    *time.Time `json:"expiryDate" gorm:"column:expiry_date"`                                  // 失效日期，使用指针表示可为空
}

// 定义各个嵌套结构体，实现Scanner和Valuer接口以支持JSONB类型

type FinancialInstitution struct {
	ID              string `json:"id"`              // 机构ID
	Name            string `json:"name"`            // 机构全称
	CustomerService string `json:"customerService"` // 客服电话
}

// 实现Scanner接口
func (f *FinancialInstitution) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-bytes value into FinancialInstitution")
	}
	return json.Unmarshal(bytes, f)
}

// 实现Valuer接口
func (f FinancialInstitution) Value() (driver.Value, error) {
	return json.Marshal(f)
}

type LoanAmountRange struct {
	Min int64 `json:"min"` // 最低贷款额
	Max int64 `json:"max"` // 最高贷款额
}

func (l *LoanAmountRange) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-bytes value into LoanAmountRange")
	}
	return json.Unmarshal(bytes, l)
}

func (l LoanAmountRange) Value() (driver.Value, error) {
	return json.Marshal(l)
}

type InterestRate struct {
	Type                int8   `json:"type"`                // 0固定/1浮动利率
	FinalRate           string `json:"finalRate"`           // 最终执行利率，使用字符串存储避免精度问题
	DiscountDescription string `json:"discountDescription"` // 利率优惠说明
}

func (i *InterestRate) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-bytes value into InterestRate")
	}
	return json.Unmarshal(bytes, i)
}

func (i InterestRate) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type LoanTerm struct {
	MinMonths int `json:"minMonths"` // 最短期限(月)
	MaxMonths int `json:"maxMonths"` // 最长期限(月)
}

func (l *LoanTerm) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-bytes value into LoanTerm")
	}
	return json.Unmarshal(bytes, l)
}

func (l LoanTerm) Value() (driver.Value, error) {
	return json.Marshal(l)
}

type Eligibility struct {
	MinOperatingYears     int                       `json:"minOperatingYears"`      // 最低经营年限
	CreditRequirement     CreditRequirementType     `json:"creditRequirement"`      // 征信要求
	CollateralRequirement CollateralRequirementType `json:"collateralRequirements"` // 担保要求
}
type CollateralRequirementType int8
type CreditRequirementType int8

const (
	CreditNoBadRecord2Years    CreditRequirementType = 0 // 近2年内无不良信用记录，当前无逾期
	CreditNoOverdue30Days1Year CreditRequirementType = 1 // 近1年无30天以上逾期，累计逾期不超过3次
	CreditNoMajorBadRecord     CreditRequirementType = 2 // 当前无重大不良记录，轻微逾期已结清可接受
	CreditNoMaliciousDefault   CreditRequirementType = 3 // 无恶意逃废债记录，政府增信项目可适当放宽
	CreditPerfect5Years        CreditRequirementType = 4 // 近5年内无任何不良信用记录，征信完美
	CreditMinorOverdueAccepted CreditRequirementType = 5 // 接受近期有轻微逾期，重点关注经营状况和还款能力
)

// 征信要求的描述映射
var creditRequirementNames = map[CreditRequirementType]string{
	CreditNoBadRecord2Years:    "近2年内无不良信用记录，当前无逾期",
	CreditNoOverdue30Days1Year: "近1年无30天以上逾期，累计逾期不超过3次",
	CreditNoMajorBadRecord:     "当前无重大不良记录，轻微逾期已结清可接受",
	CreditNoMaliciousDefault:   "无恶意逃废债记录，政府增信项目可适当放宽",
	CreditPerfect5Years:        "近5年内无任何不良信用记录，征信完美",
	CreditMinorOverdueAccepted: "接受近期有轻微逾期，重点关注经营状况和还款能力",
}

// 实现 Scanner 接口
func (c *CreditRequirementType) Scan(value interface{}) error {
	if value == nil {
		*c = 0
		return nil
	}
	var val int64
	switch v := value.(type) {
	case int64:
		val = v
	case int:
		val = int64(v)
	case int8:
		val = int64(v)
	default:
		return errors.New("cannot scan non-integer value into CreditRequirementType")
	}
	req := CreditRequirementType(val)

	*c = req
	return nil
}

// 实现 Valuer 接口
func (c CreditRequirementType) Value() (driver.Value, error) {
	return int8(c), nil
}

// 获取征信要求的描述
func (c CreditRequirementType) String() string {
	if name, ok := creditRequirementNames[c]; ok {
		return name
	}
	return "未知征信要求"
}

// 定义所有可能的担保要求常量
const (
	CollateralNone                 CollateralRequirementType = 0  // 无需抵押
	CollateralRuralHomestead       CollateralRequirementType = 1  // 农村宅基地使用权
	CollateralAgriculturalFacility CollateralRequirementType = 2  // 农业设施
	CollateralMachinery            CollateralRequirementType = 3  // 机械设备
	CollateralGreenhouse           CollateralRequirementType = 4  // 温室大棚
	CollateralTimeDeposit          CollateralRequirementType = 5  // 定期存单
	CollateralInsurancePolicy      CollateralRequirementType = 6  // 保险保单
	CollateralAccountsReceivable   CollateralRequirementType = 7  // 应收账款
	CollateralGovRiskFund          CollateralRequirementType = 8  // 政府风险补偿基金
	CollateralGuaranteeCompany     CollateralRequirementType = 9  // 融资担保公司
	CollateralLeadingEnterprise    CollateralRequirementType = 10 // 龙头企业担保
	CollateralCooperativeGuarantee CollateralRequirementType = 11 // 合作社联保
	CollateralLandRights           CollateralRequirementType = 12 // 土地经营权
	CollateralWaterRights          CollateralRequirementType = 13 // 养殖水面使用权
	CollateralForestRights         CollateralRequirementType = 14 // 林权
	CollateralWarehouseReceipt     CollateralRequirementType = 15 // 农产品期货仓单
	CollateralLivestock            CollateralRequirementType = 16 // 活体抵押（牲畜、水产等）
	CollateralIntellectualProperty CollateralRequirementType = 17 // 知识产权（农产品品牌、专利等）
)

// 担保要求的描述映射
var collateralRequirementNames = map[CollateralRequirementType]string{
	CollateralNone:                 "无需抵押",
	CollateralRuralHomestead:       "农村宅基地使用权",
	CollateralAgriculturalFacility: "农业设施",
	CollateralMachinery:            "机械设备",
	CollateralGreenhouse:           "温室大棚",
	CollateralTimeDeposit:          "定期存单",
	CollateralInsurancePolicy:      "保险保单",
	CollateralAccountsReceivable:   "应收账款",
	CollateralGovRiskFund:          "政府风险补偿基金",
	CollateralGuaranteeCompany:     "融资担保公司",
	CollateralLeadingEnterprise:    "龙头企业担保",
	CollateralCooperativeGuarantee: "合作社联保",
	CollateralLandRights:           "土地经营权",
	CollateralWaterRights:          "养殖水面使用权",
	CollateralForestRights:         "林权",
	CollateralWarehouseReceipt:     "农产品期货仓单",
	CollateralLivestock:            "活体抵押（牲畜、水产等）",
	CollateralIntellectualProperty: "知识产权（农产品品牌、专利等）",
}

// 获取担保要求的描述
func (c CollateralRequirementType) String() string {
	if name, ok := collateralRequirementNames[c]; ok {
		return name
	}
	return "未知担保要求"
}

// 实现 Scanner 接口
func (c *CollateralRequirementType) Scan(value interface{}) error {
	if value == nil {
		*c = 0
		return nil
	}

	var val int64
	switch v := value.(type) {
	case int64:
		val = v
	case int:
		val = int64(v)
	case int8:
		val = int64(v)
	default:
		return errors.New("cannot scan non-integer value into CollateralRequirementType")
	}

	req := CollateralRequirementType(val)

	*c = req
	return nil
}

// 实现 Valuer 接口
func (c CollateralRequirementType) Value() (driver.Value, error) {
	return int8(c), nil
}

func (e *Eligibility) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-bytes value into Eligibility")
	}
	return json.Unmarshal(bytes, e)
}

func (e Eligibility) Value() (driver.Value, error) {
	return json.Marshal(e)
}

type SupportedPurposes struct {
	Production     bool `json:"production"`     // 农业生产
	Equipment      bool `json:"equipment"`      // 设备购置
	Land           bool `json:"land"`           // 土地流转/租赁
	Operating      bool `json:"operating"`      // 经营周转
	Infrastructure bool `json:"infrastructure"` // 设施建设
}

func (s *SupportedPurposes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-bytes value into SupportedPurposes")
	}
	return json.Unmarshal(bytes, s)
}

func (s SupportedPurposes) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// TableName 指定表名
func (LoanProduct) TableName() string {
	return "loan_products"
}
