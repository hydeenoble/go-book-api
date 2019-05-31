package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hydeenoble/mux-rest-api/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBooks() []*schema.Book {
	// books := Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}
	var books []*schema.Book
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := booksCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for res.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var book schema.Book
		err := res.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, &book)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	res.Close(context.Background())

	return books
}

func CreateBook(book schema.Book) schema.Book {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := booksCollection.InsertOne(ctx, book)
	if err != nil {
		log.Fatal(err)
	}
	return book
}

func GetBook(id string) schema.Book {
	var book schema.Book

	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", _id}}

	err := booksCollection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}

	return book
}


func DeleteBook() {
	
}