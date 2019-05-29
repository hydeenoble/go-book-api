package main

import (
	"github.com/hydeenoble/mux-rest-api/app"
	"github.com/hydeenoble/mux-rest-api/config"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(config.GetConfig().API.Port)
}
