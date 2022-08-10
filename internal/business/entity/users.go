package entity

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

func NewUsers(
	id uint64,
	mail string,
	password string,
	firstName string,
	lastName string,
	roleID uint64,
	companyID uint64,
	createdAt time.Time,
	updatedAt time.Time,
) *Users {
	return &Users{
		ID:        id,
		Email:     mail,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		RoleID:    roleID,
		CompanyID: companyID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
