package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Product struct {
	Product_id     int             `db:"product_id" json:",omitempty"`
	Product_name   string          `db:"product_name" json:",omitempty"`
	Category_code  int             `db:"category_code" json:",omitempty"`
	Gender         string          `db:"gender" json:",omitempty"`
	Size           string          `db:"size" json:",omitempty"`
	Color          *string         `db:"color" json:",omitempty"`
	Model_code     int             `db:"model_code" json:",omitempty"`
	Price          sql.NullFloat64 `db:"price" json:",omitempty"`
	Descount_price string          `db:"descount_price" json:",omitempty"`
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

func GetProductByFilter(w http.ResponseWriter, r *http.Request) {
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
							WHERE category_code = $1,
							AND gender = $2,
							AND size = $3,
							AND price BETWEEN $4 AND $5;`,
		input.Category_code, input.Gender, input.Size, input.MinPrice, input.MaxPrice)

	if err != nil {
		log.Println("Erro ao consultar produto!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
