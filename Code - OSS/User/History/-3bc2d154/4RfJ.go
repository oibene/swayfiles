package main

import (
	"urbanAPI/controller"

	"github.com/gorilla/mux"
)

func main() {
	point := mux.NewRouter()
	prefix := "/api"

	point.HandleFunc(prefix+"customer", controller.GetCustomerByID).Methods("GET")
}
