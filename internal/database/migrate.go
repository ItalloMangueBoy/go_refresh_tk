package database

import (
	"refresh_token/internal/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		user.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
