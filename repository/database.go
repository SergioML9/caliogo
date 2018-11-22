package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBCon ...
var DBCon *sql.DB

// InitDB ...
func InitDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp("+DB_HOST+":"+DB_PORT+")/%s?parseTime=true", DB_USER, DB_PASS, DB_NAME)
	var err error
	DBCon, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

// CloseDB ...
func CloseDB() {
	DBCon.Close()
}
