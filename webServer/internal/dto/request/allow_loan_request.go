package request

type AllowLoanRequest struct {
	ApplyID string `json:"apply_id"`
	Allow   bool   `json:"allow"`
}
