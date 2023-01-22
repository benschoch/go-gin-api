package demo

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"recipes-core-api/models"
	"recipes-core-api/pkg/db"
	"time"
)

var recipeCollection = db.GetCollection(db.DB, "recipes")
var ingredientCollection = db.GetCollection(db.DB, "ingredients")
var unitCollection = db.GetCollection(db.DB, "units")

func InitDemoData() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		var i models.Ingredient
		var u models.Unit
		defer cancel()

		client := db.ConnectDB()

		log.Println("Dropping database...")
		err := client.Database("api").Drop(context.TODO())
		if err != nil {
			log.Println(err)
		}
		log.Println("Loading demo data...")

		ingredientArray := []string{"salt", "pepper", "pasta", "meat", "oil"}
		for _, s := range ingredientArray {
			newIngredient := models.Ingredient{Id: uuid.NewString(), Language: "hu_HU", Singular: s, Plural: s}
			_, err := ingredientCollection.InsertOne(context.TODO(), newIngredient)
			if err != nil {
				log.Println(err)
			}
		}

		unitArray := []string{"g", "kg", "l", "ml", "handful"}
		for _, s := range unitArray {
			newUnit := models.Unit{Id: uuid.NewString(), Language: "hu_HU", Singular: s, Plural: s}
			_, err := unitCollection.InsertOne(context.TODO(), newUnit)
			if err != nil {
				log.Println(err)
			}
		}

		for x := 1; x < 5000; x++ {
			errI := ingredientCollection.FindOne(ctx, bson.M{"singular": "salt"}).Decode(&i)
			if errI != nil {
				return
			}

			errU := unitCollection.FindOne(ctx, bson.M{"singular": "g"}).Decode(&u)
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
				Order:            1,
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
}
