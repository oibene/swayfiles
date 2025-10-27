package main

import (
	"fmt"
	"urbanAPI/controller"
)

func main() {
	data := controller.GetAllCustomer()
	fmt.Println(data)
}
