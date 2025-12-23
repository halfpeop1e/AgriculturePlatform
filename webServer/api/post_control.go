package v1

import (
	"go-agriculture/internal/dto/request"
	gorm "go-agriculture/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPostList 获取帖子列表（对应前端 /posts）
func GetPostList(c *gin.Context) {

	msg, data, code := gorm.PostServer.GetPostList()
	JsonBack(c, msg, code, data)
}

// GetPostById 获取帖子详情（对应前端 /posts/:id）
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少帖子ID",
		})
		return
	}

	userId := c.GetString("user_id")

	msg, data, code := gorm.PostServer.GetPostById(id, userId)
	JsonBack(c, msg, code, data)
}

// CreatePost 创建帖子（对应前端 POST /posts）
func CreatePost(c *gin.Context) {
	var req request.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	msg, data, code := gorm.PostServer.CreatePost(req, userId)
	JsonBack(c, msg, code, data)
}

// AddReply 添加回复（对应前端 POST /posts/:postId/replies）
func AddReply(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少帖子ID",
		})
		return
	}

	var req request.CreateReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	msg, data, code := gorm.PostServer.AddReply(postId, req, userId)
	JsonBack(c, msg, code, data)
}

// TogglePostLike 切换帖子点赞状态（对应前端 POST /posts/:postId/like）
func TogglePostLike(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少帖子ID",
		})
		return
	}

	var req request.ToggleLikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	msg, code := gorm.PostServer.TogglePostLike(postId, req, userId)
	JsonBack(c, msg, code, nil)
}

// ToggleReplyLike 切换回复点赞状态（对应前端 POST /posts/:postId/replies/:replyId/like）
func ToggleReplyLike(c *gin.Context) {
	postId := c.Param("postId")
	replyId := c.Param("replyId")

	if postId == "" || replyId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少帖子ID或回复ID",
		})
		return
	}

	var req request.ToggleLikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数解析错误: " + err.Error(),
		})
		return
	}

	userId := c.GetString("user_id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	msg, code := gorm.PostServer.ToggleReplyLike(postId, replyId, req, userId)
	JsonBack(c, msg, code, nil)
}
