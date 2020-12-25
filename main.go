package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"onlineBank/db"
	"onlineBank/pkg/core/services"
)

func main() {
	DB, err := sql.Open("sqlite3", "bank")
	if err != nil {
		log.Fatal("Cannot open database, err is:", err)
	}
	err = db.DatabaseInit(DB)
	if err != nil {
		log.Fatal("Cannot create tables, error is:", err)
	}
	Start(DB)
}
func Start(Db *sql.DB) {
	for {
		services.AuthorizeAndStart(Db)
	}
}
