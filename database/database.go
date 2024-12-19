package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Postgres")

	return db
}
