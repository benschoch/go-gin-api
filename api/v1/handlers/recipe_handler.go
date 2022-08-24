package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"recipes-core-api/api/v1/models"
	"recipes-core-api/configs"
	"time"
)

var recipeCollection = configs.GetCollection(configs.DB, "Recipe")

func CreateRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//var recipe models.Recipe
		var ingredient models.Ingredient
		defer cancel()

		ingredientObjectId, _ := primitive.ObjectIDFromHex("630522989df1f66715ab2508") // salt
		ingredientCollection.FindOne(ctx, bson.M{"_id": ingredientObjectId}).Decode(&ingredient)

		newRecipe := models.Recipe{
			Id:              primitive.NewObjectID(),
			Title:           "My first Recipe with GO!",
			PreparationTime: 20,
			Ingredient:      []models.Ingredient{ingredient},
		}

		recipeCollection.InsertOne(ctx, newRecipe)

		log.Println("Recipe created.")

	}

}
