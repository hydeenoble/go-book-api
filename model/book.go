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
	res, err := booksCollection.InsertOne(ctx, book)
	if err != nil {
		log.Fatal(err)
	}
	book.ID, _ = primitive.ObjectIDFromHex(res.InsertedID.(primitive.ObjectID).Hex())
	return book
}

func GetBook(id string) schema.Book {
	var book schema.Book

	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", _id}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := booksCollection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}

	return book
}

func DeleteBook(id string) string {
	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", _id}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := booksCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedCount == 1 {
		return "Book successfully deleted!"
	} else if res.DeletedCount == 0 {
		return "Book with id: " + id + "does not exist!"
	} else {
		return "Something wenr wrong somewhere."
	}

	// fmt.Println(res)
	// fmt.Printf("%+v\n", res)
}

func UpdateBook(id string, book schema.Book) schema.Book {
	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", _id}}
	update := bson.M{
		"$set": bson.M{
			"title": book.Title,
			"isbn":  book.Isbn,
			"author": bson.M{
				"firstname": book.Author.Firstname,
				"lastname":  book.Author.Lastname,
			},
		},
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := booksCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Fatal(err)
	}
	if res.MatchedCount == 1 {
		book.ID = _id
		return book
	} else {
		return schema.Book{}
	}
}
