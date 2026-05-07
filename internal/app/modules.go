package app

import (
	"refresh_token/internal/auth"
	"refresh_token/internal/user"
)

func (a *App) registerModules() {
	api := a.router.Group("/api")

	user.Register(api, a.db)
	auth.Register(api, a.db)
}