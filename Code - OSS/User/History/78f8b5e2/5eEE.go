package connect

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "urbanAdm"
	password = "urbansoul"
	dbname   = "urbansouldb"
)

func main() {
	ConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgresql", ConnStr)
	if err != nil {
		log.Fatal(err)
	}
}
