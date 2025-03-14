package models

import (
	"time"
)

type Stock struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Amount    uint       `json:"image"`

	ProductID string  `json:"product_id" gorm:"index"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
