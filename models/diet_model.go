package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Diet struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Language string             `json:"language,omitempty" bson:"language,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
}
