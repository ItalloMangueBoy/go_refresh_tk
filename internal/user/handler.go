package user

import (
	"errors"
	"refresh_token/internal/user/dto"
	"refresh_token/internal/validation"
	"refresh_token/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var input dto.CreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		parsedError := validation.ParseErrors(err)
		response.RespondError(c, parsedError)
		return
	}

	user, err := h.service.Create(&input)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			response.RespondConflict(c, err)
			return
		}
		response.RespondInternalError(c, err)
		return
	}

	response.RespondCreated(c, user)
}
