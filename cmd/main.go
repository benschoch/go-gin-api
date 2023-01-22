package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"recipes-core-api/pkg/db"
	"recipes-core-api/pkg/rest/demo"
	"recipes-core-api/pkg/rest/ingredient"
	"recipes-core-api/pkg/rest/recipe"
	"recipes-core-api/pkg/rest/unit"
)

func main() {
	r := gin.Default()

	// configure mongo
	db.ConnectDB()

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
