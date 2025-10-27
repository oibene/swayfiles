package database

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // postgres
)

func ConnectDB() {
	db, err := sqlx.Connect("postgres",
		"user=urbanAdm password=urbansoul dbname=urbansouldb sslmode=disable")
}
