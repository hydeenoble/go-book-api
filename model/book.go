package model

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func init() {
}

func GetBooks() {
	books := Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := booksCollection.InsertOne(ctx, books)
	fmt.Println(res.InsertedID)
	fmt.Println("I am in the model")
}

func CreateBook(book Book) Book {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := booksCollection.InsertOne(ctx, book)
	if err != nil {
		log.Fatal(err)
	}
	return book
}
