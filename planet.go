package planet

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/planet/log"
)

type App struct {
	fiber *fiber.App
}

func New() *App {
	app := fiber.New()
	return &App{
		fiber: app,
	}
}

func (app *App) Run(port string) {
	if err := app.fiber.Listen(port); err != nil {
		log.Fatal("Run app error: ", err)
	}
}

func (app *App) Fiber() *fiber.App {
	return app.fiber
}

func (app *App) Router() fiber.Router {
	return app.fiber
}
