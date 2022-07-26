package ports

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUser(id uint64) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uint64) error
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUser(id uint64) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uint64) error
}
