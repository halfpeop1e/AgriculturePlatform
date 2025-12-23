package request

type CreateQuestionRequest struct {
	ExpertId string   `json:"expertId" binding:"required"` // 专家ID
	Title    string   `json:"title" binding:"required"`    // 问题标题
	Content  string   `json:"content" binding:"required"`  // 问题内容
	AuthorId string   `json:"authorId" binding:"required"` // 作者ID（用户ID）
	Tags     []string `json:"tags"`                        // 问题标签（可选）
}

type SubmitAnswerRequest struct {
	QuestionId string `json:"questionId" binding:"required"`
	Content    string `json:"content" binding:"required"`
}
