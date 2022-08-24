package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"recipes-core-api/api/v1/models"
	"recipes-core-api/configs"
	"time"
)

var ingredientCollection = configs.GetCollection(configs.DB, "Ingredient")
var validate = validator.New()

func CreateIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var ingredient models.Ingredient
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&ingredient); err != nil {
			c.JSON(http.StatusBadRequest, models.ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&ingredient); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newIngredient := models.Ingredient{
			Id:       primitive.NewObjectID(),
			Singular: ingredient.Singular,
			Plural:   ingredient.Plural,
			Synonyms: ingredient.Synonyms,
		}

		result, err := ingredientCollection.InsertOne(ctx, newIngredient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.ApiResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result}},
		)
	}
}

func GetAllIngredients() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var ingredients []models.Ingredient
		defer cancel()

		results, err := ingredientCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var ingredient models.Ingredient
			if err = results.Decode(&ingredient); err != nil {
				c.JSON(http.StatusInternalServerError, models.ApiResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			}

			ingredients = append(ingredients, ingredient)
		}

		c.JSON(http.StatusOK, models.ApiResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": ingredients}},
		)
	}
}

func GetIngredientById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ingredientId := c.Param("ingredientId")
		var ingredient models.Ingredient
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ingredientId)
		err := ingredientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&ingredient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		c.JSON(http.StatusOK, models.ApiResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": ingredient}},
		)
	}
}
