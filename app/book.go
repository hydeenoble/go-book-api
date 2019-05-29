package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/hydeenoble/mux-rest-api/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// App help set the type for the application
type App struct {
	Router *mux.Router
}

// Initialize sets up the routes for the app
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	a.Router.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")
	a.Router.HandleFunc("/api/books", controller.CreateBook).Methods("POST")
	a.Router.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT")
	a.Router.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE")
}

// Run starts the app and serves on the specified address
func (a *App) Run(port int) {
	portString := fmt.Sprintf(":%d", port)
	loggingHandler := handlers.LoggingHandler(os.Stdout, a.Router)
	log.Fatal(http.ListenAndServe(portString, loggingHandler))
}
