package repository

import (
	"time"

	"github.com/upgradeskill/beta-team/internal/business/user"
	"gorm.io/gorm"
)

type Users struct {
	ID        uint64    `json:"id" gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	FirstName string    `json:"first_name" gorm:"column:first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name"`
	RoleID    uint64    `json:"role_id" gorm:"column:role_id"`
	CompanyID uint64    `json:"id_company" gorm:"column:company_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{DB: db}
}

func (ur *userRepository) CreateUser(user *user.Users) error {
	return ur.DB.Create(user).Error
}

func (ur *userRepository) GetUser(id uint64) (*user.Users, error) {
	var user user.Users
	err := ur.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) ListUsers() ([]*user.Users, error) {
	var users []*user.Users
	err := ur.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) UpdateUser(user *user.Users) error {
	return ur.DB.Save(user).Error
}

func (ur *userRepository) DeleteUser(id uint64) error {
	return ur.DB.Delete(&user.Users{ID: id}).Error
}
