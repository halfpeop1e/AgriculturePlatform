package v1

import (
	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetExpertList 获取专家列表
func GetExpertList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	field := c.Query("field")
	searchKey := c.Query("searchKey")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 50 {
		pageSize = 50
	}

	msg, data, code := gorm.ExpertServer.GetExpertList(page, pageSize, field, searchKey)
	JsonBack(c, msg, code, data)
}

// GetExpertDetail 获取专家详情
func GetExpertDetail(c *gin.Context) {
	expertId := c.Param("expertId")
	if expertId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少专家ID",
		})
		return
	}

	msg, data, code := gorm.ExpertServer.GetExpertDetail(expertId)
	JsonBack(c, msg, code, data)
}

// BookExpert 预约专家咨询
func BookExpert(c *gin.Context) {
	expertId := c.Param("expertId")
	if expertId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少专家ID",
		})
		return
	}

	var req request.BookExpertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	// 从上下文中获取用户ID（通过 JWT 中间件设置）
	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	msg, code := gorm.ExpertServer.BookExpert(expertId, req, userId)
	JsonBack(c, msg, code, nil)
}
