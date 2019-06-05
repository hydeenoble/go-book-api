package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne() {}

func Find() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(BackgroundContext, 5*time.Second)
	defer cancel()
	return booksCollection.Find(ctx, bson.D{{}})
}

func FindOne(filter bson.D, book Book) () {

}

func DeleteOne() {}

func UpdateOne() {}
