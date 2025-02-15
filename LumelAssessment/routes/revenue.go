package routes

import (
	"LumelAssessment/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRevenueRoutes adds revenue-related routes to the router
func RegisterRevenueRoutes(router *gin.Engine) {
	router.GET("/revenue", controllers.GetRevenueData) // Ensure function name is correct
}
