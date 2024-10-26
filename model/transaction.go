package model

type Transaction struct {
	ID         string  `json:"id"`
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
