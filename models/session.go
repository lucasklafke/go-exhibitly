package models

import "time"

// Session representa uma sessão de usuário.
type Session struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string    `gorm:"not null"`
	Token     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	ExpiresAt time.Time `gorm:"not null"`
}
