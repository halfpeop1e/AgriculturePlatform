package request

type SafeUpdateUserInfoRequest struct {
	Email           string `json:"email"`
	Currentpassword string `json:"currentpassword"`
	Newpassword     string `json:"newpassword"`
}
