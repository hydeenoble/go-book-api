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

func init() {
	// books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "433323", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
}

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
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// for index, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:index], books[index+1:]...)
	// 		var book Book
	// 		_ = json.NewDecoder(r.Body).Decode(&book)
	// 		book.ID = strconv.Itoa(rand.Intn(10000000))
	// 		books = append(books, book)
	// 		json.NewEncoder(w).Encode(book)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(books)
}

// Delete a Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// fmt.Println(params["id"])
	
	// for index, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:index], books[index+1:]...)
	// 		break
	// 	}
	// }
	json.NewEncoder(w).Encode(model.DeleteBook(params["id"]))
}
