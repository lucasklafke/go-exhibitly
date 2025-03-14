package models

import (
	"time"
)

type Product struct {
	ID          string     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Title       string     `json:"title" gorm:"unique"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Slug        string     `json:"slug" gorm:"unique"`
	CategoryID  string     `json:"category_id" gorm:"index"`
	Category    Category   `json:"category" gorm:"foreignKey:CategoryID"`
}
