package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ingredient struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Singular string             `json:"singular,omitempty" bson:"singular,omitempty"`
	Plural   string             `json:"plural,omitempty" bson:"plural,omitempty"`
	Synonyms []string           `json:"synonyms,omitempty" bson:"synonyms,omitempty"`
}
