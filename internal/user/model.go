package user

import (
	"refresh_token/internal/user/dto"
	"refresh_token/pkg/hash"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"not null;uniqueIndex:idx_email"`
	Password string    `gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return nil
}

func (u *User) VerifyPassword(password string) error {
	return hash.VerifyPassword(password, u.Password)
}

func (u *User) Public() dto.UserResponse {
	return dto.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
