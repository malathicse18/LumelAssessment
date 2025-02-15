package models

type Customer struct {
	CustomerID   string `bson:"customer_id"`
	CustomerName string `bson:"customer_name"`
	Email        string `bson:"email"`
	Address      string `bson:"address"`
}
