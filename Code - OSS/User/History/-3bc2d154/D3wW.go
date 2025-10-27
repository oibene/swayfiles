package main

import (
	"fmt"
	"urbanAPI/controller"
)

func main() {

	data, err := controller.GetCustomerByID(1)
	fmt.Println(data, err)
}
