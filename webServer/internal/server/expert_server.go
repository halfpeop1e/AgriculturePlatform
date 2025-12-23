package gorm

import (
	"encoding/json"
	"fmt"
	"go-agriculture/internal/config"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"log"
	"time"

	"gorm.io/gorm"
)

type expertServerType struct{}

var ExpertServer = &expertServerType{}

// GetExpertList 获取专家列表（支持分页、领域筛选、搜索）
func (e *expertServerType) GetExpertList(page, pageSize int, field, searchKey string) (string, *respond.ExpertListRespond, int) {
	var total int64
	var experts []model.Expert

	// 构建基础查询
	query := dao.GormDB.Model(&model.Expert{}).Where("status = ?", "active")

	// 按领域筛选
	if field != "" {
		query = query.Where("field = ?", field)
	}

	// 搜索功能：姓名或简介或教育背景包含关键词
	if searchKey != "" {
		query = query.Where("name LIKE ? OR introduction LIKE ? OR education LIKE ?",
			"%"+searchKey+"%", "%"+searchKey+"%", "%"+searchKey+"%")
	}

	// 获取总数
	if res := query.Count(&total); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}

	expertList := &respond.ExpertListRespond{
		List:     make([]respond.ExpertRespond, 0),
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		HasMore:  int64(page*pageSize) < total,
	}

	if total == 0 {
		return "获取成功", expertList, 0
	}

	// 计算偏移量
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	// 查询数据，按咨询次数和好评率排序
	if res := query.Order("consult_count DESC, rating DESC").
		Limit(pageSize).Offset(offset).
		Find(&experts); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}

	// 转换数据
	for _, expert := range experts {
		avatar := expert.Avatar
		if avatar != "" {
			avatar = "http://" + fmt.Sprint(config.GetConfig().MainConfig.ServerIP) + ":" +
				fmt.Sprint(config.GetConfig().MainConfig.Port) + "/static/avatars/" + avatar
		}

		// 解析技能 JSON
		skills := make([]string, 0)
		if expert.Skills == "" {
			skills = []string{}
		} else if err := json.Unmarshal([]byte(expert.Skills), &skills); err != nil {
			skills = []string{}
		}

		// 查询近期案例（最多3条）
		cases := e.getRecentCases(fmt.Sprintf("%d", expert.Id), 3)

		newExpert := respond.ExpertRespond{
			Id:           fmt.Sprintf("%d", expert.Id),
			Name:         expert.Name,
			Avatar:       avatar,
			Field:        expert.Field,
			Introduction: expert.Introduction,
			Skills:       skills,
			ConsultCount: expert.ConsultCount,
			Rating:       expert.Rating,
			ResponseTime: expert.ResponseTime,
			RecentCases:  cases,
		}
		expertList.List = append(expertList.List, newExpert)
	}

	// 更新 hasMore
	if len(expertList.List) == 0 {
		expertList.HasMore = false
	} else {
		expertList.HasMore = int64(page*pageSize) < total
	}

	return "获取成功", expertList, 0
}

// GetExpertDetail 获取专家详情
func (e *expertServerType) GetExpertDetail(expertId string) (string, *respond.ExpertDetailRespond, int) {
	var expert model.Expert

	if res := dao.GormDB.Where("uuid = ? AND status = ?", expertId, "active").First(&expert); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}

	// 处理头像路径
	avatar := expert.Avatar
	if avatar != "" {
		avatar = "http://" + fmt.Sprint(config.GetConfig().MainConfig.ServerIP) + ":" +
			fmt.Sprint(config.GetConfig().MainConfig.Port) + "/static/avatars/" + avatar
	}

	// 解析 JSON 字段
	skills := make([]string, 0)
	if expert.Skills == "" {
		skills = []string{}
	} else if err := json.Unmarshal([]byte(expert.Skills), &skills); err != nil {
		skills = []string{}
	}

	certification := make([]string, 0)
	if expert.Certification == "" {
		certification = []string{}
	} else if err := json.Unmarshal([]byte(expert.Certification), &certification); err != nil {
		certification = []string{}
	}

	availableTime := make([]string, 0)
	if expert.AvailableTime == "" {
		availableTime = []string{}
	} else if err := json.Unmarshal([]byte(expert.AvailableTime), &availableTime); err != nil {
		availableTime = []string{}
	}

	serviceScope := make([]string, 0)
	if expert.ServiceScope == "" {
		serviceScope = []string{}
	} else if err := json.Unmarshal([]byte(expert.ServiceScope), &serviceScope); err != nil {
		serviceScope = []string{}
	}

	// 查询近期案例
	cases := e.getRecentCases(fmt.Sprintf("%d", expert.Id), 10)

	expertDetail := &respond.ExpertDetailRespond{
		Id:            fmt.Sprintf("%d", expert.Id),
		Name:          expert.Name,
		Avatar:        avatar,
		Field:         expert.Field,
		Introduction:  expert.Introduction,
		Skills:        skills,
		Education:     expert.Education,
		Experience:    expert.Experience,
		Certification: certification,
		ConsultCount:  expert.ConsultCount,
		Rating:        expert.Rating,
		ResponseTime:  expert.ResponseTime,
		AvailableTime: availableTime,
		ServiceScope:  serviceScope,
		Price:         expert.Price,
		RecentCases:   cases,
	}

	return "获取成功", expertDetail, 0
}

// getRecentCases 获取专家的近期案例
func (e *expertServerType) getRecentCases(expertId string, limit int) []respond.CaseRespond {
	var expertCases []model.ExpertCase
	cases := make([]respond.CaseRespond, 0)

	if res := dao.GormDB.Where("expert_id = ?", expertId).
		Order("case_date DESC").
		Limit(limit).
		Find(&expertCases); res.Error != nil {
		log.Printf("查询专家案例失败: %v", res.Error)
		return cases
	}

	for _, ec := range expertCases {
		cases = append(cases, respond.CaseRespond{
			Date:     ec.CaseDate.Format("2006-01-02"),
			Question: ec.Question,
			Answer:   ec.Answer,
		})
	}

	return cases
}

func (e *expertServerType) BookExpert(expertId string, req request.BookExpertRequest, userId string) (string, int) {
	// 检查专家是否存在且状态为 active
	var expert model.Expert
	if res := dao.GormDB.Where("uuid = ? AND status = ?", expertId, "active").First(&expert); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return "专家不存在或已停用", -1
		}
		log.Printf("查询专家失败: %v", res.Error)
		return "查询失败", -1
	}

	// 检查用户是否存在
	var user model.User
	if res := dao.GormDB.Where("uuid = ?", userId).First(&user); res.Error != nil {
		log.Printf("查询用户失败: %v", res.Error)
		return "用户不存在", -1
	}

	// 检查该用户是否已对该专家有未完成的预约
	var count int64
	if res := dao.GormDB.Model(&model.ExpertBooking{}).
		Where("expert_id = ? AND user_id = ? AND status IN ?", expertId, userId, []string{"pending", "confirmed"}).
		Count(&count); res.Error != nil {
		log.Printf("查询预约记录失败: %v", res.Error)
		return "查询失败", -1
	}
	if count > 0 {
		return "您已有该专家的未完成预约", -1
	}

	// 创建预约记录
	booking := model.ExpertBooking{
		ExpertId:    expertId,
		UserId:      userId,
		UserName:    user.Name,
		Question:    req.Question,
		ContactInfo: req.ContactInfo,
		BookDate:    req.Date,
		Status:      "通过",
		CreatedAt:   time.Now(),
	}

	if res := dao.GormDB.Create(&booking); res.Error != nil {
		log.Printf("创建预约失败: %v", res.Error)
		return "预约失败", -1
	}

	// 增加专家的咨询次数
	expert.ConsultCount += 1
	if res := dao.GormDB.Save(&expert); res.Error != nil {
		log.Printf("更新专家咨询次数失败: %v", res.Error)
	}

	return "预约成功", 0
}
