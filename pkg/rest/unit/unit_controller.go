package unit

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// router.POST("/units", handlers.CreateIngredient())
	//	router.PUT("/units/:id", controllers.EditAUser())
	//	router.DELETE("/units/:id", controllers.DeleteAUser())
	// router.GET("/units/:id", handlers.GetIngredientById())
	router.GET("/units", GetAllUnits())
}
