package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	fmt.Println("I am in the run initialize")
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// Route Handlers / Endpoints
	a.Router.HandleFunc("/api/books", getBooks).Methods("GET")
	a.Router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	a.Router.HandleFunc("/api/books", createBook).Methods("POST")
	a.Router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	a.Router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
}

// Run starts the app and serves on the specified address
func (a *App) Run(port int) {
	portString := fmt.Sprintf(":%d", port)
	loggingHandler := handlers.LoggingHandler(os.Stdout, a.Router)
	log.Fatal(http.ListenAndServe(portString, loggingHandler))
}
