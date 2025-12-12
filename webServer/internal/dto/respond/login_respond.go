package respond

type LoginRespond struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	Tokens string `json:"tokens"`
}
