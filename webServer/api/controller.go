package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonBack(c *gin.Context, message string, ret int, data interface{}) {
	if ret == 0 {
		if data != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": message,
				"data":    data,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": message,
			})
		}
	} else if ret == -2 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": message,
		})
	} else if ret == -1 {
		c.JSON(500, gin.H{
			"code":    500,
			"message": message,
		})
	}
}
