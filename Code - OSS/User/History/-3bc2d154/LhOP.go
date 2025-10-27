package main

import (
	"fmt"
	"urbanAPI/controller"

	"github.com/gorilla/mux"
)

func main() {
	point := mux.NewRouter()
	prefix := "/api"

	point.HandleFunc()

	data := controller.GetCustomerByID(2)
	fmt.Println(data)
}
