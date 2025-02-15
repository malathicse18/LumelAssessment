package services

import (
	"context"
	"log"
	"time"

	"LumelAssessment/config"

	"go.mongodb.org/mongo-driver/bson"
)

// Calculate total revenue within a date range
func GetTotalRevenue(startDate, endDate string) (float64, error) {
	collection := config.GetCollection("orders")

	// Convert date strings to time.Time format
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return 0, err
	}

	// Query MongoDB to sum total revenue in the given date range
	pipeline := []bson.M{
		{"$match": bson.M{"date_of_sale": bson.M{"$gte": start, "$lte": end}}},
		{"$group": bson.M{"_id": nil, "total_revenue": bson.M{"$sum": "$total_price"}}},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Println("Error in aggregation:", err)
		return 0, err
	}
	defer cursor.Close(context.TODO())

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return 0, err
	}

	// Return total revenue
	if len(result) > 0 {
		return result[0]["total_revenue"].(float64), nil
	}
	return 0, nil
}
