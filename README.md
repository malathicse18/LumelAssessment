# Lumel Assessment

## Overview
Lumel Assessment is a Go-based web application designed to handle sales data analytics. It connects to a MongoDB database to store and retrieve sales-related data. The application provides APIs for revenue calculations, product analytics, and customer insights, as well as a data refresh feature for loading CSV data into the database.

LumelAssessment/
│── 📂 config/          # Configuration files (DB connection)
│   ├── db.go          # MongoDB connection setup
│── 📂 models/          # Structs for MongoDB documents
│   ├── order.go       # Order model
│   ├── customer.go    # Customer model
│   ├── product.go     # Product model
│── 📂 routes/          # API routes
│   ├── revenue.go     # Revenue-related routes
    ├── refresh.go
│   ├── product.go     # Product-related routes
│   ├── customer.go    # Customer-related routes
│── 📂 controllers/     # Business logic & handlers
│   ├── revenue.go     # Revenue calculations
│   ├── product.go     # Product analytics
│   ├── customer.go    # Customer analytics
│   ├── refresh.go     # Data refresh controller
│── 📂 services/        # Helper functions & services
│   ├── revenue.go     # Revenue calculations logic
│   ├── csv_loader.go  # CSV loading and parsing logic
│── 📂 utils/           # Utility functions
│   ├── logger.go      # Logging helper
│── 📂 data/            # Folder for CSV files
│   ├── data.csv # Example sales CSV
│── main.go            # Entry point
│── go.mod             # Go module file
│── README.md          # Documentation   


---

## Features

1. **Revenue Calculation API**: 
   - Endpoint: `/revenue`
   - Returns total revenue based on orders stored in the database.

2. **Product Analytics API**: 
   - Endpoint: `/products`
   - Retrieves a list of products and their details.

3. **Customer Analytics API**: 
   - Endpoint: `/customers`
   - Fetches customer data from the database.

4. **CSV Data Refresh API**: 
   - Endpoint: `/refresh`
   - Allows loading and refreshing sales data from a CSV file into the MongoDB database.

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/LumelAssessment.git
   cd LumelAssessment
