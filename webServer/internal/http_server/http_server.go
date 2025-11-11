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
	GE.Static("/static/avatars", config.GetConfig().StaticAvatarPath)
	GE.Static("/static/files", config.GetConfig().StaticFilePath)
	GE.POST("/user/register", v1.Register)
	GE.POST("/user/login", v1.Login)
	private := GE.Group("/api")
	private.Use(JWTAuthMiddleware())
	{
	}

}
