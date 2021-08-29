package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Books struct {
	Author string             `bson:"author" json:"author,omitempty"`
	Date   string             `bson:"date" json:"date" binding:"required"`
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	Title  string             `bson:"title" json:"title" binding:"required"`
}
