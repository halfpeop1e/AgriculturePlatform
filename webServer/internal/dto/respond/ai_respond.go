package respond

type LoanPlan struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Tag            string `json:"tag"`
	Rate           string `json:"rate"`
	MonthlyPayment string `json:"monthlyPayment"`
	TotalInterest  string `json:"totalInterest"`
	Description    string `json:"description"`
}

type AiRespond struct {
	AiSuggestion string     `json:"aiSuggestion"`
	LoanPlans    []LoanPlan `json:"loanPlans"`
}
