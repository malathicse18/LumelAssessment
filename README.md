# LumelAssessment Project

This project implements a web application using the Gin framework, MongoDB, and CSV file import functionality. It provides APIs for managing orders, products, and customers, as well as calculating revenue based on the order data.

## Prerequisites

Before you begin, ensure you have the following software installed:

- **Go** (version 1.22.5 or higher)
  - To check the version: `go version`
  - If not installed, follow the instructions from [Go's official site](https://go.dev/doc/install)
  
- **MongoDB** (version v7.0.9 or higher)
  - Ensure MongoDB is running locally on port `27017` or configure the database connection as needed.

- **Gin Framework** (Go web framework)
  - Install the Gin package using the following command:
    ```bash
    go get -u github.com/gin-gonic/gin
    ```

- **CSV File** (`data/data.csv`)
  - The project expects a CSV file containing order data to be loaded into MongoDB. Ensure that `data/data.csv` is available and formatted correctly.

## Setup

### Step 1: Clone the Repository

Clone the project repository to your local machine:

```bash
git clone <repository-url>
cd LumelAssessment
