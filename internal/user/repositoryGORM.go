package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepositoryGORM struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &RepositoryGORM{db: db}
}

func (repo RepositoryGORM) Create(user *User) error {
	err := repo.db.Create(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrUserAlreadyExists
		}

		return err
	}

	return nil
}

func (repo RepositoryGORM) Update(user *User) error {
	return repo.db.Save(user).Error
}

func (repo RepositoryGORM) Delete(user *User) error {
	return repo.db.Delete(user).Error
}

func (repo RepositoryGORM) GetByID(id uuid.UUID) (*User, error) {
	var user User
	err := repo.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repo RepositoryGORM) GetByEmail(email string) (*User, error) {
	var user User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repo RepositoryGORM) List() ([]User, error) {
	var users []User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
