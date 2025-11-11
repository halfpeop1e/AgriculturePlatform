package v1

import (
	request "go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	msg, resp, code := gorm.UserInfoService.Login(req)
	JsonBack(c, msg, code, resp)
}

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	msg, resp, code := gorm.UserInfoService.Register(req)
	JsonBack(c, msg, code, resp)
}
