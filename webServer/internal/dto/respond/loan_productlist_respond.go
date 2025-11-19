package respond

import (
	"time"
)

type LoanProductListRespond struct {
	LoanProductList []LoanProductRespond `json:"loanProductList"`
	Total           int64                `json:"total"`
	Page            int                  `json:"page"`
	PageSize        int                  `json:"pageSize"`
	HasMore         bool                 `json:"hasMore"`
}

type LoanProductRespond struct {
	// 产品基础信息
	ProductID     string `json:"productId"`     // 产品唯一标识
	ProductName   string `json:"productName"`   // 产品名称
	ProductAvatar string `json:"productAvatar"` // 产品头像URL

	// 机构信息
	FinancialInstitution struct {
		ID              string `json:"id"`              // 机构ID
		Name            string `json:"name"`            // 机构全称
		CustomerService string `json:"customerService"` // 客服电话
	} `json:"financialInstitution"`

	// 融资条款
	LoanAmountRange struct {
		Min float64 `json:"min"` // 最低贷款额
		Max float64 `json:"max"` // 最高贷款额
	} `json:"loanAmountRange"`

	InterestRate struct {
		Type                int8    `json:"type"`                // 0固定/1浮动利率
		FinalRate           float64 `json:"finalRate"`           // 最终执行利率，使用字符串存储避免精度问题
		DiscountDescription string  `json:"discountDescription"` // 利率优惠说明
	} `json:"interestRate"`

	LoanTerm struct {
		MinMonths int `json:"minMonths"` // 最短期限(月)
		MaxMonths int `json:"maxMonths"` // 最长期限(月)
	} `json:"loanTerm"`

	// 适用条件
	Eligibility struct {
		MinOperatingYears     int    `json:"minOperatingYears"`      // 最低经营年限
		CreditRequirement     string `json:"creditRequirement"`      // 征信要求
		CollateralRequirement string `json:"collateralRequirements"` // 担保要求
	} `json:"eligibility"`

	// 支持的业务范围
	SupportedPurposes struct {
		Production     bool `json:"production"`     // 农业生产
		Equipment      bool `json:"equipment"`      // 设备购置
		Land           bool `json:"land"`           // 土地流转/租赁
		Operating      bool `json:"operating"`      // 经营周转
		Infrastructure bool `json:"infrastructure"` // 设施建设
	} `json:"supportedPurposes"`

	// 状态信息
	EstimatedTime string     `json:"estimatedTime"` // 预计审批时间
	UpdateTime    time.Time  `json:"updateTime"`    // 最后更新时间
	EffectiveDate time.Time  `json:"effectiveDate"` // 生效日期
	ExpiryDate    *time.Time `json:"expiryDate"`    // 失效日期，使用指针表示可为空
}
