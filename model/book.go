package model

import "fmt"

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

func GetBooks() {
	fmt.Println("I am in the model")
}

func CreateBook(book *Book){
	fmt.Println("I am in the model => CreateBook method")
	fmt.Println(book)
}
