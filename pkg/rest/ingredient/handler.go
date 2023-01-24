package ingredient

import (
	"context"
	"net/http"
	"recipes-core-api/models"
	"recipes-core-api/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	dbConnection *db.Connection
	validator    *validator.Validate
}

func NewHandler(dbConnection *db.Connection, validate *validator.Validate) *Handler {
	return &Handler{dbConnection: dbConnection, validator: validate}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var ingredient models.Ingredient
		defer cancel()

		// bind request to struct
		if err := c.BindJSON(&ingredient); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		// use the validator library to validate required fields
		if validationErr := h.validator.Struct(&ingredient); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newIngredient := models.Ingredient{
			Id:       uuid.NewString(),
			Singular: ingredient.Singular,
			Plural:   ingredient.Plural,
			Synonyms: ingredient.Synonyms,
		}

		ingredientCollection := h.dbConnection.GetIngredients()
		result, err := ingredientCollection.InsertOne(ctx, newIngredient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, &models.APIResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result},
		})
	}
}

func (h *Handler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var ingredients []models.Ingredient
		defer cancel()

		ingredientCollection := h.dbConnection.GetIngredients()
		results, err := ingredientCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		// reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var ingredient models.Ingredient
			if err = results.Decode(&ingredient); err != nil {
				c.JSON(http.StatusInternalServerError, models.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			}

			ingredients = append(ingredients, ingredient)
		}

		c.JSON(http.StatusOK, models.APIResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": ingredients}},
		)
	}
}

func (h *Handler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ingredientID := c.Param("id")
		var ingredient models.Ingredient
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(ingredientID)
		ingredientCollection := h.dbConnection.GetIngredients()
		err := ingredientCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&ingredient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		c.JSON(http.StatusOK, models.APIResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": ingredient}},
		)
	}
}