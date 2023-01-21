package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"recipes-core-api/api/v1/models"
	"recipes-core-api/internal/mongo"
	"time"
)

var unitCollection = mongo.GetCollection(mongo.DB, "units")

func CreateUnit() {
	// TODO implement
}

func GetAllUnits() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var units []models.Unit
		defer cancel()

		results, err := unitCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ApiResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		for results.Next(ctx) {
			var unit models.Unit
			if err = results.Decode(&unit); err != nil {
				c.JSON(http.StatusInternalServerError, models.ApiResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			}

			units = append(units, unit)
		}

		c.JSON(http.StatusOK, models.ApiResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": units}},
		)
	}
}
