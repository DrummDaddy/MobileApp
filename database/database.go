package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Подключение к PostgreSQL
func Connect() (*sql.DB, error) {
	connStr := "user=postgres dbname=mobileapp_utf8 sslmode=disable password=159357 client_encoding=UTF8"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error to connecting to database %v", err)
	}
	return db, nil
}
