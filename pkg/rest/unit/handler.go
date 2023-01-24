package unit

import (
	"context"
	"net/http"
	"recipes-core-api/models"
	"recipes-core-api/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	dbConnection *db.Connection
}

func NewHandler(dbConnection *db.Connection) *Handler {
	return &Handler{dbConnection: dbConnection}
}

func (h *Handler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var units []models.Unit
		defer cancel()

		unitCollection := h.dbConnection.GetCollection("units")
		results, err := unitCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		for results.Next(ctx) {
			var unit models.Unit
			if err = results.Decode(&unit); err != nil {
				c.JSON(http.StatusInternalServerError, models.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			}

			units = append(units, unit)
		}

		c.JSON(http.StatusOK, models.APIResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": units}},
		)
	}
}
