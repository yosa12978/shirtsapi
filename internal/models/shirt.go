package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shirt struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Color   string             `json:"color" bson:"color"`
	Size    string             `json:"size" bson:"size"`
	Pattern string             `json:"pattern" bson:"pattern"`
	Photo   string             `json:"photo" bson:"photo"`
}
