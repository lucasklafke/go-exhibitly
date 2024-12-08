package models

import "time"

type Category struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
}
