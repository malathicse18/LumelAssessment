package models

import "time"

type Order struct {
	OrderID       string    `bson:"order_id"`
	ProductID     string    `bson:"product_id"`
	CustomerID    string    `bson:"customer_id"`
	ProductName   string    `bson:"product_name"`
	Category      string    `bson:"category"`
	Region        string    `bson:"region"`
	DateOfSale    time.Time `bson:"date_of_sale"`
	QuantitySold  int       `bson:"quantity_sold"`
	UnitPrice     float64   `bson:"unit_price"`
	Discount      float64   `bson:"discount"`
	ShippingCost  float64   `bson:"shipping_cost"`
	PaymentMethod string    `bson:"payment_method"`
	CustomerName  string    `bson:"customer_name"`
	CustomerEmail string    `bson:"customer_email"`
	CustomerAddr  string    `bson:"customer_address"`
}
