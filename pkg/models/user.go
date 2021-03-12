package models

import 	"go.mongodb.org/mongo-driver/bson/primitive"

// User is model for our users
type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name         string    `bson:"name" json:"name,omitempty"`
	Email        string    `bson:"email" json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	Hash string    `bson:"hash" json:"passwordhash,omitempty"`
}
