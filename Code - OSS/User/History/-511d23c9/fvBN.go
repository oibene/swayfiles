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
	{ID: "1", FName: "Icaro", LName: "O", Email: "icaro@email.com", Pass:"123teste123", IMG: ""},
	{ID: "2", FName: "T", LName: "Nagata", Email: "nagata@email.com", Pass:"123teste453", IMG: ""},
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func getCustomerByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range customers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}

func postCustomer(c *gin.Context) {
	var newCustomer Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

func main() {
	router := gin.Default()
	router.GET("/api/customers", getCustomers)
	router.POST("/api/customers", postCustomer)

	router.Run("localhost:3000")
}