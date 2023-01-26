package recipe

import (
	"context"
	"net/http"
	"recipes-core-api/models"
	"recipes-core-api/pkg/db"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	dbConnection *db.Connection
}

func NewHandler(dbConnection *db.Connection) *Handler {
	return &Handler{dbConnection: dbConnection}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// var recipe models.Recipe
		// var recipeIngredient models.RecipeIngredient
		var i models.Ingredient
		var u models.Unit
		defer cancel()

		// validate the request and bind to struct
		// if err := c.BindJSON(&recipe); err != nil {
		//	c.JSON(http.StatusBadRequest, models.APIResponse{
		//		Status:  http.StatusBadRequest,
		//		Message: "error",
		//		Data:    map[string]interface{}{"data": err.Error()}},
		//	)
		//	return
		// }

		ingredientCollection := h.dbConnection.GetCollection("ingredients")
		ingredientObjectID, _ := primitive.ObjectIDFromHex("63cbca0e2e2cf00250192ca2") // salt
		errI := ingredientCollection.FindOne(ctx, bson.M{"_id": ingredientObjectID}).Decode(&i)
		if errI != nil {
			return
		}

		unitObjectID, _ := primitive.ObjectIDFromHex("63cbca0e2e2cf00250192ca7") // g
		unitCollection := h.dbConnection.GetUnits()
		errU := unitCollection.FindOne(ctx, bson.M{"_id": unitObjectID}).Decode(&u)
		if errU != nil {
			return
		}

		newRecipeIngredient1 := models.RecipeIngredient{
			Ingredient: i,
			Unit:       u,
			AmountFrom: 5,
		}

		newRecipeIngredient2 := models.RecipeIngredient{
			Ingredient: i,
			Unit:       u,
		}

		newRecipeIngredientGroup := models.RecipeIngredientGroup{
			Name:             "default",
			Order:            12,
			RecipeIngredient: []models.RecipeIngredient{newRecipeIngredient1, newRecipeIngredient2},
		}

		newRecipe := models.Recipe{
			ID:                    uuid.NewString(),
			Language:              "hu_HU",
			IsPublished:           true,
			Title:                 "My first Recipe with GO!",
			Slug:                  "my-first-recipe-with-go",
			PreparationTime:       20,
			CookingTime:           60,
			Difficulty:            1,
			YoutubeVideoID:        "someFancyYoutubeVideoId",
			RecipeIngredientGroup: []models.RecipeIngredientGroup{newRecipeIngredientGroup},
		}
		recipeCollection := h.dbConnection.GetRecipes()
		_, err := recipeCollection.InsertOne(ctx, newRecipe)
		if err != nil {
			h.sendErrorJSON(c, err)
			return
		}
	}
}

func (h *Handler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var recipes []models.Recipe
		defer cancel()

		recipeCollection := h.dbConnection.GetRecipes()
		results, err := recipeCollection.Find(ctx, bson.M{})

		if err != nil {
			h.sendErrorJSON(c, err)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var recipe models.Recipe
			if err = results.Decode(&recipe); err != nil {
				h.sendErrorJSON(c, err)
			}
			recipes = append(recipes, recipe)
		}

		c.JSON(http.StatusOK, models.APIResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": recipes}},
		)
	}
}

func (h *Handler) GetAllPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var recipes []models.Recipe
		defer cancel()

		pageStr := c.DefaultQuery("page", "1")
		sizeStr := c.DefaultQuery("size", "12")

		page, _ := strconv.ParseInt(pageStr, 10, 64)
		size, _ := strconv.ParseInt(sizeStr, 10, 64)

		pageOptions := options.Find().SetSkip((page - 1) * size).SetLimit(size)

		recipeCollection := h.dbConnection.GetRecipes()
		documents, _ := recipeCollection.EstimatedDocumentCount(ctx)
		results, err := recipeCollection.Find(ctx, bson.M{}, pageOptions)
		if err != nil {
			return
		}

		defer results.Close(ctx)

		for results.Next(ctx) {
			var recipe models.Recipe
			if err = results.Decode(&recipe); err != nil {
				h.sendErrorJSON(c, err)
			}
			recipes = append(recipes, recipe)
		}

		c.JSON(http.StatusOK, models.APIResponse2{
			Status:         http.StatusOK,
			Message:        "success",
			Total:          documents,
			RecipesPerPage: size,
			PageNumber:     page,
			Data:           map[string]interface{}{"recipes": recipes}},
		)
	}
}

func (h *Handler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		recipeID := c.Param("id")
		var recipe models.Recipe
		defer cancel()

		// objId, _ := primitive.ObjectIDFromHex(recipeID)
		recipeCollection := h.dbConnection.GetRecipes()
		err := recipeCollection.FindOne(ctx, bson.M{"_id": recipeID}).Decode(&recipe)
		if err != nil {
			h.sendErrorJSON(c, err)
			return
		}

		c.JSON(http.StatusOK, models.APIResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": recipe}},
		)
	}
}

func (h *Handler) GetByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		title := c.Param("title")
		var recipe models.Recipe
		defer cancel()

		recipeCollection := h.dbConnection.GetRecipes()
		err := recipeCollection.FindOne(ctx, bson.M{"title": title}).Decode(&recipe)
		if err != nil {
			h.sendErrorJSON(c, err)
			return
		}

		c.JSON(http.StatusOK, models.APIResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": recipe}},
		)
	}
}

func (h *Handler) sendErrorJSON(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, models.APIResponse{
		Status:  http.StatusInternalServerError,
		Message: "error",
		Data:    map[string]interface{}{"data": err.Error()},
	})
}
