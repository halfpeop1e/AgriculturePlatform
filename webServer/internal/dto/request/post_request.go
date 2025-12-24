package request

// CreatePostRequest 创建帖子请求（匹配前端 CreatePostInput）
type CreatePostRequest struct {
	Title   string      `json:"title" binding:"required"`   // 帖子标题
	Content string      `json:"content" binding:"required"` // 帖子内容
	Author  AuthorInput `json:"author" binding:"required"`  // 作者信息
}

// AuthorInput 作者信息输入
type AuthorInput struct {
	Id       string `json:"id" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

// CreateReplyRequest 创建回复请求（匹配前端 CreateReplyInput）
type CreateReplyRequest struct {
	Content string      `json:"content" binding:"required"` // 回复内容
	Author  AuthorInput `json:"author" binding:"required"`  // 作者信息
}

// ToggleLikeRequest 切换点赞状态请求
type ToggleLikeRequest struct {
	Liked bool `json:"liked"` // 目标点赞状态
}
