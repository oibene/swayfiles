package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Product struct {
	Customer_id int     `db:"customer_id" json:",omitempty"`
	Name        string  `db:"name" json:",omitempty"`
	Last_name   *string `db:"last_name" json:",omitempty"`
	Email       string  `db:"email" json:",omitempty"`
	Password    string  `db:"password" json:",omitempty"`
	Img_URL     *string `db:"img_URL" json:",omitempty"`
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	var p []Product
	data := database.ConnectDB()

	var input models.GetProductInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&p, `SELECT *
							FROM product 
							WHERE product_id = $1`, input.Product_id)

	if err != nil {
		log.Println("Erro ao consultar produto!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
