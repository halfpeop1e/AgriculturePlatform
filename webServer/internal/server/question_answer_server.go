package gorm

import (
	"encoding/json"
	"fmt"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"log"
	"time"
)

type questionServerType struct{}

var QuestionServer = &questionServerType{}

func (q *questionServerType) GetQuestionList() (string, *[]respond.QuestionRespond, int) {
	var questions []model.ExpertCase

	// 查询所有未删除的问题，按时间倒序
	if res := dao.GormDB.Find(&questions); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败1", nil, -1
	}
	// 转换为响应格式
	questionList := make([]respond.QuestionRespond, 0, len(questions))

	for _, question := range questions {
		var tags []string
		if question.Tags != "" {
			json.Unmarshal([]byte(question.Tags), &tags)
		}
		var expert model.Expert
		log.Print("expId:" + question.ExpertId)
		if res := dao.GormDB.Where("id = ?", question.ExpertId).First(&expert); res.Error != nil {

			log.Printf("Database error: %v", res.Error)

			return "查询失败3", nil, -1

		}

		var user model.User
		log.Print("userId:" + question.AuthorId)
		if res := dao.GormDB.Where("uuid = ?", question.AuthorId).First(&user); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败4", nil, -1
		}

		newQuestion := respond.QuestionRespond{
			Id:         fmt.Sprintf("%d", question.Id),
			ExpertId:   question.ExpertId,
			ExpertUuid: expert.Uuid,
			Expert:     expert.Name,
			Title:      question.Tittle,
			Author:     user.Name,
			AuthorId:   question.AuthorId,
			Tags:       tags,
			Question:   question.Question,
			Answer:     question.Answer,
			CaseDate:   question.CaseDate.Format("2006-01-02 15:04:05"),
		}
		questionList = append(questionList, newQuestion)
	}

	return "获取成功", &questionList, 0
}

func (q *questionServerType) GetQuestionListByUser(userId string, role string) (string, *[]respond.QuestionRespond, int) {
	var questions []model.ExpertCase

	if role == "normal" {
		if res := dao.GormDB.Where("author_id = ?", userId).Find(&questions); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败10", nil, -1
		}
	} else {
		var expert model.Expert
		if res := dao.GormDB.Where("uuid = ?", userId).First(&expert); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败10", nil, -1
		}
		if res := dao.GormDB.Where("expert_id = ?", fmt.Sprintf("%d", expert.Id)).Find(&questions); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败11", nil, -1
		}
	}

	// 转换为响应格式
	questionList := make([]respond.QuestionRespond, 0, len(questions))

	for _, question := range questions {
		var tags []string
		if question.Tags != "" {
			json.Unmarshal([]byte(question.Tags), &tags)
		}
		var expert model.Expert
		log.Print("expId:" + question.ExpertId)
		if res := dao.GormDB.Where("id = ?", question.ExpertId).First(&expert); res.Error != nil {

			log.Printf("Database error: %v", res.Error)

			return "查询失败3", nil, -1

		}

		var user model.User
		log.Print("userId:" + question.AuthorId)
		if res := dao.GormDB.Where("uuid = ?", question.AuthorId).First(&user); res.Error != nil {
			log.Printf("Database error: %v", res.Error)
			return "查询失败4", nil, -1
		}

		newQuestion := respond.QuestionRespond{
			Id:         fmt.Sprintf("%d", question.Id),
			ExpertId:   question.ExpertId,
			ExpertUuid: expert.Uuid,
			Expert:     expert.Name,
			Title:      question.Tittle,
			Author:     user.Name,
			AuthorId:   question.AuthorId,
			Tags:       tags,
			Question:   question.Question,
			Answer:     question.Answer,
			CaseDate:   question.CaseDate.Format("2006-01-02 15:04:05"),
		}
		questionList = append(questionList, newQuestion)
	}

	return "获取成功", &questionList, 0
}
func (q *questionServerType) CreateQuestion(req request.CreateQuestionRequest) (string, int) {

	var user model.User
	if res := dao.GormDB.Where("uuid = ?", req.AuthorId).First(&user); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "用户不存在", -1
	}
	useTags, _ := json.Marshal(req.Tags)
	question := model.ExpertCase{
		ExpertId:  req.ExpertId,
		Question:  req.Content,
		Tittle:    req.Title,
		Tags:      string(useTags),
		AuthorId:  req.AuthorId,
		Answer:    "",
		CaseDate:  time.Now(),
		CreatedAt: time.Now(),
	}

	if res := dao.GormDB.Create(&question); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "创建失败", -1
	}

	return "创建成功", 0
}

func (q *questionServerType) AnswerQuestion(req request.SubmitAnswerRequest, userId string) (string, int) {
	// 检查问题是否存在
	var question model.ExpertCase
	if res := dao.GormDB.Where("id = ?", req.QuestionId).First(&question); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "问题不存在", -1
	}

	question.Answer = req.Content

	dao.GormDB.Save(&question)

	return "回答成功", 0
}

func (q *questionServerType) GetQuestionDetail(questionId string) (string, *respond.QuestionRespond, int) {
	var question model.ExpertCase
	if res := dao.GormDB.Where("id = ?", questionId).First(&question); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "问题不存在", nil, -1
	}

	var expert model.Expert

	if res := dao.GormDB.Where("id = ?", question.ExpertId).First(&expert); res.Error != nil {

		log.Printf("Database error: %v", res.Error)

		return "专家不存在", nil, -1

	}

	var user model.User
	if res := dao.GormDB.Where("uuid = ?", question.AuthorId).First(&user); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "用户不存在", nil, -1
	}
	var tags []string
	if question.Tags != "" {
		json.Unmarshal([]byte(question.Tags), &tags)
	}
	questionDetail := &respond.QuestionRespond{
		Id:         questionId,
		ExpertId:   question.ExpertId,
		Expert:     expert.Name,
		ExpertUuid: expert.Uuid,
		Title:      question.Tittle,
		Author:     user.Name,
		AuthorId:   question.AuthorId,
		Question:   question.Question,
		Answer:     question.Answer,
		Tags:       tags,
		CaseDate:   question.CaseDate.Format("2006-01-02"),
	}
	return "获取成功", questionDetail, 0
}
