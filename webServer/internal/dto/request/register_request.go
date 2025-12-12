package request

type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	ComPassword string `json:"compassword"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}
