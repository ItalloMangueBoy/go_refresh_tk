package user

import (
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
	return repo.db.Create(user).Error
}

func (repo RepositoryGORM) Update(user *User) error {
	return repo.db.Save(user).Error
}

func (repo RepositoryGORM) Delete(user *User) error {
	return repo.db.Delete(user).Error
}

func (repo RepositoryGORM) FindByID(user *User, id uuid.UUID) error {
	return repo.db.First(user, id).Error
}

func (repo RepositoryGORM) FindByEmail(user *User, email string) error {
	return repo.db.Where("email = ?", email).First(user).Error
}

func (repo RepositoryGORM) List(users *[]User) error {
	return repo.db.Find(users).Error
}
