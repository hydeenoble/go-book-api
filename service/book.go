package service

import (
	"log"

	"github.com/hydeenoble/mux-rest-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetBooks - gets all books from the DB
func GetBooks() []*model.Book {
	var books []*model.Book
	res, err := model.Find()
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for res.Next(model.BackgroundContext) {
		// create a value into which the single document can be decoded
		var book model.Book
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
	res.Close(model.BackgroundContext)

	return books
}

// CreateBook - Creates a new book in the DB
func CreateBook(book model.Book) model.Book {
	res, err := model.InsertOne(book)
	if err != nil {
		log.Fatal(err)
	}
	book.ID, _ = primitive.ObjectIDFromHex(res.InsertedID.(primitive.ObjectID).Hex())
	return book
}

// GetBook - gets a single Book from the DB corresponding to the ID speciffied
func GetBook(id string) model.Book {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", _id}}
	book, err := model.FindOne(filter)
	if err != nil {
		log.Fatal(err)
	}
	return book
}

// DeleteBook - Deletes an exisiting Book in the DB
func DeleteBook(id string) string {
	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", _id}}

	res, err := model.DeleteOne(filter)
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
}

// // UpdateBook - updates an existing Book in the DB
// func UpdateBook(id string, book schema.Book) schema.Book {
// 	_id, _ := primitive.ObjectIDFromHex(id)

// 	filter := bson.D{{"_id", _id}}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"title": book.Title,
// 			"isbn":  book.Isbn,
// 			"author": bson.M{
// 				"firstname": book.Author.Firstname,
// 				"lastname":  book.Author.Lastname,
// 			},
// 		},
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	res, err := booksCollection.UpdateOne(ctx, filter, update)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if res.MatchedCount == 1 {
// 		book.ID = _id
// 		return book
// 	}
// 	return schema.Book{}

// }
