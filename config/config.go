package config

import (
	"github.com/Random7-JF/xpense-tracker/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type App struct {
	Web   *fiber.App
	Store *session.Store
	Db    db.Service
}
