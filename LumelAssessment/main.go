package main

import (
	"LumelAssessment/config"
	"LumelAssessment/routes"
	"LumelAssessment/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Open a log file (create it if it doesn't exist)
	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("⚠ Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set the log output to the log file and also to the console
	multiLogger := log.New(logFile, "", log.Ldate|log.Ltime)
	log.SetOutput(multiLogger.Writer()) // This ensures logs go to both the file and console

	// Log to the file and console
	log.Println("Starting application...")

	// Connect to MongoDB
	config.ConnectDB()
	log.Println("Connected to MongoDB")

	// Load CSV data into MongoDB (if needed)
	err = services.LoadCSVToDB("data/data.csv") // Ensure correct file path
	if err != nil {
		log.Printf("⚠ CSV Load Error: %v", err)
	} else {
		log.Println("CSV data successfully loaded into MongoDB!")
	}

	// Initialize Gin router
	router := gin.Default()

	// Register all routes
	routes.RegisterRevenueRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterCustomerRoutes(router)
	routes.RefreshRoutes(router)

	// Log that routes have been successfully registered
	log.Println("Routes successfully registered")

	// Start the server on port 8080
	log.Println("Server running on http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("⚠ Error starting server: %v", err)
	}
}
