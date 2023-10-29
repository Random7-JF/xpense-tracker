package main

import (
	"github.com/Random7-JF/xpense-tracker/config"
	"github.com/Random7-JF/xpense-tracker/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var App config.App

func main() {
	App.Db = db.ConnectSqliteDb("xpense.db")

	App.Web = fiber.New()
	createWeb(&App)
	App.Store = session.New()

	App.Web.Listen(":3000")
}
