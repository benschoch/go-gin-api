package main

import (
	"log"
	"recipes-core-api/pkg/db"
	"recipes-core-api/pkg/middleware"
	"recipes-core-api/pkg/rest/demo"
	"recipes-core-api/pkg/rest/ingredient"
	"recipes-core-api/pkg/rest/recipe"
	"recipes-core-api/pkg/rest/unit"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// configure mongo
	db.ConnectDB()

	// middleware
	r.Use(middleware.LoggerMiddleware())

	// register routes
	ingredient.RegisterRoutes(r)
	recipe.RegisterRoutes(r)
	unit.RegisterRoutes(r)

	demo.RegisterRoutes(r)

	err := r.Run(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
