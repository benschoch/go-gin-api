package recipe

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"recipes-core-api/models"
	"recipes-core-api/pkg/db"
	"strconv"
	"time"
)

var recipeCollection = db.GetCollection(db.DB, "recipes")
var ingredientCollection = db.GetCollection(db.DB, "ingredients")
var unitCollection = db.GetCollection(db.DB, "units")

func CreateRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//var recipe models.Recipe
		//var recipeIngredient models.RecipeIngredient
		var i models.Ingredient
		var u models.Unit
		defer cancel()

		// validate the request and bind to struct
		//if err := c.BindJSON(&recipe); err != nil {
		//	c.JSON(http.StatusBadRequest, models.ApiResponse{
		//		Status:  http.StatusBadRequest,
		//		Message: "error",
		//		Data:    map[string]interface{}{"data": err.Error()}},
		//	)
		//	return
		//}

		ingredientObjectId, _ := primitive.ObjectIDFromHex("63cbca0e2e2cf00250192ca2") // salt
		errI := ingredientCollection.FindOne(ctx, bson.M{"_id": ingredientObjectId}).Decode(&i)
		if errI != nil {
			return
		}

		unitObjectId, _ := primitive.ObjectIDFromHex("63cbca0e2e2cf00250192ca7") // g
		errU := unitCollection.FindOne(ctx, bson.M{"_id": unitObjectId}).Decode(&u)
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
			Id:                    uuid.NewString(),
			Language:              "hu_HU",
			IsPublished:           true,
			Title:                 "My first Recipe with GO!",
			Slug:                  "my-first-recipe-with-go",
			PreparationTime:       20,
			CookingTime:           60,
			Difficulty:            1,
			YoutubeVideoId:        "someFancyYoutubeVideoId",
			RecipeIngredientGroup: []models.RecipeIngredientGroup{newRecipeIngredientGroup},
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

func GetRecipesPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var recipes []models.Recipe
		defer cancel()

		pageStr := c.DefaultQuery("page", "0")
		sizeStr := c.DefaultQuery("size", "12")

		page, _ := strconv.Atoi(pageStr)
		size, _ := strconv.Atoi(sizeStr)

		pageOptions := options.Find().SetSkip(int64(page)).SetLimit(int64(size))

		results, err := recipeCollection.Find(ctx, bson.M{}, pageOptions)
		if err != nil {
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

func GetRecipeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		recipeId := c.Param("id")
		var recipe models.Recipe
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(recipeId)
		err := recipeCollection.FindOne(ctx, bson.M{"_id": recipeId}).Decode(&recipe)
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
			Data:    map[string]interface{}{"data": recipe}},
		)
	}
}

func GetRecipeByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		title := c.Param("title")
		var recipe models.Recipe
		defer cancel()

		err := recipeCollection.FindOne(ctx, bson.M{"title": title}).Decode(&recipe)
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
			Data:    map[string]interface{}{"data": recipe}},
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
