package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"urbanAPI/database"
	"urbanAPI/models"
)

type Comments struct {
	Comment_id    int    `db:"customer_id" json:",omitempty"`
	Customer_id   int    `db:"customer_id" json:",omitempty"`
	Customer_name string `db:"customer_name" json:",omitempty"`
	Comment       string `db:"comment" json:",omitempty"`
	Rating        int    `db:"rating" json:",omitempty"`
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	var cm []Comments
	data := database.ConnectDB()

	var input models.GetCommentsInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = data.Select(&cm, `SELECT cm.comment, cm.rating,
								ct.name, ct.last_name

							FROM comments cm

							INNER JOIN customer ct on cm.customer_id = ct.customer_id

							WHERE cm.product_id = $1`, input.Product_id)

	if err != nil {
		log.Println("Erro ao consultar usuario!", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cm)
}
