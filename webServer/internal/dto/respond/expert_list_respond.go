package respond

// ExpertRespond 专家列表响应
type ExpertRespond struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Avatar       string        `json:"avatar"`
	Field        string        `json:"field"`
	Introduction string        `json:"introduction"`
	Skills       []string      `json:"skills"`
	ConsultCount int           `json:"consultCount"`
	Rating       float64       `json:"rating"`
	ResponseTime string        `json:"responseTime,omitempty"`
	RecentCases  []CaseRespond `json:"recentCases,omitempty"`
}

type ExpertListRespond struct {
	List     []ExpertRespond `json:"list"`
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	HasMore  bool            `json:"hasMore"`
}

// ExpertDetailRespond 专家详情响应
type ExpertDetailRespond struct {
	Id            string        `json:"id"`
	Name          string        `json:"name"`
	Avatar        string        `json:"avatar"`
	Field         string        `json:"field"`
	Introduction  string        `json:"introduction"`
	Skills        []string      `json:"skills"`
	Education     string        `json:"education,omitempty"`
	Experience    string        `json:"experience,omitempty"`
	Certification []string      `json:"certification,omitempty"`
	ConsultCount  int           `json:"consultCount"`
	Rating        float64       `json:"rating"`
	ResponseTime  string        `json:"responseTime,omitempty"`
	AvailableTime []string      `json:"availableTime,omitempty"`
	ServiceScope  []string      `json:"serviceScope,omitempty"`
	Price         float64       `json:"price,omitempty"`
	RecentCases   []CaseRespond `json:"recentCases,omitempty"`
}

// CaseRespond 专家案例响应
type CaseRespond struct {
	Date     string `json:"date"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
