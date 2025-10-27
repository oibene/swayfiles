package controller

import (
	"log"
	"urbanAPI/database"
)

type Customer struct {
	id         int
	first_name string
	last_name  string
	email      string
	password   string
	img_URL    string
}

func GetCustomerByID(id int) (ct *Customer) {
	data := database.ConnectDB()
	err := data.Select(&ct, `SELECT * FROM customer WHERE customer_id = ?`, id)

	if err != nil {
		log.Println(id)
		log.Println("Customer not found!", err)
		return nil
	}
	return ct
}

func GetAllCustomers() (cts *[]Customer) {

	data := database.ConnectDB()
	err := data.Select(&cts, `SELECT * FROM customer`)

	if err != nil {
		log.Println("Customers not found!", err)
		return nil
	}
	return cts
}
