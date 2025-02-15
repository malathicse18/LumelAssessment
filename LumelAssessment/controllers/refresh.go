package controllers

import (
	"LumelAssessment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RefreshDataFromCSV handles API requests to load CSV data into MongoDB
func RefreshDataFromCSV(c *gin.Context) {
	err := services.LoadCSVToDB("data/sales_data.csv") // Adjust file path if needed
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CSV data successfully loaded into MongoDB!"})
}
