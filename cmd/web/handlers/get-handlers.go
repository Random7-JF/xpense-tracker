package handlers

import (
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
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

func Logout(c *fiber.Ctx) error {
	session, _ := h.App.Store.Get(c)
	noAuth := server.Auth{Valid: false}
	session.Set("Auth", noAuth)
	session.Save()
	return c.Redirect("/")
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

func ExpenseList(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	selectedRange := c.Query("dateRange")
	log.Printf("The selected range is: %s", selectedRange)

	var expenses []model.Expense
	switch selectedRange {
	case "pastWeek":
		expenses, _ = h.App.Db.GetExpenseByFreq("Weekly")
	case "pastMonth":
		expenses, _ = h.App.Db.GetExpenseByFreq("Monthly")
	case "oneTime":
		expenses, _ = h.App.Db.GetExpenseByFreq("Once")
	default:
		log.Printf("unsupported date range: %s", selectedRange)
	}

	data["Expense"] = expenses

	return c.Render("expense-table-overview", data)
}
