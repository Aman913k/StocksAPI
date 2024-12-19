package model

type Stock struct {
	StockID int64  `json:"stockID"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}
