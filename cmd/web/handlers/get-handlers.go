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
	selectedRange := c.Query("freq")
	search := c.Query("search")
	session, _ := h.App.Store.Get(c)
	auth := session.Get("Auth")

	log.Printf("The selected range is: %s", selectedRange)
	log.Printf("The search term is: %s", search)
	var expenses []model.Expense

	if selectedRange != "" {
		switch selectedRange {
		case "Weekly":
			expenses, _ = h.App.Db.GetExpenseByFreq("Weekly")
		case "Monthly":
			expenses, _ = h.App.Db.GetExpenseByFreq("Monthly")
		case "Yearly":
			expenses, _ = h.App.Db.GetExpenseByFreq("Yearly")
		case "oneTime":
			expenses, _ = h.App.Db.GetExpenseByFreq("Once")
		case "----":
			expenses, _ = h.App.Db.GetExpense(h.App.Db.GetUserId(auth.(server.Auth).Username))
		default:
			log.Printf("unsupported date range: %s", selectedRange)
		}
	} else if search != "" {
		expenses, _ = h.App.Db.GetExpenseBySearch("%" + search + "%")
	} else {
		expenses, _ = h.App.Db.GetExpense(h.App.Db.GetUserId(auth.(server.Auth).Username))
	}

	data["Expense"] = expenses

	return c.Render("expense-table-overview", data)
}
