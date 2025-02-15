package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB not responding:", err)
	}

	log.Println("✅ Connected to MongoDB")
	DB = client.Database("salesdb")

	// Ensure collection exists by inserting a dummy record
	_, err = DB.Collection("orders").InsertOne(context.TODO(), map[string]interface{}{"init": "test"})
	if err != nil {
		log.Println("⚠ Collection 'orders' already exists, skipping init")
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
