package services

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"LumelAssessment/config"
	"LumelAssessment/models"
)

// LoadCSVToDB reads a CSV file from the "data" folder and inserts data into MongoDB
func LoadCSVToDB(filename string) error {
	collection := config.GetCollection("orders")

	// Get the absolute path of the CSV file
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("Error getting absolute file path: %v", err)
	}

	// Open the CSV file
	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading CSV file: %v", err)
	}

	var orders []interface{}
	for i, record := range records {
		if i == 0 {
			continue 
		}

		// Convert numeric fields
		quantity, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)

		// Convert date format 
		dateOfSale, err := time.Parse("2006-01-02", record[6])
		if err != nil {
			log.Printf("⚠ Skipping record due to invalid date: %v", err)
			continue
		}

		order := models.Order{
			OrderID:       record[0],
			ProductID:     record[1],
			CustomerID:    record[2],
			ProductName:   record[3],
			Category:      record[4],
			Region:        record[5],
			DateOfSale:    dateOfSale,
			QuantitySold:  quantity,
			UnitPrice:     unitPrice,
			Discount:      discount,
			ShippingCost:  shippingCost,
			PaymentMethod: record[11],
			CustomerName:  record[12],
			CustomerEmail: record[13],
			CustomerAddr:  record[14],
		}

		orders = append(orders, order)
	}

	// Insert all orders into MongoDB
	if len(orders) > 0 {
		_, err := collection.InsertMany(context.TODO(), orders)
		if err != nil {
			return fmt.Errorf("❌ Error inserting records: %v", err)
		}
		fmt.Println("CSV data successfully inserted into MongoDB!")
	} else {
		fmt.Println("⚠ No valid records to insert.")
	}

	return nil
}
