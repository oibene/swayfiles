package controller

import (
	"log"
	"urbanAPI/database"
)

type Customer struct {
	Customer_id int     `db:"customer_id"`
	First_name  string  `db:"first_name"`
	Last_name   string  `db:"last_name"`
	Email       string  `db:"email"`
	Password    string  `db:"password"`
	Img_URL     *string `db:"img_URL"`
}

func GetCustomerByID(id int) (ct *Customer) {
	data := database.ConnectDB()
	err := data.Select(&ct, `SELECT * FROM customer WHERE customer_id = ?`, id)

	if err != nil {
		log.Println("Erro ao consultar usuario!", err)
		return nil
	}
	return ct
}

func GetAllCustomers() (cts []Customer) {

	data := database.ConnectDB()
	err := data.Select(&cts, `SELECT * FROM customer`)

	if err != nil {
		log.Println("Erro ao consultar usuarios!", err)
		return nil
	}
	return cts
}
