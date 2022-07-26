package user

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
)

type Users struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type UserList []Users

func (us *Users) FromDomain(user *domain.User) {
	if us == nil {
		us = &Users{}
	}

	us.ID = user.ID
	us.Name = user.Name
	us.Role = user.Role
}

func (us *Users) ToDomain() *domain.User {
	if us == nil {
		us = &Users{}
	}

	return &domain.User{
		ID:   us.ID,
		Name: us.Name,
		Role: us.Role,
	}
}

func (us UserList) FromDomain(users []domain.User) UserList {
	for _, u := range users {
		user := Users{}
		user.FromDomain(&u)
		us = append(us, user)
	}

	return us
}
