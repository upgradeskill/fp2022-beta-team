package domain

import (
	"time"
)

type Items struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	PurchasePrice string    `json:"purchase_price"`
	SellingPrice  string    `json:"selling_price"`
	ItemSKU       string    `json:"item_sku"`
	ItemUnit      string    `json:"item_unit"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Items) TableName() string {
	return "items"
}

func NewItems(
	id uint64,

	createdAt time.Time,
	updatedAt time.Time,
) *Items {
	return &Items{
		ID: id,

		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
