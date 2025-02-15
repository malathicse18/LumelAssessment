package controllers

import (
	"LumelAssessment/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetCustomers handles the API request for retrieving customers
func GetCustomers(c *gin.Context) {
	var customers []bson.M

	// Aggregation query to get distinct customer details from the orders collection
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": "$customer_id", // Group by customer_id
				"customer_name": bson.M{
					"$first": "$customer_name", // Take the first customer_name
				},
				"customer_email": bson.M{
					"$first": "$customer_email", // Take the first customer_email
				},
				"customer_address": bson.M{
					"$first": "$customer_address", // Take the first customer_address
				},
			},
		},
	}

	// Execute the aggregation
	cursor, err := config.DB.Collection("orders").Aggregate(c, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving customers",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(c)

	// Decode the result into the customers slice
	if err := cursor.All(c, &customers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding customers",
			"error":   err.Error(),
		})
		return
	}

	// Return the list of customers
	c.JSON(http.StatusOK, gin.H{
		"data":    customers,
		"message": "Customers retrieved successfully",
	})
}
