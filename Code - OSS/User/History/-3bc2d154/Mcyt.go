package main

import (
	"fmt"
	"urbanAPI/controller"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func main() {
	point := mux.NewRouter()

	data := controller.GetCustomerByID(2)
	fmt.Println(data)
}
