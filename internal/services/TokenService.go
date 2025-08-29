package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenServiceInterface interface {
	GenerateToken(userId int64) (string, error)
	ParseToken(tokenString string) (*jwt.Token, error)
	GetDuration() time.Duration
}

type TokenService struct {
	SecretKey string
	Duration time.Duration
}

func NewTokenService(secretKey string, duration time.Duration) *TokenService {
	return &TokenService{SecretKey: secretKey, Duration: duration}
}

func (t *TokenService) GenerateToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(t.Duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.SecretKey))
}

func (t *TokenService) ParseToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(t.SecretKey), nil
    })
}

func (t *TokenService) GetDuration() time.Duration {
	return t.Duration
}