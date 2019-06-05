package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(book Book) (*mongo.InsertOneResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return booksCollection.InsertOne(ctx, book)
}

func Find() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(BackgroundContext, 5*time.Second)
	defer cancel()
	return booksCollection.Find(ctx, bson.D{{}})
}

func FindOne(filter bson.D) (Book, error) {
	var book Book
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := booksCollection.FindOne(ctx, filter).Decode(&book)
	return book, err
}

func DeleteOne() {}

func UpdateOne() {}
