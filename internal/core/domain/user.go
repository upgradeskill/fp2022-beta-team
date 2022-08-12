package domain

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

func NewPerson(
	id uint64,
	email string,
	password string,
	firstname string,
	lastname string,
	roleid uint64,
	companyid uint64,
	createdat time.Time,
	updatedat time.Time,
) *Users {
	return &Users{
		ID:        id,
		Email:     email,
		Password:  password,
		FirstName: firstname,
		LastName:  lastname,
		RoleID:    roleid,
		CompanyID: companyid,
		CreatedAt: createdat,
		UpdatedAt: updatedat,
	}
}

func (u *Users) GetEmail() string {
	return u.Email
}
