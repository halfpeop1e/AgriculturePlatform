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
		product.PUT("/edit/:productId", v1.EditerProduct)
		product.DELETE("/delete/:productId", v1.DeleteProduct)
	}
	order := GE.Group("/order")
	order.Use(JWTAuthMiddleware())
	{
		order.GET("/list", v1.GetOrderList)
	}
	user := GE.Group("/user")
	user.GET("/profile/:userId", v1.GetUserInfo)
	user.Use(JWTAuthMiddleware())
	{
		user.POST("/profile/:userId/update", v1.UpdateUserInfo)
		user.POST("/security/:userId/update", v1.SafeUpdateInfo)
	}
	loan := GE.Group("/loan")
	loan.Use(JWTAuthMiddleware())
	{
		loan.GET("/apply/list", v1.GetApplyList)
		loan.GET("/list", v1.GetLoanProductList)
		loan.POST("/add", v1.AddLoanProduct)
		loan.POST("/apply", v1.ApplyLoan)
	}
	question := GE.Group("/question")
	question.Use(JWTAuthMiddleware())
	{
		question.GET("/list", v1.GetQuestionList)
		question.POST("/create", v1.CreateQuestion)
		question.GET("/detail/:questionId", v1.GetQuestionDetail)
		question.POST("/answer/:questionId", v1.AnswerQuestion)
	}

	expert := GE.Group("/expert")
	expert.Use(JWTAuthMiddleware())
	{
		expert.GET("/list", v1.GetExpertList)
		expert.GET("/detail/:expertId", v1.GetExpertDetail)
		expert.POST("/contact/:expertId/book", v1.BookExpert)
	}
}
