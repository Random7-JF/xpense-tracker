package main

import (
	"github.com/Random7-JF/xpense-tracker/cmd/web/handlers"
	"github.com/Random7-JF/xpense-tracker/cmd/web/middleware"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func addRoutes() {
	mw := middleware.NewMiddleware(&App)
	App.Web.Use(mw.SetupSession(), logger.New())
	App.Web.Get("/", handlers.Index)
	App.Web.Get("/login", handlers.Login)
	App.Web.Post("/login", handlers.PostLogin)
	App.Web.Get("/logout", handlers.Logout)
	App.Web.Get("/register", handlers.Register)
	App.Web.Post("/register", handlers.PostRegister)

	app := App.Web.Group("/app")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000/*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(mw.SaveSession(), mw.Authenticate())
	app.Get("/admin/control-panel", handlers.Admin)
	app.Get("/expense/reports", handlers.Reports)
	app.Get("/expense/add", handlers.ExpenseModify)
	app.Get("/expense/list", handlers.ExpenseList)
	app.Get("/expense/fill", handlers.ExpenseFill)
	app.Post("/expense/fill", handlers.PostExpenseFill)
	app.Post("/expense/drop", handlers.PostExpenseDrop)
	app.Post("/expense/add", handlers.PostExpenseAdd)
	app.Post("/expense/remove", handlers.PostExpenseRemove)
	app.Post("/expense/modify", handlers.PostExpenseModify)
	app.Post("/expense/update", handlers.PostExpenseUpdate)
	app.Get("/expense/dashboard", handlers.ExpenseDashboard)

	App.Web.Static("/", "./views/static")
}
