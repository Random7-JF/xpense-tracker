package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Random7-JF/xpense-tracker/model"
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

	err := h.App.Db.CreateUser(regForm.Username, regForm.Password, regForm.Email)
	if err != nil {
		log.Printf("Register error: %s", err)
	}

	return c.Render("partials/form/register-response", fiber.Map{"User": regForm})
}

func PostLogin(c *fiber.Ctx) error {
	var loginForm server.LoginForm
	var authInfo server.Auth
	loginForm.Username = c.FormValue("username")
	loginForm.Password = c.FormValue("password")

	auth := h.App.Db.AuthUser(loginForm.Username, loginForm.Password)
	authInfo.Username = loginForm.Username

	if !auth {
		authInfo.Message = "Bad Password"
		authInfo.Valid = auth
		return c.Redirect("/app/expense/dashboard", http.StatusSeeOther)
	} else {
		authInfo.Message = "Logged In!"
		authInfo.Valid = auth
		authInfo.UserId = h.App.Db.GetUserId(authInfo.Username)
		session, err := h.App.Store.Get(c)

		if err != nil {
			log.Println("Session error", err)
			return c.Redirect("/app/expense/dashboard", http.StatusSeeOther)
		}

		session.Set("Auth", authInfo)
		err = session.Save()
		if err != nil {
			log.Println("Session error", err)
			return c.Redirect("/app/expense/dashboard", http.StatusSeeOther)
		}
	}
	return c.Redirect("/app/expense/dashboard", http.StatusSeeOther)
}

func PostExpenseRemove(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	session, _ := h.App.Store.Get(c)
	auth := session.Get("Auth")
	id := c.FormValue("remove-expense-id")
	idint, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Printf("strconv error %s", err)
	}
	h.App.Db.RemoveExpense(int(idint))

	updatedTable, _ := h.App.Db.GetExpense(h.App.Db.GetUserId(auth.(server.Auth).Username))
	data["Expense"] = updatedTable

	return c.Render("expense-table-overview", data)
}

func PostExpenseAdd(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["Auth"] = server.GetAuthStatus(c, h.App)

	var ExpenseForm model.Expense
	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		log.Println("Error with string convert", err)
		ExpenseForm.Error = fmt.Sprintf("Error with string convert: %s", err)
		data["Form"] = ExpenseForm
		return c.Render("partials/form/app/expense/modify-response", data)
	}
	ExpenseForm.Label = c.FormValue("label")
	ExpenseForm.Amount = amount
	ExpenseForm.Frequency = c.FormValue("frequency")
	ExpenseForm.Tag = c.FormValue("tags")
	ExpenseForm.SubmissionDate = time.Now().String()
	ExpenseForm.ExpenseDate = time.Now().String()
	session, _ := h.App.Store.Get(c)
	auth := session.Get("Auth")
	ExpenseForm.UserId = h.App.Db.GetUserId(auth.(server.Auth).Username)

	err = h.App.Db.AddExpense(ExpenseForm)
	if err != nil {
		log.Println("Error with db call", err)
		ExpenseForm.Error = fmt.Sprintf("Error with db call: %s", err)
		data["Form"] = ExpenseForm
		return c.Render("partials/form/app/expense/add-response", data)
	}
	data["Form"] = ExpenseForm
	return c.Render("partials/form/app/expense/add-response", data)
}

func PostExpenseModify(c *fiber.Ctx) error {
	data := make(map[string]interface{})

	id := c.FormValue("modify-expense-id")
	expenseModify, _ := h.App.Db.GetExpenseByID(id)
	data["Expense"] = expenseModify

	return c.Render("expense-add-form", data)
}

func PostExpenseUpdate(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Printf("AtoI Error in expense update: %s", err)
	}
	var expense model.Expense
	expense.Id = id
	expense.Label = c.FormValue("label")
	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		log.Printf("Parsefloat error in post expense update: %s", err)
		return c.Redirect("/app/expense/dashboard")

	}
	expense.Amount = amount
	expense.Frequency = c.FormValue("frequency")
	expense.Tag = c.FormValue("tags")

	err = h.App.Db.UpdateExpenseById(expense)
	if err != nil {
		log.Printf("SQL call error: %s", err)
		return c.Redirect("/app/expense/dashboard")
	}

	data["Auth"] = server.GetAuthStatus(c, h.App)
	session, _ := h.App.Store.Get(c)
	auth := session.Get("Auth")
	expenses, err := h.App.Db.GetExpense(h.App.Db.GetUserId(auth.(server.Auth).Username))
	if err != nil {
		log.Println("Error in getting expenses", err)
		return c.Render("expense-table-overview", data)
	}
	data["Expense"] = expenses

	return c.Render("expense-table-overview", data)
}

func PostExpenseFill(c *fiber.Ctx) error {
	userid := c.Query("userid")
	log.Printf("userid: %s", userid)
	err := h.App.Db.ExpenseFill(userid)
	if err != nil {
		log.Printf("expensefill error: %s", err)
	}
	return c.Render("partials/form/fill-response", fiber.Map{})
}

func PostExpenseDrop(c *fiber.Ctx) error {
	userid := c.Query("userid")
	err := h.App.Db.ExpenseDrop(userid)
	if err != nil {
		log.Printf("expensedrop error: %s", err)
	}

	return c.Render("partials/form/drop-response", fiber.Map{})
}
