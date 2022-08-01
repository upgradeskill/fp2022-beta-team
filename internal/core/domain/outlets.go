package domain

import (
	"time"
)

type Outlets struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	UserID    uint64    `json:"user_id"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Outlets) TableName() string {
	return "outlets"
}

func NewOutlets(
	id uint64,
	name string,
	userID uint64,
	isPrimary bool,
	createdAt time.Time,
	updatedAt time.Time,
) *Outlets {
	return &Outlets{
		ID:        id,
		Name:      name,
		UserID:    userID,
		IsPrimary: isPrimary,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
