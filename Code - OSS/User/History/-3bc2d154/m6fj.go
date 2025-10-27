package main

import (
	"net/http"
	"urbanAPI/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	corsOpt := cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodOptions,
				http.MethodHead,
			},
			AllowedHeaders: []string{"*"},
		},
	)

	point := mux.NewRouter()
	prefix := "/api"

	point.HandleFunc(prefix+"customer", controller.GetCustomerByID).Methods("GET")

	http.Handle("/", point)

	http.ListenAndServe(":3000", corsOpt.Handler(point))
}
