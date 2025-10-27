package main

import (
	"fmt"
	"urbanAPI/controller"
)

func main() {
	data := controller.GetCustomerByID()
	fmt.Println(data)
}
