package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // postgres
)

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		"user=urbanAdm password=urbansoul dbname=urbansouldb sslmode=disable")

	if err != nil {
		log.Println("Not Connected!", err)
		return nil, err
	}
	return db, nil
}
