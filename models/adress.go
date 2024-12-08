package models

import "time"

type Address struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null;index"`
	PostalCode   string    `json:"postal_code" gorm:"type:varchar(10);not null;unique"`
	Street       string    `json:"street" gorm:"type:varchar(100);not null"`
	Number       string    `json:"number" gorm:"type:varchar(20);not null"`
	Complement   string    `json:"complement" gorm:"type:varchar(100)"`
	Neighborhood string    `json:"neighborhood" gorm:"type:varchar(100);not null"`
	City         string    `json:"city" gorm:"type:varchar(100);not null"`
	State        string    `json:"state" gorm:"type:char(2);not null"`
	Country      string    `json:"country" gorm:"type:varchar(50);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
