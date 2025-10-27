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
	Img_URL     *string `db:"img_url"`
}

func GetCustomerByID(id int) (ct *Customer) {
	data := database.ConnectDB()
	err := data.Select(&ct, `SELECT * FROM customer WHERE customer_id = id=$1`, id)

	if err != nil {
		log.Println("Erro ao consultar usuario!", err)
		return nil
	}
	return ct
}
