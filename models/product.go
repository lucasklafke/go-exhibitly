package models

import (
	"time"
)

type Product struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Title       string     `json:"title"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	Image       string     `json:"image"`

	// AuthorID string `json:"author"`
	// Author   User
}
