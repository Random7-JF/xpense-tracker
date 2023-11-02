package handlers

import (
	"fmt"
	"log"
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

	h.App.Db.CreateUser(regForm.Username, regForm.Password, regForm.Email)

	return c.Render("partials/form/register-response", fiber.Map{"User": regForm})
}

func PostLogin(c *fiber.Ctx) error {
	var loginForm server.LoginForm
	loginForm.Username = c.FormValue("username")
	loginForm.Password = c.FormValue("password")

	auth := h.App.Db.AuthUser(loginForm.Username, loginForm.Password)
	if !auth {
		loginForm.Error = "Bad Password"
		return c.Render("partials/form/login-response", fiber.Map{"User": loginForm})
	}

	session, err := h.App.Store.Get(c)
	if err != nil {
		log.Println("Session error", err)
	}

	session.Set("Auth", server.Auth{Valid: true, Username: loginForm.Username})
	log.Println("Postlogin - ", auth, " - ", loginForm)
	authed := session.Get("Auth")

	err = session.Save()
	if err != nil {
		log.Println("Session error", err)
	}

	log.Println("Authed = ", authed.(server.Auth).Valid)
	return c.Render("partials/form/login-response", fiber.Map{"User": loginForm})
}

func PostExpenseModify(c *fiber.Ctx) error {
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
	ExpenseForm.Frequency = c.FormValue("freq")
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
		return c.Render("partials/form/app/expense/modify-response", data)
	}
	data["Form"] = ExpenseForm
	return c.Render("partials/form/app/expense/modify-response", data)
}
