package v1

import (
	"log"
	"net/http"
	"strconv"

	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"

	"github.com/gin-gonic/gin"
)

func GetProductList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	salerId := c.DefaultQuery("salerId", "all")

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

	msg, data, code := gorm.ProductServer.GetProductList(page, pageSize, salerId)

	JsonBack(c, msg, code, data)
}

func PostProduct(c *gin.Context) {
	// 先绑定到包装结构体
	var wrapper request.PostProductRequestWrapper
	if err := c.ShouldBind(&wrapper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	log.Printf("wrapper: %+v", wrapper)
	// 提取实际的数据
	req := wrapper.FormData
	if c.GetString("user_id") != req.SalerId {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "这id是你吗",
		})
		return
	}
	msg, code := gorm.ProductServer.PostProduct(req)
	JsonBack(c, msg, code, nil)
}

func BuyProduct(c *gin.Context) {
	var req request.BuyProductRequestWrapper
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}
	msg, code := gorm.ProductServer.BuyProduct(req.Order, c.GetString("user_id"))
	JsonBack(c, msg, code, nil)
}

func GetOrderList(c *gin.Context) {
	msg, code, data := gorm.ProductServer.GetOrderList()
	JsonBack(c, msg, code, data)
}

func EditerProduct(c *gin.Context) {
	productId := c.Param("productId")
	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少产品ID",
		})
		return
	}
	var wrapper request.EditerProductRequestWrapper
	if err := c.ShouldBind(&wrapper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}
	req := wrapper.FormData
	msg, code := gorm.ProductServer.EditerProduct(req, productId)
	JsonBack(c, msg, code, nil)
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("productId")
	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少产品ID",
		})
		return
	}
	msg, code := gorm.ProductServer.DeleteProduct(productId)
	JsonBack(c, msg, code, nil)
}
