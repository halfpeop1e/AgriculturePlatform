package v1

import (
	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLoanProductList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	msg, data, code := gorm.GetLoanProductList(page, pageSize)
	JsonBack(c, msg, code, data)
}

func AddLoanProduct(c *gin.Context) {
	var req request.AddLoanProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	msg, code := gorm.AddLoanProduct(req)

	JsonBack(c, msg, code, nil)
}

func ApplyLoan(c *gin.Context) {
	var req request.ApplyLoanRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}
	msg, code := gorm.ApllyLoanProduct(req)
	JsonBack(c, msg, code, nil)
}

func GetApplyList(c *gin.Context) {
	var userId = c.Query("user_id")
	msg, code, data := gorm.GetApplyList(userId)
	JsonBack(c, msg, code, data)
}

func AllowLoan(c *gin.Context) {
	var req request.AllowLoanRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}
	msg, code := gorm.AllowLoan(req)
	JsonBack(c, msg, code, nil)
}
