package v1

import (
	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"

	"github.com/gin-gonic/gin"
)

func GetQuestionList(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		msg, data, code := gorm.QuestionServer.GetQuestionList()
		JsonBack(c, msg, code, data)
	} else {
		role := c.Query("role")
		msg, data, code := gorm.QuestionServer.GetQuestionListByUser(userId, role)
		JsonBack(c, msg, code, data)
	}

}

func CreateQuestion(c *gin.Context) {
	var req request.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonBack(c, "参数错误", -1, nil)
		return
	}
	msg, code := gorm.QuestionServer.CreateQuestion(req)
	JsonBack(c, msg, code, nil)
}

func AnswerQuestion(c *gin.Context) {
	var req request.SubmitAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonBack(c, "参数错误", -1, nil)
		return
	}
	userId := c.GetString("user_id")
	msg, code := gorm.QuestionServer.AnswerQuestion(req, userId)
	JsonBack(c, msg, code, nil)
}

func GetQuestionDetail(c *gin.Context) {

	questionId := c.Param("questionId")

	msg, data, code := gorm.QuestionServer.GetQuestionDetail(questionId)

	JsonBack(c, msg, code, data)

}
