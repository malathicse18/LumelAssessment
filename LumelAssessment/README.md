  LumelAssessment Project

LumelAssessment Project
=======================

Welcome to the LumelAssessment Project! This is a Go-based backend service that interacts with MongoDB, loads CSV data into the database, and provides APIs to retrieve revenue, product, and customer data. This project uses the Gin web framework and MongoDB.

Table of Contents
-----------------

*   [Prerequisites](#prerequisites)
*   [Step 1: Clone the Repository](#step-1-clone-the-repository)
*   [Step 2: Install Dependencies](#step-2-install-dependencies)
*   [Step 3: Configure MongoDB](#step-3-configure-mongodb)
*   [Step 4: Running the Application](#step-4-running-the-application)
*   [Step 5: API Endpoints](#step-5-api-endpoints)
*   [Step 6: Troubleshooting](#step-6-troubleshooting)

Prerequisites
-------------

Before you begin, ensure you have the following installed on your machine:

*   **Go**: Version 1.18 or higher
    *   Check your Go version with: `go version`
    *   Install Go from [here](https://golang.org/dl/).
*   **MongoDB**: Version 5.0 or higher
    *   Install MongoDB from [here](https://www.mongodb.com/try/download/community).
*   **Gin Web Framework**: This project uses the Gin framework, which is a lightweight web framework for Go.
*   **MongoDB Go Driver**: To interact with MongoDB.

Step 1: Clone the Repository
----------------------------

Clone the repository to your local machine using the following command:

    git clone https://github.com/yourusername/LumelAssessment.git

Navigate to the project directory:

    cd LumelAssessment

Step 2: Install Dependencies
----------------------------

To set up the necessary dependencies for the project, follow these steps.

### Install Go Dependencies

Before you can run the application, you need to install the required Go dependencies. The primary dependency for this project is the Gin web framework.

#### Step 1: Install Gin Web Framework

Use the following `go get` command to install the Gin framework:

    go get -u github.com/gin-gonic/gin

#### Step 2: Install MongoDB Go Driver

You also need to install the MongoDB Go driver to connect to the MongoDB database:

    go get go.mongodb.org/mongo-driver/mongo

Step 3: Configure MongoDB
-------------------------

### 1\. Install MongoDB

If you don't have MongoDB installed, follow these steps:

*   **Windows**: Download the installer from MongoDB's [official site](https://www.mongodb.com/try/download/community) and follow the installation instructions.
*   **Linux**: You can install MongoDB by following the guide on the [MongoDB Linux installation page](https://docs.mongodb.com/manual/administration/install-on-linux/).
*   **macOS**: Install MongoDB via Homebrew by running the command:
    
        brew tap mongodb/brew
        brew install mongodb-community
    

### 2\. Set Up MongoDB

Ensure MongoDB is running before proceeding. The default port for MongoDB is 27017. If you need to change the default settings, modify the `main.go` file where MongoDB settings are configured (usually in a configuration section).

Step 4: Running the Application
-------------------------------

### 1\. Ensure MongoDB is Running

Before you start the application, make sure your MongoDB server is running. If MongoDB is not running, start it by following the instructions in Step 3: Configure MongoDB.

### 2\. Navigate to the Project Directory

In your terminal or command prompt, navigate to the root folder of your project where the `main.go` file is located:

    cd path/to/LumelAssessment

### 3\. Run the Application

Now, you can run the application using the following command:

    go run main.go

If the application starts successfully, you should see the following output:

    ✅ Connected to MongoDB
    🚀 Server running on http://localhost:8080

Step 5: API Endpoints
---------------------

The following API endpoints are available in the application:

*   `GET /revenue`: Fetches the revenue data.
*   `GET /products`: Fetches the product data.
*   `GET /customers`: Fetches customer data.

Each of these endpoints provides a response with the relevant data retrieved from the MongoDB database.

Step 6: Troubleshooting
-----------------------

In case you encounter any issues, here are some common problems and their solutions:

### 1\. MongoDB Connection Issues

If the application fails to connect to MongoDB or you receive an error related to the database connection, follow these steps:

Ensure that MongoDB is running on your machine. If it is not, start MongoDB using the following command:

*   **Windows (using MongoDB installed as a service)**:
    
        net start MongoDB
    
*   **Windows (running MongoDB manually)**:
    
        mongod
    
*   **Linux/macOS**:
    
        sudo systemctl start mongod
    

### 2\. CSV Loading Errors

If you encounter issues loading the CSV file into MongoDB, verify that the file path is correct. The default file path is `data/data.csv`, but this can be adjusted in the `main.go` file. Make sure the CSV file is in the correct format and contains the expected data.

### 3\. API Not Responding

If the API is not responding as expected, check if the server is running and ensure that MongoDB is properly connected. If issues persist, check the application logs for errors.
