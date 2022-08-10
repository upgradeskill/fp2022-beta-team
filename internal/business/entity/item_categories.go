package entity

import (
	"time"
)

type ItemCategories struct {
	ID         uint64    `json:"id"`
	ItemID     uint64    `json:"item_id"`
	CategoryID uint64    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (ItemCategories) TableName() string {
	return "item_categories"
}

func NewItemCategories(
	id uint64,
	ItemID uint64,
	CategoryID uint64,
	createdAt time.Time,
) *ItemCategories {
	return &ItemCategories{
		ID:         id,
		ItemID:     ItemID,
		CategoryID: CategoryID,
		CreatedAt:  createdAt,
	}
}
