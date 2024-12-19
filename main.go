package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aman913k/STOCKSAPI/database"
	"github.com/Aman913k/STOCKSAPI/router"
)

func main() {
	r := router.Router()

	db := database.GetConnection()
	defer db.Close()

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS stocksdb (
			stock_id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			price BIGINT NOT NULL,
			company TEXT NOT NULL
		);`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create stock table: %v", err)
	}
	log.Println("Stock table is ready.")

	http.ListenAndServe(":5000", r)
	fmt.Println("Server is running....")

}
