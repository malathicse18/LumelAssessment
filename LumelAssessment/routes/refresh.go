package routes

import (
	"LumelAssessment/controllers"

	"github.com/gin-gonic/gin"
)

// RefreshRoutes registers the CSV loader API route
func RefreshRoutes(router *gin.Engine) {
	router.GET("/refresh", controllers.RefreshDataFromCSV)
}
