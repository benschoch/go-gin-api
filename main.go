package main

import (
	"github.com/gin-gonic/gin"
	"recipes-core-api/api/v1/routes"
	"recipes-core-api/configs"
)

func main() {
	router := gin.Default()
	// configure mongo
	configs.ConnectDB()
	// add routes
	routes.IngredientRoute(router)
	routes.RecipeRoute(router)

	router.Run("localhost:8000")
}
