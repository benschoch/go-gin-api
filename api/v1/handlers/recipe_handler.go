package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"recipes-core-api/api/v1/models"
	"recipes-core-api/internal/mongo"
	"time"
)

var recipeCollection = mongo.GetCollection(mongo.DB, "Recipe")

func CreateRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var recipe models.Recipe
		//var recipeIngredient models.RecipeIngredient
		var ingredient models.Ingredient
		var unit models.Unit
		defer cancel()

		// validate the request and bind to struct
		if err := c.BindJSON(&recipe); err != nil {
			c.JSON(http.StatusBadRequest, models.ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		ingredientObjectId, _ := primitive.ObjectIDFromHex("63076fdfd134ccb4cc3c182d") // salt
		errI := ingredientCollection.FindOne(ctx, bson.M{"_id": ingredientObjectId}).Decode(&ingredient)
		if errI != nil {
			return
		}

		unitObjectId, _ := primitive.ObjectIDFromHex("63076fdfd134ccb4cc3c1832") // g
		errU := unitCollection.FindOne(ctx, bson.M{"_id": unitObjectId}).Decode(&unit)
		if errU != nil {
			return
		}

		newRecipeIngredient1 := models.RecipeIngredient{
			Id:         primitive.NewObjectID(),
			Ingredient: ingredient,
			Unit:       unit,
			Amount:     5,
		}

		newRecipeIngredient2 := models.RecipeIngredient{
			Id:         primitive.NewObjectID(),
			Ingredient: ingredient,
			Unit:       unit,
			Amount:     5,
		}

		newRecipe := models.Recipe{
			Id:               primitive.NewObjectID(),
			Title:            "My first Recipe with GO!",
			PreparationTime:  20,
			RecipeIngredient: []models.RecipeIngredient{newRecipeIngredient1, newRecipeIngredient2},
		}

		recipeCollection.InsertOne(ctx, newRecipe)

	}

}

func GetAllRecipes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var recipes []models.Recipe
		defer cancel()

		results, err := recipeCollection.Find(ctx, bson.M{})

		if err != nil {
			sendErrorJson(c, err)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var recipe models.Recipe
			if err = results.Decode(&recipe); err != nil {
				sendErrorJson(c, err)
			}
			recipes = append(recipes, recipe)
		}

		c.JSON(http.StatusOK, models.ApiResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": recipes}},
		)
	}
}

func sendErrorJson(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, models.ApiResponse{
		Status:  http.StatusInternalServerError,
		Message: "error",
		Data:    map[string]interface{}{"data": err.Error()},
	})
}
