package auth

import (
	"time"

	"refresh_token/pkg/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewGormRepository(db)

	// Em produção, estas chaves/configs virão do os.Getenv() ou viper
	accessMgr := token.NewAccessTokenManager("secretapikey", 15*time.Minute)
	refreshMgr := token.NewRefreshTokenManager(32)

	service := NewService(repo, accessMgr, refreshMgr)
	handler := NewHandler(service)

	authRouter := rg.Group("/auth")

	RegisterRoutes(authRouter, handler)
}
