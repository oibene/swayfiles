package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Customer struct {
	Customer_id int     `db:"customer_id"`
	First_name  string  `db:"first_name"`
	Last_name   string  `db:"last_name"`
	Email       string  `db:"email"`
	Password    string  `db:"password"`
	Img_URL     *string `db:"img_url"`
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) (ct []Customer) {
	data := database.ConnectDB()

	var input models.GetCustomerInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&ct, `SELECT * FROM customer WHERE customer_id = $1`, id)

	if err != nil {
		log.Println("Erro ao consultar usuario!", err)
		return nil
	}
	return ct
}
