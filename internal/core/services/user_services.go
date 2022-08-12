package services

import (
	"context"
	"errors"

	"github.com/upgradeskill/beta-team/internal/core/ports"
	"github.com/upgradeskill/beta-team/internal/requests"
)

type UserService struct {
	userRepository    ports.IUserRepository
	companyRepository ports.ICompanyRepository
}

func NewUserService(repository ports.IUserRepository, companyRepo ports.ICompanyRepository) *UserService {
	return &UserService{
		userRepository:    repository,
		companyRepository: companyRepo,
	}
}

func (us *UserService) Register(ctx context.Context, input *requests.RegisterRequest) error {
	if input.Password != input.ConfirmPassword {
		return errors.New("password and confirm password not match")
	}

	company := &requests.CreateCompanyRequest{
		Name: input.CompanyName,
		Desc: input.CompanyDesc,
	}

	companyID, err := us.companyRepository.CreateCompany(ctx, company)
	if err != nil {
		return errors.New("failed to create company")
	}

	user := requests.CreateUserRequest{
		FirstName:       input.FirstName,
		LastName:        input.LastName,
		Email:           input.Email,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
		RoleID:          input.RoleID,
		CompanyID:       companyID,
	}

	err = us.userRepository.CreateUser(ctx, &user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

// func (us *UserService) Login(ctx context.Context, input *requests.LoginRequest) (*string, error) {
// 	token, err := us.userRepository.Login(ctx, input)
// 	if err != nil {
// 		return nil, errors.New("failed to login")
// 	}

// 	return &token, nil
// }
