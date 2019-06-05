package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hydeenoble/mux-rest-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	host     = config.GetConfig().Mongo.Host
	port     = config.GetConfig().Mongo.Port
	user     = config.GetConfig().Mongo.Username
	password = config.GetConfig().Mongo.Password
	dbname   = config.GetConfig().Mongo.DbName
)

var (
	booksCollection *mongo.Collection
)

var BackgroundContext = context.Background()

// Initiates a DB connection
func init() {
	ctx, cancel := context.WithTimeout(BackgroundContext, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, password, host, port, dbname)))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel = context.WithTimeout(BackgroundContext, 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	booksCollection = client.Database(dbname).Collection("books")

	fmt.Println("Successfully connected!")
}
