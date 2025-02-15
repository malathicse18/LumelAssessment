package routes // Change from 'package controllers' to 'package routes'

import (
	"LumelAssessment/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(router *gin.Engine) {
	customerRoutes := router.Group("/customers")
	{
		customerRoutes.GET("/", controllers.GetCustomers)
	}
}
