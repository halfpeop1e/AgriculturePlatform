package httpserver

import (
	v1 "go-agriculture/api"
	"go-agriculture/internal/util"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Print("token : " + authHeader)
		if authHeader == "" {
			v1.JsonBack(c, "请先登录", -2, nil)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			v1.JsonBack(c, "无效的Token", -2, nil)
			c.Abort()
			return
		}

		// 解析 Token
		claims, err := util.ParseToken(tokenParts[1])
		if err != nil {
			v1.JsonBack(c, "无效的Token", -2, nil)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
