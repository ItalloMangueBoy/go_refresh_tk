package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenManager interface {
	GenerateToken(payload Payload) (string, error)
	ValidateToken(tokenStr string) (*Payload, error)
}

type jwtManager struct {
	secretKey []byte
	ttl       time.Duration
}

func NewAccessTokenManager(secretKey string, ttl time.Duration) AccessTokenManager {
	return &jwtManager{
		secretKey: []byte(secretKey),
		ttl:       ttl,
	}
}

func (m *jwtManager) GenerateToken(payload Payload) (string, error) {
	claims := Payload{
		UserID: payload.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.ttl)),
			Issuer:    "auth_service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (m *jwtManager) ValidateToken(tokenStr string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
