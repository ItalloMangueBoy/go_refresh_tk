package user

import "gorm.io/gorm"

type RepositoryGORM struct {
	db *gorm.DB	
}

func NewGormRepository(db *gorm.DB) Repository {
	return &RepositoryGORM{db: db}
}
