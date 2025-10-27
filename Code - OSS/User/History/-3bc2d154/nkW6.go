package main

import (
	"fmt"
	"urbanAPI/controller"
)

func main() {
	data := controller.GetAllCustomers()
	fmt.Println(data)
}
