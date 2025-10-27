package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Product_Images struct {
	Product_id int    `db:"product_id" json:",omitempty"`
	Img_url    string `db:"img_url" json:",omitempty"`
}

type Product_Category struct {
	Category_code int `db:"category_code" json:,omitempty`
	Description   int `db:"description" json:,omitempty`
}

type Product_Model struct {
	Model_code  int `db:"model_code" json:,omitempty`
	Description int `db:"description" json:,omitempty`
	Notes       int `db:"notes" json:,omitempty`
	Composition int `db:"composition" json:,omitempty`
}

type Product struct {
	Product_id     int     `db:"product_id" json:",omitempty"`
	Product_name   string  `db:"product_name" json:",omitempty"`
	Gender         string  `db:"gender" json:",omitempty"`
	Size           string  `db:"size" json:",omitempty"`
	Color          *string `db:"color" json:",omitempty"`
	Price          string  `db:"price" json:",omitempty"`
	Descount_price string  `db:"descount_price" json:",omitempty"`

	Category    []Product_Category
	Images      []Product_Images
	Description []Product_Model
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

	err = data.Select(&p, `SELECT p.product_name, p.gender, p.size, p.color, p.price, p.descount_price,
						c.description as category,
						d.description, d.notes, d.composition,
						i.img_url

						FROM product p 

						inner join category c
						on c.category_code = p.category_code

						inner join description d
						on d.model_code = p.model_code


						WHERE p.product_id = $1`,
		input.Product_id)

	if err != nil {
		log.Println("Erro ao consultar produto!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
