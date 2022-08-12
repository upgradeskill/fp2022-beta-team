package domain

import (
	"time"
)

type Companies struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCompany(
	id uint64,
	name string,
	desc string,
	createdat time.Time,
	updatedat time.Time,
) *Companies {
	return &Companies{
		ID:        id,
		Name:      name,
		Desc:      desc,
		CreatedAt: createdat,
		UpdatedAt: updatedat,
	}
}
