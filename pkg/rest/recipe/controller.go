package recipe

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.POST("/recipes", handler.Create())
	//	router.PUT("/recipes/:id", controllers.EditAUser())
	//	router.DELETE("/recipes/:id", controllers.DeleteAUser())
	router.GET("/recipes/:id", handler.GetByID())
	router.GET("/recipes/title/:title", handler.GetByTitle())
	router.GET("/recipes/slug/:slug", handler.GetByTitle())
	router.GET("/recipes", handler.GetAllPaginated())
}
