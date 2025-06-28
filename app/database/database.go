package database

// Package database provides functions to initialize and manage the database connection.

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
		log.Fatal("Error to connect to database:", err)
	}
	return db, err
}

func GetDB() *sql.DB {
	return db
}