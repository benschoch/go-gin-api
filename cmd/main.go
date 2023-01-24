package main

import (
	"log"
	"recipes-core-api/pkg/db"
	"recipes-core-api/pkg/middleware"
	"recipes-core-api/pkg/rest/demo"
	"recipes-core-api/pkg/rest/ingredient"
	"recipes-core-api/pkg/rest/recipe"
	"recipes-core-api/pkg/rest/unit"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// read db config from environment
	dbConfig, err := db.NewConfigFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// setup db connection
	dbConnection, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// setup validator
	validate := validator.New()

	// middleware
	r.Use(middleware.LoggerMiddleware())

	// register routes
	ingredientHandler := ingredient.NewHandler(dbConnection, validate)
	ingredient.RegisterRoutes(r, ingredientHandler)

	recipeHandler := recipe.NewHandler(dbConnection)
	recipe.RegisterRoutes(r, recipeHandler)

	unitHandler := unit.NewHandler(dbConnection)
	unit.RegisterRoutes(r, unitHandler)

	demoHandler := demo.NewHandler(dbConnection)
	demo.RegisterRoutes(r, demoHandler)

	err = r.Run(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
