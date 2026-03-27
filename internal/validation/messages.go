package validation

import "github.com/go-playground/validator/v10"

func message(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Required field"
	case "email":
		return "Invalid email address"
	default:
		return "Invalid field"
	}
}