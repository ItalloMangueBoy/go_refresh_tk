package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Error   *APIError   `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, APIResponse{
		Success: true,
		Data:    data,
	})
}

func RespondCreated(c *gin.Context, data interface{}) {
	c.JSON(201, APIResponse{
		Success: true,
		Data:    data,
	})
}
