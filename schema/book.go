package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// Book Struct (Model)
type Book struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn"`
	Title  string             `json:"title"`
	Author *Author            `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}