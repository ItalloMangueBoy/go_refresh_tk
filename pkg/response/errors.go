package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func RespondError(c *gin.Context, err *APIError) {
	c.JSON(err.Code, APIResponse{
		Success: false,
		Error:   err,
	})
}

func RespondInternalError(c *gin.Context, err error) {
	RespondError(c, &APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Details: err.Error(),
	})
}

func RespondConflict(c *gin.Context, err error) {
	RespondError(c, &APIError{
		Code:    http.StatusConflict,
		Message: "Conflict",
		Details: err.Error(),
	})
}

func RespondBadRequest(c *gin.Context, err error) {
	RespondError(c, &APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
		Details: err.Error(),
	})
}

func RespondNotFound(c *gin.Context, err error) {
	RespondError(c, &APIError{
		Code:    http.StatusNotFound,
		Message: "Not found",
		Details: err.Error(),
	})
}
