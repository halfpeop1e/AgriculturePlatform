package request

import "time"

type AddLoanProductRequest struct {
	// 产品基础信息
	ProductID     string `json:"productId" binding:"required"`   // 产品唯一标识
	ProductName   string `json:"productName" binding:"required"` // 产品名称
	ProductAvatar string `json:"productAvatar"`                  // 产品头像URL

	// 机构信息
	FinancialInstitution struct {
		ID              string `json:"id" binding:"required"`              // 机构ID
		Name            string `json:"name" binding:"required"`            // 机构全称
		CustomerService string `json:"customerService" binding:"required"` // 客服电话
	} `json:"financialInstitution" binding:"required"`

	// 融资条款
	LoanAmountRange struct {
		Min int64 `json:"min" binding:"required,gt=0"` // 最低贷款额
		Max int64 `json:"max" binding:"required,gt=0"` // 最高贷款额
	} `json:"loanAmountRange" binding:"required"`

	InterestRate struct {
		Type                int8   `json:"type" binding:"required,oneof=0 1"` // 0固定/1浮动利率
		FinalRate           string `json:"finalRate" binding:"required"`      // 最终执行利率
		DiscountDescription string `json:"discountDescription"`               // 利率优惠说明
	} `json:"interestRate" binding:"required"`

	LoanTerm struct {
		MinMonths int `json:"minMonths" binding:"required,gt=0"` // 最短期限(月)
		MaxMonths int `json:"maxMonths" binding:"required,gt=0"` // 最长期限(月)
	} `json:"loanTerm" binding:"required"`

	// 适用条件
	Eligibility struct {
		MinOperatingYears     int    `json:"minOperatingYears" binding:"required,gt=0"` // 最低经营年限
		CreditRequirement     string `json:"creditRequirement" binding:"required"`      // 征信要求
		CollateralRequirement string `json:"collateralRequirements" binding:"required"` // 担保要求
	} `json:"eligibility" binding:"required"`

	// 支持的业务范围
	SupportedPurposes struct {
		Production     bool `json:"production"`     // 农业生产
		Equipment      bool `json:"equipment"`      // 设备购置
		Land           bool `json:"land"`           // 土地流转/租赁
		Operating      bool `json:"operating"`      // 经营周转
		Infrastructure bool `json:"infrastructure"` // 设施建设
	} `json:"supportedPurposes" binding:"required"`

	// 状态信息
	EstimatedTime string     `json:"estimatedTime" binding:"required"` // 预计审批时间
	UpdateTime    time.Time  `json:"updateTime" binding:"required"`    // 最后更新时间
	EffectiveDate time.Time  `json:"effectiveDate" binding:"required"` // 生效日期
	ExpiryDate    *time.Time `json:"expiryDate"`                       // 失效日期，使用指针表示可为空
}
