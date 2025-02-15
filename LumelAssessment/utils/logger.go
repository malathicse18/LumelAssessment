package utils

import (
	"log"
	"os"
)

// Initialize a global logger to be used across the application
var Logger *log.Logger

// Initialize function for setting up logger
func InitLogger() {
	file, err := os.OpenFile("logs/data_refresh.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	Logger = log.New(file, "", log.LstdFlags)
}

// Log function for logging messages
func LogMessage(message string) {
	if Logger != nil {
		Logger.Println(message)
	}
}
