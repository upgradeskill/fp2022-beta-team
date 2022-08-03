package requests

type RegisterRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	CompanyID       uint64 `json:"company_id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
