package demo

import (
	"context"
	"fmt"
	"log"
	"recipes-core-api/models"
	"recipes-core-api/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	dbConnection *db.Connection
}

func NewHandler(dbConnection *db.Connection) *Handler {
	return &Handler{dbConnection: dbConnection}
}

func (h *Handler) Demo() gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerTimeout := 60 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), handlerTimeout)

		var ing models.Ingredient
		var uni models.Unit
		var ser models.ServingType
		var reg models.Region
		var cou1, cou2 models.Course
		var diet1, diet2 models.Diet

		defer cancel()

		log.Println("Dropping database...")
		err := h.dbConnection.DropDB()
		if err != nil {
			log.Println(err)
		}
		log.Println("Creating demo data...")

		courseCollection := h.dbConnection.GetCourses()
		courseArray := []string{"Breakfast", "Brunch", "Dinner", "Lunch"}
		for _, s := range courseArray {
			newC := models.Course{Id: uuid.NewString(), Language: "en_LD", Name: s}
			_, err := courseCollection.InsertOne(context.TODO(), newC)
			if err != nil {
				log.Println(err)
			}
		}

		dietCollection := h.dbConnection.GetDiets()
		dietArray := []string{"Gluten-free", "Lactose-free", "Low Carb", "Vegan"}
		for _, s := range dietArray {
			newD := models.Diet{Id: uuid.NewString(), Language: "en_LD", Name: s}
			_, err := dietCollection.InsertOne(context.TODO(), newD)
			if err != nil {
				log.Println(err)
			}
		}

		ingredientCollection := h.dbConnection.GetIngredients()
		ingredientArray := []string{"salt", "pepper", "pasta", "meat", "oil"}
		for _, s := range ingredientArray {
			newIngredient := models.Ingredient{Id: uuid.NewString(), Language: "en_LD", Singular: s, Plural: s}
			_, err := ingredientCollection.InsertOne(context.TODO(), newIngredient)
			if err != nil {
				log.Println(err)
			}
		}

		unitCollection := h.dbConnection.GetUnits()
		unitArray := []string{"g", "kg", "l", "ml", "handful"}
		for _, s := range unitArray {
			newUnit := models.Unit{Id: uuid.NewString(), Language: "en_LD", Singular: s, Plural: s}
			_, err := unitCollection.InsertOne(context.TODO(), newUnit)
			if err != nil {
				log.Println(err)
			}
		}

		regionCollection := h.dbConnection.GetRegions()
		regionArray := []string{"Asia", "Spain", "Hungarian"}
		for _, s := range regionArray {
			newRegion := models.Region{ID: uuid.NewString(), Language: "en_LD", Name: s}
			_, err := regionCollection.InsertOne(context.TODO(), newRegion)
			if err != nil {
				log.Println(err)
			}
		}

		foodTypeCollection := h.dbConnection.GetFoodTypes()
		foodTypeArray := []string{"Pasta", "Fruit", "Meat", "Fish"}
		for _, s := range foodTypeArray {
			newFoodType := models.FoodType{Id: uuid.NewString(), Language: "en_LD", Name: s}
			_, err := foodTypeCollection.InsertOne(context.TODO(), newFoodType)
			if err != nil {
				log.Println(err)
			}
		}

		servingTypeCollection := h.dbConnection.GetServingTypes()
		servingTypeArray := []string{"Cup", "Piece", "Ball", "Bar"}
		for _, s := range servingTypeArray {
			newST := models.ServingType{Id: uuid.NewString(), Language: "en_LD", Singular: s, Plural: s}
			_, err := servingTypeCollection.InsertOne(context.TODO(), newST)
			if err != nil {
				log.Println(err)
			}
		}

		recipeCollection := h.dbConnection.GetRecipes()
		for x := 1; x < 5000; x++ {
			errI := ingredientCollection.FindOne(ctx, bson.M{"singular": "salt"}).Decode(&ing)
			if errI != nil {
				return
			}

			errU := unitCollection.FindOne(ctx, bson.M{"singular": "g"}).Decode(&uni)
			if errU != nil {
				return
			}

			// ingredient groups + ingredient
			newRecipeIngredient1 := models.RecipeIngredient{
				Ingredient: ing,
				Unit:       uni,
				AmountFrom: 5,
			}
			newRecipeIngredient2 := models.RecipeIngredient{
				Ingredient: ing,
				Unit:       uni,
			}
			newRecipeIngredientGroup := models.RecipeIngredientGroup{
				Name:             "default",
				Order:            1,
				RecipeIngredient: []models.RecipeIngredient{newRecipeIngredient1, newRecipeIngredient2},
			}

			// image + variations
			newImageVar1 := models.ImageVariation{
				Url:    "https://path/to/our/storage/nice.jpg",
				Filter: "420x420",
			}
			newImageVar2 := models.ImageVariation{
				Url:    "https://path/to/our/storage/nice.jpg",
				Filter: "420x420",
			}
			newImage := models.Image{
				URL:            "https://path/to/our/storage/nice.jpg",
				ImageVariation: []models.ImageVariation{newImageVar1, newImageVar2},
			}

			// serving type
			errS := servingTypeCollection.FindOne(ctx, bson.M{"singular": "Cup"}).Decode(&ser)
			if errS != nil {
				return
			}

			// regions
			errR := regionCollection.FindOne(ctx, bson.M{"name": "Asia"}).Decode(&reg)
			if errR != nil {
				return
			}

			// course
			errC1 := courseCollection.FindOne(ctx, bson.M{"name": "Breakfast"}).Decode(&cou1)
			if errC1 != nil {
				return
			}
			errC2 := courseCollection.FindOne(ctx, bson.M{"name": "Lunch"}).Decode(&cou2)
			if errC2 != nil {
				return
			}

			// diet
			errD1 := dietCollection.FindOne(ctx, bson.M{"name": "Gluten-free"}).Decode(&diet1)
			if errD1 != nil {
				return
			}
			errD2 := dietCollection.FindOne(ctx, bson.M{"name": "Vegan"}).Decode(&diet2)
			if errD2 != nil {
				return
			}

			newRecipe := models.Recipe{
				ID:                    uuid.NewString(),
				Language:              "en_LD",
				IsPublished:           true,
				Title:                 fmt.Sprintf("%s%d", "Recipe-", x),
				Slug:                  "my-first-recipe-with-go",
				PreparationTime:       20,
				CookingTime:           60,
				Difficulty:            1,
				YoutubeVideoID:        "someFancyYoutubeVideoId",
				Course:                []models.Course{cou1, cou2},
				Diet:                  []models.Diet{diet1, diet2},
				RecipeIngredientGroup: []models.RecipeIngredientGroup{newRecipeIngredientGroup},
				Image:                 newImage,
				ServingType:           ser,
				Region:                []models.Region{reg},
			}

			recipeCollection.InsertOne(ctx, newRecipe)
		}
	}
}
