package user

import (
	"time"
)

type Users struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	RoleID    uint64    `json:"role_id"`
	CompanyID uint64    `json:"id_company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
