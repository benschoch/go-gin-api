package ingredient

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	r.POST("/ingredients", handler.Create())
	//	router.PUT("/ingredients/:userId", controllers.EditAUser())
	//	router.DELETE("/ingredients/:userId", controllers.DeleteAUser())
	r.GET("/ingredients/:id", handler.GetByID())
	r.GET("/ingredients", handler.GetAll())
}
