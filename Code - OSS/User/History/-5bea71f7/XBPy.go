package repository

import (
	"log"
	"urbanAPI/database"
)

func GetProductByCategory() {
	data := database.ConnectDB()

	err = data.Select(&p, `SELECT *
							FROM product 
							WHERE product_id = $1`, input.Product_id)

	if err != nil {
		log.Println("Erro ao consultar produto!", err)
		return
	}

}
