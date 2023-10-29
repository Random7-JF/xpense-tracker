package handlers

import (
	"github.com/Random7-JF/xpense-tracker/server"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense"
	return c.Render("pages/index", data, "layouts/main")
}

func Login(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense - Login"
	return c.Render("pages/login", data, "layouts/main")
}

func Register(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense - Register"
	return c.Render("pages/register", data, "layouts/main")
}

func ExpenseModify(c *fiber.Ctx) error {
	data := make(map[string]interface{})

	data["Auth"] = server.GetAuthStatus(c, h.App)
	return c.Render("pages/app/expense/modify", data, "layouts/main")
}
