package main

import (
	"log"
	"urbanAPI/database"
)

func main() {
	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

}
