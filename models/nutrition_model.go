package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nutrition struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Kcal     int                `json:"kcal,omitempty" bson:"kcal,omitempty"`
	Protein  int                `json:"protein,omitempty" bson:"protein,omitempty"`
	Carb     int                `json:"carb,omitempty" bson:"carb,omitempty"`
	Fat      int                `json:"fat,omitempty" bson:"fat,omitempty"`
	Kj       int                `json:"kj,omitempty" bson:"kj,omitempty"`
	Roughage int                `json:"roughage,omitempty" bson:"roughage,omitempty"`
	Iron     int                `json:"iron,omitempty" bson:"iron,omitempty"`
}
