package main

import (
	"fmt"
	"urbanAPI/controller"
)

func main() {
	data := controller.GetCustomerByID(2)
	fmt.Println(data)
}
