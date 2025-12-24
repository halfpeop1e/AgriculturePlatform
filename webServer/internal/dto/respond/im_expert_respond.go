package respond

// ExpertProfileRespond 专家资料响应（用于用户获取自己的专家资料）
type ExpertProfileRespond struct {
	Name          string   `json:"name"`
	Avatar        string   `json:"avatar"`
	Field         string   `json:"field"`
	Introduction  string   `json:"introduction"`
	Skills        []string `json:"skills"`
	Education     string   `json:"education,omitempty"`
	Experience    string   `json:"experience,omitempty"`
	Certification []string `json:"certification,omitempty"`
	AvailableTime []string `json:"availableTime,omitempty"`
	ServiceScope  []string `json:"serviceScope,omitempty"`
	Price         float64  `json:"price,omitempty"`
	Completed     bool     `json:"completed"`
}
