package repository

import (
	"context"
	"time"

	"github.com/upgradeskill/beta-team/internal/core/domain"
	"github.com/upgradeskill/beta-team/internal/requests"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{DB: db}
}

// func (ur *userRepository) Login(ctx context.Context, input *requests.LoginRequest) (string, error) {

// 	return ur.DB.Create(user).Error
// }

func (ur *userRepository) CreateUser(ctx context.Context, input *requests.CreateUserRequest) error {
	user := &domain.Users{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		RoleID:    input.RoleID,
		CompanyID: input.CompanyID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return ur.DB.Create(user).Error
}

// func (ur *userRepository) CreateUser(user *user.Users) error {
// 	return ur.DB.Create(user).Error
// }

// func (ur *userRepository) GetUser(id uint64) (*user.Users, error) {
// 	var user user.Users
// 	err := ur.DB.First(&user, id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (ur *userRepository) ListUsers() ([]*user.Users, error) {
// 	var users []*user.Users
// 	err := ur.DB.Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func (ur *userRepository) UpdateUser(user *user.Users) error {
// 	return ur.DB.Save(user).Error
// }

// func (ur *userRepository) DeleteUser(id uint64) error {
// 	return ur.DB.Delete(&user.Users{ID: id}).Error
// }
