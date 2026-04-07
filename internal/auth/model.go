package auth

import (
	"refresh_token/internal/user"
	"refresh_token/pkg/encrypt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	User      user.User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Token     string         `gorm:"not null;uniqueIndex:idx_token" json:"-"`
	Revoked   bool           `gorm:"not null;default:false"`
	ExpiresAt time.Time      `gorm:"not null;index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if rt.ID == uuid.Nil {
		rt.ID = uuid.New()
	}
	return nil
}

func (rt *RefreshToken) SetToken(token string) error {
	hashedToken, err := encrypt.Hash(token)
	if err != nil {
		return err
	}
	rt.Token = hashedToken
	return nil
}

func (rt *RefreshToken) VerifyToken(token string) error {
	return encrypt.Verify(token, rt.Token)
}

func (rt *RefreshToken) Revoke() {
	rt.Revoked = true
}

func (rt *RefreshToken) TimeToLive() time.Duration {
	ttl := time.Until(rt.ExpiresAt)
	if ttl < 0 {
		return 0
	}
	return ttl
}

func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

func (rt *RefreshToken) IsRevoked() bool {
	return rt.Revoked
}

func (rt *RefreshToken) IsValid() bool {
	return !rt.IsExpired() && !rt.IsRevoked()
}

func (rt *RefreshToken) ToResponse() ResponseDTO {
	return ResponseDTO{
		ID:        rt.ID,
		UserID:    rt.UserID,
		Revoked:   rt.Revoked,
		ExpiresAt: rt.ExpiresAt,
		CreatedAt: rt.CreatedAt,
		UpdatedAt: rt.UpdatedAt,
	}
}
