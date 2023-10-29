package server

import "github.com/Random7-JF/xpense-tracker/config"

type server struct {
	App *config.App
}

var s server

func NewServer(NewApp *config.App) {
	s = server{
		App: NewApp,
	}
}
