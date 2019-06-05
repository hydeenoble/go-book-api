package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hydeenoble/mux-rest-api/model"
	"github.com/hydeenoble/mux-rest-api/service"
)

// GetBooks - gets all books from the DB
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetBooks())
}

// GetBook - gets a single Book from the DB corresponding to the ID speciffied
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	json.NewEncoder(w).Encode(service.GetBook(params["id"]))
}

// CreateBook - Creates a new book in the DB
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	json.NewEncoder(w).Encode(service.CreateBook(book))
}

// UpdateBook - updates an existing Book in the DB
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	json.NewEncoder(w).Encode(service.UpdateBook(params["id"], book))
}

// DeleteBook - Deletes an exisiting Book in the DB
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(service.DeleteBook(params["id"]))
}
