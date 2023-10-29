package handlers

import (
	"github.com/Random7-JF/xpense-tracker/server"
	"github.com/gofiber/fiber/v2"
)

func PostRegister(c *fiber.Ctx) error {
	var regForm server.RegisterForm
	regForm.Username = c.FormValue("username")
	regForm.Email = c.FormValue("email")
	regForm.Password = c.FormValue("password")
	regForm.PasswordConfirm = c.FormValue("passwordconfirm")

	if regForm.Password != regForm.PasswordConfirm {
		regForm.Error = "Passwords do not match."
		return c.Render("partials/form/register-response", fiber.Map{"User": regForm})
	}

	if h.App.Db.CheckIfUserExists(regForm.Username) {
		regForm.Error = "Username in use."
		return c.Render("partials/form/register-response", fiber.Map{"User": regForm})
	}

	h.App.Db.CreateUser(regForm.Username, regForm.Password, regForm.Email)

	return c.Render("partials/form/register-response", fiber.Map{"User": regForm})
}
