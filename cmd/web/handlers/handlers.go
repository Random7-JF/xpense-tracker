package handlers

import (
	"text/template"

	"github.com/Random7-JF/xpense-tracker/config"
	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	App *config.App
}

var h handlers

func NewHandlers(webapp *config.App) {
	h = handlers{
		App: webapp,
	}
}

func renderBlock(c *fiber.Ctx, templateFile string, block string, data interface{}) error {
	c.Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("views/" + templateFile + ".html"))
	return tmpl.ExecuteTemplate(c, block, data)
}
