# StocksAPI

StocksAPI is a RESTful API designed to manage stock data using Go (Golang) and PostgreSQL. It provides endpoints for CRUD operations on stock records, enabling efficient and scalable stock management.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete stock records.
- **PostgreSQL Integration**: Seamless integration with PostgreSQL for persistent storage.
- **Environment Configuration**: Easily configurable via `.env` file.
- **Middleware**: Includes middleware for database connections and error handling.
- **Scalable Design**: Structured to support scalability and maintainability.

## Getting Started

Follow the instructions below to set up and run the project on your local machine.

### Prerequisites

- [Go](https://golang.org/doc/install) (v1.20 or higher recommended)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Aman913k/StocksAPI.git
   cd StocksAPI

2. **Set up the database**:
   Create a PostgreSQL database (e.g., stocksdb).
   Update the .env file with your database credentials:
   POSTGRES_URL="postgres://username:password@localhost:5432/stocksdb"

4. **Install dependencies**:
   ```bash
   go mod tidy

6. **Run the application**:
    ```bash
    go run main.go
