package main

import (
	"fmt"
	"urbanAPI/controller"
)

func main() {
	data := controller.GetCustomerByID(1)
	fmt.Println(data)
}
