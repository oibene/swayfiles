package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Customer struct {
	ID string `json:"id_customer"`
	FName string `json:"first_name"`
	LName string `json:"last_name"`
	Email string `json:"email"`
	Pass string `json:"password"`
	IMG string `json:"img_URL"`
}

var customers = []Customer{
	{id_customer: "1", first_name: "Icaro", last_name: "O", email: "icaro@email.com", password:"123teste123", img_URL: ""},
	{id_customer: "2", first_name: "T", last_name: "Nagata", email: "nagata@email.com", password:"123teste453", img_URL: ""},
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func main() {
	router := gin.Default()
	router.GET("/api/customers", getCustomers)

	router.Run("localhost:3000")
}