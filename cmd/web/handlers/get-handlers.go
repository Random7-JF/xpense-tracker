package handlers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{"Title": "Xpense"}, "layouts/main")
}

func Login(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{"Title": "Xpense - Login"}, "layouts/main")
}

func Register(c *fiber.Ctx) error {
	return c.Render("pages/register", fiber.Map{"Title": "Xpense - Register"}, "layouts/main")
}
