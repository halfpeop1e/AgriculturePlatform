package respond

type CheckMyLoanRespond struct {
	LoanedSum float64     `json:"loanedSum"` //所有贷款金额
	LoanSum   float64     `json:"loanSum"`   //所有待还金额
	LoanList  []LoanOrder `json:"loanList"`
}

type LoanOrder struct {
	Year       int     `json:"year"`
	Month      int     `json:"month"`
	Amount     float64 `json:"amount"`
	LoanName   string  `json:"loanName"`
	LoanStatus string  `json:"loanStatus"`
}
