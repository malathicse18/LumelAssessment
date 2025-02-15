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
	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("⚠ Error opening log file: %v", err)
	}
	defer logFile.Close()

	multiLogger := log.New(logFile, "", log.Ldate|log.Ltime)
	log.SetOutput(multiLogger.Writer())

	log.Println("Starting application...")

	config.ConnectDB()
	log.Println("Connected to MongoDB")

	err = services.LoadCSVToDB("data/data.csv")
	if err != nil {
		log.Printf("⚠ CSV Load Error: %v", err)
	} else {
		log.Println("CSV data successfully loaded into MongoDB!")
	}

	router := gin.Default()

	routes.RegisterRevenueRoutes(router)
	routes.RegisterProductRoutes(router)
	routes.RegisterCustomerRoutes(router)
	routes.RefreshRoutes(router)

	log.Println("Routes successfully registered")

	log.Println("Server running on http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("⚠ Error starting server: %v", err)
	}
}
