package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title           string             `json:"title,omitempty" bson:"title,omitempty"`
	PreparationTime int                `json:"preparation_time,omitempty" bson:"preparation_time,omitempty"`
	Ingredient      []Ingredient       `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
}
