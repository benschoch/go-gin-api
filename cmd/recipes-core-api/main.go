package main

import (
	"github.com/gin-gonic/gin"
	"recipes-core-api/api/v1/routes"
	"recipes-core-api/internal/mongo"
)

func main() {
	router := gin.Default()
	// configure mongo
	mongo.ConnectDB()
	// add routes
	routes.IngredientRoute(router)
	routes.RecipeRoute(router)
	routes.UnitRoute(router)

	router.Run(":9000")
}
