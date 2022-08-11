package user

import (
	"context"
	"time"

	"github.com/upgradeskill/beta-team/internal/data/user"
	"github.com/upgradeskill/beta-team/internal/errs"
)

type UserServiceImpl struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u *UserServiceImpl) RegisterUser(ctx context.Context, input user.RegisterRequest) error {
	if input.Password != input.ConfirmPassword {
		return errs.New("password and confirm password not match", 400)
	}
	user := &Users{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		RoleID:    input.RoleID,
		CompanyID: input.CompanyID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserServiceImpl) GetUser(id uint64) (*Users, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) ListUsers() ([]Users, error) {
	listUser, err := u.userRepo.ListUsers()
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (u *UserServiceImpl) UpdateUser(user *Users) (*Users, error) {
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
