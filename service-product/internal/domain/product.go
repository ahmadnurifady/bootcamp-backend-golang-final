package domain

type Product struct {
	ID          string  `json:"id"`
	ProductName string  `json:"productName"`
	Stock       int     `json:"stock"`
	Price       float32 `json:"price"`
}
