package models

type Product struct {
	ProductID   string  `bson:"product_id"`
	ProductName string  `bson:"product_name"`
	Category    string  `bson:"category"`
	UnitPrice   float64 `bson:"unit_price"`
}
