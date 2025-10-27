package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	customer_id int
	first_name  string
	last_name   string
	email       string
	password    string
	img_URL     string
}

func GetAllUsers(cnn *gin.Context) {
	ct := Customer{}

	data, err := cnn.GetRawData()

	if err != nil {
		cnn.AbortWithStatusJSON(400, "Customer Not Defined")
		return
	}

	err = json.Unmarshal(data, &ct)
	if err != nil {
		cnn.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	_, err = database.DB.Exec("SELECT * FROM customer")
	if err != nil {
		cnn.AbortWithStatusJSON(400, "Error on select customers")
		panic(err)
	} else {
		cnn.JSON(http.StatusOK, "Customer found!")
	}
}
