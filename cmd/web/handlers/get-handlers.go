package handlers

import (
	"log"

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
	return c.Render("pages/app/expense/add", data, "layouts/main")
}

func ExpenseDashboard(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	session, _ := h.App.Store.Get(c)
	auth := session.Get("Auth")

	expense, err := h.App.Db.GetExpense(h.App.Db.GetUserId(auth.(server.Auth).Username))
	if err != nil {
		log.Println("Error in getting expenses", err)
		return c.Render("pages/app/expense/overview", data, "layouts/main")
	}
	data["Expense"] = expense

	return c.Render("pages/app/expense/overview", data, "layouts/main")
}
