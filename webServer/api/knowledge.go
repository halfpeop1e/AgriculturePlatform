package v1

import (
	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"

	"github.com/gin-gonic/gin"
)

func GetKnowledgeList(c *gin.Context) {
	msg, data, code := gorm.KnowledgeServer.GetKnowledgeList()
	JsonBack(c, msg, code, data)
}

func AddKnowledge(c *gin.Context) {
	var req request.KnowledgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonBack(c, "参数错误", 400, nil)
		return
	}
	msg, code := gorm.KnowledgeServer.AddKnowledge(req)
	JsonBack(c, msg, code, nil)
}
