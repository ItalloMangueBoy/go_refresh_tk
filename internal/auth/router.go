package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *handler) {
	rg.POST("/login", h.Login)
	// rg.POST("/refresh", h.Refresh)
	// rg.POST("/logout", h.Logout)

	// rg.GET("/tokens/:id", h.GetTokenByID)
	// rg.GET("/tokens/user/:id", h.ListTokensByUserID)

	// rg.PUT("/tokens/:id/revoke", h.RevokeToken)
	// rg.PUT("/tokens/user/:id/revoke-all", h.RevokeAllTokens)
	// rg.DELETE("/tokens/:id", h.DeleteToken)
}
