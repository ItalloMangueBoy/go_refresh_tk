package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.POST("/", h.Create)
	// rg.PUT("/:id", h.Update)
	// rg.DELETE("/:id", h.Delete)
	// rg.GET("/", h.GetAll)
	// rg.GET("/:id", h.GetByID)
	// rg.GET("/email/:email", h.GetByEmail)
}