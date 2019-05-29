package main

import (
	"github.com/hydeenoble/mux-rest-api/app"
)

func main() {
	// testing external function
	app := &app.App{}
	app.Initialize()
	// Init Router
	// r := mux.NewRouter()

	// Mock Data - @todo - implement DB

	
	// log.Fatal(http.ListenAndServe(":8000", r))
	app.Run(8000)

}
