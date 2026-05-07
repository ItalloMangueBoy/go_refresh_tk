package auth

import (
	"errors"
	"refresh_token/internal/user"
	"refresh_token/internal/validation"
	"refresh_token/pkg/response"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (n handler) Login(c *gin.Context) {
	var input LoginRequestDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		parsedError := validation.ParseErrors(err)
		response.RespondError(c, parsedError)
		return
	}

	result, err := n.service.Login(&input)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) || errors.Is(err, ErrInvalidCredentials) {
			response.RespondUnauthorized(c, err)
			return
		}

		response.RespondInternalError(c, err)
		return
	}

	response.RespondSuccess(c, result)
}
