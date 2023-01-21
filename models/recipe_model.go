package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title            string             `json:"title,omitempty" bson:"title,omitempty"`
	PreparationTime  int                `json:"preparation_time,omitempty" bson:"preparation_time,omitempty"`
	RecipeIngredient []RecipeIngredient `json:"recipe_ingredients,omitempty" bson:"recipe_ingredients,omitempty"`
}

type RecipeIngredient struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Ingredient Ingredient         `json:"ingredient,omitempty" bson:"ingredient,omitempty"`
	Unit       Unit               `json:"unit,omitempty" bson:"unit,omitempty"`
	Amount     int                `json:"amount,omitempty" bson:"amount,omitempty"`
}
