package connect

import (
	"database/sql"
	"fmt"
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
	cfg := postgresql.NewConfig()

	cfg.User = "urbanAdm"
	cfg.Passwd - "urbansoul"
	cfg.Net = "tcp"
	cfg.Addr = "localhost:3000"
	cfg.DBName = "urbansouldb"

	ConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		localhost)
}
