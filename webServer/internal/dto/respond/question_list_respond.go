package respond

type QuestionRespond struct {
	Id       string `json:"id"`
	ExpertId string `json:"expertId"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	CaseDate string `json:"caseDate"`
}

type Answer struct {
	Id         string `json:"id"`
	QuestionId string `json:"questionId"`
	Content    string `json:"content"`
	Expert     string `json:"expert"`
	ExpertId   string `json:"expertId"`
	Date       string `json:"date"`
	IsAccpeted bool   `json:"isAccpeted"`
}

type QuestionDetailRespond struct {
	Id          string   `json:"id"`
	ExpertId    string   `json:"expertId"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Author      string   `json:"author"`
	AuthorId    string   `json:"authorId"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
	AnswerCount int      `json:"answerCount"`
	IsAnswered  bool     `json:"isAnswered"`
	Answers     []Answer `json:"answers"`
}
