package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username"  bson:"username"`
	Password string             `json:"-" bson:"password"`
	Salt     string             `json:"-" bson:"salt"`
	Token    string             `json:"-" bson:"token"`
	Role     string             `json:"role" bson:"role"`
}
