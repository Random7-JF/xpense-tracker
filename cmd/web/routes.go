package main

import (
	"github.com/Random7-JF/xpense-tracker/cmd/web/handlers"
	"github.com/Random7-JF/xpense-tracker/cmd/web/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000/*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(mw.SaveSession(), mw.Authenticate())
	app.Get("/expense/reports", func(c *fiber.Ctx) error { return c.SendString("Expense endpoint hit") })
	app.Get("/expense/modify", handlers.ExpenseModify)
	app.Post("/expense/modify", handlers.PostExpenseModify)
	app.Get("/expense/dashboard", func(c *fiber.Ctx) error { return c.SendString("Expense endpoint hit") })

	App.Web.Static("/", "./views/static")
}
