package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Fname     string             `json:"fname" bson:"fname"`
	Lname     string             `json:"lname" bson:"lname"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Birthdate time.Time          `json:"birthdate" bson:"birthdate"`
	CreateAt  time.Time          `json:"createat" bson:"createat"`
}
