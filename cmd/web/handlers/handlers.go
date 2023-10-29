package handlers

import (
	"github.com/Random7-JF/xpense-tracker/config"
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
