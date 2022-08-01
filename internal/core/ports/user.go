package ports

import (
	"context"

	"github.com/upgradeskill/beta-team/internal/core/domain"
	"github.com/upgradeskill/beta-team/internal/requests"
)

type IUserRepository interface {
	CreateUser(user *domain.Users) error
	GetUser(id uint64) (*domain.Users, error)
	ListUsers() ([]*domain.Users, error)
	UpdateUser(user *domain.Users) error
	DeleteUser(id uint64) error
}

type IUserUsecase interface {
	RegisterUser(ctx context.Context, req *requests.RegisterRequest) error
	CreateUser(user *domain.Users) error
	GetUser(id uint64) (*domain.Users, error)
	ListUsers() ([]*domain.Users, error)
	UpdateUser(user *domain.Users) error
	DeleteUser(id uint64) error
}
