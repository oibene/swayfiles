package models

type GetCustomerInput struct {
	Customer_id int
}

type GetProductInput struct {
	Product_id    int
	Order_Item_id int
}

type GetCommentsInput struct {
	Customer_id int
	Product_id  int
}
