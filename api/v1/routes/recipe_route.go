package routes

import (
	"github.com/gin-gonic/gin"
	"recipes-core-api/api/v1/handlers"
)

func RecipeRoute(router *gin.Engine) {
	// router.POST("/recipes", handlers.CreateIngredient())
	//	router.PUT("/ingredients/:userId", controllers.EditAUser())
	//	router.DELETE("/ingredients/:userId", controllers.DeleteAUser())
	// router.GET("/recipes/:id", handlers.GetIngredientById())
	router.GET("/recipes", handlers.CreateRecipe())
}
