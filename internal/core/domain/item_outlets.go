package domain

import (
	"time"
)

type ItemOutlets struct {
	ID        uint64    `json:"id"`
	ItemID    uint64    `json:"item_id"`
	OutletID  uint64    `json:"outlet_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (ItemOutlets) TableName() string {
	return "item_outlets"
}

func NewItemOutlets(
	id uint64,
	ItemID uint64,
	OutletID uint64,
	createdAt time.Time,
) *ItemOutlets {
	return &ItemOutlets{
		ID:        id,
		ItemID:    ItemID,
		OutletID:  OutletID,
		CreatedAt: createdAt,
	}
}
