package webs

type ProductDTO struct {
	ProductName  string `json:"product_name" form:"product_name"`
	ProductPrice int    `json:"product_price" form:"product_price"`
	ProductStock int    `json:"product_stock" form:"product_stock"`
	UserID       uint   `json:"user_id" form:"user_id"`
	CompanyID    uint   `json:"company_id" form:"company_id"`
}

type ProductBodyResponse struct {
	ID           uint        `json:"id"`
	ProductName  string      `json:"product_name" form:"product_name"`
	ProductSlug  string      `json:"product_slug" form:"product_slug"`
	ProductPrice int         `json:"product_price" form:"product_price"`
	ProductStock int         `json:"product_stock" form:"product_stock"`
	User         interface{} `json:"user"`
	Company      interface{} `json:"company"`
}
