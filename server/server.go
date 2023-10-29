package server

import (
	"encoding/gob"
	"errors"
	"log"

	"github.com/Random7-JF/xpense-tracker/config"
	"github.com/gofiber/fiber/v2"
)

type server struct {
	App *config.App
}

var s server

func NewServer(NewApp *config.App) {
	s = server{
		App: NewApp,
	}
}

func RegisterGobs() {
	gob.Register(Auth{})
}

func GetKey(key string, c *fiber.Ctx, app *config.App) (interface{}, error) {
	session, err := app.Store.Get(c)
	if err != nil {
		return nil, err
	}
	reqkey := session.Get(key)
	log.Println("reg key: ", reqkey)
	return reqkey, nil
}

func UpdateSessionKey(key string, value interface{}, app *config.App, c *fiber.Ctx) error {
	session, err := app.Store.Get(c)
	if err != nil {
		return errors.New("unable to get session store")
	}
	session.Set(key, value)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
