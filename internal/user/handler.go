package user

import (
	"errors"
	"refresh_token/internal/validation"
	"refresh_token/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var input CreateDTO

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

func (h *Handler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.RespondBadRequest(c, err)
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.RespondNotFound(c, err)
			return
		}
		response.RespondInternalError(c, err)
		return
	}

	response.RespondSuccess(c, user)
}

func (h *Handler) GetByEmail(c *gin.Context) {
	var input GetByEmailDTO
	if err := c.ShouldBindQuery(&input); err != nil {
		parsedError := validation.ParseErrors(err)
		response.RespondError(c, parsedError)
		return
	}

	user, err := h.service.GetByEmail(&input)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.RespondNotFound(c, err)
			return
		}
		response.RespondInternalError(c, err)
		return
	}

	response.RespondSuccess(c, user)
}
