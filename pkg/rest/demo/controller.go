package demo

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	r.Any("/demo", handler.Demo())
}
