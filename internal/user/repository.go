package user

import "github.com/google/uuid"

type Repository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	FindByID(user *User, id uuid.UUID) error
	FindByEmail(user *User, email string) error
	List(users *[]User) error
}
