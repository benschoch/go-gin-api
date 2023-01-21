package recipe

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/recipes", CreateRecipe())
	//	router.PUT("/recipes/:id", controllers.EditAUser())
	//	router.DELETE("/recipes/:id", controllers.DeleteAUser())
	// router.GET("/recipes/:id", handlers.GetIngredientById())
	router.GET("/recipes", GetAllRecipes())
}
