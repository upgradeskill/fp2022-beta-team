package ports

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
)

type IAuthRepository interface {
	CreateUser(user *domain.Users) error
	GetUser(id uint64) (*domain.Users, error)
	ListUsers() ([]*domain.Users, error)
	UpdateUser(user *domain.Users) error
	DeleteUser(id uint64) error
}

type IAuthUsecase interface {
	CreateUser(user *domain.Users) error
	GetUser(id uint64) (*domain.Users, error)
	ListUsers() ([]*domain.Users, error)
	UpdateUser(user *domain.Users) error
	DeleteUser(id uint64) error
}
