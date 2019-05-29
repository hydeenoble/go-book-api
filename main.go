package main

import (
	"github.com/hydeenoble/mux-rest-api/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(8000)
}
