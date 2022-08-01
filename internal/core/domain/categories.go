package domain

import (
	"time"
)

type Categories struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	UserID    uint64    `json:"id_user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Categories) TableName() string {
	return "categories"
}

func NewCategories(
	id uint64,
	name string,
	userID uint64,
	createdAt time.Time,
	updatedAt time.Time,
) *Categories {
	return &Categories{
		ID:        id,
		Name:      name,
		UserID:    userID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
