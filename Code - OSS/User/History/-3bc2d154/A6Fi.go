package main

import (
	"urbanAPI/database"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDB()

}
