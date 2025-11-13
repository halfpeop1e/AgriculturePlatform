package httpserver

import (
	v1 "go-agriculture/api"
	"go-agriculture/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GE *gin.Engine

func init() {

	GE = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	GE.Use(cors.New(corsConfig))
	GE.POST("/user/register", v1.Register)
	GE.POST("/user/login", v1.Login)
	static := GE.Group("")
	static.Use()
	{
		static.Static("/static/avatars", config.GetConfig().StaticAvatarPath)
		static.Static("/static/files", config.GetConfig().StaticFilePath)
	}
	product := GE.Group("/product")
	product.Use(JWTAuthMiddleware())
	{
		product.GET("/list", v1.GetProductList)
		product.POST("/sell", v1.PostProduct)
		product.POST("/buy", v1.BuyProduct)
	}
	order := GE.Group("/order")
	order.Use(JWTAuthMiddleware())
	{
		order.GET("/list", v1.GetOrderList)
	}
	user := GE.Group("/user")
	user.Use(JWTAuthMiddleware())
	{
		user.GET("/profile/:userId", v1.GetUserInfo)
		user.POST("/profile/:userId/update", v1.UpdateUserInfo)
		user.POST("/security/:userId/update", v1.SafeUpdateInfo)
	}
}
