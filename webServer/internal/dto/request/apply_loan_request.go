package request

type ApplyLoanRequest struct {
	UserID    string  `json:"user_id"`
	ProductID string  `json:"product_id"`
	Amount    float64 `json:"amount"`
	Deadline  int64   `json:"deadline"`
}
