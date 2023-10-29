package handlers

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/server"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	auth, err := server.GetKey("Auth", c, h.App)
	if err != nil {
		log.Println("login page error", err)
	}
	data["Auth"] = auth
	data["Title"] = "Xpense"
	return c.Render("pages/index", data, "layouts/main")
}

func Login(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	auth, err := server.GetKey("Auth", c, h.App)
	if err != nil {
		log.Println("login page error", err)
	}
	data["Auth"] = auth
	data["Title"] = "Xpense - Login"
	return c.Render("pages/login", data, "layouts/main")
}

func Register(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	auth, err := server.GetKey("Auth", c, h.App)
	if err != nil {
		log.Println("login page error", err)
	}
	data["Auth"] = auth
	data["Title"] = "Xpense - Register"
	return c.Render("pages/register", data, "layouts/main")
}
