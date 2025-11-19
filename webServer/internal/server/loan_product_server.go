package gorm

import (
	"errors"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
)

func GetLoanProductList(page, pageSize int) (string, *respond.LoanProductListRespond, int) {
	var total int64
	if res := dao.GormDB.Model(&model.LoanProduct{}).Count(&total); res.Error != nil {
		return "查询失败", nil, -1
	}
	productList := &respond.LoanProductListRespond{
		LoanProductList: make([]respond.LoanProductRespond, 0),
		Total:           total,
		Page:            page,
		PageSize:        pageSize,
		HasMore:         int64(page*pageSize) < total,
	}
	if total == 0 {
		return "获取成功", productList, 0
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	var products []model.LoanProduct
	if res := dao.GormDB.Model(&model.LoanProduct{}).Offset(offset).Order("update_time DESC").Limit(pageSize).Find(&products); res.Error != nil {
		return "查询失败", nil, -1
	}
	for _, product := range products {
		productRespond := respond.LoanProductRespond{
			ProductID:            product.ProductID,
			ProductName:          product.ProductName,
			ProductAvatar:        product.ProductAvatar,
			FinancialInstitution: product.FinancialInstitution,
			LoanAmountRange:      product.LoanAmountRange,
			InterestRate:         product.InterestRate,
			LoanTerm:             product.LoanTerm,
			Eligibility: struct {
				MinOperatingYears     int    `json:"minOperatingYears"`
				CreditRequirement     string `json:"creditRequirement"`
				CollateralRequirement string `json:"collateralRequirements"`
			}{
				MinOperatingYears:     product.Eligibility.MinOperatingYears,
				CreditRequirement:     product.Eligibility.CreditRequirement.String(),
				CollateralRequirement: product.Eligibility.CollateralRequirement.String(),
			},
			SupportedPurposes: product.SupportedPurposes,
			EstimatedTime:     product.EstimatedTime,
			UpdateTime:        product.UpdateTime,
			EffectiveDate:     product.EffectiveDate,
			ExpiryDate:        &product.EffectiveDate,
		}
		productList.LoanProductList = append(productList.LoanProductList, productRespond)
	}
	if len(productList.LoanProductList) == 0 {
		productList.HasMore = false
	} else {
		productList.HasMore = int64(page*pageSize) < total
	}
	return "获取成功", productList, 0
}

func AddLoanProduct(req request.AddLoanProductRequest) (string, int) {
	var product model.LoanProduct

	// 产品基础信息
	product.ProductID = req.ProductID
	product.ProductName = req.ProductName
	product.ProductAvatar = req.ProductAvatar

	// 机构信息
	product.FinancialInstitution.ID = req.FinancialInstitution.ID
	product.FinancialInstitution.Name = req.FinancialInstitution.Name
	product.FinancialInstitution.CustomerService = req.FinancialInstitution.CustomerService

	// 融资条款
	product.LoanAmountRange.Min = req.LoanAmountRange.Min
	product.LoanAmountRange.Max = req.LoanAmountRange.Max

	product.InterestRate.Type = req.InterestRate.Type
	product.InterestRate.FinalRate = req.InterestRate.FinalRate
	product.InterestRate.DiscountDescription = req.InterestRate.DiscountDescription

	product.LoanTerm.MinMonths = req.LoanTerm.MinMonths
	product.LoanTerm.MaxMonths = req.LoanTerm.MaxMonths

	// 适用条件 - 这里需要将字符串转换为枚举类型
	product.Eligibility.MinOperatingYears = req.Eligibility.MinOperatingYears

	// 将征信要求的字符串转换为枚举类型
	creditReq, err := StringToCreditRequirementType(req.Eligibility.CreditRequirement)
	if err != nil {
		return "无效的征信要求: " + err.Error(), -1
	}
	product.Eligibility.CreditRequirement = creditReq

	// 将担保要求的字符串转换为枚举类型
	collateralReq, err := StringToCollateralRequirementType(req.Eligibility.CollateralRequirement)
	if err != nil {
		return "无效的担保要求: " + err.Error(), -1
	}
	product.Eligibility.CollateralRequirement = collateralReq

	// 支持的业务范围
	product.SupportedPurposes.Production = req.SupportedPurposes.Production
	product.SupportedPurposes.Equipment = req.SupportedPurposes.Equipment
	product.SupportedPurposes.Land = req.SupportedPurposes.Land
	product.SupportedPurposes.Operating = req.SupportedPurposes.Operating
	product.SupportedPurposes.Infrastructure = req.SupportedPurposes.Infrastructure

	// 状态信息
	product.EstimatedTime = req.EstimatedTime
	product.UpdateTime = req.UpdateTime
	product.EffectiveDate = req.EffectiveDate
	product.ExpiryDate = req.ExpiryDate

	// 保存到数据库
	if res := dao.GormDB.Create(&product); res.Error != nil {
		return "保存失败: " + res.Error.Error(), -1
	}

	return "添加成功", 0
}

// 将字符串转换为征信要求枚举类型
func StringToCreditRequirementType(req string) (model.CreditRequirementType, error) {
	switch req {
	case "近2年内无不良信用记录，当前无逾期":
		return model.CreditNoBadRecord2Years, nil
	case "近1年无30天以上逾期，累计逾期不超过3次":
		return model.CreditNoOverdue30Days1Year, nil
	case "当前无重大不良记录，轻微逾期已结清可接受":
		return model.CreditNoMajorBadRecord, nil
	case "无恶意逃废债记录，政府增信项目可适当放宽":
		return model.CreditNoMaliciousDefault, nil
	case "近5年内无任何不良信用记录，征信完美":
		return model.CreditPerfect5Years, nil
	case "接受近期有轻微逾期，重点关注经营状况和还款能力":
		return model.CreditMinorOverdueAccepted, nil
	default:
		return 0, errors.New("未知的征信要求")
	}
}

// 将字符串转换为担保要求枚举类型
func StringToCollateralRequirementType(req string) (model.CollateralRequirementType, error) {
	switch req {
	case "无需抵押":
		return model.CollateralNone, nil
	case "农村宅基地使用权":
		return model.CollateralRuralHomestead, nil
	case "农业设施":
		return model.CollateralAgriculturalFacility, nil
	case "机械设备":
		return model.CollateralMachinery, nil
	case "温室大棚":
		return model.CollateralGreenhouse, nil
	case "定期存单":
		return model.CollateralTimeDeposit, nil
	case "保险保单":
		return model.CollateralInsurancePolicy, nil
	case "应收账款":
		return model.CollateralAccountsReceivable, nil
	case "政府风险补偿基金":
		return model.CollateralGovRiskFund, nil
	case "融资担保公司":
		return model.CollateralGuaranteeCompany, nil
	case "龙头企业担保":
		return model.CollateralLeadingEnterprise, nil
	case "合作社联保":
		return model.CollateralCooperativeGuarantee, nil
	case "土地经营权":
		return model.CollateralLandRights, nil
	case "养殖水面使用权":
		return model.CollateralWaterRights, nil
	case "林权":
		return model.CollateralForestRights, nil
	case "农产品期货仓单":
		return model.CollateralWarehouseReceipt, nil
	case "活体抵押（牲畜、水产等）":
		return model.CollateralLivestock, nil
	case "知识产权（农产品品牌、专利等）":
		return model.CollateralIntellectualProperty, nil
	default:
		return 0, errors.New("未知的担保要求")
	}
}
