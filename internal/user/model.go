package user

import (
	"refresh_token/pkg/encrypt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"not null;uniqueIndex:idx_email"`
	Password  string         `gorm:"not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := encrypt.Hash(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return nil
}

func (u *User) VerifyPassword(password string) error {
	return encrypt.Verify(password, u.Password)
}

func (u *User) ToResponse() ResponseDTO {
	return ResponseDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
