package utils

import (
	"Go-Starter-Template/pkg/entities/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

func GenerateToken(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(domain.JwtSecret))
}
