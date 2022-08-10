package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/upgradeskill/beta-team/internal/core/domain"
	port "github.com/upgradeskill/beta-team/internal/core/ports"
	"github.com/upgradeskill/beta-team/internal/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	userRepo    port.IUserRepository
	companyRepo port.ICompanyRepository
}

func NewUserService(userRepo port.IUserRepository, companyRepo port.ICompanyRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userRepo:    userRepo,
		companyRepo: companyRepo,
	}
}

func (u *UserUsecaseImpl) RegisterUser(ctx context.Context, input *requests.RegisterRequest) error {
	if input.Password != input.ConfirmPassword {
		return errors.New("the passwords are not equal")
	}

	checkCompany, err := u.companyRepo.GetCompanyByID(input.CompanyID)
	if err != nil {
		return err
	}

	if checkCompany.ID == 0 {
		return errors.New("company not found")
	}

	user := domain.NewUsers()

	return nil
}

func (u *UserUsecaseImpl) GetUser(id uint64) (*domain.Users, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecaseImpl) ListUsers() ([]*domain.Users, error) {
	listUser, err := u.userRepo.ListUsers()
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (u *UserUsecaseImpl) CreateUser(user *domain.Users) (*domain.Users, error) {
	user = domain.NewUsers(
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.RoleID,
		user.CompanyID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecaseImpl) UpdateUser(user *domain.Users) (*domain.Users, error) {
	user = domain.NewUsers(
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.RoleID,
		user.CompanyID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecaseImpl) DeleteUser(id uint64) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}
