package controller

import (
	"encoding/json"
	// "math/rand"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
	"github.com/gorilla/mux"
	"github.com/hydeenoble/mux-rest-api/model"
	"github.com/hydeenoble/mux-rest-api/schema"
)

// Init books var as a slice Book Struct

var books []schema.Book

// Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GetBooks())
}

// Get single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	json.NewEncoder(w).Encode(model.GetBook(params["id"]))
}

//Create a New Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book schema.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	json.NewEncoder(w).Encode(model.CreateBook(book))
}

// Update a Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book schema.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	json.NewEncoder(w).Encode(model.UpdateBook(params["id"], book))
}

// Delete a Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(model.DeleteBook(params["id"]))
}
