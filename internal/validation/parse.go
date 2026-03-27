package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"refresh_token/pkg/response"

	"github.com/go-playground/validator/v10"
)

func ParseErrors(err error) *response.APIError {
	ve := validator.ValidationErrors{}

	if errors.As(err, &ve) {
		fields := make(map[string]string)

		for _, fe := range ve {
			fields[fe.Field()] = message(fe)
		}

		return &response.APIError{
			Code:    http.StatusUnprocessableEntity,
			Message: "Validation failed",
			Details: fields,
		}
	}

	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return &response.APIError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid type for field %s", ute.Field),
		}
	}

	return &response.APIError{
		Code:    http.StatusBadRequest,
		Message: "Invalid request body",
	}
}
