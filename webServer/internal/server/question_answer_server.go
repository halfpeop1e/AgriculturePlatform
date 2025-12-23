package gorm

import (
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
	var questions []model.Question

	// 查询所有未删除的问题，按时间倒序
	if res := dao.GormDB.Where("delete_at IS NULL").Order("date DESC").Find(&questions); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询失败", nil, -1
	}

	// 转换为响应格式
	questionList := make([]respond.QuestionRespond, 0, len(questions))

	for _, question := range questions {
		tags := question.Tags
		if tags == nil {
			tags = []string{}
		}

		newQuestion := respond.QuestionRespond{
			Id:          fmt.Sprintf("%d", question.Id),
			ExpertId:    question.ExpertId,
			Title:       question.Title,
			Content:     question.Content,
			Author:      question.Author,
			AuthorId:    question.AuthorId,
			Date:        question.Date.Format("2006-01-02 15:04:05"),
			Tags:        tags,
			AnswerCount: question.AnswerCount,
			IsAnswered:  question.IsAnswered,
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
	question := model.Question{
		ExpertId:    req.ExpertId,
		Title:       req.Title,
		Content:     req.Content,
		Author:      user.Name,
		AuthorId:    req.AuthorId,
		Date:        time.Now(),
		Tags:        req.Tags,
		AnswerCount: 0,
		IsAnswered:  false,
	}

	if res := dao.GormDB.Create(&question); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "创建失败", -1
	}

	return "创建成功", 0
}

func (q *questionServerType) AnswerQuestion(req request.SubmitAnswerRequest, userId string) (string, int) {
	// 检查问题是否存在
	var question model.Question
	if res := dao.GormDB.Where("id = ?", req.QuestionId).First(&question); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "问题不存在", -1
	}

	answer := model.Answer{
		QuestionId: question.Id,
		Content:    req.Content,
		Date:       time.Now(),
	}

	if res := dao.GormDB.Create(&answer); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "回答失败", -1
	}
	question.IsAnswered = true
	question.AnswerCount++
	dao.GormDB.Save(&question)

	return "回答成功", 0
}

func (q *questionServerType) GetQuestionDetail(questionId string) (string, *respond.QuestionDetailRespond, int) {
	var question model.Question
	if res := dao.GormDB.Where("id = ?", questionId).First(&question); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "问题不存在", nil, -1
	}
	var answers []model.Answer
	if res := dao.GormDB.Where("question_id = ?", question.Id).Order("date DESC").Find(&answers); res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		return "查询回答失败", nil, -1
	}
	answerList := make([]respond.Answer, 0, len(answers))
	for _, answer := range answers {
		newAnswer := respond.Answer{
			Id:         fmt.Sprintf("%d", answer.Id),
			QuestionId: fmt.Sprintf("%d", answer.QuestionId),
			Content:    answer.Content,
			Expert:     "", // 这里可以根据需要查询专家信息填充
			ExpertId:   "",
			Date:       answer.Date.Format("2006-01-02 15:04:05"),
			IsAccpeted: answer.IsAccpet,
		}
		answerList = append(answerList, newAnswer)
	}
	questionDetail := &respond.QuestionDetailRespond{
		Id:          fmt.Sprintf("%d", question.Id),
		ExpertId:    question.ExpertId,
		Title:       question.Title,
		Content:     question.Content,
		Author:      question.Author,
		AuthorId:    question.AuthorId,
		Date:        question.Date.Format("2006-01-02 15:04:05"),
		Tags:        question.Tags,
		AnswerCount: question.AnswerCount,
		IsAnswered:  question.IsAnswered,
		Answers:     answerList,
	}
	return "获取成功", questionDetail, 0
}
