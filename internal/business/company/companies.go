package company

import (
	"time"
)

type Companies struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCompanies(
	id uint64,
	name string,
	description string,
	createdAt time.Time,
	updatedAt time.Time,
) *Companies {
	return &Companies{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
