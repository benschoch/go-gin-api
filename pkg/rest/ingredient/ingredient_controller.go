package ingredient

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.POST("/ingredients", CreateIngredient())
	//	router.PUT("/ingredients/:userId", controllers.EditAUser())
	//	router.DELETE("/ingredients/:userId", controllers.DeleteAUser())
	r.GET("/ingredients/:id", GetIngredientByID())
	r.GET("/ingredients", GetAllIngredients())
}
