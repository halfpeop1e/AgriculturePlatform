package respond

type ApplyListRespond struct {
	Id          int64  `json:"id"`
	Uuid        string `json:"uuid"`
	UserId      string `json:"user_id"`
	ProductID   string `json:"product_id"`
	Avatar      string `json:"avatar"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
}
