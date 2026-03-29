package user

import "github.com/google/uuid"

type Repository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	GetByID(id uuid.UUID) (user *User, err error)
	GetByEmail(email string) (user *User, err error)
	List() (users []User, err error)
}
