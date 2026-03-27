package dto

type CreateInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	// ConfirmPassword string `json:"confirm_password" binding:"required"`
}
