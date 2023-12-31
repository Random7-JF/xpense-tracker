package handlers

import (
	"fmt"
	"log"

	"github.com/Random7-JF/xpense-tracker/model"
	"github.com/Random7-JF/xpense-tracker/server"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense"
	data["Test"] = c.IP()
	log.Printf("IP Found: %s", data["Test"])
	return c.Render("pages/index", data, "layouts/main")
}

func Login(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense - Login"
	return c.Render("pages/login", data, "layouts/main")
}

func Reports(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense - Reports"
	return c.Render("pages/app/expense/reports", data, "layouts/main")
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

	expense, err := h.App.Db.GetExpense(h.App.Db.GetUserId(data["Auth"].(server.Auth).Username))
	if err != nil {
		log.Println("Error in getting expenses", err)
		return c.Render("pages/app/expense/dashboard", data, "layouts/main")
	}
	data["Expense"] = expense

	return c.Render("pages/app/expense/dashboard", data, "layouts/main")
}

func ExpenseFill(c *fiber.Ctx) error {
	userid := c.Query("userid")
	h.App.Db.ExpenseFill(userid)
	return c.Redirect("/app/expense/dashboard")
}

func ExpenseList(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)

	selectedRange := c.Query("freq")
	search := c.Query("search")
	tagSearch := c.Query("tag")
	var expenses []model.Expense
	userid := h.App.Db.GetUserId(data["Auth"].(server.Auth).Username)

	if selectedRange != "" {
		switch selectedRange {
		case "Weekly":
			expenses, _ = h.App.Db.GetExpenseByFreq("Weekly", userid)
		case "Monthly":
			expenses, _ = h.App.Db.GetExpenseByFreq("Monthly", userid)
		case "Yearly":
			expenses, _ = h.App.Db.GetExpenseByFreq("Yearly", userid)
		case "oneTime":
			expenses, _ = h.App.Db.GetExpenseByFreq("Once", userid)
		case "----":
			expenses, _ = h.App.Db.GetExpense(h.App.Db.GetUserId(data["Auth"].(server.Auth).Username))
		default:
			log.Printf("unsupported date range: %s", selectedRange)
		}
	} else if search != "" {
		expenses, _ = h.App.Db.GetExpenseBySearch("%"+search+"%", userid)
	} else if tagSearch != "" {
		expenses, _ = h.App.Db.GetExpenseByTag("%"+tagSearch+"%", userid)
	} else {
		expenses, _ = h.App.Db.GetExpense(h.App.Db.GetUserId(data["Auth"].(server.Auth).Username))
	}

	data["Expense"] = expenses

	return c.Render("expense-table-overview", data)
}

func Admin(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)
	data["Title"] = "Xpense - Admin"

	users := h.App.Db.GetUsers()
	type userCounts struct {
		User  model.User
		Count int
	}
	var combined []userCounts
	for _, user := range users {
		var u userCounts
		u.User.Username = user.Username
		u.User.Email = user.Email
		u.User.Id = user.Id
		count, err := h.App.Db.GetExpenseCountByUser(fmt.Sprintf("%d", user.Id))
		u.Count = count
		if err != nil {
			log.Printf("error in get expense count by user for user: %d, %s", user.Id, err)
		}
		combined = append(combined, u)
	}
	data["Users"] = combined

	return c.Render("pages/app/admin/control-panel", data, "layouts/main")
}
