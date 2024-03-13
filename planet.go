package planet

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/planet/errors"
	"github.com/lowl11/planet/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	fiber    *fiber.App
	shutdown chan struct{}
}

func New(cfg ...fiber.Config) *App {
	var handlerCfg *fiber.Config
	if len(cfg) > 0 {
		handlerCfg = &cfg[0]
	} else {
		handlerCfg = &fiber.Config{}
	}

	handlerCfg.ReadBufferSize = 20480
	handlerCfg.ErrorHandler = func(ctx *fiber.Ctx, err error) error {
		log.Error(err)

		if planetErr, ok := err.(errors.Error); ok {
			ctx.Response().Header.Set("Content-Type", "application/json")
			ctx.Response().SetStatusCode(planetErr.HttpCode())
			return ctx.Send(planetErr.Output())
		}

		return errors.
			New(err.Error()).
			SetHTTP(http.StatusInternalServerError)
	}

	app := fiber.New(*handlerCfg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	shutdown := make(chan struct{})

	go func() {
		<-c

		log.Info("Gracefully shutting down...")

		shutdownErr := app.Shutdown()
		if shutdownErr != nil {
			panic(shutdownErr)
		}

		log.Info("API server stopped")

		shutdown <- struct{}{}
	}()

	return &App{
		fiber:    app,
		shutdown: shutdown,
	}
}

func (app *App) Run(port string) {
	go func() {
		if err := app.fiber.Listen(port); err != nil {
			log.Fatal("Run app error: ", err)
		}
	}()

	<-app.shutdown
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
