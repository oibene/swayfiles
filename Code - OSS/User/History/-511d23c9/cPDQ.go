package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Customer struct {
	id_customer string `json:id_customer`
	first_name string `json:first_name`
	last_name string `json:last_name`
	email string `json:email`
	password string `json:password`
	img_URL string `json:img_URL`
}

var customers = []Customer {
	{id_customer: "1", first_name: "Icaro", last_name: "O", email: "icaro@email.com", senha:"123teste123", img_URL: ""},
	{id_customer: "2", first_name: "T", last_name: "Nagata", email: "nagata@email.com", senha:"123teste453", img_URL: ""},
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func main() {
	router := gin.Default()
	router.GET("/customers", customers)

	router.Run("localhost:3000/api")
}