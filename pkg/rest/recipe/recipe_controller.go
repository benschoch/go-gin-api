package recipe

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/recipes", CreateRecipe())
	//	router.PUT("/recipes/:id", controllers.EditAUser())
	//	router.DELETE("/recipes/:id", controllers.DeleteAUser())
	r.GET("/recipes/:id", GetRecipeById())
	r.GET("/recipes/title/:title", GetRecipeByTitle())
	r.GET("/recipes", GetAllRecipes())
}
