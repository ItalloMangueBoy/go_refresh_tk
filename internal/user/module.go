package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewGormRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	userRouter := rg.Group("/users")

	RegisterRoutes(userRouter, handler)
}
