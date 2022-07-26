package service

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
	userPort "github.com/upgradeskill/beta-team/internal/core/ports/user"
)

type UserServiceImpl struct {
	userRepo userPort.UserRepository
}

func NewUserService(userRepo userPort.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u *UserServiceImpl) GetUser(id uint64) (*domain.User, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) ListUsers() ([]*domain.User, error) {
	listUser, err := u.userRepo.ListUsers()
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (u *UserServiceImpl) CreateUser(user *domain.User) (*domain.User, error) {
	user = domain.NewUser(user.ID, user.Name, user.Role)

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) UpdateUser(user *domain.User) (*domain.User, error) {
	user = domain.NewUser(user.ID, user.Name, user.Role)

	err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) DeleteUser(id uint64) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
