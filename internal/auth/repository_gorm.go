package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(refreshToken *RefreshToken) error {
	return r.db.Create(refreshToken).Error
}

func (r *gormRepository) GetByID(id uuid.UUID) (*RefreshToken, error) {
	var token RefreshToken
	if err := r.db.First(&token, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *gormRepository) ListByUserID(userID uuid.UUID) ([]RefreshToken, error) {
	var tokens []RefreshToken
	if err := r.db.Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *gormRepository) Revoke(id uuid.UUID) error {
	return r.db.Model(&RefreshToken{}).Where("id = ? AND revoked = ?", id, false).Update("revoked", true).Error
}

func (r *gormRepository) RevokeAllByUserID(userID uuid.UUID) error {
	return r.db.Model(&RefreshToken{}).Where("user_id = ? AND revoked = ?", userID, false).Update("revoked", true).Error
}

func (r *gormRepository) Replace(oldTokenID uuid.UUID, newToken *RefreshToken) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&RefreshToken{}).Where("id = ?", oldTokenID).Update("revoked", true).Error; err != nil {
			return err
		}

		if err := tx.Create(newToken).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *gormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&RefreshToken{}, "id = ?", id).Error
}

func (r *gormRepository) DeleteExpiredTokens() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&RefreshToken{}).Error
}
