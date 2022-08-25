package routes

import (
	"github.com/gin-gonic/gin"
	"recipes-core-api/api/v1/handlers"
)

func RecipeRoute(router *gin.Engine) {
	router.POST("/recipes", handlers.CreateRecipe())
	//	router.PUT("/recipes/:id", controllers.EditAUser())
	//	router.DELETE("/recipes/:id", controllers.DeleteAUser())
	// router.GET("/recipes/:id", handlers.GetIngredientById())
	router.GET("/recipes", handlers.GetAllRecipes())
}
