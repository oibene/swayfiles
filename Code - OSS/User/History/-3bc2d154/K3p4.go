package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) //to int
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	connStr := fmt.Sprintf("host=%s port%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(err)
	}

	DB = db
	fmt.Printf("Connected!")

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	route.SetTrustedProxies([]string{"localhost"})

	router := gin.Default()
	router.Run("localhost:3000")
}
