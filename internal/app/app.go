package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB) *App {
	router := gin.Default()

	return &App{
		db:     db,
		router: router,
	}
}
