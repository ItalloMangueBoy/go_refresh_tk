package auth

import (
	"refresh_token/config"
	"refresh_token/internal/user"
	"refresh_token/pkg/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewGormRepository(db)
	userRepo := user.NewGormRepository(db)

	accessMgr := token.NewAccessTokenManager(config.AccessTokenSecret, config.AccessTokenTTL)
	refreshMgr := token.NewRefreshTokenManager(config.RefreshTokenLength)

	service := NewService(repo, userRepo, accessMgr, refreshMgr)
	handler := NewHandler(service)

	authRouter := rg.Group("/auth")

	RegisterRoutes(authRouter, handler)
}
