package routes

import (
	"github.com/gin-gonic/gin"
	"recipes-core-api/api/v1/handlers"
)

func UnitRoute(router *gin.Engine) {
	// router.POST("/units", handlers.CreateIngredient())
	//	router.PUT("/units/:id", controllers.EditAUser())
	//	router.DELETE("/units/:id", controllers.DeleteAUser())
	// router.GET("/units/:id", handlers.GetIngredientById())
	router.GET("/units", handlers.GetAllUnits())
}
