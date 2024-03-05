package planet

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/planet/log"
)

type App struct {
	fiber *fiber.App
}

func New(cfg ...fiber.Config) *App {
	app := fiber.New(cfg...)
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

func (app *App) With(wrap func(app *fiber.App)) {
	wrap(app.fiber)
}

func (app *App) Router() fiber.Router {
	return app.fiber
}
