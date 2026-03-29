package user

import "github.com/google/uuid"

type ResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type CreateDTO struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

func (dto CreateDTO) ToModel() (*User, error) {
	user := &User{
		Name:  dto.Name,
		Email: dto.Email,
	}

	if err := user.SetPassword(dto.Password); err != nil {
		return nil, err
	}

	return user, nil
}
