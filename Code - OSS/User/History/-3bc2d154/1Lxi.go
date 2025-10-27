package main

import (
	"urbanAPI/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	corsOpt := cors.New(
		cors.Options{
			AllowedOrigins: []string
		}
	)

	point := mux.NewRouter()
	prefix := "/api"

	point.HandleFunc(prefix+"customer", controller.GetCustomerByID).Methods("GET")
}
