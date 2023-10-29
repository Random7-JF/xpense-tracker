package main

import (
	"github.com/Random7-JF/xpense-tracker/cmd/web/handlers"
	"github.com/Random7-JF/xpense-tracker/cmd/web/middleware"
	"github.com/Random7-JF/xpense-tracker/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func createWeb(app *config.App) {
	addEngine()
	addRoutes()
	handlers.NewHandlers(app)
	middleware.NewMiddleware(app)
}

func addEngine() {
	engine := html.New("./views", ".html")
	engine.Reload(true)

	App.Web = fiber.New(fiber.Config{
		Views: engine,
	})
}
