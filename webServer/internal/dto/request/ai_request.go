package request

type AiRequest struct {
	Purpose string  `json:"purpose"`
	Amount  float64 `json:"amount"`
	Term    int     `json:"term"`
}
