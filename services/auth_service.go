package services

import (
	"exhibitly/repositories"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	SessionRepo *repositories.SessionRepository
}

func (s *AuthService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(os.Getenv("secretKey"))
	fmt.Println(tokenString, err)
	return tokenString, err
}
