package request

type BookExpertRequest struct {
	Question    string `json:"question" binding:"required"`    // 咨询问题
	Date        string `json:"date" binding:"required"`        // 预约时间
	ContactInfo string `json:"contactInfo" binding:"required"` // 联系方式
}
