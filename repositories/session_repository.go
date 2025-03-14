package repositories

import (
	"exhibitly/models"
	"time"

	"gorm.io/gorm"
)

type SessionRepository struct {
	DB *gorm.DB
}

// CreateSession cria uma nova sessão.
func (r *SessionRepository) CreateSession(session *models.Session) error {
	return r.DB.Create(session).Error
}

// FindSessionByToken busca uma sessão pelo token.
func (r *SessionRepository) FindSessionByToken(token string) (*models.Session, error) {
	var session models.Session
	err := r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// DeleteSessionByToken remove uma sessão pelo token.
func (r *SessionRepository) DeleteSessionByToken(token string) error {
	return r.DB.Where("token = ?", token).Delete(&models.Session{}).Error
}

// DeleteExpiredSessions limpa sessões expiradas.
func (r *SessionRepository) DeleteExpiredSessions() error {
	return r.DB.Where("expires_at < ?", time.Now()).Delete(&models.Session{}).Error
}
