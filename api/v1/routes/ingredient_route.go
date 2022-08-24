package routes

import (
	"github.com/gin-gonic/gin"
	"recipes-core-api/api/v1/handlers"
)

func IngredientRoute(router *gin.Engine) {
	router.POST("/ingredients", handlers.CreateIngredient())
	//	router.PUT("/ingredients/:userId", controllers.EditAUser())
	//	router.DELETE("/ingredients/:userId", controllers.DeleteAUser())
	router.GET("/ingredients/:ingredientId", handlers.GetIngredientById())
	router.GET("/ingredients", handlers.GetAllIngredients())
}
