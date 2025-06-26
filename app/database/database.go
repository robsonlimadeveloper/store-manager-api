package database

import (
	"database/sql"
	"log"
	"store-manager-api/app/database/config"
)

var db *sql.DB

func Init() (*sql.DB, error) {
	dbURL := config.GetStringDatabaseURL()

	var err error
	db, err := sql.Open("postgres", dbURL)
	
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	return db, err
}

func GetDB() *sql.DB {
	return db
}