package user

type RegisterRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	RoleID          uint64 `json:"role_id"`
	CompanyID       uint64 `json:"company_id"`
}
