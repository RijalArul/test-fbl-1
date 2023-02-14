package webs

import "time"

type TransactionDTO struct {
	TotalQuantity int `json:"total_quantity" form:"total_quantity"`
	TotalPrice    int `json:"total_price" form:"total_price"`
	UserID        uint
	CompanyID     uint `json:"company_id" form:"company_id"`
	ProductID     uint `json:"product_id" form:"product_id"`
}

type TransactionResponse struct {
	CreatedAt     *time.Time `json:"created_at"`
	CompanyName   string     `json:"company_name"`
	ProductName   string     `json:"product_name"`
	TotalQuantity int        `json:"total_quantity"`
	Price         int        `json:"price"`
	TotalPrice    int        `json:"total_price"`
	RestStock     int        `json:"rest_stock"`
}
