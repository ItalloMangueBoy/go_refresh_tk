package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.POST("/", h.Create)
	rg.GET("/", h.GetSearch)
	rg.GET("/:id", h.GetByID)
	rg.GET("/email/:email", h.GetByEmail)
	// rg.PUT("/:id", h.Update)
	// rg.DELETE("/:id", h.Delete)
}
