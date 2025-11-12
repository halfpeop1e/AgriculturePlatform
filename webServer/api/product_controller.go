package v1

import (
	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductList(c *gin.Context) {

	msg, data, code := gorm.ProductServer.GetProductList()

	JsonBack(c, msg, code, data)
}

func PostProduct(c *gin.Context) {
	var req request.PostProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}
	// if c.GetString("user_id") != req.SalerId {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code":    400,
	// 		"message": "这id是你吗",
	// 	})
	// 	return
	// }
	msg, code := gorm.ProductServer.PostProduct(req)
	JsonBack(c, msg, code, nil)
}

func BuyProduct(c *gin.Context) {
	var req request.BuyProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}
	msg, code := gorm.ProductServer.BuyProduct(req, c.GetString("user_id"))
	JsonBack(c, msg, code, nil)
}
