package app

func (a *App) Start() error {
	a.registerModules()

	return a.router.Run(":8080")
}
