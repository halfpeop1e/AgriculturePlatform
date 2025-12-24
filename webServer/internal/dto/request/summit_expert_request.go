package request

type SubmitExpertProfileRequest struct {
	Name          string   `json:"name" binding:"required"`
	Avatar        string   `json:"avatar" binding:"required"`
	Field         string   `json:"field" binding:"required"`
	Introduction  string   `json:"introduction" binding:"required"`
	Skills        []string `json:"skills" binding:"required"`
	Education     string   `json:"education"`
	Experience    string   `json:"experience"`
	Certification []string `json:"certification"`
	AvailableTime []string `json:"availableTime"`
	ServiceScope  []string `json:"serviceScope"`
	Price         *float64 `json:"price"`
	Completed     bool     `json:"completed"`
}
