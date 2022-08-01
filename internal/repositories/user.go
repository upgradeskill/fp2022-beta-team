package user

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{DB: db}
}

func (ur *userRepository) CreateUser(user *domain.Users) error {
	return ur.DB.Create(user).Error
}

func (ur *userRepository) GetUser(id uint64) (*domain.Users, error) {
	var user domain.Users
	err := ur.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) ListUsers() ([]*domain.Users, error) {
	var users []*domain.Users
	err := ur.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) UpdateUser(user *domain.Users) error {
	return ur.DB.Save(user).Error
}

func (ur *userRepository) DeleteUser(id uint64) error {
	return ur.DB.Delete(&domain.Users{ID: id}).Error
}
