package demo

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.Any("/demo", CreateDemoHandler())
}
