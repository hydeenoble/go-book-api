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

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	fmt.Println("I am in the run initialize")
	// var books []controller.Book
	// books = append(books, controller.Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &controller.Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, controller.Book{ID: "2", Isbn: "433323", Title: "Book Two", Author: &controller.Author{Firstname: "Steve", Lastname: "Smith"}})
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// Route Handlers / Endpoints
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
