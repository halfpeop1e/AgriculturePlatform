package gorm

import (
	"errors"
	"fmt"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"go-agriculture/internal/util"
	"math"
	"strconv"
	"time"
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
	product.ProductID = util.GenerateUUID()
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

func ApllyLoanProduct(req request.ApplyLoanRequest) (string, int) {
	var apply model.Loan
	apply.Uuid = util.GenerateUUID()
	apply.ProductID = req.ProductID
	apply.UserId = req.UserID
	apply.Status = 0
	apply.Amount = req.Amount
	apply.Deadline = req.Deadline
	apply.CreatAt = time.Now()
	dao.GormDB.Create(&apply)
	return "申请成功", 0
}

func GetApplyList(userId string) (string, int, *[]respond.ApplyListRespond) {

	statusToS := func(status int) string {
		switch status {
		case 0:
			return "申请中"
		case 1:
			return "已通过"
		case 2:
			return "已拒绝"
		case 3:
			return "已撤销"
		case 4:
			return "已还款"
		default:
			return "未知状态"
		}
	}

	var responds []respond.ApplyListRespond
	var applys []model.Loan
	if res := dao.GormDB.Model(&model.Loan{}).Where("user_id = ?", userId).Find(&applys); res.Error != nil {
		return "查询失败", -1, nil
	}
	for _, apply := range applys {
		var product model.LoanProduct
		if res := dao.GormDB.Model(&model.LoanProduct{}).Where("product_id = ?", apply.ProductID).First(&product); res.Error != nil {
			continue
		}
		respondItem := respond.ApplyListRespond{
			Id:          apply.Id,
			Uuid:        apply.Uuid,
			UserId:      apply.UserId,
			ProductID:   apply.ProductID,
			Avatar:      product.ProductAvatar,
			ProductName: product.ProductName,
			Status:      statusToS(int(apply.Status)),
		}
		responds = append(responds, respondItem)
	}
	return "获取成功", 0, &responds
}

func AllowLoan(req request.AllowLoanRequest) (string, int) {
	var apply model.Loan
	if res := dao.GormDB.Model(&model.Loan{}).Where("uuid = ?", req.ApplyID).First(&apply); res.Error != nil {
		return "申请记录不存在", -1
	}
	if req.Allow {
		apply.Status = 1
		var product model.LoanProduct

		if res := dao.GormDB.Model(&model.LoanProduct{}).Where("product_id = ?", apply.ProductID).First(&product); res.Error != nil {
			return "查询失败", -1
		}

		// 计算等额本息月供
		monthlyRate := product.InterestRate.FinalRate / 100 / 12
		months := int(apply.Deadline)
		principal := apply.Amount

		var monthlyPayment float64
		if monthlyRate == 0 {
			monthlyPayment = principal / float64(months)
		} else {
			monthlyPayment = (principal * monthlyRate * math.Pow(1+monthlyRate, float64(months))) /
				(math.Pow(1+monthlyRate, float64(months)) - 1)
		}

		// 生成分期记录（等额本息）
		remainingPrincipal := principal
		for i := 0; i < int(apply.Deadline); i++ {
			var AllOrder model.AllOrder

			// 计算当期利息和本金
			var currentInterest float64
			if remainingPrincipal > 0 {
				currentInterest = remainingPrincipal * monthlyRate
			} else {
				currentInterest = 0
			}

			// 最后一期可能金额略有不同，确保总金额准确
			currentPrincipal := monthlyPayment - currentInterest
			if i == int(apply.Deadline)-1 {
				// 最后一期处理剩余本金
				currentPrincipal = remainingPrincipal
				currentPayment := currentPrincipal + currentInterest
				if math.Abs(currentPayment-monthlyPayment) > 0.01 {
					monthlyPayment = currentPayment
				}
			}

			// 计算年月
			currentTime := apply.CreatAt.AddDate(0, i, 0)
			year := currentTime.Year()
			month := currentTime.Month()

			newOrder := respond.LoanOrder{
				Year:     year,
				Month:    int(month),
				Amount:   monthlyPayment,
				LoanName: product.ProductName + "第" + strconv.Itoa(i+1) + "期",
			}

			newOrder.LoanStatus = "Unpaid"
			AllOrder.Amount = newOrder.Amount
			AllOrder.LoanName = newOrder.LoanName
			AllOrder.LoanStatus = newOrder.LoanStatus
			AllOrder.Month = newOrder.Month
			AllOrder.Year = newOrder.Year
			AllOrder.LoanId = apply.Uuid
			AllOrder.CreatAt = time.Now()
			dao.GormDB.Create(&AllOrder)

			// 更新剩余本金
			remainingPrincipal -= currentPrincipal
			if remainingPrincipal < 0.01 {
				remainingPrincipal = 0
			}
		}
	} else {
		apply.Status = 2
	}

	dao.GormDB.Save(&apply)
	return "更新成功", 0
}

func GetMyLoan(userId string) (string, int, *respond.CheckMyLoanRespond) {
	var responds respond.CheckMyLoanRespond
	var loanedSum float64 = 0
	var loanSum float64 = 0
	var loans []model.Loan
	if res := dao.GormDB.Model(&model.Loan{}).
		Where("user_id = ?", userId).
		Where("phone = ? OR phone = ?", 1, 4).Find(&loans); res.Error != nil {
		return "查询失败", -1, nil
	}
	for _, loan := range loans {
		var AllOrder []model.AllOrder
		if res := dao.GormDB.Model(&model.AllOrder{}).Where("loan_id = ?", loan.Uuid).Find(&AllOrder); res.Error != nil {
			return "查询失败", -1, nil
		}
		for _, order := range AllOrder {
			loanedSum += order.Amount
			if order.LoanStatus == "Unpaid" {
				loanSum += order.Amount
			}
			responds.LoanList = append(responds.LoanList, respond.LoanOrder{
				ID:         order.Id,
				Year:       order.Year,
				Month:      order.Month,
				Amount:     order.Amount,
				LoanName:   order.LoanName,
				LoanStatus: order.LoanStatus,
			})
		}
		responds.LoanedSum = loanedSum
		responds.LoanSum = loanSum
	}
	return "获取成功", 0, &responds
}

func GiveMoney(req request.GiveMoneyRequest) (string, int) {
	var order model.AllOrder

	if res := dao.GormDB.Model(&model.AllOrder{}).Where("id = ?", req.LoanId).First(&order); res.Error != nil {

		return "查询失败", -1

	}

	if order.LoanStatus == "Unpaid" {

		order.LoanStatus = "Paid"

		dao.GormDB.Save(&order)
		var orders []model.AllOrder
		if res := dao.GormDB.Model(&model.AllOrder{}).Where("loan_id = ?", order.LoanId).Find(&orders); res.Error != nil {
			return "查询失败", -1
		}
		for _, order := range orders {
			if order.LoanStatus == "Unpaid" {
				return "支付成功", 0
			}
		}
		var loan model.Loan
		if res := dao.GormDB.Model(&model.Loan{}).Where("uuid = ?", order.LoanId).First(&loan); res.Error != nil {
			return "查询失败", -1
		}
		loan.Status = 4
		dao.GormDB.Save(&loan)
		return "支付成功", 0

	} else {

		return "订单已支付", -1

	}
}

func AskQwen(req request.AiRequest) (string, int, *respond.AiRespond) {
	var products []model.LoanProduct
	var usefulproducts []model.LoanProduct
	if res := dao.GormDB.Model(&model.LoanProduct{}).Find(&products); res.Error != nil {
		return "查询失败", -1, nil
	}

	// 筛选符合条件的产品
	for _, product := range products {
		if product.LoanTerm.MinMonths <= req.Term && product.LoanTerm.MaxMonths >= req.Term &&
			product.LoanAmountRange.Min <= req.Amount && product.LoanAmountRange.Max >= req.Amount {
			usefulproducts = append(usefulproducts, product)
		}
	}

	if len(usefulproducts) == 0 {
		return "未找到合适的产品", -1, &respond.AiRespond{
			AiSuggestion: "未找到合适的产品",
			LoanPlans:    []respond.LoanPlan{},
		}
	}

	// 按利率排序（从低到高）
	for i := 0; i < len(usefulproducts)-1; i++ {
		for j := 0; j < len(usefulproducts)-i-1; j++ {
			if usefulproducts[j].InterestRate.FinalRate > usefulproducts[j+1].InterestRate.FinalRate {
				usefulproducts[j], usefulproducts[j+1] = usefulproducts[j+1], usefulproducts[j]
			}
		}
	}

	// 转换为前端需要的格式
	var loanPlans []respond.LoanPlan
	for i, product := range usefulproducts {
		// 计算月供（等额本息）
		monthlyRate := product.InterestRate.FinalRate / 100 / 12
		months := int(req.Term)
		principal := req.Amount

		var monthlyPayment float64
		if monthlyRate == 0 {
			monthlyPayment = principal / float64(months)
		} else {
			monthlyPayment = (principal * monthlyRate * math.Pow(1+monthlyRate, float64(months))) /
				(math.Pow(1+monthlyRate, float64(months)) - 1)
		}

		// 计算总利息
		totalInterest := monthlyPayment*float64(months) - principal

		// 设置标签
		tag := "可选"
		if i == 0 {
			tag = "推荐" // 利率最低的标记为推荐
		}

		plan := respond.LoanPlan{
			Id:             product.ProductID,
			Name:           product.ProductName,
			Tag:            tag,
			Rate:           fmt.Sprintf("年化%.2f%%", product.InterestRate.FinalRate),
			MonthlyPayment: fmt.Sprintf("%.2f", monthlyPayment),
			TotalInterest:  fmt.Sprintf("%.2f", totalInterest),
			Description: fmt.Sprintf("%s提供，支持%s等用途，预计审批时间%s",
				product.FinancialInstitution.Name,
				req.Purpose,
				product.EstimatedTime),
		}
		loanPlans = append(loanPlans, plan)
	}

	// 生成智能AI建议
	var aiSuggestion string
	if len(usefulproducts) == 1 {
		// 只有一个产品时
		plan := loanPlans[0]
		monthlyPayment, _ := strconv.ParseFloat(plan.MonthlyPayment, 64)
		totalInterest, _ := strconv.ParseFloat(plan.TotalInterest, 64)

		if monthlyPayment > 1000 && req.Term > 12 {
			aiSuggestion = fmt.Sprintf("AI风险评估提示：虽然%s是您的最佳选择，但月供%.2f元且期限较长，建议您仔细评估长期还款能力。该产品利率%s合理，但长期还款压力需要重点关注。建议您确保月收入稳定，并预留应急资金。总利息%.2f元相对可控，审批时间%s。",
				plan.Name, monthlyPayment, plan.Rate, totalInterest, plan.Description[len(plan.Description)-8:len(plan.Description)-4])
		} else if monthlyPayment > 1000 {
			aiSuggestion = fmt.Sprintf("根据您的贷款需求分析，我们为您找到了1个最优选择：%s。月供%.2f元相对较高，建议您确保有稳定的还款能力。该产品利率为%s，总利息成本%.2f元，审批时间%s，是当前最适合您需求的产品。",
				plan.Name, monthlyPayment, plan.Rate, totalInterest, plan.Description[len(plan.Description)-8:len(plan.Description)-4])
		} else if totalInterest > req.Amount*0.15 {
			aiSuggestion = fmt.Sprintf("AI成本分析提醒：%s的总利息%.2f元较高，占本金的%.1f%%。虽然该产品审批效率优势明显，适合紧急用款，但建议您评估是否可接受较高融资成本。如非紧急，建议等待更优产品。利率%s，月供%.2f元。",
				plan.Name, totalInterest, (totalInterest/req.Amount)*100, plan.Rate, monthlyPayment)
		} else if totalInterest > req.Amount*0.1 {
			aiSuggestion = fmt.Sprintf("AI智能分析显示，针对您的贷款需求，目前最匹配的是：%s。虽然总利息成本%.2f元相对较高，但该产品审批速度快，适合急需资金的客户。建议您优先考虑资金周转效率，利率为%s，月供%.2f元。",
				plan.Name, totalInterest, plan.Rate, monthlyPayment)
		} else if req.Amount > 100000 {
			aiSuggestion = fmt.Sprintf("针对您的大额贷款需求，AI强烈推荐：%s。虽然产品选择单一，但该产品利率%s在大额贷款中极具竞争力，总利息%.2f元控制良好。考虑到贷款规模，建议您重点关注该产品的审批条件和额度限制。月供%.2f元，建议确保充足现金流。",
				plan.Name, plan.Rate, totalInterest, monthlyPayment)
		} else {
			aiSuggestion = fmt.Sprintf("AI为您推荐了最佳贷款选择：%s。该产品利率%s在同类产品中具有明显优势，月供仅%.2f元，总利息%.2f元，还款压力较小。基于您的信用状况和资金需求，这是当前最优方案，审批预计%s。",
				plan.Name, plan.Rate, monthlyPayment, totalInterest, plan.Description[len(plan.Description)-8:len(plan.Description)-4])
		}
	} else if len(usefulproducts) == 2 {
		// 两个产品时进行对比
		firstPlan := loanPlans[0]
		secondPlan := loanPlans[1]
		firstRate, _ := strconv.ParseFloat(firstPlan.Rate[3:len(firstPlan.Rate)-1], 64)
		secondRate, _ := strconv.ParseFloat(secondPlan.Rate[3:len(secondPlan.Rate)-1], 64)

		firstMonthlyPayment, _ := strconv.ParseFloat(firstPlan.MonthlyPayment, 64)
		secondMonthlyPayment, _ := strconv.ParseFloat(secondPlan.MonthlyPayment, 64)
		firstTotalInterest, _ := strconv.ParseFloat(firstPlan.TotalInterest, 64)
		secondTotalInterest, _ := strconv.ParseFloat(secondPlan.TotalInterest, 64)

		if firstRate < secondRate*0.8 {
			monthlyDiff := firstMonthlyPayment - secondMonthlyPayment
			interestDiff := secondTotalInterest - firstTotalInterest
			aiSuggestion = fmt.Sprintf("AI深度分析显示，我们为您找到2个可选产品。强烈推荐%s，利率%s相比第二款产品%s优势明显，可为您节省利息成本%.2f元。虽然月供相差%.2f元，但长期看总成本更低。第二款产品可作为紧急情况备选。建议您优先选择推荐方案以获得最佳经济效益。",
				firstPlan.Name, firstPlan.Rate, secondPlan.Rate, interestDiff, monthlyDiff)
		} else if req.Amount > 50000 {
			aiSuggestion = fmt.Sprintf("针对您的大额贷款需求，AI为您对比了2个优质产品。虽然两款产品利率接近，但考虑到贷款规模，建议您选择%s。即使利率差异微小，长期累积也能节省可观的利息支出。建议您详细比较审批条件和服务质量，选择最符合您需求的方案。",
				firstPlan.Name)
		} else if firstMonthlyPayment > secondMonthlyPayment*1.2 {
			aiSuggestion = fmt.Sprintf("AI还款压力分析：我们为您找到2个产品，但需注意月供差异。虽然%s利率%s更低，但月供%.2f元显著高于%s的%.2f元。如果您的现金流紧张，建议优先考虑月供负担较小的产品；如果资金充足，则推荐低利率的%s以减少总成本。请根据实际还款能力谨慎选择。",
				firstPlan.Name, firstPlan.Rate, firstMonthlyPayment, secondPlan.Name, secondMonthlyPayment, firstPlan.Name)
		} else {
			aiSuggestion = fmt.Sprintf("根据您的贷款条件分析，我们筛选出2个合适产品。两款产品利率相近，推荐您根据审批时间和机构服务偏好进行选择。建议优先考虑%s，审批更快捷，能满足您的紧急资金需求。",
				firstPlan.Name)
		}
	} else if len(usefulproducts) >= 3 {
		// 多个产品时提供选择建议
		bestPlan := loanPlans[0]
		worstRate, _ := strconv.ParseFloat(loanPlans[len(loanPlans)-1].Rate[3:len(loanPlans[len(loanPlans)-1].Rate)-1], 64)
		bestRate, _ := strconv.ParseFloat(bestPlan.Rate[3:len(bestPlan.Rate)-1], 64)

		if bestRate < worstRate*0.5 {
			aiSuggestion = fmt.Sprintf("AI重磅发现！在%d个产品中发现超优惠选择：%s利率%s比最差方案低50%%以上，这是难得的低成本融资机会！建议您立即把握，该产品可为您节省大量利息支出。其他产品仅作参考，强烈建议您优先申请此优惠方案。",
				len(usefulproducts), bestPlan.Name, bestPlan.Rate)
		} else if bestRate < worstRate*0.7 {
			aiSuggestion = fmt.Sprintf("AI智能筛选完成！从众多产品中为您精选出%d个方案。特别推荐%s，其利率%s远低于市场平均水平，可为您大幅降低融资成本。其他产品可作为参考对比。建议您优先考虑推荐方案，享受最优惠的贷款条件。",
				len(usefulproducts), bestPlan.Name, bestPlan.Rate)
		} else if req.Amount > 100000 {
			bestMonthlyPayment, _ := strconv.ParseFloat(bestPlan.MonthlyPayment, 64)
			aiSuggestion = fmt.Sprintf("针对您%.2f元的大额贷款需求，AI为您筛选了%d个优质产品。重点推荐%s，月供%.2f元相对合理。在大额贷款中，即使微小利率差异也会产生显著影响。建议您选择%sg降低总成本，同时确保有充足的月度现金流。",
				req.Amount, len(usefulproducts), bestPlan.Name, bestMonthlyPayment, bestPlan.Rate)
		} else if req.Amount > 50000 {
			aiSuggestion = fmt.Sprintf("针对您的较高额度贷款需求，AI为您匹配了%d个优质产品。考虑到贷款规模较大，建议您重点关注利率差异。推荐选择%s，虽然各产品利率差异不大，但长期累积的利息节省效果显著。建议详细比较各方案后再做决定。",
				len(usefulproducts), bestPlan.Name)
		} else if req.Term > 36 {
			aiSuggestion = fmt.Sprintf("针对您长达%d个月的贷款周期，AI进行了长期成本分析。强烈推荐%s，虽然各产品利率看似差异不大，但在3年以上的贷款周期中，低利率的优势会被显著放大。%s能为您节省可观的利息支出，建议您优先考虑。",
				req.Term, bestPlan.Name, bestPlan.Name)
		} else if req.Term > 24 {
			aiSuggestion = fmt.Sprintf("针对您较长期限（%d个月）的贷款需求，AI分析认为您更适合选择%s。长期贷款中利率差异会被放大，建议您优先考虑低利率产品以控制总成本。虽然其他产品审批可能更快，但长期来看%s能为您节省更多资金。建议仔细权衡时间成本与资金成本。",
				req.Term, bestPlan.Name, bestPlan.Name)
		} else if req.Amount < 10000 && req.Term < 12 {
			aiSuggestion = fmt.Sprintf("针对您的短期小额贷款需求，AI为您匹配了%d个灵活方案。考虑到金额小期限短，建议您优先关注审批效率而非微小利率差异。推荐%s，审批流程简化，快速到账满足您紧急用款需求。虽然利率不是最低，但时间成本往往更有价值。",
				len(usefulproducts), bestPlan.Name)
		} else if req.Amount < 10000 {
			aiSuggestion = fmt.Sprintf("针对您的短期小额贷款需求，AI为您匹配了%d个灵活方案。考虑到金额较小，建议您重点考虑%s的审批效率。该产品不仅利率合理，审批流程也较为简化，适合您的紧急资金周转。其他产品可作为备选，建议根据实际用款紧迫度决定。",
				len(usefulproducts), bestPlan.Name)
		} else {
			aiSuggestion = fmt.Sprintf("基于您的贷款需求分析，AI为您智能推荐%d个合适方案。考虑到您的资金需求和还款能力，建议您优先选择利率最低的%s，同时关注审批时效。不同方案在服务质量和审批速度上各有优势，可根据您的紧急程度灵活选择。",
				len(usefulproducts), bestPlan.Name)
		}
	} else {
		// 备用建议（不会执行到这里，但保留以防万一）
		aiSuggestion = fmt.Sprintf("AI智能分析完成！根据您的贷款条件，我们为您找到了%d个优质方案。建议您选择%s，该产品在利率、审批效率和服务质量方面综合表现最佳。如有特殊需求，也可以考虑其他备选方案。",
			len(usefulproducts), loanPlans[0].Name)
	}

	return "获取成功", 0, &respond.AiRespond{
		AiSuggestion: aiSuggestion,
		LoanPlans:    loanPlans,
	}
}
