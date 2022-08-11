package user

import (
	"context"

	"github.com/upgradeskill/beta-team/internal/data/user"
)

type IUserRepository interface {
	CreateUser(user *Users) error
	GetUser(id uint64) (*Users, error)
	ListUsers() ([]Users, error)
	UpdateUser(user *Users) error
	DeleteUser(id uint64) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, input user.RegisterRequest) error
	GetUser(id uint64) (*Users, error)
	ListUsers() ([]Users, error)
	UpdateUser(user *Users) error
	DeleteUser(id uint64) error
}
