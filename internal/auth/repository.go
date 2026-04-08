package auth

import (
	"github.com/google/uuid"
)

type Repository interface {
	Create(refreshToken *RefreshToken) error
	GetByID(id uuid.UUID) (*RefreshToken, error)
	ListByUserID(userID uuid.UUID) ([]RefreshToken, error)
	Revoke(id uuid.UUID) error
	RevokeAllByUserID(userID uuid.UUID) error
	Replace(oldTokenID uuid.UUID, newToken *RefreshToken) error
	Delete(id uuid.UUID) error
	DeleteExpiredTokens() error
}
