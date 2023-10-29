package middleware

import (
	"errors"
	"log"

	"github.com/Random7-JF/xpense-tracker/config"
	"github.com/Random7-JF/xpense-tracker/helper"
	"github.com/Random7-JF/xpense-tracker/server"
	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	App *config.App
}

func NewMiddleware(app *config.App) middleware {
	return middleware{
		App: app,
	}
}

func (mw *middleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth, err := server.GetKey("Auth", c, mw.App)
		if err != nil {
			return c.Redirect("/login")
		}
		if auth == nil {
			log.Println("Middleware Auth Error:", err)
			return c.Redirect("/login")
		}

		if auth.(server.Auth).Valid {
			return c.Next()
		}
		return c.Redirect("/login")
	}
}

func (mw *middleware) SaveSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := mw.App.Store.Get(c)
		if err != nil {
			return errors.New("unable to get session store")
		}
		if err := session.Save(); err != nil {
			return err
		}
		return c.Next()
	}
}

func (mw *middleware) SetupSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := mw.App.Store.Get(c)
		if err != nil {
			return errors.New("unable to get session store")
		}

		auth := session.Get("Auth")
		if auth == nil {
			helper.UpdateSessionKey(mw.App, c, "Auth", server.Auth{Valid: false, Message: ""})
		}

		return c.Next()
	}
}
