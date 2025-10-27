package models

type GetCustomerInput struct {
	Customer_id int
}

type GetProductInput struct {
	Product_id    int
	Category_code int
	Gender        string
	Size          string
	MinPrice      string
	MaxPrice      string
}
