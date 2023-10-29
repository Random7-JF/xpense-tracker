package main

import (
	"github.com/Random7-JF/xpense-tracker/cmd/web/handlers"
	"github.com/Random7-JF/xpense-tracker/cmd/web/middleware"
	"github.com/gofiber/fiber/v2"
)

func addRoutes() {
	mw := middleware.NewMiddleware(&App)
	App.Web.Use(mw.SetupSession())

	App.Web.Get("/", handlers.Index)
	App.Web.Get("/login", handlers.Login)
	App.Web.Post("/login", handlers.PostLogin)
	App.Web.Get("/register", handlers.Register)
	App.Web.Post("/register", handlers.PostRegister)

	app := App.Web.Group("/app")
	app.Use(mw.SaveSession(), mw.Authenticate())
	app.Get("/expenses", func(c *fiber.Ctx) error { return c.SendString("Expense endpoint hit") })

	App.Web.Static("/", "./views/static")
}
