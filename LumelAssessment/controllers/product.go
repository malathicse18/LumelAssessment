package controllers

import (
	"LumelAssessment/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetProducts handles the API request for retrieving products
func GetProducts(c *gin.Context) {
	var products []bson.M

	// Aggregation query to get distinct products from the orders collection
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": "$product_name", // Group by product name
				"category": bson.M{
					"$first": "$category", // Take the first category from the grouped data
				},
				"product_id": bson.M{
					"$first": "$product_id", // Take the first product_id
				},
			},
		},
	}

	// Execute the aggregation
	cursor, err := config.DB.Collection("orders").Aggregate(c, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving products",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(c)

	// Decode the result into the products slice
	if err := cursor.All(c, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding products",
			"error":   err.Error(),
		})
		return
	}

	// Return the list of products
	c.JSON(http.StatusOK, gin.H{
		"data":    products,
		"message": "Products retrieved successfully",
	})
}
