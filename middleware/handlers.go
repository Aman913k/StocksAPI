package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Aman913k/STOCKSAPI/database"
	"github.com/Aman913k/STOCKSAPI/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id"`
	Message string `json:"message,omitempty"`
}



func createStock(stock model.Stock) int64 {
	db := database.GetConnection()
	defer db.Close()

	query := `INSERT INTO stocksdb(name, price, company) VALUES($1, $2, $3) RETURNING stock_id`
	var id int64

	err := db.QueryRow(query, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}







func getStock(id int64) (*model.Stock, error) {
	db := database.GetConnection()
	defer db.Close()
	query := `SELECT *FROM stocksdb WHERE stock_id = $1`

	var stock model.Stock
	row := db.QueryRow(query, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return &stock, nil

	case nil:
		return &stock, nil

	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return &stock, err

}

func getAllStocks() (*[]model.Stock, error) {
	db := database.GetConnection()
	defer db.Close()

	query := `SELECT *FROM stocksdb`

	var stocks []model.Stock

	rows, err := db.Query(query)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var stock model.Stock

		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		stocks = append(stocks, stock)
	}

	return &stocks, err

}


func updateStock(id int64, stock model.Stock) int64 {
	db := database.GetConnection()
	defer db.Close()

	query := `UPDATE stocksdb SET name=$2, price=$3, company=$4 where stock_id=$1`

	res, err := db.Exec(query, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)

	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := database.GetConnection()

	defer db.Close()
	query := `DELETE FROM stocksdb WHERE stock_id=$1`

	res, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)

	return rowsAffected

}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////





func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("Unable to get all stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)

}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Unable to get stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock model.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := createStock(stock)

	res := response{
		ID:      insertID,
		Message: "stock created Successfully",
	}

	json.NewEncoder(w).Encode(res)

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	var stock model.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock completed successfuly. Total rows affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock deleted seccessfully. Total rows affected %v", deletedRows)

	res := response {
		ID: int64(id),
		Message: msg,
	}
	
	json.NewEncoder(w).Encode(res)

}
