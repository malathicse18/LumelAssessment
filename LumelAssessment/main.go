package main

import (
	"LumelAssessment/config"
	"LumelAssessment/routes"
	"LumelAssessment/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()
	fmt.Println("✅ Connected to MongoDB")

	// Load CSV data into MongoDB (if needed)
	err := services.LoadCSVToDB("data/data.csv") // Ensure correct file path
	if err != nil {
		log.Printf("⚠ CSV Load Error: %v", err)
	} else {
		fmt.Println("✅ CSV data successfully loaded into MongoDB!")
	}

	// Initialize Gin router
	router := gin.Default()

	// Register all routes
	routes.RegisterRevenueRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterCustomerRoutes(router)
	routes.RefreshRoutes(router)

	// Log that routes have been successfully registered
	fmt.Println("📡 Routes successfully registered")

	// Start the server on port 8080
	fmt.Println("🚀 Server running on http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("⚠ Error starting server: %v", err)
	}
}
