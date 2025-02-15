package routes

import (
	"LumelAssessment/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes adds product-related routes
func RegisterProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts)
}
