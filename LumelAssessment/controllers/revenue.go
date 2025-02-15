package controllers

import (
	"LumelAssessment/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetRevenueData handles the API request for revenue data
func GetRevenueData(c *gin.Context) {
	var totalRevenue float64

	// Aggregation pipeline for calculating total revenue
	pipeline := bson.A{
		// Project step to calculate revenue per order
		bson.M{
			"$project": bson.M{
				"revenue": bson.M{
					"$add": []interface{}{
						bson.M{"$multiply": []interface{}{"$quantity_sold", "$unit_price"}},              // Quantity * Unit Price
						bson.M{"$multiply": []interface{}{"$quantity_sold", "$unit_price", "$discount"}}, // Discount
						bson.M{"$multiply": []interface{}{"$quantity_sold", "$shipping_cost"}},           // Shipping Cost
					},
				},
			},
		},
		// Group step to sum up the revenue from all orders
		bson.M{
			"$group": bson.M{
				"_id":          nil,
				"totalRevenue": bson.M{"$sum": "$revenue"},
			},
		},
	}

	cursor, err := config.DB.Collection("orders").Aggregate(c, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving revenue data",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(c)

	// If revenue is available
	if cursor.Next(c) {
		var result struct {
			TotalRevenue float64 `bson:"totalRevenue"`
		}
		if err := cursor.Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error decoding revenue data",
				"error":   err.Error(),
			})
			return
		}
		totalRevenue = result.TotalRevenue
	}

	// Return the calculated revenue
	c.JSON(http.StatusOK, gin.H{
		"data":    totalRevenue,
		"message": "Revenue data retrieved successfully",
	})
}
